package relay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestICloudPrivateRelay(t *testing.T) {
	act, err := ICloudPrivateRelay("172.225.18.12")
	assert.NoError(t, err)

	exp := location{
		CountryCode: "MX",
		State:       "MX-CHH",
		City:        "Chihuahua"}
	assert.Equal(t, exp, act)
}

func TestIsICloudPrivateRelayAddress(t *testing.T) {
	type args struct {
		ipAddress string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "iCloud private relay address",
			args: args{
				ipAddress: "172.225.18.12",
			},
			want: true,
		},
		{
			name: "Google IP address",
			args: args{
				ipAddress: "142.251.39.14",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsICloudPrivateRelayAddress(tt.args.ipAddress); got != tt.want {
				t.Errorf("IsICloudPrivateRelayAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
