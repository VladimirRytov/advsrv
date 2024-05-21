package datatransferobjects

type Message struct {
	Code    int
	Message string
}

type ListenParams struct {
	Adress string
	Port   uint
	Tls    *Tls
}

type Tls struct {
	PathToCertificate string `yaml:"pathToCertificate"`
	PathToKey         string `yaml:"pathToKey"`
}

type FileStorage struct {
	Path string
}
