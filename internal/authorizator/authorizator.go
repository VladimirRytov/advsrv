package authorizator

const (
	records int32 = iota * 4
	users
	databases
	files

	CreatePerm int32 = 3
	WritePerm  int32 = 2
	ReadPerm   int32 = 1
	RemovePerm int32 = 0

	AdminPermissions int32 = int32(^uint32(0) >> 1)
)

type Authorizstor struct{}

func NewAuthrizator() *Authorizstor {
	return &Authorizstor{}
}

func (a *Authorizstor) AdminPermissions() int32 {
	return AdminPermissions
}
