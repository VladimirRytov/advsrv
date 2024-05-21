package requests

import (
	"bytes"
	"context"
	"strings"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
)

const (
	Base int = iota + 1
	Bearer
	Token

	Miniature = "miniature"
	Small     = "small"
	Large     = "large"
)

func (rh *Requesting) AuthenticateUser(ctx context.Context, rawHeader string) ([]byte, error) {
	authType, authParams := rh.handleAuthToken(rawHeader)
	if authType != Base {
		return nil, ErrNeedBasicMethod
	}
	user, err := rh.b64.FromBase64([]byte(authParams))
	if err != nil {
		return nil, err
	}
	splited := bytes.Split(user, []byte{':'})
	if len(splited) != 2 {
		return nil, ErrQuery
	}
	userDto := datatransferobjects.UserDTO{
		Name:     string(splited[0]),
		Password: splited[1],
	}
	return rh.userRepo.Authenticate(ctx, &userDto)
}

func (rh *Requesting) CheckToken(ctx context.Context, head string) ([]byte, error) {
	authType, token := rh.handleAuthToken(head)
	if authType != Bearer && authType != Token {
		return nil, ErrNeedBearerMethod
	}
	err := rh.validator.Validate([]byte(token))
	return []byte(token), err
}

func (rh *Requesting) handleAuthToken(rawHead string) (int, string) {
	splited := strings.Split(rawHead, " ")
	if len(splited) != 2 {
		return Token, rawHead
	}
	switch splited[0] {
	case "Bearer":
		return Bearer, splited[1]
	case "Basic":
		return Base, splited[1]
	default:
		return 0, rawHead
	}
}

func (rh *Requesting) checkAndFetchToken(ctx context.Context, rawHead string) (datatransferobjects.UserToken, error) {
	token, err := rh.CheckToken(ctx, rawHead)
	if err != nil {
		return datatransferobjects.UserToken{}, ErrValidate
	}
	return rh.validator.FetchPayload(token)
}
