package preflight

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/wtxue/kok-operator/pkg/constants"
	"github.com/wtxue/kok-operator/pkg/controllers/common"
	"github.com/wtxue/kok-operator/pkg/util/ssh"
	"k8s.io/klog"
)

const (
	ipv4Forward = "/proc/sys/net/ipv4/ip_forward"
)

var tools = []string{"ipvsadm", "modprobe", "modinfo", "ip", "awk", "iptables"}

func newCommonChecks(s ssh.Interface) []Checker {
	return []Checker{
		IsPrivilegedUserCheck{Interface: s},
		CPUArchCeck{Interface: s, Arch: 64},
		KernelCheck{Interface: s, MinKernelVersion: 4, MinMajorVersion: 10},
		// KernelModuleCheck{Interface: s, Module: "iptable_nat"},
		FileContentCheck{Interface: s, Path: ipv4Forward, Content: []byte{'1'}},
		// FileAvailableCheck{Interface: s, Path: constants.KubectlConfigFile},
		DirAvailableCheck{Interface: s, Path: constants.CNIConfDIr},
		DirAvailableCheck{Interface: s, Path: constants.CNIDataDir},
		// PortOpenCheck{Interface: s, port: constants.ProxyHealthzPort},
		// PortOpenCheck{Interface: s, port: constants.ProxyStatusPort},
		PortOpenCheck{Interface: s, port: constants.KubeletPort},
	}
}

// RunMasterChecks checks for master
func RunMasterChecks(s ssh.Interface, c *common.Cluster) error {
	checks := newCommonChecks(s)
	checks = append(checks, []Checker{
		NumCPUCheck{Interface: s, NumCPU: 1},
		DirAvailableCheck{Interface: s, Path: constants.EtcdDataDir},
		PortOpenCheck{Interface: s, port: 6443}, // kube-apiserver
		PortOpenCheck{Interface: s, port: constants.ProxyHealthzPort},
		// PortOpenCheck{Interface: s, port: constants.ProxyStatusPort},
		PortOpenCheck{Interface: s, port: constants.EtcdListenClientPort},
		// PortOpenCheck{Interface: s, port: constants.EtcdListenPeerPort},
	}...)

	for _, tool := range tools {
		checks = append(checks, InPathCheck{Interface: s, executable: tool})
	}

	return RunChecks(checks)
}

// RunNodeChecks checks for node
func RunNodeChecks(s ssh.Interface) error {
	checks := newCommonChecks(s)
	checks = append(checks, []Checker{}...)

	for _, tool := range tools {
		checks = append(checks, InPathCheck{Interface: s, executable: tool})
	}

	return RunChecks(checks)
}

// RunChecks runs each check, displays it's warnings/errors, and once all
// are processed will exit if any errors occurred.
func RunChecks(checks []Checker) error {
	var errsBuffer bytes.Buffer

	for _, c := range checks {
		name := c.Name()
		warnings, errs := c.Check()

		for _, w := range warnings {
			errsBuffer.WriteString(fmt.Sprintf("\t[WARNING %s]: %v\n", name, w))
		}
		for _, i := range errs {
			errsBuffer.WriteString(fmt.Sprintf("\t[ERROR %s]: %v\n", name, i.Error()))
		}
	}
	if errsBuffer.Len() > 0 {
		return &Error{Msg: errsBuffer.String()}
	}
	return nil
}

// Error defines struct for communicating error messages generated by preflight checks
type Error struct {
	Msg string
}

// Error implements the standard error interface
func (e *Error) Error() string {
	return fmt.Sprintf("[preflight] Some fatal errors occurred:\n%s", e.Msg)
}

// Preflight identifies this error as a preflight error
func (e *Error) Preflight() bool {
	return true
}

// Checker validates the state of the system to ensure kubeadm will be
// successful as often as possible.
type Checker interface {
	Check() (warnings, errorList []error)
	Name() string
}

// NumCPUCheck checks if current number of CPUs is not less than required
type NumCPUCheck struct {
	ssh.Interface
	NumCPU int
}

// Name returns the label for NumCPUCheck
func (NumCPUCheck) Name() string {
	return "NumCPU"
}

// Check number of CPUs required by kubeadm
func (ncc NumCPUCheck) Check() (warnings, errorList []error) {
	result, err := ncc.CombinedOutput("getconf _NPROCESSORS_ONLN")
	if err != nil {
		errorList = append(errorList, err)
		return
	}
	numCPU, err := strconv.Atoi(strings.TrimSpace(string(result)))
	if err != nil {
		errorList = append(errorList, err)
		return
	}

	if numCPU < ncc.NumCPU {
		errorList = append(errorList, errors.Errorf("the number of available CPUs %d is less than the required %d", numCPU, ncc.NumCPU))
	}
	return warnings, errorList
}

