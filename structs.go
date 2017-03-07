package nfsd

// Basic XDR type structs to enable partial xdr.Pack()/xdr.Unpack() of ONC RPC's using "union/switch" and "void"
//
// Where possible (i.e. when the corresponding ONC RPC specification doesn't use union/switch or void),
// these structs are tagged with the package xdr annotations enabling direct xdr.Pack()/xdr.Unpack().

type StatusOnlyResultsStruct struct {
	Status uint32 `XDR_Name:"Enumeration"` // enum nfsstat3
}

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
	SetMode  bool           //
	Mode     uint32         // only used/valid if SetMode == true
	SetUID   bool           //
	UID      uint32         // only used/valid if SetUID == true
	SetGID   bool           //
	GID      uint32         // only used/valid if SetGID == true
	SetSize  bool           //
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

type PostOpFh3Struct struct { // union post_op_fh3
	HandleFollows bool
	Handle        []byte // only used/valid if HandleFollows == true
}

type WCCDataStruct struct { // struct wcc_data
	Before PreOpAttrStruct
	After  PostOpAttrStruct
}

type DirOpArgs3Struct struct { // struct diropargs3
	Dir  []byte `XDR_Name:"Variable-Length Opaque Data" XDR_MaxSize:"64"`
	Name string `XDR_Name:"String" XDR_MaxSize:"255"`
}

type CreateHowStruct struct { // union CreateHowStruct
	Mode          uint32                   // enum createmode3
	ObjAttributes SAttr3Struct             // only used/valid if Mode == Unchecked || Mode == Guarded
	Verf          [NFS3CreateVerfSize]byte // only used/valid if Mode == Exclusive
}

type DirListEntryStruct struct { // struct entry3
	FileID uint64 `XDR_Name:"Unsigned Hyper Integer"`
	Name   string `XDR_Name:"String" XDR_MaxSize:"255"`
	Cookie uint64 `XDR_Name:"Unsigned Hyper Integer"`
}

type DirListEntryPlusStruct struct { // struct entryplus3
	FileID         uint64 `XDR_Name:"Unsigned Hyper Integer"`
	Name           string `XDR_Name:"String" XDR_MaxSize:"255"`
	Cookie         uint64 `XDR_Name:"Unsigned Hyper Integer"`
	NameAttributes PostOpAttrStruct
	NameHandle     PostOpFh3Struct
}

// Mount V3 API call/reply structs

type MountProc3MntArgsStruct struct {
	DirPath string `XDR_Name:"String" XDR_MaxSize:"1024"`
}

type MountProc3MntResultsStruct struct { // union mountres3
	Status      uint32   `XDR_Name:"Enumeration"`                                  // OK or enum mountstat3
	FHandle     []byte   `XDR_Name:"Variable-Length Opaque Data" XDR_MaxSize:"64"` // only used/valid if Status == OK
	AuthFlavors []uint32 `XDR_Name:"Variable-Length Array"`                        // only used/valid if Status == OK; enum auth_flavor; must == AuthSys
}

type MountProc3UmntArgsStruct struct {
	DirPath string `XDR_Name:"String" XDR_MaxSize:"1024"`
}

// NFSv3 API call/reply structs

type NFSProc3GetAttrArgsStruct struct {
	Object []byte `XDR_Name:"Variable-Length Opaque Data" XDR_MaxSize:"64"`
}

type NFSProc3GetAttrResultsStruct struct {
	Status     uint32       `XDR_Name:"Enumeration"` // OK or enum nfsstat3
	Attributes FAttr3Struct `XDR_Name:"Structure"`   // only used/valid if Status == OK
}

type NFSProc3SetAttrArgsStruct struct {
	Object        []byte
	NewAttributes SAttr3Struct
	Guard         SAttrGuard3Struct
}

type NFSProc3SetAttrResultsStruct struct {
	Status uint32 // OK or enum nfsstat3
	WCC    WCCDataStruct
}

type NFSProc3LookupArgsStruct struct {
	What DirOpArgs3Struct `XDR_Name:"Structure"`
}

type NFSProc3LookupResultsStruct struct {
	Status        uint32           // OK or enum nfsstat3
	Object        []byte           // only used/valid if Status == OK
	ObjAttributes PostOpAttrStruct // only used/valid if Status == OK
	DirAttributes PostOpAttrStruct
}

type NFSProc3AccessArgsStruct struct {
	Object []byte `XDR_Name:"Variable-Length Opaque Data" XDR_MaxSize:"64"`
	Access uint32 `XDR_Name:"Unsigned Integer"`
}

type NFSProc3AccessResultsStruct struct {
	Status        uint32           // OK or enum nfsstat3
	ObjAttributes PostOpAttrStruct //
	Access        uint32           // only used/valid if Status == OK
}

