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
	NFSProc3Lookup(nfsProc3LookupArgs *NFSProc3LookupArgsStruct) (nfsProc3LookupResults *NFSProc3LookupResultsStruct)
	NFSProc3Access(nfsProc3AccessArgs *NFSProc3AccessArgsStruct) (nfsProc3AccessResults *NFSProc3AccessResultsStruct)
	NFSProc3ReadLink(nfsProc3ReadLinkArgs *NFSProc3ReadLinkArgsStruct) (nfsProc3ReadLinkResults *NFSProc3ReadLinkResultsStruct)
	NFSProc3Read(nfsProc3ReadArgs *NFSProc3ReadArgsStruct) (nfsProc3ReadResults *NFSProc3ReadResultsStruct)
	NFSProc3Write(nfsProc3WriteArgs *NFSProc3WriteArgsStruct) (nfsProc3WriteResults *NFSProc3WriteResultsStruct)
	NFSProc3Create(nfsProc3CreateArgs *NFSProc3CreateArgsStruct) (nfsProc3CreateResults *NFSProc3CreateResultsStruct)
	NFSProc3MKDir(nfsProc3MKDirArgs *NFSProc3MKDirArgsStruct) (nfsProc3MKDirResults *NFSProc3MKDirResultsStruct)
	NFSProc3SymLink(nfsProc3SymLinkArgs *NFSProc3SymLinkArgsStruct) (nfsProc3SymLinkResults *NFSProc3SymLinkResultsStruct)
	NFSProc3Remove(nfsProc3RemoveArgs *NFSProc3RemoveArgsStruct) (nfsProc3RemoveResults *NFSProc3RemoveResultsStruct)
	NFSProc3RMDir(nfsProc3RMDirArgs *NFSProc3RMDirArgsStruct) (nfsProc3RMDirResults *NFSProc3RMDirResultsStruct)
	NFSProc3Rename(nfsProc3RenameArgs *NFSProc3RenameArgsStruct) (nfsProc3RenameResults *NFSProc3RenameResultsStruct)
	NFSProc3Link(nfsProc3LinkArgs *NFSProc3LinkArgsStruct) (nfsProc3LinkResults *NFSProc3LinkResultsStruct)
	NFSProc3ReadDir(nfsProc3ReadDirArgs *NFSProc3ReadDirArgsStruct) (nfsProc3ReadDirResults *NFSProc3ReadDirResultsStruct)
	NFSProc3ReadDirPlus(nfsProc3ReadDirPlusArgs *NFSProc3ReadDirPlusArgsStruct) (nfsProc3ReadDirPlusResults *NFSProc3ReadDirPlusResultsStruct)
	NFSProc3FSStat(nfsProc3FSStatArgs *NFSProc3FSStatArgsStruct) (nfsProc3FSStatResults *NFSProc3FSStatResultsStruct)
	NFSProc3FSInfo(nfsProc3FSInfoArgs *NFSProc3FSInfoArgsStruct) (nfsProc3FSInfoResults *NFSProc3FSInfoResultsStruct)
	NFSProc3PathConf(nfsProc3PathConfArgs *NFSProc3PathConfArgsStruct) (nfsProc3PathConfResults *NFSProc3PathConfResultsStruct)
	NFSProc3Commit(nfsProc3CommitArgs *NFSProc3CommitArgsStruct) (nfsProc3CommitResults *NFSProc3CommitResultsStruct)
}
