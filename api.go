package nfsd

// See also consts.go and structs.go for exported constants and structures referenced by this API

type MountV3Interface interface {
	MountProc3Null()
	MountProc3Mnt(mountProc3MntArgs *MountProc3MntArgsStruct) (mountProc3MntResults *MountProc3MntResultsStruct)
	MountProc3Umnt(mountProc3UmntArgs *MountProc3UmntArgsStruct)
}

type NFSV3Interface interface {
	NFSProc3Null()
	NFSProc3GetAttr(nfsProc3GetAttrArgs *NFSProc3GetAttrArgsStruct) (nfsProc3GetAttrResults *NFSProc3GetAttrResultsStruct)
	NFSProc3SetAttr(nfsProc3SetAttrArgs *NFSProc3SetAttrArgsStruct) (nfsProc3SetAttrResults *NFSProc3SetAttrResultsStruct)
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
