package nfsd

// Basic XDR type structs to enable partial xdr.Pack()/xdr.Unpack() of ONC RPC's using "union/switch" and "void"
//
// Where possible (i.e. when the corresponding ONC RPC specification doesn't use union/switch or void),
// these structs are tagged with the package xdr annotations enabling direct xdr.Pack()/xdr.Unpack().

type UnsignedIntegerStruct struct {
	UnsignedInteger uint32 `XDR_Name:"Unsigned Integer"`
}

type EnumerationStruct struct {
	Enumeration uint32 `XDR_Name:"Enumeration"`
}

type BooleanStruct struct {
	Bool bool `XDR_Name:"Boolean"`
}

type UnsignedHyperIntegerStruct struct {
	UnsignedHyperInteger uint64 `XDR_Name:"Unsigned Hyper Integer"`
}

// Mount V3 / NFSv3 API embedded structs

type SpecData3Struct struct { // struct specdata3
	SpecData1 uint32 `XDR_Name:"Unsigned Integer"` // should be == 0
	SpecData2 uint32 `XDR_Name:"Unsigned Integer"` // should be == 0
}

type NFSTime3Struct struct { // struct nfstime3
	Seconds  uint32 `XDR_Name:"Unsigned Integer"`
	NSeconds uint32 `XDR_Name:"Unsigned Integer"`
}

type FAttr3Struct struct { // struct fattr3
	Type   uint32          `XDR_Name:"Enumeration"` // enum ftype3
	Mode   uint32          `XDR_Name:"Unsigned Integer"`
	NLink  uint32          `XDR_Name:"Unsigned Integer"`
	UID    uint32          `XDR_Name:"Unsigned Integer"`
	GID    uint32          `XDR_Name:"Unsigned Integer"`
	Size   uint64          `XDR_Name:"Hyper Integer"`
	Used   uint64          `XDR_Name:"Hyper Integer"`
	RDev   SpecData3Struct `XDR_Name:"Structure"`
	FSID   uint64          `XDR_Name:"Hyper Integer"`
	FileID uint64          `XDR_Name:"Hyper Integer"`
	ATime  NFSTime3Struct  `XDR_Name:"Structure"`
	MTime  NFSTime3Struct  `XDR_Name:"Structure"`
	CTime  NFSTime3Struct  `XDR_Name:"Structure"`
}

type SAttr3Struct struct { // struct sattr3
	SetMode  bool
	Mode     uint32 //         only used/valid if SetMode == true
	SetUID   bool
	UID      uint32 //         only used/valid if SetUID == true
	SetGID   bool
	GID      uint32 //         only used/valid if SetGID == true
	SetSize  bool
	Size     uint64         // only used/valid if SetSize == true
	SetATime uint32         // enum time_how
	ATime    NFSTime3Struct // only used/valid if time_how == SetToClientTime
	SetMTime uint32         // enum time_how
	MTime    NFSTime3Struct // only used/valid if time_how == SetToClientTime
}

type SAttrGuard3Struct struct { // union sattrguard3
	CheckCTime bool
	CTime      NFSTime3Struct // only used/valid if CheckCTime == true
}

type WCCAttrStruct struct { // struct wcc_attr
	Size  uint64         `XDR_Name:"Hyper Integer"`
	MTime NFSTime3Struct `XDR_Name:"Structure"`
	CTime NFSTime3Struct `XDR_Name:"Structure"`
}

type PreOpAttrStruct struct { // union pre_op_attr
	AttributesFollow bool
	Attributes       WCCAttrStruct // only used/valid if AttributesFollow == true
}

type PostOpAttrStruct struct { // union post_op_attr
	AttributesFollow bool
	Attributes       FAttr3Struct // only used/valid if AttributesFollow == true
}

type WCCDataStruct struct { // struct wcc_data
	Before PreOpAttrStruct
	After  PostOpAttrStruct
}

// Mount V3 API call/reply structs

type MountProc3MntArgsStruct struct {
	DirPath string `XDR_Name:"String" XDR_MaxSize:"1024"`
}

type MountProc3MntResultsStruct struct { // union mountres3
	Status  uint32 `XDR_Name:"Enumeration"`                                  // OK or enum mountstat3
	FHandle []byte `XDR_Name:"Variable-Length Opaque Data" XDR_MaxSize:"64"` // only used/valid if Status == OK
}

type MountProc3UmntArgsStruct struct {
	DirPath string `XDR_Name:"String" XDR_MaxSize:"1024"`
}

// NFSv3 API call/reply structs

type NFSProc3GetAttrArgsStruct struct {
	FHandle []byte `XDR_Name:"Variable-Length Opaque Data" XDR_MaxSize:"64"`
}

type NFSProc3GetAttrResultsStruct struct {
	Status     uint32       `XDR_Name:"Enumeration"` // OK or enum nfsstat3
	Attributes FAttr3Struct `XDR_Name:"Structure"`   // only used/valid if Status == OK
}

type NFSProc3SetAttrArgsStruct struct {
	FHandle       []byte
	NewAttributes SAttr3Struct
	Guard         SAttrGuard3Struct
}

type NFSProc3SetAttrResultsStruct struct {
	WCC WCCDataStruct
}
