package nfsd

import "github.com/swiftstack/onc"

// See also consts.go and structs.go for exported constants and structures referenced by this API

type MountV3Interface interface {
	MountProc3Null(authSysBody *onc.AuthSysBodyStruct)
	MountProc3Mnt(authSysBody *onc.AuthSysBodyStruct, mountProc3MntArgs *MountProc3MntArgsStruct) (mountProc3MntResults *MountProc3MntResultsStruct)
	MountProc3Umnt(authSysBody *onc.AuthSysBodyStruct, mountProc3UmntArgs *MountProc3UmntArgsStruct)
}

// StartIPv4TCPMountV3Server launches a Mount V3 server on the specified IPv4 TCP Port
//
// Arguments:
//   port      specifies the TCP port # upon which to serve Mount V3 via IPv4
//   publish   indicates whether or not to publish the Mount V3 server via portmapper/rpcbind
//   callbacks specifies the receiver of the API "up calls" as listed in MountV3Interface
//
// Returns:
//   published indicates whether or not portmapper/rpcbind successfully registered the program:version:port tuple
//   err       is non-nil on failure (but published is valid either way)
func StartIPv4TCPMountV3Server(port uint16, publish bool, callbacks MountV3Interface) (published bool, err error) {
	published, err = startIPv4TCPMountV3Server(port, publish, callbacks)
	return
}

// StartIPv4UDPMountV3Server launches a Mount V3 server on the specified IPv4 UDP Port
//
// Arguments:
//   port      specifies the UDP port # upon which to serve Mount V3 via IPv4
//   publish   indicates whether or not to publish the Mount V3 server via portmapper/rpcbind
//   callbacks specifies the receiver of the API "up calls" as listed in MountV3Interface
//
// Returns:
//   published indicates whether or not portmapper/rpcbind successfully registered the program:version:port tuple
//   err       is non-nil on failure (but published is valid either way)
func StartIPv4UDPMountV3Server(port uint16, publish bool, callbacks MountV3Interface) (published bool, err error) {
	published, err = startIPv4UDPMountV3Server(port, publish, callbacks)
	return
}

// StopIPv4TCPMountV3Server stops an Mount V3 server
//
// Arguments:
//   port      specifies the TCP port # upon which Mount V3 servicing via IPv4 should be halted
//   unpublish indicates whether or not to remove a previously published Mount V3 server via portmapper/rpcbind
//
// Returns:
//   unpublished indicates whether or not portmapper/rpcbind successfully unregistered the program:version:port tuple
//   err         is non-nil on failure (but unpublished is valid either way)
func StopIPv4TCPMountV3Server(port uint16, unpublish bool) (unpublished bool, err error) {
	unpublished, err = stopIPv4TCPMountV3Server(port, unpublish)
	return
}

// StopIPv4UDPMountV3Server stops an Mount V3 server
//
// Arguments:
//   port      specifies the UDP port # upon which Mount V3 servicing via IPv4 should be halted
//   unpublish indicates whether or not to remove a previously published Mount V3 server via portmapper/rpcbind
//
// Returns:
//   unpublished indicates whether or not portmapper/rpcbind successfully unregistered the program:version:port tuple
//   err         is non-nil on failure (but unpublished is valid either way)
func StopIPv4UDPMountV3Server(port uint16, unpublish bool) (unpublished bool, err error) {
	unpublished, err = stopIPv4UDPMountV3Server(port, unpublish)
	return
}