// CPUArchCeck checks cpu arch
type CPUArchCeck struct {
	ssh.Interface
	Arch int
}

// Name returns the label for CPUArchCeck
func (CPUArchCeck) Name() string {
	return "NumCPU"
}

// Check checks cpu arch
func (cac CPUArchCeck) Check() (warnings, errorList []error) {
	result, err := cac.CombinedOutput("getconf LONG_BIT")
	if err != nil {
		errorList = append(errorList, err)
		return
	}
	arch, err := strconv.Atoi(strings.TrimSpace(string(result)))
	if err != nil {
		errorList = append(errorList, err)
		return
	}

	if arch != cac.Arch {
		errorList = append(errorList, errors.Errorf("only support CPU arch %d, but current is %d", cac.Arch, arch))
	}
	return warnings, errorList
}

// PortOpenCheck ensures the given port is available for use.
type PortOpenCheck struct {
	ssh.Interface
	port  int
	label string
}

// Name returns name for PortOpenCheck. If not known, will return "PortXXXX" based on port number
func (poc PortOpenCheck) Name() string {
	if poc.label != "" {
		return poc.label
	}
	return fmt.Sprintf("Port-%d", poc.port)
}

// Check validates if the particular port is available.
func (poc PortOpenCheck) Check() (warnings, errorList []error) {
	klog.Infof("validating availability of port %d", poc.port)
	errorList = []error{}
	stdout, _, exit, err := poc.Exec(fmt.Sprintf("ss -tl | grep ':%d'", poc.port))
	if err != nil {
		errorList = append(errorList, err)
		return
	}
	if exit == 0 {
		errorList = append(errorList, errors.Errorf("Port %d is in use:%s", poc.port, stdout))
	}

	return nil, errorList
}

// FileAvailableCheck checks that the given file does not already exist.
type FileAvailableCheck struct {
	ssh.Interface
	Path  string
	Label string
}

// Name returns label for individual FileAvailableChecks. If not known, will return based on path.
func (fac FileAvailableCheck) Name() string {
	if fac.Label != "" {
		return fac.Label
	}
	return fmt.Sprintf("FileAvailable-%s", strings.Replace(fac.Path, "/", "-", -1))
}

// Check validates if the given file does not already exist.
func (fac FileAvailableCheck) Check() (warnings, errorList []error) {
	klog.Infof("validating the existence of file %s", fac.Path)
	if _, err := fac.Stat(fac.Path); err == nil {
		errorList = append(errorList, errors.Errorf("%s already exists", fac.Path))
	}
	return nil, errorList
}

// FileContentCheck checks that the given file contains the string Content.
type FileContentCheck struct {
	ssh.Interface
	Path    string
	Content []byte
	Label   string
}

// Name returns label for individual FileContentChecks. If not known, will return based on path.
func (fcc FileContentCheck) Name() string {
	if fcc.Label != "" {
		return fcc.Label
	}
	return fmt.Sprintf("FileContent-%s", strings.Replace(fcc.Path, "/", "-", -1))
}

// Check validates if the given file contains the given content.
func (fcc FileContentCheck) Check() (warnings, errorList []error) {
	klog.Infof("validating the contents of file %s", fcc.Path)
	data, err := fcc.ReadFile(fcc.Path)
	if err != nil {
		return nil, []error{err}
	}

	if !bytes.Equal(data[:len(fcc.Content)], fcc.Content) {
		return nil, []error{errors.Errorf("%s contents are not set to %s", fcc.Path, fcc.Content)}
	}
	return nil, []error{}
}

// InPathCheck checks if the given executable is present in $PATH
type InPathCheck struct {
	ssh.Interface

	executable string
	mandatory  bool
	label      string
	suggestion string
}

// Name returns label for individual InPathCheck. If not known, will return based on path.
func (ipc InPathCheck) Name() string {
	if ipc.label != "" {
		return ipc.label
	}
	return fmt.Sprintf("FileExisting-%s", strings.Replace(ipc.executable, "/", "-", -1))
}

