package nfsd

import (
	"fmt"

	"github.com/swiftstack/onc"
	"github.com/swiftstack/onc/oncserver"
	"github.com/swiftstack/xdr"
)

type mountRequestHandlerStruct struct {
	callbacks MountV3Interface
	prot      uint32 // either onc.IPProtoTCP or onc.IPProtoUDP
	port      uint16
}

type nfsRequestHandlerStruct struct {
	callbacks NFSv3Interface
	prot      uint32 // either onc.IPProtoTCP or onc.IPProtoUDP
	port      uint16
}

func (mountRequestHandler *mountRequestHandlerStruct) ONCRequest(connHandle oncserver.ConnHandle, xid uint32, prog uint32, vers uint32, proc uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		err error
	)

	if onc.ProgNumMount != prog {
		err = fmt.Errorf("prog was %v... expected it to be onc.ProgNumMount (%v)", prog, onc.ProgNumMount)
		mountRequestHandler.callbacks.ErrorLog(err)
		panic(err) // i.e. this shouldn't have happened if oncserver "dispatcher" is functioning correctly
	}

	if 3 != vers {
		err = fmt.Errorf("vers was %v... expected it to be 3", vers)
		mountRequestHandler.callbacks.ErrorLog(err)
		panic(err) // i.e. this shouldn't have happened if oncserver "dispatcher" is functioning correctly
	}

	switch proc {
	case ProcNULL:
		mountRequestHandler.null(connHandle, xid, authSysBody, parms)
	case MOUNTPROC3MNT:
		mountRequestHandler.mnt(connHandle, xid, authSysBody, parms)
	case MOUNTPROC3UMNT:
		mountRequestHandler.umnt(connHandle, xid, authSysBody, parms)
	default:
		err = fmt.Errorf("proc %v not available", proc)
		mountRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.ProcUnavail)
		if nil != err {
			mountRequestHandler.callbacks.ErrorLog(err)
		}
	}
}

func (mountRequestHandler *mountRequestHandlerStruct) null(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		err error
	)

	if 0 != len(parms) {
		err = fmt.Errorf("ProcNULL(...parms) should have been void")
		mountRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			mountRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}

	mountRequestHandler.callbacks.MountProc3Null(authSysBody)

	err = oncserver.SendAcceptedSuccess(connHandle, xid, nil)
	if nil != err {
		mountRequestHandler.callbacks.ErrorLog(err)
	}
}

func (mountRequestHandler *mountRequestHandlerStruct) mnt(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		bytesConsumed        uint64
		err                  error
		mountProc3MntArgs    MountProc3MntArgsStruct
		mountProc3MntResults *MountProc3MntResultsStruct
		results              []byte
		statusOnlyResults    StatusOnlyResultsStruct
	)

	bytesConsumed, err = xdr.Unpack(parms, &mountProc3MntArgs)
	if nil != err {
		mountRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			mountRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}
	if uint64(len(parms)) != bytesConsumed {
		err = fmt.Errorf("xdr.Unpack() failed to consume all of parms")
		mountRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			mountRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}

	mountProc3MntResults = mountRequestHandler.callbacks.MountProc3Mnt(authSysBody, &mountProc3MntArgs)

	if OK == mountProc3MntResults.Status {
		if (1 != len(mountProc3MntResults.AuthFlavors)) || (onc.AuthSys != mountProc3MntResults.AuthFlavors[0]) {
			err = fmt.Errorf("mountProc3MntResults.AuthFlavors must == []{onc.AuthSys}")
			mountRequestHandler.callbacks.ErrorLog(err)
			err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.SystemErr)
			if nil != err {
				mountRequestHandler.callbacks.ErrorLog(err)
			}
			return
		}
		results, err = xdr.Pack(mountProc3MntResults)
		if nil != err {
			mountRequestHandler.callbacks.ErrorLog(err)
			err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.SystemErr)
			if nil != err {
				mountRequestHandler.callbacks.ErrorLog(err)
			}
			return
		}
	} else {
		statusOnlyResults.Status = mountProc3MntResults.Status
		results, err = xdr.Pack(statusOnlyResults)
		if nil != err {
			mountRequestHandler.callbacks.ErrorLog(err)
			err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.SystemErr)
			if nil != err {
				mountRequestHandler.callbacks.ErrorLog(err)
			}
			return
		}
	}

	err = oncserver.SendAcceptedSuccess(connHandle, xid, results)
	if nil != err {
		mountRequestHandler.callbacks.ErrorLog(err)
	}
}

