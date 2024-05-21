package authorizator

const (
	RecordWrite  = records + WritePerm
	RecordRead   = records + ReadPerm
	RecordCreate = records + CreatePerm
	RecordDelete = records + RemovePerm
)

func (a *Authorizstor) CanWriteRecords(perm int32) bool {
	return perm>>RecordWrite&1 == 1
}

func (a *Authorizstor) SetCanWriteRecords() int32 {
	return 1 << RecordWrite
}

func (a *Authorizstor) CanReadRecords(perm int32) bool {
	return perm>>RecordRead&1 == 1
}

func (a *Authorizstor) SetCanReadRecords() int32 {
	return 1 << RecordRead
}

func (a *Authorizstor) CanCreateRecords(perm int32) bool {
	return perm>>RecordCreate&1 == 1
}

func (a *Authorizstor) SetCanCreateRecords() int32 {
	return 1 << RecordCreate
}

func (a *Authorizstor) CanRemoveRecords(perm int32) bool {
	return perm>>RecordDelete&1 == 1
}

func (a *Authorizstor) SetCanRemoveRecords() int32 {
	return 1 << RecordDelete
}
