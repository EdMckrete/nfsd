package nfsd

// See also consts.go and structs.go for exported constants and structures referenced by this API

type MountV3Interface interface {
	MountProc3Null()
	MountProc3Mnt(dirpath string) (fhsStatus uint32, fhandle []byte)
	MountProc3Umnt(dirpath string)
}

type NFSV3Interface interface {
	NFSProc3Null()
	NFSProc3GetAttr(fh []byte) (status uint32, fattr3 *FAttr3Struct)
	NFSProc3SetAttr(fh []byte, newAttributes *SAttr3Struct, guard *SAttrGuard3Struct) (status uint32) // , SETATTR3res
	/*
		NFSProc3Lookup      = uint32(3)
		NFSProc3Access      = uint32(4)
		NFSProc3ReadLink    = uint32(5)
		NFSProc3Read        = uint32(6)
		NFSProc3Write       = uint32(7)
		NFSProc3Create      = uint32(8)
		NFSProc3MKDir       = uint32(9)
		NFSProc3SymLink     = uint32(10)
		NFSProc3Rempve      = uint32(12)
		NFSProc3RMDir       = uint32(13)
		NFSProc3Rename      = uint32(14)
		NFSProc3Link        = uint32(15)
		NFSProc3ReadDir     = uint32(16)
		NFSProc3ReadDirPlus = uint32(17)
		NFSProc3FSStat      = uint32(18)
		NFSProc3FSInfo      = uint32(19)
		NFSProc3PathConf    = uint32(20)
		NFSProc3Commit      = uint32(21)
	*/
}
