package zookeeper

import (
	"reflect"
	"testing"

	"github.com/natemarks/zoochecker/input"
)

func TestSendToClusterNode(t *testing.T) {
	type args struct {
		node    input.ClusterNode
		message string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "RUOK",
			args: args{
				node: input.ClusterNode{
					Host:    "localhost",
					Port:    21811,
					Timeout: 5,
				},
				message: "ruok",
			},
			want:    "imok",
			wantErr: false,
		},
		//{
		//	name: "SRVR",
		//	args: args{
		//		node: input.ClusterNode{
		//			Host:    "localhost",
		//			Port:    21811,
		//			Timeout: 5,
		//		},
		//		message: "srvr",
		//	},
		//	want:    "...",
		//	wantErr: false,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SendToClusterNode(tt.args.node, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendToClusterNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SendToClusterNode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeIsOk(t *testing.T) {
	type args struct {
		node input.ClusterNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "RUOK",
			args: args{
				node: input.ClusterNode{
					Host:    "localhost",
					Port:    21811,
					Timeout: 5,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NodeIsOk(tt.args.node); got != tt.want {
				t.Errorf("NodeIsOk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClusterStatus(t *testing.T) {
	type args struct {
		node input.ClusterNode
	}
	tests := []struct {
		name string
		args args
		want Status
	}{
		{
			name: "node1",
			args: args{
				node: input.ClusterNode{
					Host:    "localhost",
					Port:    21811,
					Timeout: 5,
				},
			},
			want: Status{
				Mode:      "follower",
				Followers: 0,
			},
		},
		{
			name: "node2",
			args: args{
				node: input.ClusterNode{
					Host:    "localhost",
					Port:    21812,
					Timeout: 5,
				},
			},
			want: Status{
				Mode:      "follower",
				Followers: 0,
			},
		},
		{
			name: "node3",
			args: args{
				node: input.ClusterNode{
					Host:    "localhost",
					Port:    21813,
					Timeout: 5,
				},
			},
			want: Status{
				Mode:      "leader",
				Followers: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NodeStatus(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NodeStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