// Check validates if the given executable is present in the path.
func (ipc InPathCheck) Check() (warnings, errs []error) {
	klog.Infof("validating the presence of executable %s", ipc.executable)
	_, err := ipc.LookPath(ipc.executable)
	if err != nil {
		if ipc.mandatory {
			// Return as an error:
			return nil, []error{errors.Errorf("%s not found in system path", ipc.executable)}
		}
		// Return as a warning:
		warningMessage := fmt.Sprintf("%s not found in system path", ipc.executable)
		if ipc.suggestion != "" {
			warningMessage += fmt.Sprintf("\nSuggestion: %s", ipc.suggestion)
		}
		return []error{errors.New(warningMessage)}, nil
	}
	return nil, nil
}

// IsPrivilegedUserCheck verifies user is privileged (linux - root, windows - Administrator)
type IsPrivilegedUserCheck struct {
	ssh.Interface
}

// Name returns name for IsPrivilegedUserCheck
func (IsPrivilegedUserCheck) Name() string {
	return "IsPrivilegedUser"
}

// Check validates if an user has elevated (root) privileges.
func (ipuc IsPrivilegedUserCheck) Check() (warnings, errorList []error) {
	errorList = []error{}

	result, err := ipuc.CombinedOutput("id -u")
	if err != nil {
		errorList = append(errorList, err)
		return
	}
	uid, err := strconv.Atoi(strings.TrimSpace(string(result)))
	if err != nil {
		errorList = append(errorList, err)
		return
	}

	if uid != 0 {
		errorList = append(errorList, errors.New("user is not running as root"))
	}

	return nil, errorList
}

// DirAvailableCheck checks if the given directory either does not exist, or is empty.
type DirAvailableCheck struct {
	ssh.Interface
	Path  string
	Label string
}

// Name returns label for individual DirAvailableChecks. If not known, will return based on path.
func (dac DirAvailableCheck) Name() string {
	if dac.Label != "" {
		return dac.Label
	}
	return fmt.Sprintf("DirAvailable-%s", strings.Replace(dac.Path, "/", "-", -1))
}

// Check validates if a directory does not exist or empty.
func (dac DirAvailableCheck) Check() (warnings, errorList []error) {
	klog.Infof("validating the existence of directory %s", dac.Path)
	if _, err := dac.Stat(dac.Path); err == nil {
		errorList = append(errorList, errors.Errorf("%s already exists", dac.Path))
	}

	return nil, errorList
}

// KernelCheck checks if kernel meet requires
type KernelCheck struct {
	ssh.Interface
	MinKernelVersion int
	MinMajorVersion  int
}

// Name returns label for KernelCheck
func (kc KernelCheck) Name() string {
	return fmt.Sprintf("KernelCheck-%d-%d", kc.MinKernelVersion, kc.MinMajorVersion)
}

// Check validates kernel version
func (kc KernelCheck) Check() (warnings, errorList []error) {
	result, err := kc.CombinedOutput("uname -r")
	if err != nil {
		errorList = append(errorList, err)
		return
	}
	versionStr := strings.TrimSpace(string(result))
	versions := strings.Split(strings.TrimSpace(string(result)), ".")
	if len(versions) < 2 {
		errorList = append(errorList, errors.Errorf("parse version error:%s", versionStr))
		return
	}
	kernelVersion, err := strconv.Atoi(versions[0])
	if err != nil {
		errorList = append(errorList, err)
		return
	}
	majorVersion, err := strconv.Atoi(versions[1])
	if err != nil {
		errorList = append(errorList, err)
		return
	}
	if (kernelVersion < kc.MinKernelVersion) ||
		(kernelVersion == kc.MinKernelVersion && majorVersion < kc.MinMajorVersion) {
		errorList = append(errorList, errors.Errorf("kernel version(%s) must not lower than %d.%d", versionStr, kc.MinKernelVersion, kc.MinMajorVersion))
	}

	return nil, errorList
}

// KernelModuleCheck checks that the given kernel module wheather exists
type KernelModuleCheck struct {
	ssh.Interface
	Module string
	Label  string
}

// Name returns label for individual FileAvailableChecks. If not known, will return based on path.
func (kmc KernelModuleCheck) Name() string {
	if kmc.Label != "" {
		return kmc.Label
	}
	return fmt.Sprintf("KernelModule-%s", strings.Replace(kmc.Module, "/", "-", -1))
}

// Check validates if the given file does not already exist.
func (kmc KernelModuleCheck) Check() (warnings, errorList []error) {
	_, _, exit, err := kmc.Execf("modinfo %s", kmc.Module)
	if err != nil || exit != 0 {
		errorList = append(errorList, errors.Errorf("%s is required", kmc.Module))
	}

	return nil, errorList
}
