package nfsd

import (
	"fmt"
)

func startIPv4TCPMountV3Server(port uint16, publish bool, callbacks MountV3Interface) (err error) {
	err = fmt.Errorf("TODO")
	return
}

func startIPv4UDPMountV3Server(port uint16, publish bool, callbacks MountV3Interface) (err error) {
	err = fmt.Errorf("TODO")
	return
}

func stopIPv4TCPMountV3Server(port uint16, unpublish bool) (err error) {
	err = fmt.Errorf("TODO")
	return
}

func stopIPv4UDPMountV3Server(port uint16, unpublish bool) (err error) {
	err = fmt.Errorf("TODO")
	return
}

func startIPv4TCPNFSv3Server(port uint16, publish bool, callbacks NFSv3Interface) (err error) {
	err = fmt.Errorf("TODO")
	return
}

func startIPv4UDPNFSv3Server(port uint16, publish bool, callbacks NFSv3Interface) (err error) {
	err = fmt.Errorf("TODO")
	return
}

func stopIPv4TCPNFSv3Server(port uint16, unpublish bool) (err error) {
	err = fmt.Errorf("TODO")
	return
}

func stopIPv4UDPNFSv3Server(port uint16, unpublish bool) (err error) {
	err = fmt.Errorf("TODO")
	return
}
