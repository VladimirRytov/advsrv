package configparser

import (
	"os"
	"path/filepath"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"

	"github.com/spf13/viper"
)

const Template = `listenParams:
  adress: '127.0.0.1'
  port: 4567
#  tls:
#    pathToCertificate: ""
#    pathToKey: ""

spec:
  fileStorage:
    path: ""

# you can use local or network params or both to connect to database
# first the server will try to connect to network databases, then will try to connect to local databases
# the server will use the first successful connection 
  recordsDatabases: # options for records
    networkDatabases:
    - type: "Postgres" # Postgress, Mysql, Sql Server
      port: 5432
      adress: '127.0.0.1' 
      dbName: 'gorm_test'
      userName: 'gorm_test'
      password: 'gorm_test'

#    localDatabases:
#    - type: 'Sqlite'
#      path: ''  
#      name: 'name.sqlite'

  userAuth: # options for users
    networkDatabases:
    - type: 'Postgres' # Postgress, Mysql, Sql Server
      port: 5432
      adress: '127.0.0.1' 
      dbName: 'gorm_test'
      userName: 'gorm_test'
      password: 'gorm_test'  
#    localDatabases:
#    - type: 'Sqlite'
#      path: ''  
#      name: ':memory:'

  users:
  - name: 'admin'
# password will be salted after successful connection
    password: '1234'
    permissions:
      databases: 
      - "all" # all, none (default), read, write, remove, create
      - "create"
      users: 
      - "all"
      records: 
      - "all"
  - name: 'user'  
    allowEmptyPassword: true  
`
const (
	ConfigFileFormat = "yaml"
	ConfigFileName   = "config"
	Records          = "recordsDatabases"
	Users            = "userAuth"
)

func (c *ConfigHandler) Records() string { return Records }
func (c *ConfigHandler) Users() string   { return Users }

type Base64 interface {
	ToBase64String(in []byte) string
	FromBase64String(source string) ([]byte, error)
}

type Authorizator interface {
	SetCanWriteUsers() int32
	SetCanReadUsers() int32
	SetCanCreateUsers() int32
	SetCanRemoveUsers() int32

	SetCanWriteFiles() int32
	SetCanReadFiles() int32
	SetCanCreateFiles() int32
	SetCanRemoveFiles() int32

	SetCanWriteDatabases() int32
	SetCanReadDatabases() int32
	SetCanCreateDatabases() int32
	SetCanRemoveDatabases() int32

	SetCanWriteRecords() int32
	SetCanReadRecords() int32
	SetCanCreateRecords() int32
	SetCanRemoveRecords() int32
}

type ConfigHandler struct {
	b64  Base64
	v    *viper.Viper
	auth Authorizator
}

func NewConfigparser(au Authorizator, b64 Base64) *ConfigHandler {
	return &ConfigHandler{v: viper.New(), auth: au, b64: b64}
}

func (c *ConfigHandler) ParseConfig(path string) error {
	c.v.SetConfigType(ConfigFileFormat)
	c.v.SetConfigFile(path)
	c.v.AddConfigPath(path)
	return c.v.ReadInConfig()
}

func (c *ConfigHandler) ListenParams() (datatransferobjects.ListenParams, error) {
	var params datatransferobjects.ListenParams
	err := c.v.UnmarshalKey("listenParams", &params)
	return params, err
}

func (c *ConfigHandler) FileStorage() (datatransferobjects.FileStorage, error) {
	var params datatransferobjects.FileStorage
	err := c.v.UnmarshalKey("spec.fileStorage", &params)
	if err != nil {
		return params, err
	}
	params.Path, err = filepath.Abs(params.Path)
	return params, err
}

func (c *ConfigHandler) users() ([]User, error) {
	var users []User
	err := c.v.UnmarshalKey("spec.users", &users)
	return users, err
}

func (c *ConfigHandler) Databases(dbType string) ([]datatransferobjects.NetworkDataBaseDSN, []datatransferobjects.LocalDSN, error) {
	var databses RecordsDatabases
	err := c.v.UnmarshalKey("spec."+dbType, &databses)
	return databses.NetworkDatabases, databses.LocalDatabases, err
}

func (c *ConfigHandler) UsersDTO() ([]datatransferobjects.UserDTO, error) {
	var usersDTO []datatransferobjects.UserDTO = make([]datatransferobjects.UserDTO, 0)
	users, err := c.users()
	if err != nil {
		return nil, err
	}
	for i := range users {
		if users[i].Salted {
			pass, err := c.b64.FromBase64String(users[i].Password)
			if err != nil {
				return nil, err
			}
			users[i].Password = string(pass)
		}
		user, err := c.userToDTO(users[i])
		if err != nil {
			return nil, err
		}
		usersDTO = append(usersDTO, user)
		if users[i].AllowEmptyPassword {
			users[i].Password = ""
		} else {
			users[i].Password = c.b64.ToBase64String(user.Password)
			users[i].Salted = true
		}
	}
	c.v.Set("spec.users", users)
	return usersDTO, nil
}

func (c *ConfigHandler) RewriteConfig() error {
	return c.v.WriteConfig()
}

func (c *ConfigHandler) GenerateTemplate() error {
	_, err := os.Stdout.Write([]byte(Template))
	return err
}
