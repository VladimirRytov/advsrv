package authorizatorhandler

type PermHandler interface {
	FilePerms
	DatabasePerms
	RecordsPerm
	UsersPerm
}

type FilePerms interface {
	CanWriteFiles(perm int32) bool
	CanReadFiles(perm int32) bool
	CanCreateFiles(perm int32) bool
	CanRemoveFiles(perm int32) bool
}

type DatabasePerms interface {
	CanWriteDatabases(perm int32) bool
	CanReadDatabases(perm int32) bool
	CanCreateDatabases(perm int32) bool
	CanRemoveDatabases(perm int32) bool
}

type RecordsPerm interface {
	CanWriteRecords(perm int32) bool
	CanReadRecords(perm int32) bool
	CanCreateRecords(perm int32) bool
	CanRemoveRecords(perm int32) bool
}

type UsersPerm interface {
	CanWriteUsers(perm int32) bool
	CanReadUsers(perm int32) bool
	CanCreateUsers(perm int32) bool
	CanRemoveUsers(perm int32) bool
}