type NFSProc3ReadLinkArgsStruct struct {
	SymLink []byte `XDR_Name:"Variable-Length Opaque Data" XDR_MaxSize:"64"`
}

type NFSProc3ReadLinkResultsStruct struct {
	Status            uint32           // OK or enum nfsstat3
	SymLinkAttributes PostOpAttrStruct //
	Path              []byte           // only used/valid if Status == OK
}

type NFSProc3ReadArgsStruct struct {
	File   []byte `XDR_Name:"Variable-Length Opaque Data" XDR_MaxSize:"64"`
	Offset uint64 `XDR_Name:"Hyper Integer"`
	Count  uint32 `XDR_Name:"Unsigned Integer"`
}

type NFSProc3ReadResultsStruct struct {
	Status         uint32           // OK or enum nfsstat3
	FileAttributes PostOpAttrStruct //
	Count          uint32           // only used/valid if Status == OK
	EOF            bool             // only used/valid if Status == OK
	Data           []byte           // only used/valid if Status == OK
}

type NFSProc3WriteArgsStruct struct {
	File   []byte `XDR_Name:"Variable-Length Opaque Data" XDR_MaxSize:"64"` //
	Offset uint64 `XDR_Name:"Hyper Integer"`                                //
	Count  uint32 `XDR_Name:"Unsigned Integer"`                             //
	Stable uint32 `XDR_Name:"Enumeration"`                                  // enum stable_how
	Data   []byte `XDR_Name:"Variable-Length Opaque Data"`                  //
}

type NFSProc3WriteResultsStruct struct {
	Status    uint32                  // OK or enum nfsstat3
	FileWCC   WCCDataStruct           //
	Count     uint32                  // only used/valid if Status == OK
	Committed uint32                  // only used/valid if Status == OK; enum stable_how
	Verf      [NFS3WriteVersSize]byte // only used/valid if Status == OK
}

type NFSProc3CreateArgsStruct struct {
	Where DirOpArgs3Struct
	How   CreateHowStruct
}

type NFSProc3CreateResultsStruct struct {
	Status        uint32           // OK or enum nfsstat3
	Obj           PostOpFh3Struct  // only used/valid if Status == OK
	ObjAttributes PostOpAttrStruct // only used/valid if Status == OK
	DirWCC        WCCDataStruct
}

type NFSProc3MKDirArgsStruct struct {
	Where      DirOpArgs3Struct
	Attributes SAttr3Struct
}

type NFSProc3MKDirResultsStruct struct {
	Status        uint32           // OK or enum nfsstat3
	Obj           PostOpFh3Struct  // only used/valid if Status == OK
	ObjAttributes PostOpAttrStruct // only used/valid if Status == OK
	DirWCC        WCCDataStruct
}

type NFSProc3SymLinkArgsStruct struct {
	Where             DirOpArgs3Struct
	SymLinkAttributes SAttr3Struct
	SymLinkData       []byte
}

type NFSProc3SymLinkResultsStruct struct {
	Status        uint32           // OK or enum nfsstat3
	Obj           PostOpFh3Struct  // only used/valid if Status == OK
	ObjAttributes PostOpAttrStruct // only used/valid if Status == OK
	DirWCC        WCCDataStruct
}

type NFSProc3RemoveArgsStruct struct {
	Where DirOpArgs3Struct `XDR_Name:"Structure"`
}

type NFSProc3RemoveResultsStruct struct {
	Status uint32 // OK or enum nfsstat3
	DirWCC WCCDataStruct
}

type NFSProc3RMDirArgsStruct struct {
	Where DirOpArgs3Struct `XDR_Name:"Structure"`
}

type NFSProc3RMDirResultsStruct struct {
	Status uint32 // OK or enum nfsstat3
	DirWCC WCCDataStruct
}

type NFSProc3RenameArgsStruct struct {
	From DirOpArgs3Struct `XDR_Name:"Structure"`
	To   DirOpArgs3Struct `XDR_Name:"Structure"`
}

type NFSProc3RenameResultsStruct struct {
	Status     uint32 // OK or enum nfsstat3
	FromDirWCC WCCDataStruct
	ToDirWCC   WCCDataStruct
}

type NFSProc3LinkArgsStruct struct {
	File []byte           `XDR_Name:"Variable-Length Opaque Data" XDR_MaxSize:"64"`
	Link DirOpArgs3Struct `XDR_Name:"Structure"`
}

type NFSProc3LinkResultsStruct struct {
	Status         uint32 // OK or enum nfsstat3
	FileAttributes PostOpAttrStruct
	LinkDirWCC     WCCDataStruct
}

