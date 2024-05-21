package configparser

import "github.com/VladimirRytov/advsrv/internal/datatransferobjects"

type Config struct {
	ListenParams datatransferobjects.ListenParams `yaml:"listenParams"`
	Spec         Spec
}

type Spec struct {
	RecordsDatabases RecordsDatabases `yaml:"recordsDatabases"`
	UserAuth         UserAuth         `yaml:"userAuth"`
	Users            []User
	FileStorage      datatransferobjects.FileStorage
}

type RecordsDatabases struct {
	NetworkDatabases []datatransferobjects.NetworkDataBaseDSN `yaml:"networkDatabases"`
	LocalDatabases   []datatransferobjects.LocalDSN           `yaml:"localDatabases"`
}

type UserAuth struct {
	NetworkDatabases []datatransferobjects.NetworkDataBaseDSN `yaml:"networkDatabases"`
	LocalDatabases   []datatransferobjects.LocalDSN           `yaml:"localDatabases"`
}

type User struct {
	Name               string
	Salted             bool
	AllowEmptyPassword bool `yaml:"allowEmptyPassword"`
	Password           string
	Permissions        Permissions
}

type Permissions struct {
	Databases []string
	Users     []string
	Records   []string
	Files     []string
}
