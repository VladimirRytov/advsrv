package encryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"io"

	"github.com/VladimirRytov/advsrv/internal/systeminfo"
)

type AesEncryptor struct {
	iv  [aes.BlockSize]byte
	key cipher.Block
}

func AesInit() (*AesEncryptor, error) {
	Aes := new(AesEncryptor)
	cpu, err := systeminfo.CPU()
	if err != nil {
		return nil, err
	}
	sha := sha256.Sum256([]byte(cpu + systeminfo.ComputerName()))

	Aes.key, err = aes.NewCipher(sha[:])
	if err != nil {
		return nil, err
	}
	return Aes, nil
}

func (r *AesEncryptor) Decrypt(out io.Writer, in io.Reader) error {
	writer := &cipher.StreamWriter{S: cipher.NewOFB(r.key, r.iv[:]), W: out}
	_, err := io.Copy(writer, in)

	if err != nil {
		return err
	}
	return nil
}

func (r *AesEncryptor) Encrypt(out io.Writer, in io.Reader) error {
	reader := &cipher.StreamReader{S: cipher.NewOFB(r.key, r.iv[:]), R: in}
	_, err := io.Copy(out, reader)
	if err != nil {
		return err
	}
	return nil
}
