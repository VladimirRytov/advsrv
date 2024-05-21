package authorizator

const (
	FileWrite  = WritePerm + files
	FileRead   = files + ReadPerm
	FileCreate = files + CreatePerm
	FileDelete = files + RemovePerm
)

func (a *Authorizstor) CanWriteFiles(perm int32) bool {
	return perm>>FileWrite&1 == 1
}

func (a *Authorizstor) SetCanWriteFiles() int32 {
	return 1 << FileWrite
}

func (a *Authorizstor) CanReadFiles(perm int32) bool {
	return perm>>FileRead&1 == 1
}

func (a *Authorizstor) SetCanReadFiles() int32 {
	return 1 << FileRead
}

func (a *Authorizstor) CanCreateFiles(perm int32) bool {
	return perm>>FileCreate&1 == 1
}

func (a *Authorizstor) SetCanCreateFiles() int32 {
	return 1 << FileCreate
}

func (a *Authorizstor) CanRemoveFiles(perm int32) bool {
	return perm>>FileDelete&1 == 1
}

func (a *Authorizstor) SetCanRemoveFiles() int32 {
	return 1 << FileDelete
}