func (mountRequestHandler *mountRequestHandlerStruct) umnt(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		bytesConsumed      uint64
		err                error
		mountProc3UmntArgs MountProc3UmntArgsStruct
	)

	bytesConsumed, err = xdr.Unpack(parms, &mountProc3UmntArgs)
	if nil != err {
		mountRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			mountRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}
	if uint64(len(parms)) != bytesConsumed {
		err = fmt.Errorf("xdr.Unpack() failed to consume all of parms")
		mountRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			mountRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}

	mountRequestHandler.callbacks.MountProc3Umnt(authSysBody, &mountProc3UmntArgs)

	err = oncserver.SendAcceptedSuccess(connHandle, xid, nil)
	if nil != err {
		mountRequestHandler.callbacks.ErrorLog(err)
	}
}

func (nfsRequestHandler *nfsRequestHandlerStruct) ONCRequest(connHandle oncserver.ConnHandle, xid uint32, prog uint32, vers uint32, proc uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		err error
	)

	if onc.ProgNumNFS != prog {
		err = fmt.Errorf("prog was %v... expected it to be onc.ProgNumNFS (%v)", prog, onc.ProgNumNFS)
		nfsRequestHandler.callbacks.ErrorLog(err)
		panic(err) // i.e. this shouldn't have happened if oncserver "dispatcher" is functioning correctly
	}

	if 3 != vers {
		err = fmt.Errorf("vers was %v... expected it to be 3", vers)
		nfsRequestHandler.callbacks.ErrorLog(err)
		panic(err) // i.e. this shouldn't have happened if oncserver "dispatcher" is functioning correctly
	}

	switch proc {
	case ProcNULL:
		nfsRequestHandler.null(connHandle, xid, authSysBody, parms)
	case NFSPROC3GETATTR:
		nfsRequestHandler.getattr(connHandle, xid, authSysBody, parms)
	case NFSPROC3SETATTR:
		nfsRequestHandler.setattr(connHandle, xid, authSysBody, parms)
	case NFSPROC3LOOKUP:
		nfsRequestHandler.lookup(connHandle, xid, authSysBody, parms)
	case NFSPROC3ACCESS:
		nfsRequestHandler.access(connHandle, xid, authSysBody, parms)
	case NFSPROC3READLINK:
		nfsRequestHandler.readlink(connHandle, xid, authSysBody, parms)
	case NFSPROC3READ:
		nfsRequestHandler.read(connHandle, xid, authSysBody, parms)
	case NFSPROC3WRITE:
		nfsRequestHandler.write(connHandle, xid, authSysBody, parms)
	case NFSPROC3CREATE:
		nfsRequestHandler.create(connHandle, xid, authSysBody, parms)
	case NFSPROC3MKDIR:
		nfsRequestHandler.mkdir(connHandle, xid, authSysBody, parms)
	case NFSPROC3SYMLINK:
		nfsRequestHandler.symlink(connHandle, xid, authSysBody, parms)
	case NFSPROC3REMOVE:
		nfsRequestHandler.remove(connHandle, xid, authSysBody, parms)
	case NFSPROC3RMDIR:
		nfsRequestHandler.rmdir(connHandle, xid, authSysBody, parms)
	case NFSPROC3RENAME:
		nfsRequestHandler.rename(connHandle, xid, authSysBody, parms)
	case NFSPROC3LINK:
		nfsRequestHandler.link(connHandle, xid, authSysBody, parms)
	case NFSPROC3READDIR:
		nfsRequestHandler.readdir(connHandle, xid, authSysBody, parms)
	case NFSPROC3READDIRPLUS:
		nfsRequestHandler.readdirplus(connHandle, xid, authSysBody, parms)
	case NFSPROC3FSSTAT:
		nfsRequestHandler.fsstat(connHandle, xid, authSysBody, parms)
	case NFSPROC3FSINFO:
		nfsRequestHandler.fsinfo(connHandle, xid, authSysBody, parms)
	case NFSPROC3PATHCONF:
		nfsRequestHandler.pathconf(connHandle, xid, authSysBody, parms)
	case NFSPROC3COMMIT:
		nfsRequestHandler.commit(connHandle, xid, authSysBody, parms)
	default:
		err = fmt.Errorf("proc %v not available", proc)
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.ProcUnavail)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
	}
}

