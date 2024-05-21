package requests

import (
	"context"
	"io"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
)

func (rh *Requesting) FileGetRequest(ctx context.Context, token string, name string) (string, error) {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return "", err
	}
	err = rh.authorizator.CanReadRecords(&user)
	if err != nil {
		return "", err
	}
	return rh.files.Path(name)
}

func (rh *Requesting) FileGetFormatedRequest(ctx context.Context, token string, name, size string) (datatransferobjects.File, error) {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return datatransferobjects.File{}, err
	}
	err = rh.authorizator.CanReadRecords(&user)
	if err != nil {
		return datatransferobjects.File{}, err
	}
	switch size {
	case Miniature:
		return rh.files.GetMiniature(name, 64)
	case Small:
		return rh.files.GetMiniature(name, 128)
	case Large:
		return rh.files.GetMiniature(name, 450)
	}
	return rh.files.Get(name)
}

func (rh *Requesting) FilesGetRequest(ctx context.Context, token string, params map[string]string) ([]datatransferobjects.File, error) {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return nil, err
	}
	err = rh.authorizator.CanReadRecords(&user)
	if err != nil {
		return nil, err
	}
	fp, err := rh.CheckFileQueries(ctx, params)
	if err != nil {
		return nil, err
	}
	switch fp.Size {
	case Miniature:
		return rh.files.ListWithMiniatures(64)
	case Small:
		return rh.files.ListWithMiniatures(128)
	case Large:
		return rh.files.ListWithMiniatures(450)
	}
	if fp.Miniatures {
		return rh.files.ListWithMiniatures(64)
	}
	return rh.files.List()
}

func (rh *Requesting) FilePostRequest(ctx context.Context, token string, params map[string]string, name string, file io.ReadSeekCloser) (string, error) {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return "", err
	}
	err = rh.authorizator.CanCreateRecords(&user)
	if err != nil {
		return "", err
	}
	return rh.files.Set(name, file)
}

func (rh *Requesting) FilesPostRequest(ctx context.Context, token string, params map[string]string, name datatransferobjects.Files) ([]datatransferobjects.File, error) {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return nil, err
	}
	err = rh.authorizator.CanCreateRecords(&user)
	if err != nil {
		return nil, err
	}
	return rh.files.List()
}

func (rh *Requesting) FileDeleteRequest(ctx context.Context, token string, name string) error {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return err
	}
	err = rh.authorizator.CanDeleteRecords(&user)
	if err != nil {
		return err
	}
	return rh.files.Remove(name)
}
