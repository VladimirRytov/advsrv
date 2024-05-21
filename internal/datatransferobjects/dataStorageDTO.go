package datatransferobjects

type LocalDSN struct {
	Path string `json:"path"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type NetworkDataBaseDSN struct {
	Type     string `json:"type"`
	Adress   string `json:"adress"`
	DbName   string `json:"dbName"`
	UserName string `json:"userName"`
	Password string `json:"password,omitempty"`
	Port     uint   `json:"port,omitempty"`
}
