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
	if onc.ProgNumMount != prog {
		err := fmt.Errorf("prog was %v... expected it to be onc.ProgNumMount (%v)", prog, onc.ProgNumMount)
		panic(err)
	}

	if 3 != vers {
		err := fmt.Errorf("vers was %v... expected it to be 3", vers)
		panic(err)
	}

	switch proc {
	case ProcNULL:
		mountRequestHandler.null(connHandle, xid, authSysBody, parms)
	case MOUNTPROC3MNT:
		mountRequestHandler.mnt(connHandle, xid, authSysBody, parms)
	case MOUNTPROC3UMNT:
		mountRequestHandler.umnt(connHandle, xid, authSysBody, parms)
	default:
		_ = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.ProcUnavail)
	}
}

func (mountRequestHandler *mountRequestHandlerStruct) null(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	if 0 != len(parms) {
		_ = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		return
	}

	mountRequestHandler.callbacks.MountProc3Null(authSysBody)

	_ = oncserver.SendAcceptedSuccess(connHandle, xid, nil)
}

func (mountRequestHandler *mountRequestHandlerStruct) mnt(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		bytesConsumed        uint64
		err                  error
		mountProc3MntArgs    MountProc3MntArgsStruct
		mountProc3MntResults *MountProc3MntResultsStruct
		results              []byte
	)

	bytesConsumed, err = xdr.Unpack(parms, &mountProc3MntArgs)
	if nil != err {
		_ = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		return
	}
	if uint64(len(parms)) != bytesConsumed {
		_ = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		return
	}

	mountProc3MntResults = mountRequestHandler.callbacks.MountProc3Mnt(authSysBody, &mountProc3MntArgs)

	results, err = xdr.Pack(mountProc3MntResults)
	if nil != err {
		_ = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.SystemErr)
		return
	}

	_ = oncserver.SendAcceptedSuccess(connHandle, xid, results)
}

func (mountRequestHandler *mountRequestHandlerStruct) umnt(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	var (
		bytesConsumed      uint64
		err                error
		mountProc3UmntArgs MountProc3UmntArgsStruct
	)

	bytesConsumed, err = xdr.Unpack(parms, &mountProc3UmntArgs)
	if nil != err {
		_ = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		return
	}
	if uint64(len(parms)) != bytesConsumed {
		_ = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		return
	}

	mountRequestHandler.callbacks.MountProc3Umnt(authSysBody, &mountProc3UmntArgs)

	_ = oncserver.SendAcceptedSuccess(connHandle, xid, nil)
}

func (nfsRequestHandler *nfsRequestHandlerStruct) ONCRequest(connHandle oncserver.ConnHandle, xid uint32, prog uint32, vers uint32, proc uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	if onc.ProgNumNFS != prog {
		err := fmt.Errorf("prog was %v... expected it to be onc.ProgNumNFS (%v)", prog, onc.ProgNumNFS)
		panic(err)
	}

	if 3 != vers {
		err := fmt.Errorf("vers was %v... expected it to be 3", vers)
		panic(err)
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
		_ = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.ProcUnavail)
	}
}

func (nfsRequestHandler *nfsRequestHandlerStruct) null(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	if 0 != len(parms) {
		_ = oncserver.SendAcceptedOtherErrorReply(connHandle, xid, onc.GarbageArgs)
		return
	}

	nfsRequestHandler.callbacks.NFSProc3Null(authSysBody)

	_ = oncserver.SendAcceptedSuccess(connHandle, xid, nil)
}

func (nfsRequestHandler *nfsRequestHandlerStruct) getattr(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) setattr(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) lookup(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) access(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) readlink(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) read(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) write(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) create(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) mkdir(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) symlink(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) remove(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) rmdir(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) rename(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) link(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) readdir(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) readdirplus(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) fsstat(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) fsinfo(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) pathconf(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	// TODO
}

func (nfsRequestHandler *nfsRequestHandlerStruct) commit(connHandle oncserver.ConnHandle, xid uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	// TODO
}
