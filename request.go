package nfsd

import (
	"github.com/swiftstack/onc"
	"github.com/swiftstack/onc/oncserver"
)

type requestHandlerStruct struct {
	oncserver.RequestCallbacks
}

func (requestHandler *requestHandlerStruct) ONCRequest(connHandle oncserver.ConnHandle, xid uint32, prog uint32, vers uint32, proc uint32, authSysBody *onc.AuthSysBodyStruct, parms []byte) {
	// TODO
}
