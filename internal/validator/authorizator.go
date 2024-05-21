package validator

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/encodedecoder"
)

var ErrValidate = errValidate()

func errValidate() error {
	return errors.New("validation error")
}

type TokenWorker struct {
	header []byte
	key    *rsa.PrivateKey
	b64    B64Enc
}

func NewWorker(b64 B64Enc) (*TokenWorker, error) {
	var err error
	w := &TokenWorker{}
	w.header, err = encodePayload(&Header{Alg: "RS256", Typ: "JWT"})
	if err != nil {
		return nil, err
	}
	w.key, err = rsa.GenerateKey(rand.Reader, 2048)
	w.b64 = b64
	return w, err
}

func (tw *TokenWorker) publicKeyBytes() []byte {
	return x509.MarshalPKCS1PublicKey(&tw.key.PublicKey)
}

func (tw *TokenWorker) privateKeyBytes() []byte {
	return x509.MarshalPKCS1PrivateKey(tw.key)
}

func (tw *TokenWorker) SignDate(userName string, perm int32) ([]byte, error) {
	var b bytes.Buffer

	payload, err := encodePayload(&datatransferobjects.UserToken{Log: userName, Perm: perm})
	if err != nil {
		return nil, err
	}
	b.Grow(len(tw.header) + len(payload) + 2)
	b.Write(tw.header)
	b.WriteByte('.')
	b.Write(payload)

	s := sha256.Sum256(b.Bytes())
	rawSignature, err := rsa.SignPKCS1v15(nil, tw.key, crypto.SHA256, s[:])
	if err != nil {
		return nil, err
	}
	b.WriteByte('.')

	sig := tw.b64.ToBase64URL(rawSignature)
	if err != nil {
		return nil, err
	}
	b.Write(sig)

	return b.Bytes(), nil
}

func encodePayload[T Header | datatransferobjects.UserToken](payload *T) ([]byte, error) {
	rawpayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	enc := base64.RawURLEncoding.WithPadding(base64.NoPadding)

	return []byte(enc.EncodeToString(rawpayload)), nil
}

func decodePayload[T Header | datatransferobjects.UserToken](out *T, raw []byte, b64 B64Enc) error {
	decoded, err := b64.FromBase64URL(raw)
	if err != nil {
		return err
	}
	r := bytes.NewReader(decoded)
	err = encodedecoder.FromJSON(out, r)
	if err != nil {
		return err
	}
	return nil
}

func (tw *TokenWorker) Validate(token []byte) error {
	startSign := bytes.LastIndexByte(token, '.')
	if startSign < 0 {
		return errors.New("неверный формат")
	}
	endHeader := bytes.IndexByte(token, '.')
	if startSign < 0 || endHeader >= startSign {
		return errors.New("неверный формат")
	}

	var head Header
	err := decodePayload(&head, token[:endHeader], tw.b64)
	if err != nil {
		return err
	}

	sig, err := tw.b64.FromBase64URL(token[startSign+1:])
	if err != nil {
		return err
	}
	s := sha256.Sum256(token[:startSign])
	return rsa.VerifyPKCS1v15(&tw.key.PublicKey, crypto.SHA256, s[:], sig)
}

func (tw *TokenWorker) FetchPayload(token []byte) (datatransferobjects.UserToken, error) {
	splited := bytes.Split(token, []byte{'.'})
	if len(splited) != 3 {
		return datatransferobjects.UserToken{}, errors.New("неверный формат")
	}
	sig, err := tw.b64.FromBase64URL(splited[1])
	if err != nil {
		return datatransferobjects.UserToken{}, err
	}
	var userT datatransferobjects.UserToken
	err = json.Unmarshal(sig, &userT)
	if err != nil {
		return datatransferobjects.UserToken{}, err
	}
	return userT, err
}
