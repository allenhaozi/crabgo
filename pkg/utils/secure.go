package utils

import (
	log "github.com/sirupsen/logrus"
	"strings"
)

// super hack way do subnet match
func IsFromSameSubnet(remoteAddr, podIP string) bool {
	// TODO support ipv6
	remoteSubnet := strings.Join(strings.Split(remoteAddr, ".")[:2], ".")
	log.Info("remoteSubnet ", remoteSubnet)
	podIPSubnet := strings.Join(strings.Split(podIP, ".")[:2], ".")
	log.Info("podIPSubnet ", podIPSubnet)
	return remoteSubnet == podIPSubnet
}
