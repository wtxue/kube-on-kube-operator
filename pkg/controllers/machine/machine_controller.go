/*
Copyright 2020 wtxue.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package machine

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"time"

	"github.com/pkg/errors"
	devopsv1 "github.com/wtxue/kok-operator/pkg/apis/devops/v1"
	"github.com/wtxue/kok-operator/pkg/gmanager"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	controllerName = "machine"
)

// machineReconciler reconciles a machine object
type machineReconciler struct {
	client.Client
	Log    logr.Logger
	Mgr    manager.Manager
	Scheme *runtime.Scheme
	*gmanager.GManager
}

type manchineContext struct {
	Ctx context.Context
	Key types.NamespacedName
	*devopsv1.Cluster
	*devopsv1.Machine
	*devopsv1.ClusterCredential
	logr.Logger
}

func Add(mgr manager.Manager, pMgr *gmanager.GManager) error {
	reconciler := &machineReconciler{
		Client:   mgr.GetClient(),
		Mgr:      mgr,
		Log:      ctrl.Log.WithName("controller").WithName(controllerName),
		Scheme:   mgr.GetScheme(),
		GManager: pMgr,
	}

	err := reconciler.SetupWithManager(mgr)
	if err != nil {
		return errors.Wrapf(err, "unable to create machine controller")
	}

	return nil
}

func (r *machineReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&devopsv1.Machine{}).
		Complete(r)
}

// +kubebuilder:rbac:groups=devops.k8s.io,resources=virtulclusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=devops.k8s.io,resources=virtulclusters/status,verbs=get;update;patch

func (r *machineReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	logger := r.Log.WithValues(controllerName, req.NamespacedName.String())

	startTime := time.Now()
	defer func() {
		diffTime := time.Since(startTime)
		var logLevel klog.Level
		if diffTime > 1*time.Second {
			logLevel = 1
		} else if diffTime > 100*time.Millisecond {
			logLevel = 2
		} else {
			logLevel = 4
		}
		klog.V(logLevel).Infof("##### [%s] reconciling is finished. time taken: %v. ", req.NamespacedName, diffTime)
	}()

	m := &devopsv1.Machine{}
	err := r.Client.Get(ctx, req.NamespacedName, m)
	if err != nil {
		if apierrors.IsNotFound(err) {
			logger.Error(err, "not find machine")
			return reconcile.Result{}, nil
		}

		logger.Error(err, "failed to get machine")
		return reconcile.Result{}, err
	}

	if m.Spec.Pause == true {
		logger.Info("machine is Pause")
		return reconcile.Result{}, nil
	}

	if len(string(m.Status.Phase)) == 0 {
		m.Status.Phase = devopsv1.MachineInitializing
		err = r.Client.Status().Update(ctx, m)
		if err != nil {
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, nil
	}

	cluster := &devopsv1.Cluster{}
	err = r.Client.Get(ctx, types.NamespacedName{Name: m.Spec.ClusterName, Namespace: m.Namespace}, cluster)
	if err != nil {
		if apierrors.IsNotFound(err) {
			logger.Error(err, "not find cluster")
			return reconcile.Result{}, nil
		}

		logger.Error(err, "failed to get cluster")
		return reconcile.Result{}, err
	}

	if cluster.Status.Phase != devopsv1.ClusterRunning {
		return reconcile.Result{
			Requeue:      true,
			RequeueAfter: 30 * time.Second,
		}, nil
	}

	credential := &devopsv1.ClusterCredential{}
	err = r.Client.Get(ctx, types.NamespacedName{Name: m.Spec.ClusterName, Namespace: m.Namespace}, credential)
	if err != nil {
		if apierrors.IsNotFound(err) {
			klog.V(3).Infof("not find ClusterCredential with name [%q]", req.NamespacedName)
			return reconcile.Result{}, nil
		}

		logger.Error(err, "failed to get ClusterCredential")
		return reconcile.Result{}, err
	}

	klog.Infof("name: %s", cluster.Name)

	r.reconcile(ctx, &manchineContext{
		Key:               req.NamespacedName,
		Logger:            logger,
		Machine:           m,
		Cluster:           cluster,
		ClusterCredential: credential,
	})
	return ctrl.Result{}, nil
}