type NFSv3Interface interface {
	NFSProc3Null(authSysBody *onc.AuthSysBodyStruct)
	NFSProc3GetAttr(authSysBody *onc.AuthSysBodyStruct, nfsProc3GetAttrArgs *NFSProc3GetAttrArgsStruct) (nfsProc3GetAttrResults *NFSProc3GetAttrResultsStruct)
	NFSProc3SetAttr(authSysBody *onc.AuthSysBodyStruct, nfsProc3SetAttrArgs *NFSProc3SetAttrArgsStruct) (nfsProc3SetAttrResults *NFSProc3SetAttrResultsStruct)
	NFSProc3Lookup(authSysBody *onc.AuthSysBodyStruct, nfsProc3LookupArgs *NFSProc3LookupArgsStruct) (nfsProc3LookupResults *NFSProc3LookupResultsStruct)
	NFSProc3Access(authSysBody *onc.AuthSysBodyStruct, nfsProc3AccessArgs *NFSProc3AccessArgsStruct) (nfsProc3AccessResults *NFSProc3AccessResultsStruct)
	NFSProc3ReadLink(authSysBody *onc.AuthSysBodyStruct, nfsProc3ReadLinkArgs *NFSProc3ReadLinkArgsStruct) (nfsProc3ReadLinkResults *NFSProc3ReadLinkResultsStruct)
	NFSProc3Read(authSysBody *onc.AuthSysBodyStruct, nfsProc3ReadArgs *NFSProc3ReadArgsStruct) (nfsProc3ReadResults *NFSProc3ReadResultsStruct)
	NFSProc3Write(authSysBody *onc.AuthSysBodyStruct, nfsProc3WriteArgs *NFSProc3WriteArgsStruct) (nfsProc3WriteResults *NFSProc3WriteResultsStruct)
	NFSProc3Create(authSysBody *onc.AuthSysBodyStruct, nfsProc3CreateArgs *NFSProc3CreateArgsStruct) (nfsProc3CreateResults *NFSProc3CreateResultsStruct)
	NFSProc3MKDir(authSysBody *onc.AuthSysBodyStruct, nfsProc3MKDirArgs *NFSProc3MKDirArgsStruct) (nfsProc3MKDirResults *NFSProc3MKDirResultsStruct)
	NFSProc3SymLink(authSysBody *onc.AuthSysBodyStruct, nfsProc3SymLinkArgs *NFSProc3SymLinkArgsStruct) (nfsProc3SymLinkResults *NFSProc3SymLinkResultsStruct)
	NFSProc3Remove(authSysBody *onc.AuthSysBodyStruct, nfsProc3RemoveArgs *NFSProc3RemoveArgsStruct) (nfsProc3RemoveResults *NFSProc3RemoveResultsStruct)
	NFSProc3RMDir(authSysBody *onc.AuthSysBodyStruct, nfsProc3RMDirArgs *NFSProc3RMDirArgsStruct) (nfsProc3RMDirResults *NFSProc3RMDirResultsStruct)
	NFSProc3Rename(authSysBody *onc.AuthSysBodyStruct, nfsProc3RenameArgs *NFSProc3RenameArgsStruct) (nfsProc3RenameResults *NFSProc3RenameResultsStruct)
	NFSProc3Link(authSysBody *onc.AuthSysBodyStruct, nfsProc3LinkArgs *NFSProc3LinkArgsStruct) (nfsProc3LinkResults *NFSProc3LinkResultsStruct)
	NFSProc3ReadDir(authSysBody *onc.AuthSysBodyStruct, nfsProc3ReadDirArgs *NFSProc3ReadDirArgsStruct) (nfsProc3ReadDirResults *NFSProc3ReadDirResultsStruct)
	NFSProc3ReadDirPlus(authSysBody *onc.AuthSysBodyStruct, nfsProc3ReadDirPlusArgs *NFSProc3ReadDirPlusArgsStruct) (nfsProc3ReadDirPlusResults *NFSProc3ReadDirPlusResultsStruct)
	NFSProc3FSStat(authSysBody *onc.AuthSysBodyStruct, nfsProc3FSStatArgs *NFSProc3FSStatArgsStruct) (nfsProc3FSStatResults *NFSProc3FSStatResultsStruct)
	NFSProc3FSInfo(authSysBody *onc.AuthSysBodyStruct, nfsProc3FSInfoArgs *NFSProc3FSInfoArgsStruct) (nfsProc3FSInfoResults *NFSProc3FSInfoResultsStruct)
	NFSProc3PathConf(authSysBody *onc.AuthSysBodyStruct, nfsProc3PathConfArgs *NFSProc3PathConfArgsStruct) (nfsProc3PathConfResults *NFSProc3PathConfResultsStruct)
	NFSProc3Commit(authSysBody *onc.AuthSysBodyStruct, nfsProc3CommitArgs *NFSProc3CommitArgsStruct) (nfsProc3CommitResults *NFSProc3CommitResultsStruct)
}

// StartIPv4TCPNFSv3Server launches an NFSv3 server on the specified IPv4 TCP Port
//
// Arguments:
//   port      specifies the TCP port # upon which to serve NFSv3 via IPv4
//   publish   indicates whether or not to publish the NFSv3 server via portmapper/rpcbind
//   callbacks specifies the receiver of the API "up calls" as listed in NFSv3Interface
//
// Returns:
//   published indicates whether or not portmapper/rpcbind successfully registered the program:version:port tuple
//   err       is non-nil on failure (but published is valid either way)
func StartIPv4TCPNFSv3Server(port uint16, publish bool, callbacks NFSv3Interface) (published bool, err error) {
	published, err = startIPv4TCPNFSv3Server(port, publish, callbacks)
	return
}

// StartIPv4UDPNFSv3Server launches an NFSv3 server on the specified IPv4 TCP Port
//
// Arguments:
//   port      specifies the UDP port # upon which to serve NFSv3 via IPv4
//   publish   indicates whether or not to publish the NFSv3 server via portmapper/rpcbind
//   callbacks specifies the receiver of the API "up calls" as listed in NFSv3Interface
//
// Returns:
//   published indicates whether or not portmapper/rpcbind successfully registered the program:version:port tuple
//   err       is non-nil on failure (but published is valid either way)
func StartIPv4UDPNFSv3Server(port uint16, publish bool, callbacks NFSv3Interface) (published bool, err error) {
	published, err = startIPv4UDPNFSv3Server(port, publish, callbacks)
	return
}

// StopIPv4TCPNFSv3Server stops an NFSv3 server
//
// Arguments:
//   port      specifies the TCP port # upon which NFSv3 servicing via IPv4 should be halted
//   unpublish indicates whether or not to remove a previously published NFSv3 server via portmapper/rpcbind
//
// Returns:
//   unpublished indicates whether or not portmapper/rpcbind successfully unregistered the program:version:port tuple
//   err         is non-nil on failure (but unpublished is valid either way)
func StopIPv4TCPNFSv3Server(port uint16, unpublish bool) (unpublished bool, err error) {
	unpublished, err = stopIPv4TCPNFSv3Server(port, unpublish)
	return
}

// StopIPv4UDPNFSv3Server stops an NFSv3 server
//
// Arguments:
//   port      specifies the UDP port # upon which NFSv3 servicing via IPv4 should be halted
//   unpublish indicates whether or not to remove a previously published NFSv3 server via portmapper/rpcbind
//
// Returns:
//   unpublished indicates whether or not portmapper/rpcbind successfully unregistered the program:version:port tuple
//   err         is non-nil on failure (but unpublished is valid either way)
func StopIPv4UDPNFSv3Server(port uint16, unpublish bool) (unpublished bool, err error) {
	unpublished, err = stopIPv4UDPNFSv3Server(port, unpublish)
	return
}
