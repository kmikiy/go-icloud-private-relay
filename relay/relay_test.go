package relay

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplePrivateRelay(t *testing.T) {
	act, err := ApplePrivateRelay("172.225.18.12")
	assert.NoError(t, err)

	_, ipNet, _ := net.ParseCIDR("172.225.18.0/28")
	exp := egress{
		IPNet:       *ipNet,
		CIDR:        ipNet.String(),
		CountryCode: "MX",
		State:       "MX-CHH",
		City:        "Chihuahua"}
	assert.Equal(t, exp, act)
}

func TestIsApplePrivateRelayAddress(t *testing.T) {
	type args struct {
		ipAddress string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Apple private relay address",
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
			if got := IsApplePrivateRelayAddress(tt.args.ipAddress); got != tt.want {
				t.Errorf("IsApplePrivateRelayAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}