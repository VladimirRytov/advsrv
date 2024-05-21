package handlers

type DBManager interface {
	AvailableLocalDatabases()
	AvailableNetworkDatabases()
	DefaultPort(string)
}
