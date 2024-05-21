package validator

type B64Enc interface {
	ToBase64URL([]byte) []byte
	FromBase64URL([]byte) ([]byte, error)
}

type Storage interface {
	Load(name string) ([]byte, error)
	SaveConfig(name string, data []byte) error
}
