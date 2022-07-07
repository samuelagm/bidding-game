package contract

import (
	"reflect"
	"testing"
)

func TestNewBiddingWarEventHelper(t *testing.T) {
	type args struct {
		wsshost string
		address string
	}
	tests := []struct {
		name string
		args args
		want *BiddingWarHelper
	}{
		{
			name: "Basic Test",
			args: args{
				wsshost: cfg.WSS_Host,
				address: cfg.ContractAddress,
			},
			want: &BiddingWarHelper{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBiddingWarEventHelper(tt.args.wsshost, tt.args.address); reflect.DeepEqual(got, nil) {
				t.Errorf("NewBiddingWarEventHelper() = %v, want %v", got, tt.want)
			}
		})
	}
}
