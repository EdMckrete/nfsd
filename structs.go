package nfsd

type SpecData3Struct struct {
	SpecData1 uint32 `XDR_Name:"Unsigned Integer"` // should be == 0
	SpecData2 uint32 `XDR_Name:"Unsigned Integer"` // should be == 0
}

type NFSTime3Struct struct {
	Seconds  uint32 `XDR_Name:"Unsigned Integer"`
	NSeconds uint32 `XDR_Name:"Unsigned Integer"`
}

type FAttr3Struct struct {
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

type SAttr3Struct struct {
	SetMode  bool
	Mode     uint32
	SetUID   bool
	UID      uint32
	SetGID   bool
	GID      uint32
	SetSize  bool
	Size     uint64
	SetATime uint32 // enum time_how
	ATime    NFSTime3Struct
	SetMTime uint32 // enum time_how
	MTime    NFSTime3Struct
}

type SAttrGuard3Struct struct {
	CheckCTime bool
	CTime      NFSTime3Struct
}