type NFSProc3ReadDirArgsStruct struct {
	Dir        []byte                   `XDR_Name:"Variable-Length Opaque Data" XDR_MaxSize:"64"`
	Cookie     uint64                   `XDR_Name:"Unsigned Hyper Integer"`
	CookieVerf [NFS3CookieVerfSize]byte `XDR_Name:"Fixed-Length Opaque Data"`
	Count      uint32                   `XDR_Name:"Unsigned Integer"`
}

type NFSProc3ReadDirResultsStruct struct {
	Status        uint32                   // OK or enum nfsstat3
	DirAttributes PostOpAttrStruct         //
	CookieVerf    [NFS3CookieVerfSize]byte // only used/valid if Status == OK
	Entries       []DirListEntryStruct     // only used/valid if Status == OK
	EOF           bool                     // only used/valid if Status == OK
}

type NFSProc3ReadDirPlusArgsStruct struct {
	Dir        []byte                   `XDR_Name:"Variable-Length Opaque Data" XDR_MaxSize:"64"`
	Cookie     uint64                   `XDR_Name:"Unsigned Hyper Integer"`
	CookieVerf [NFS3CookieVerfSize]byte `XDR_Name:"Fixed-Length Opaque Data"`
	DirCount   uint32                   `XDR_Name:"Unsigned Integer"`
	MaxCount   uint32                   `XDR_Name:"Unsigned Integer"`
}

type NFSProc3ReadDirPlusResultsStruct struct {
	Status        uint32                   // OK or enum nfsstat3
	DirAttributes PostOpAttrStruct         //
	CookieVerf    [NFS3CookieVerfSize]byte // only used/valid if Status == OK
	Entries       []DirListEntryPlusStruct // only used/valid if Status == OK
	EOF           bool                     // only used/valid if Status == OK
}

type NFSProc3FSStatArgsStruct struct {
	FSRoot []byte `XDR_Name:"Variable-Length Opaque Data" XDR_MaxSize:"64"`
}

type NFSProc3FSStatResultsStruct struct {
	Status        uint32           // OK or enum nfsstat3
	ObjAttributes PostOpAttrStruct //
	TBytes        uint64           // only used/valid if Status == OK
	FBytes        uint64           // only used/valid if Status == OK
	ABytes        uint64           // only used/valid if Status == OK
	TFiles        uint64           // only used/valid if Status == OK
	FFiles        uint64           // only used/valid if Status == OK
	AFiles        uint64           // only used/valid if Status == OK
	InvarSec      uint32           // only used/valid if Status == OK
}

type NFSProc3FSInfoArgsStruct struct {
	FSRoot []byte `XDR_Name:"Variable-Length Opaque Data" XDR_MaxSize:"64"`
}

type NFSProc3FSInfoResultsStruct struct {
	Status        uint32           // OK or enum nfsstat3
	ObjAttributes PostOpAttrStruct //
	RTMax         uint32           // only used/valid if Status == OK
	RTPref        uint32           // only used/valid if Status == OK
	RTMult        uint32           // only used/valid if Status == OK
	WTMax         uint32           // only used/valid if Status == OK
	WTMult        uint32           // only used/valid if Status == OK
	DTPref        uint32           // only used/valid if Status == OK
	MaxFileSize   uint64           // only used/valid if Status == OK
	TimeDelta     NFSTime3Struct   // only used/valid if Status == OK
	Properties    uint32           // only used/valid if Status == OK
}

type NFSProc3PathConfArgsStruct struct {
	Object []byte `XDR_Name:"Variable-Length Opaque Data" XDR_MaxSize:"64"`
}

type NFSProc3PathConfResultsStruct struct {
	Status          uint32           // OK or enum nfsstat3
	ObjAttributes   PostOpAttrStruct //
	LinkMax         uint32           // only used/valid if Status == OK
	NameMax         uint32           // only used/valid if Status == OK
	NoTrunc         bool             // only used/valid if Status == OK
	ChOwnRestricted bool             // only used/valid if Status == OK
	CaseInsensitive bool             // only used/valid if Status == OK
	CasePreserving  bool             // only used/valid if Status == OK
}

type NFSProc3CommitArgsStruct struct {
	File   []byte `XDR_Name:"Variable-Length Opaque Data" XDR_MaxSize:"64"`
	Offset uint64 `XDR_Name:"Unsigned Hyper Integer"`
	Count  uint32 `XDR_Name:"Unsigned Integer"`
}

type NFSProc3CommitResultsStruct struct {
	Status  uint32                  // OK or enum nfsstat3
	FileWCC WCCDataStruct           //
	Verf    [NFS3WriteVersSize]byte // only used/valid if Status == OK
}
