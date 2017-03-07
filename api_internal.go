package nfsd

import (
	"github.com/swiftstack/onc"
	"github.com/swiftstack/onc/oncclient"
	"github.com/swiftstack/onc/oncserver"
)

func startIPv4TCPMountV3Server(port uint16, publish bool, callbacks MountV3Interface) (published bool, err error) {
	published = false

	err = oncserver.StartServer(onc.IPProtoTCP, port, []oncserver.ProgVersStruct{{Prog: onc.ProgNumMount, VersList: []uint32{3}}}, &mountRequestHandlerStruct{prot: onc.IPProtoTCP, port: port})
	if nil != err {
		return
	}

	if publish {
		publishErr := oncclient.DoPmapProcSet(onc.ProgNumMount, 3, onc.IPProtoTCP, port)
		published = (nil == publishErr)
	}

	return
}

func startIPv4UDPMountV3Server(port uint16, publish bool, callbacks MountV3Interface) (published bool, err error) {
	published = false

	err = oncserver.StartServer(onc.IPProtoUDP, port, []oncserver.ProgVersStruct{{Prog: onc.ProgNumMount, VersList: []uint32{3}}}, &mountRequestHandlerStruct{prot: onc.IPProtoUDP, port: port})
	if nil != err {
		return
	}

	if publish {
		publishErr := oncclient.DoPmapProcSet(onc.ProgNumMount, 3, onc.IPProtoUDP, port)
		published = (nil == publishErr)
	}

	return
}

func stopIPv4TCPMountV3Server(port uint16, unpublish bool) (unpublished bool, err error) {
	if unpublish {
		unpublishErr := oncclient.DoPmapProcUnset(onc.ProgNumMount, 3, onc.IPProtoTCP)
		unpublished = (nil == unpublishErr)
	} else {
		unpublished = false
	}

	err = oncserver.StopServer(onc.IPProtoTCP, port)

	return
}

func stopIPv4UDPMountV3Server(port uint16, unpublish bool) (unpublished bool, err error) {
	if unpublish {
		unpublishErr := oncclient.DoPmapProcUnset(onc.ProgNumMount, 3, onc.IPProtoUDP)
		unpublished = (nil == unpublishErr)
	} else {
		unpublished = false
	}

	err = oncserver.StopServer(onc.IPProtoUDP, port)

	return
}

func startIPv4TCPNFSv3Server(port uint16, publish bool, callbacks NFSv3Interface) (published bool, err error) {
	published = false

	err = oncserver.StartServer(onc.IPProtoTCP, port, []oncserver.ProgVersStruct{{Prog: onc.ProgNumNFS, VersList: []uint32{3}}}, &nfsRequestHandlerStruct{prot: onc.IPProtoTCP, port: port})
	if nil != err {
		return
	}

	if publish {
		publishErr := oncclient.DoPmapProcSet(onc.ProgNumNFS, 3, onc.IPProtoTCP, port)
		published = (nil == publishErr)
	}

	return
}

func startIPv4UDPNFSv3Server(port uint16, publish bool, callbacks NFSv3Interface) (published bool, err error) {
	published = false

	err = oncserver.StartServer(onc.IPProtoUDP, port, []oncserver.ProgVersStruct{{Prog: onc.ProgNumNFS, VersList: []uint32{3}}}, &nfsRequestHandlerStruct{prot: onc.IPProtoUDP, port: port})
	if nil != err {
		return
	}

	if publish {
		publishErr := oncclient.DoPmapProcSet(onc.ProgNumNFS, 3, onc.IPProtoUDP, port)
		published = (nil == publishErr)
	}

	return
}

func stopIPv4TCPNFSv3Server(port uint16, unpublish bool) (unpublished bool, err error) {
	if unpublish {
		unpublishErr := oncclient.DoPmapProcUnset(onc.ProgNumNFS, 3, onc.IPProtoTCP)
		unpublished = (nil == unpublishErr)
	} else {
		unpublished = false
	}

	err = oncserver.StopServer(onc.IPProtoTCP, port)

	return
}

func stopIPv4UDPNFSv3Server(port uint16, unpublish bool) (unpublished bool, err error) {
	if unpublish {
		unpublishErr := oncclient.DoPmapProcUnset(onc.ProgNumNFS, 3, onc.IPProtoUDP)
		unpublished = (nil == unpublishErr)
	} else {
		unpublished = false
	}

	err = oncserver.StopServer(onc.IPProtoUDP, port)

	return
}
