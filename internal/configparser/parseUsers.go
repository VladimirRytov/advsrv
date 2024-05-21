package configparser

import (
	"errors"
	"strings"
)

func (c *ConfigHandler) parseUserPermission(permissions Permissions) (int32, error) {
	dbPermissions, err := c.parseDatabasePermission(permissions.Databases)
	if err != nil {
		return 0, err
	}
	usersPermissions, err := c.parseUsersPermission(permissions.Users)
	if err != nil {
		return 0, err
	}
	recordsPermissions, err := c.parseRecordsPermission(permissions.Records)
	if err != nil {
		return 0, err
	}
	filesPermissions, err := c.parseFilesPermission(permissions.Files)
	if err != nil {
		return 0, err
	}
	return dbPermissions | usersPermissions | recordsPermissions | filesPermissions, nil
}

func (c *ConfigHandler) parseDatabasePermission(perm []string) (int32, error) {
	var totalPerms int32
	for _, v := range perm {
		switch strings.ToLower(v) {
		case "all":
			totalPerms = c.auth.SetCanCreateDatabases() | c.auth.SetCanReadDatabases() | c.auth.SetCanWriteDatabases() | c.auth.SetCanRemoveDatabases()
		case "none":
			totalPerms = 0
		case "create":
			totalPerms |= c.auth.SetCanCreateDatabases()
		case "read":
			totalPerms |= c.auth.SetCanReadDatabases()
		case "write":
			totalPerms |= c.auth.SetCanWriteDatabases()
		case "remove":
			totalPerms |= c.auth.SetCanRemoveDatabases()
		default:
			return 0, errors.New("unknown permission " + v)
		}
	}
	return totalPerms, nil
}

func (c *ConfigHandler) parseUsersPermission(perm []string) (int32, error) {
	var totalPerms int32
	for _, v := range perm {
		switch strings.ToLower(v) {
		case "all":
			totalPerms = c.auth.SetCanCreateUsers() | c.auth.SetCanReadUsers() | c.auth.SetCanWriteUsers() | c.auth.SetCanRemoveUsers()
		case "none":
			totalPerms = 0
		case "create":
			totalPerms |= c.auth.SetCanCreateUsers()
		case "read":
			totalPerms |= c.auth.SetCanReadUsers()
		case "write":
			totalPerms |= c.auth.SetCanWriteUsers()
		case "remove":
			totalPerms |= c.auth.SetCanRemoveUsers()
		default:
			return 0, errors.New("unknown permission " + v)
		}
	}
	return totalPerms, nil
}

func (c *ConfigHandler) parseFilesPermission(perm []string) (int32, error) {
	var totalPerms int32
	for _, v := range perm {
		switch strings.ToLower(v) {
		case "all":
			totalPerms = c.auth.SetCanCreateFiles() | c.auth.SetCanReadFiles() | c.auth.SetCanWriteFiles() | c.auth.SetCanRemoveFiles()
		case "none":
			totalPerms = 0
		case "create":
			totalPerms |= c.auth.SetCanCreateFiles()
		case "read":
			totalPerms |= c.auth.SetCanReadFiles()
		case "write":
			totalPerms |= c.auth.SetCanWriteFiles()
		case "remove":
			totalPerms |= c.auth.SetCanRemoveFiles()
		default:
			return 0, errors.New("unknown permission " + v)
		}
	}
	return totalPerms, nil
}

func (c *ConfigHandler) parseRecordsPermission(perm []string) (int32, error) {
	var totalPerms int32
	for _, v := range perm {
		switch strings.ToLower(v) {
		case "all":
			totalPerms = c.auth.SetCanCreateRecords() | c.auth.SetCanReadRecords() | c.auth.SetCanWriteRecords() | c.auth.SetCanRemoveRecords()
		case "none":
			totalPerms = 0
		case "create":
			totalPerms |= c.auth.SetCanCreateRecords()
		case "read":
			totalPerms |= c.auth.SetCanReadRecords()
		case "write":
			totalPerms |= c.auth.SetCanWriteRecords()
		case "remove":
			totalPerms |= c.auth.SetCanRemoveRecords()
		default:
			return 0, errors.New("unknown permission " + v)
		}
	}
	return totalPerms, nil
}
