package authorizator

const (
	UserWrite  = users + WritePerm
	UserRead   = users + ReadPerm
	UserCreate = users + CreatePerm
	UserDelete = users + RemovePerm
)

func (a *Authorizstor) CanWriteUsers(perm int32) bool {
	return perm>>UserWrite&1 == 1
}

func (a *Authorizstor) SetCanWriteUsers() int32 {
	return 1 << UserWrite
}

func (a *Authorizstor) CanReadUsers(perm int32) bool {
	return perm>>UserRead&1 == 1
}

func (a *Authorizstor) SetCanReadUsers() int32 {
	return 1 << UserRead
}

func (a *Authorizstor) CanCreateUsers(perm int32) bool {
	return perm>>UserCreate&1 == 1
}

func (a *Authorizstor) SetCanCreateUsers() int32 {
	return 1 << UserCreate
}

func (a *Authorizstor) CanRemoveUsers(perm int32) bool {
	return perm>>UserDelete&1 == 1
}

func (a *Authorizstor) SetCanRemoveUsers() int32 {
	return 1 << UserDelete
}