func (nfsRequestHandler *nfsRequestHandlerStruct) null(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		err error
	)

	if 0 != len(parms) {
		err = fmt.Errorf("ProcNULL(...parms) should have been void")
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}

	nfsRequestHandler.callbacks.NFSProc3Null(authSysBody)

	err = oncserver.SendAcceptedSuccess(connHandle, xid, nil)
	if nil != err {
		nfsRequestHandler.callbacks.ErrorLog(err)
	}
}

func (nfsRequestHandler *nfsRequestHandlerStruct) getattr(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		bytesConsumed          uint64
		err                    error
		nfsProc3GetAttrArgs    NFSProc3GetAttrArgsStruct
		nfsProc3GetAttrResults *NFSProc3GetAttrResultsStruct
		results                []byte
		statusOnlyResults      StatusOnlyResultsStruct
	)

	bytesConsumed, err = xdr.Unpack(parms, &nfsProc3GetAttrArgs)
	if nil != err {
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}
	if uint64(len(parms)) != bytesConsumed {
		err = fmt.Errorf("xdr.Unpack() failed to consume all of parms")
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}

	nfsProc3GetAttrResults = nfsRequestHandler.callbacks.NFSProc3GetAttr(authSysBody, &nfsProc3GetAttrArgs)

	// TODO
	/*
		if OK == mountProc3MntResults.Status {
			if (1 != len(mountProc3MntResults.AuthFlavors)) || (onc.AuthSys != mountProc3MntResults.AuthFlavors[0]) {
				err = fmt.Errorf("mountProc3MntResults.AuthFlavors must == []{onc.AuthSys}")
				mountRequestHandler.callbacks.ErrorLog(err)
				err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.SystemErr)
				if nil != err {
					mountRequestHandler.callbacks.ErrorLog(err)
				}
				return
			}
			results, err = xdr.Pack(mountProc3MntResults)
			if nil != err {
				mountRequestHandler.callbacks.ErrorLog(err)
				err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.SystemErr)
				if nil != err {
					mountRequestHandler.callbacks.ErrorLog(err)
				}
				return
			}
		} else {
			statusOnlyResults.Status = mountProc3MntResults.Status
			results, err = xdr.Pack(statusOnlyResults)
			if nil != err {
				mountRequestHandler.callbacks.ErrorLog(err)
				err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.SystemErr)
				if nil != err {
					mountRequestHandler.callbacks.ErrorLog(err)
				}
				return
			}
		}

		err = oncserver.SendAcceptedSuccess(connHandle, xid, results)
		if nil != err {
			mountRequestHandler.callbacks.ErrorLog(err)
		}
	*/
}

