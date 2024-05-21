package authorizator

const (
	DatabaseWrite  = WritePerm + databases
	DatabaseRead   = ReadPerm + databases
	DatabaseCreate = CreatePerm + databases
	DatabaseDelete = RemovePerm + databases
)

func (a *Authorizstor) CanWriteDatabases(perm int32) bool {
	return perm>>DatabaseWrite&1 == 1
}

func (a *Authorizstor) SetCanWriteDatabases() int32 {
	return 1 << DatabaseWrite
}

func (a *Authorizstor) CanReadDatabases(perm int32) bool {
	return perm>>DatabaseRead&1 == 1
}

func (a *Authorizstor) SetCanReadDatabases() int32 {
	return 1 << DatabaseRead
}

func (a *Authorizstor) CanCreateDatabases(perm int32) bool {
	return perm>>DatabaseCreate&1 == 1
}

func (a *Authorizstor) SetCanCreateDatabases() int32 {
	return 1 << DatabaseCreate
}

func (a *Authorizstor) CanRemoveDatabases(perm int32) bool {
	return perm>>DatabaseDelete&1 == 1
}

func (a *Authorizstor) SetCanRemoveDatabases() int32 {
	return 1 << DatabaseDelete
}
