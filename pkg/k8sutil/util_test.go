package k8sutil

import "testing"

func TestGetServiceCIDRAndNodeCIDRMaskSize(t *testing.T) {
	type args struct {
		clusterCIDR          string
		maxClusterServiceNum int32
		maxNodePodNum        int32
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   int32
		wantErr bool
	}{
		{
			name: "maxClusterServiceNum == 0",
			args: args{
				clusterCIDR:          "192.168.0.0/24",
				maxClusterServiceNum: 0,
			},
			wantErr: true,
		},
		{
			name: "maxNodePodNum == 0",
			args: args{
				clusterCIDR:          "192.168.0.0/24",
				maxClusterServiceNum: 0,
			},
			wantErr: true,
		},
		{
			name: "maxClusterServiceNum maxNodePodNum < clusterCIDR size",
			args: args{
				clusterCIDR:          "192.168.0.0/24",
				maxClusterServiceNum: 32,
				maxNodePodNum:        64,
			},
			want:    "192.168.0.224/27",
			want1:   26,
			wantErr: false,
		},
		{
			name: "maxClusterServiceNum == clusterCIDR size",
			args: args{
				clusterCIDR:          "192.168.0.0/24",
				maxClusterServiceNum: 256,
				maxNodePodNum:        64,
			},
			want:    "192.168.0.0/24",
			want1:   26,
			wantErr: false,
		},
		{
			name: "maxClusterServiceNum > clusterCIDR size",
			args: args{
				clusterCIDR:          "192.168.0.0/24",
				maxClusterServiceNum: 257,
				maxNodePodNum:        64,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := GetServiceCIDRAndNodeCIDRMaskSize(tt.args.clusterCIDR, tt.args.maxClusterServiceNum, tt.args.maxNodePodNum)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetServiceCIDRAndNodeCIDRMaskSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetServiceCIDRAndNodeCIDRMaskSize() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetServiceCIDRAndNodeCIDRMaskSize() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGetNodeCIDRMaskSize(t *testing.T) {
	type args struct {
		clusterCIDR   string
		maxNodePodNum int32
	}
	tests := []struct {
		name    string
		args    args
		want    int32
		wantErr bool
	}{
		{
			name: "maxNodePodNum == 0",
			args: args{
				clusterCIDR: "192.168.0.0/24",
			},
			wantErr: true,
		},
		{
			name: "maxNodePodNum < clusterCIDR size",
			args: args{
				clusterCIDR:   "192.168.0.0/24",
				maxNodePodNum: 64,
			},
			want:    26,
			wantErr: false,
		},
		{
			name: "maxNodePodNum == clusterCIDR size",
			args: args{
				clusterCIDR:   "192.168.0.0/24",
				maxNodePodNum: 256,
			},
			want:    24,
			wantErr: false,
		},
		{
			name: "maxNodePodNum > clusterCIDR size",
			args: args{
				clusterCIDR:   "192.168.0.0/25",
				maxNodePodNum: 256,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNodeCIDRMaskSize(tt.args.clusterCIDR, tt.args.maxNodePodNum)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNodeCIDRMaskSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetNodeCIDRMaskSize() got = %v, want %v", got, tt.want)
			}
		})
	}
}