func (nfsRequestHandler *nfsRequestHandlerStruct) setattr(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		bytesConsumed          uint64
		err                    error
		nfsProc3SetAttrArgs    NFSProc3SetAttrArgsStruct
		nfsProc3SetAttrResults *NFSProc3SetAttrResultsStruct
		results                []byte
		statusOnlyResults      StatusOnlyResultsStruct
	)

	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) lookup(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		bytesConsumed         uint64
		err                   error
		nfsProc3LookupArgs    NFSProc3LookupArgsStruct
		nfsProc3LookupResults *NFSProc3LookupResultsStruct
		results               []byte
		statusOnlyResults     StatusOnlyResultsStruct
	)

	bytesConsumed, err = xdr.Unpack(parms, &nfsProc3LookupArgs)
	if nil != err {
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}
	if uint64(len(parms)) != bytesConsumed {
		err = fmt.Errorf("xdr.Unpack() failed to consume all of parms")
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}

	nfsProc3LookupResults = nfsRequestHandler.callbacks.NFSProc3Lookup(authSysBody, &nfsProc3LookupArgs)

	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) access(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		bytesConsumed         uint64
		err                   error
		nfsProc3AccessArgs    NFSProc3AccessArgsStruct
		nfsProc3AccessResults *NFSProc3AccessResultsStruct
		results               []byte
		statusOnlyResults     StatusOnlyResultsStruct
	)

	bytesConsumed, err = xdr.Unpack(parms, &nfsProc3AccessArgs)
	if nil != err {
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}
	if uint64(len(parms)) != bytesConsumed {
		err = fmt.Errorf("xdr.Unpack() failed to consume all of parms")
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}

	nfsProc3AccessResults = nfsRequestHandler.callbacks.NFSProc3Access(authSysBody, &nfsProc3AccessArgs)

	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) readlink(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		bytesConsumed           uint64
		err                     error
		nfsProc3ReadLinkArgs    NFSProc3ReadLinkArgsStruct
		nfsProc3ReadLinkResults *NFSProc3ReadLinkResultsStruct
		results                 []byte
		statusOnlyResults       StatusOnlyResultsStruct
	)

	bytesConsumed, err = xdr.Unpack(parms, &nfsProc3ReadLinkArgs)
	if nil != err {
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}
	if uint64(len(parms)) != bytesConsumed {
		err = fmt.Errorf("xdr.Unpack() failed to consume all of parms")
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}

	nfsProc3ReadLinkResults = nfsRequestHandler.callbacks.NFSProc3ReadLink(authSysBody, &nfsProc3ReadLinkArgs)

	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) read(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		bytesConsumed       uint64
		err                 error
		nfsProc3ReadArgs    NFSProc3ReadArgsStruct
		nfsProc3ReadResults *NFSProc3ReadResultsStruct
		results             []byte
		statusOnlyResults   StatusOnlyResultsStruct
	)

	bytesConsumed, err = xdr.Unpack(parms, &nfsProc3ReadArgs)
	if nil != err {
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}
	if uint64(len(parms)) != bytesConsumed {
		err = fmt.Errorf("xdr.Unpack() failed to consume all of parms")
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}

	nfsProc3ReadResults = nfsRequestHandler.callbacks.NFSProc3Read(authSysBody, &nfsProc3ReadArgs)

	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) write(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		bytesConsumed        uint64
		err                  error
		nfsProc3WriteArgs    NFSProc3WriteArgsStruct
		nfsProc3WriteResults *NFSProc3WriteResultsStruct
		results              []byte
		statusOnlyResults    StatusOnlyResultsStruct
	)

	bytesConsumed, err = xdr.Unpack(parms, &nfsProc3WriteArgs)
	if nil != err {
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}
	if uint64(len(parms)) != bytesConsumed {
		err = fmt.Errorf("xdr.Unpack() failed to consume all of parms")
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}

	nfsProc3WriteResults = nfsRequestHandler.callbacks.NFSProc3Write(authSysBody, &nfsProc3WriteArgs)

	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) create(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		bytesConsumed         uint64
		err                   error
		nfsProc3CreateArgs    NFSProc3CreateArgsStruct
		nfsProc3CreateResults *NFSProc3CreateResultsStruct
		results               []byte
		statusOnlyResults     StatusOnlyResultsStruct
	)

	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) mkdir(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		bytesConsumed        uint64
		err                  error
		nfsProc3MKDirArgs    NFSProc3MKDirArgsStruct
		nfsProc3MKDirResults *NFSProc3MKDirResultsStruct
		results              []byte
		statusOnlyResults    StatusOnlyResultsStruct
	)

	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) symlink(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		bytesConsumed          uint64
		err                    error
		nfsProc3SymLinkArgs    NFSProc3SymLinkArgsStruct
		nfsProc3SymLinkResults *NFSProc3SymLinkResultsStruct
		results                []byte
		statusOnlyResults      StatusOnlyResultsStruct
	)

	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) remove(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		bytesConsumed         uint64
		err                   error
		nfsProc3RemoveArgs    NFSProc3RemoveArgsStruct
		nfsProc3RemoveResults *NFSProc3RemoveResultsStruct
		results               []byte
		statusOnlyResults     StatusOnlyResultsStruct
	)

	bytesConsumed, err = xdr.Unpack(parms, &nfsProc3RemoveArgs)
	if nil != err {
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}
	if uint64(len(parms)) != bytesConsumed {
		err = fmt.Errorf("xdr.Unpack() failed to consume all of parms")
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}

	nfsProc3RemoveResults = nfsRequestHandler.callbacks.NFSProc3Remove(authSysBody, &nfsProc3RemoveArgs)

	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) rmdir(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		bytesConsumed        uint64
		err                  error
		nfsProc3RMDirArgs    NFSProc3RMDirArgsStruct
		nfsProc3RMDirResults *NFSProc3RMDirResultsStruct
		results              []byte
		statusOnlyResults    StatusOnlyResultsStruct
	)

	bytesConsumed, err = xdr.Unpack(parms, &nfsProc3RMDirArgs)
	if nil != err {
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}
	if uint64(len(parms)) != bytesConsumed {
		err = fmt.Errorf("xdr.Unpack() failed to consume all of parms")
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}

	nfsProc3RMDirResults = nfsRequestHandler.callbacks.NFSProc3RMDir(authSysBody, &nfsProc3RMDirArgs)

	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) rename(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		bytesConsumed         uint64
		err                   error
		nfsProc3RenameArgs    NFSProc3RenameArgsStruct
		nfsProc3RenameResults *NFSProc3RenameResultsStruct
		results               []byte
		statusOnlyResults     StatusOnlyResultsStruct
	)

	bytesConsumed, err = xdr.Unpack(parms, &nfsProc3RenameArgs)
	if nil != err {
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}
	if uint64(len(parms)) != bytesConsumed {
		err = fmt.Errorf("xdr.Unpack() failed to consume all of parms")
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}

	nfsProc3RenameResults = nfsRequestHandler.callbacks.NFSProc3Rename(authSysBody, &nfsProc3RenameArgs)

	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) link(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		bytesConsumed       uint64
		err                 error
		nfsProc3LinkArgs    NFSProc3LinkArgsStruct
		nfsProc3LinkResults *NFSProc3LinkResultsStruct
		results             []byte
		statusOnlyResults   StatusOnlyResultsStruct
	)

	bytesConsumed, err = xdr.Unpack(parms, &nfsProc3LinkArgs)
	if nil != err {
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}
	if uint64(len(parms)) != bytesConsumed {
		err = fmt.Errorf("xdr.Unpack() failed to consume all of parms")
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}

	nfsProc3LinkResults = nfsRequestHandler.callbacks.NFSProc3Link(authSysBody, &nfsProc3LinkArgs)

	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) readdir(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		bytesConsumed          uint64
		err                    error
		nfsProc3ReadDirArgs    NFSProc3ReadDirArgsStruct
		nfsProc3ReadDirResults *NFSProc3ReadDirResultsStruct
		results                []byte
		statusOnlyResults      StatusOnlyResultsStruct
	)

	bytesConsumed, err = xdr.Unpack(parms, &nfsProc3ReadDirArgs)
	if nil != err {
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}
	if uint64(len(parms)) != bytesConsumed {
		err = fmt.Errorf("xdr.Unpack() failed to consume all of parms")
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}

	nfsProc3ReadDirResults = nfsRequestHandler.callbacks.NFSProc3ReadDir(authSysBody, &nfsProc3ReadDirArgs)

	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) readdirplus(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		bytesConsumed              uint64
		err                        error
		nfsProc3ReadDirPlusArgs    NFSProc3ReadDirPlusArgsStruct
		nfsProc3ReadDirPlusResults *NFSProc3ReadDirPlusResultsStruct
		results                    []byte
		statusOnlyResults          StatusOnlyResultsStruct
	)

	bytesConsumed, err = xdr.Unpack(parms, &nfsProc3ReadDirPlusArgs)
	if nil != err {
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}
	if uint64(len(parms)) != bytesConsumed {
		err = fmt.Errorf("xdr.Unpack() failed to consume all of parms")
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}

	nfsProc3ReadDirPlusResults = nfsRequestHandler.callbacks.NFSProc3ReadDirPlus(authSysBody, &nfsProc3ReadDirPlusArgs)

	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) fsstat(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		bytesConsumed         uint64
		err                   error
		nfsProc3FSStatArgs    NFSProc3FSStatArgsStruct
		nfsProc3FSStatResults *NFSProc3FSStatResultsStruct
		results               []byte
		statusOnlyResults     StatusOnlyResultsStruct
	)

	bytesConsumed, err = xdr.Unpack(parms, &nfsProc3FSStatArgs)
	if nil != err {
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}
	if uint64(len(parms)) != bytesConsumed {
		err = fmt.Errorf("xdr.Unpack() failed to consume all of parms")
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}

	nfsProc3FSStatResults = nfsRequestHandler.callbacks.NFSProc3FSStat(authSysBody, &nfsProc3FSStatArgs)

	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) fsinfo(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		bytesConsumed         uint64
		err                   error
		nfsProc3FSInfoArgs    NFSProc3FSInfoArgsStruct
		nfsProc3FSInfoResults *NFSProc3FSInfoResultsStruct
		results               []byte
		statusOnlyResults     StatusOnlyResultsStruct
	)

	bytesConsumed, err = xdr.Unpack(parms, &nfsProc3FSInfoArgs)
	if nil != err {
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}
	if uint64(len(parms)) != bytesConsumed {
		err = fmt.Errorf("xdr.Unpack() failed to consume all of parms")
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}

	nfsProc3FSInfoResults = nfsRequestHandler.callbacks.NFSProc3FSInfo(authSysBody, &nfsProc3FSInfoArgs)

	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) pathconf(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		bytesConsumed           uint64
		err                     error
		nfsProc3PathConfArgs    NFSProc3PathConfArgsStruct
		nfsProc3PathConfResults *NFSProc3PathConfResultsStruct
		results                 []byte
		statusOnlyResults       StatusOnlyResultsStruct
	)

	bytesConsumed, err = xdr.Unpack(parms, &nfsProc3PathConfArgs)
	if nil != err {
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}
	if uint64(len(parms)) != bytesConsumed {
		err = fmt.Errorf("xdr.Unpack() failed to consume all of parms")
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}

	nfsProc3PathConfResults = nfsRequestHandler.callbacks.NFSProc3PathConf(authSysBody, &nfsProc3PathConfArgs)

	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) commit(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		bytesConsumed         uint64
		err                   error
		nfsProc3CommitArgs    NFSProc3CommitArgsStruct
		nfsProc3CommitResults *NFSProc3CommitResultsStruct
		results               []byte
		statusOnlyResults     StatusOnlyResultsStruct
	)

	bytesConsumed, err = xdr.Unpack(parms, &nfsProc3CommitArgs)
	if nil != err {
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}
	if uint64(len(parms)) != bytesConsumed {
		err = fmt.Errorf("xdr.Unpack() failed to consume all of parms")
		nfsRequestHandler.callbacks.ErrorLog(err)
		err = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		if nil != err {
			nfsRequestHandler.callbacks.ErrorLog(err)
		}
		return
	}

	nfsProc3CommitResults = nfsRequestHandler.callbacks.NFSProc3Commit(authSysBody, &nfsProc3CommitArgs)

	// TODO
}
