package requests

import (
	"context"
	"net/url"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
)

func (rh *Requesting) TagsGetRequest(ctx context.Context, token string) ([]datatransferobjects.TagDTO, error) {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return nil, err
	}
	err = rh.authorizator.CanReadRecords(&user)
	if err != nil {
		return nil, err
	}
	return rh.advRepo.AllTags(ctx)
}

func (rh *Requesting) TagGetRequest(ctx context.Context, token string, name string) (datatransferobjects.TagDTO, error) {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return datatransferobjects.TagDTO{}, err
	}
	err = rh.authorizator.CanReadRecords(&user)
	if err != nil {
		return datatransferobjects.TagDTO{}, err
	}

	parsedName, err := url.QueryUnescape(name)
	if err != nil {
		return datatransferobjects.TagDTO{}, err
	}
	return rh.advRepo.TagByName(ctx, parsedName)
}

func (rh *Requesting) TagPostRequest(ctx context.Context, token string, tag *datatransferobjects.TagDTO) error {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return err
	}
	err = rh.authorizator.CanCreateRecords(&user)
	if err != nil {
		return err
	}

	return rh.advRepo.NewTag(ctx, tag)
}

func (rh *Requesting) TagPutRequest(ctx context.Context, token string, tag *datatransferobjects.TagDTO) error {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return err
	}
	err = rh.authorizator.CanWriteRecords(&user)
	if err != nil {
		return err
	}

	return rh.advRepo.UpdateTag(ctx, tag)
}

func (rh *Requesting) TagDeleteRequest(ctx context.Context, token string, name string) error {
	user, err := rh.checkAndFetchToken(ctx, token)
	if err != nil {
		return err
	}
	err = rh.authorizator.CanDeleteRecords(&user)
	if err != nil {
		return err
	}

	return rh.advRepo.RemoveTagByName(ctx, name)
}
