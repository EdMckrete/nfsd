package nfsd

const ( // ONC RPC Prog/Vers
	MountProgram = uint32(100005) // program MOUNT_PROGRAM
	MountVersion = uint32(3)      // version Mount_V3

	NFSProgram = uint32(100003) //   program NFS_PROGRAM
	NFSVersion = uint32(3)      //   version NFS_V3
)

const ( // Common
	FHSize3 = uint32(64) // Max size in bytes of a file handle

	OK = uint32(0)

	ProcNULL = uint32(0)
)

const ( // Mount-specific
	MntPathLen = uint32(1024) // Maximum bytes in a path name
	MntNameLen = uint32(255)  // Maximum bytes in a name
)

const ( // NFSv3-specific
	NFS3CookieVerfSize = uint32(8) // The size in bytes of the opaque cookie verifier passed by READDIR and READDIRPLUS
	NFS3CreateVerfSize = uint32(8) // The size in bytes of the opaque verifier used for exclusive CREATE
	NFS3WriteVersSize  = uint32(8) // The size in butes of the opaque verifier used for asynchronous WRITE
)

const ( // enum mountstat3
	MNT3ErrPERM        = uint32(1)
	MNT3ErrNOENT       = uint32(2)
	MNT3ErrIO          = uint32(5)
	MNT3ErrACCES       = uint32(13)
	MNT3ErrNOTDIR      = uint32(20)
	MNT3ErrINVAL       = uint32(22)
	MNT3ErrNAMETOOLONG = uint32(63)
	MNT3ErrNOTSUPP     = uint32(10004)
	MNT3ErrSERVERFAULT = uint32(10006)
)

const ( // enum nfsstat3
	NFS3ErrPERM        = uint32(1)
	NFS3ErrNOENT       = uint32(2)
	NFS3ErrIO          = uint32(5)
	NFS3ErrNXIO        = uint32(6)
	NFS3ErrACCES       = uint32(13)
	NFS3ErrEXIST       = uint32(17)
	NFS3ErrXDEV        = uint32(18)
	NFS3ErrNODEV       = uint32(19)
	NFS3ErrNOTDIR      = uint32(20)
	NFS3ErrISDIR       = uint32(21)
	NFS3ErrINVAL       = uint32(22)
	NFS3ErrFBIG        = uint32(27)
	NFS3ErrNOSPC       = uint32(28)
	NFS3ErrROFS        = uint32(30)
	NFS3ErrMLINK       = uint32(31)
	NFS3ErrNAMETOOLONG = uint32(63)
	NFS3ErrNOTEMPTY    = uint32(66)
	NFS3ErrDQUOT       = uint32(69)
	NFS3ErrSTALE       = uint32(70)
	NFS3ErrREMOTE      = uint32(71)
	NFS3ErrBADHANDLE   = uint32(10001)
	NFS3ErrNOTSYNC     = uint32(10002)
	NFS3ErrBADCOOKIE   = uint32(10003)
	NFS3ErrNOTSUPP     = uint32(10004)
	NFS3ErrTOOSMALL    = uint32(10005)
	NFS3ErrSERVERFAULT = uint32(10006)
	NFS3ErrBADTYPE     = uint32(10007)
	NFS3ErrJUKEBOX     = uint32(10008)
)

const ( // program MOUNT_PROGRAM version MOUNT_V3
	MOUNTPROC3MNT  = uint32(1)
	MOUNTPROC3UMNT = uint32(3)
)

const ( // program NFS_PROGRAM version NFS_V3
	NFSPROC3GETATTR     = uint32(1)
	NFSPROC3SETATTR     = uint32(2)
	NFSPROC3LOOKUP      = uint32(3)
	NFSPROC3ACCESS      = uint32(4)
	NFSPROC3READLINK    = uint32(5)
	NFSPROC3READ        = uint32(6)
	NFSPROC3WRITE       = uint32(7)
	NFSPROC3CREATE      = uint32(8)
	NFSPROC3MKDIR       = uint32(9)
	NFSPROC3SYMLINK     = uint32(10)
	NFSPROC3REMOVE      = uint32(12)
	NFSPROC3RMDIR       = uint32(13)
	NFSPROC3RENAME      = uint32(14)
	NFSPROC3LINK        = uint32(15)
	NFSPROC3READDIR     = uint32(16)
	NFSPROC3READDIRPLUS = uint32(17)
	NFSPROC3FSSTAT      = uint32(18)
	NFSPROC3FSINFO      = uint32(19)
	NFSPROC3PATHCONF    = uint32(20)
	NFSPROC3COMMIT      = uint32(21)
)

const ( // enum ftype3
	FTypeREG  = uint32(1)
	FTypeDIR  = uint32(2)
	FTypeBLK  = uint32(3)
	FTypeCHR  = uint32(4)
	FTypeLNK  = uint32(5)
	FTypeSOCK = uint32(6)
	FTypeFIFO = uint32(7)
)

const ( // enum time_how
	DontChange      = uint32(0)
	SetToServerTime = uint32(1)
	SetToClientTime = uint32(2)
)

const (
	Access3Read    = uint32(0x0001)
	Access3Lookup  = uint32(0x0002)
	Access3Modify  = uint32(0x0004)
	Access3Extend  = uint32(0x0008)
	Access3Delete  = uint32(0x0010)
	Access3Execute = uint32(0x0020)
)

const ( // enum stable_how
	Unstable = uint32(0)
	DataSync = uint32(1)
	FileSync = uint32(2)
)

const ( // enum createmode3
	Unchecked = uint32(0)
	Guarded   = uint32(1)
	Exclusive = uint32(2)
)

const (
	FSF3Link        = uint32(0x0001)
	FSF3SymLink     = uint32(0x0002)
	FSF3Homogeneous = uint32(0x0008)
	FSF3CanSetTime  = uint32(0x0010)
)
