package front

import (
	"context"
	"io"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/front/converter"
)

type Requests interface {
	ClientRequests
	OrderRequests
	AdvertisementQueries
	BlockAdvertisementRequests
	LineAdvertisementRequests
	TagRequests
	ExtraChargeRequests
	CostRateRequests
	UserRequests
	SubScribeRequests
	FileRequests
}

type ClientRequests interface {
	CheckClientQueries(ctx context.Context, params map[string]string) (datatransferobjects.ClientParams, error)

	ClientsGetRequest(ctx context.Context, authHeader string, params datatransferobjects.ClientParams) ([]datatransferobjects.ClientDTO, error)
	ClientGetRequest(ctx context.Context, authHeader string, params datatransferobjects.ClientParams, name string) (datatransferobjects.ClientDTO, error)
	ClientPostRequest(ctx context.Context, authHeader string, params datatransferobjects.ClientParams, client *datatransferobjects.ClientDTO) error
	ClientPutRequest(ctx context.Context, authHeader string, client *datatransferobjects.ClientDTO) error
	ClientDeleteRequest(ctx context.Context, authHeader string, name string) error
}

type OrderRequests interface {
	CheckGetSeveralOrdersQueries(ctx context.Context, params map[string]string) (datatransferobjects.OrderParams, error)
	CheckOrderQueries(ctx context.Context, params map[string]string) (datatransferobjects.OrderParams, error)
	CalculateOrderCost(ctx context.Context, token string, params datatransferobjects.OrderParams,
		order *datatransferobjects.OrderDTO) (datatransferobjects.OrderDTO, error)
	OrderGetRequest(ctx context.Context, authHeader string, params datatransferobjects.OrderParams, id int) (datatransferobjects.OrderDTO, error)
	OrdersGetRequest(ctx context.Context, authHeader string, params datatransferobjects.OrderParams) ([]datatransferobjects.OrderDTO, error)
	OrderPostRequest(ctx context.Context, authHeader string, params datatransferobjects.OrderParams, order *datatransferobjects.OrderDTO) error
	OrderPutRequest(ctx context.Context, authHeader string, order *datatransferobjects.OrderDTO) error
	OrderDeleteRequest(ctx context.Context, authHeader string, id int) error
}

type AdvertisementQueries interface {
	CheckAdvertisementQueries(ctx context.Context, params map[string]string) (datatransferobjects.AdvertisementParams, error)
}

type BlockAdvertisementRequests interface {
	CalculateBlockAdvertisementCost(ctx context.Context, token string, params datatransferobjects.AdvertisementParams,
		blockAdv *datatransferobjects.BlockAdvertisementDTO) (datatransferobjects.BlockAdvertisementDTO, error)
	BlockAdvertisementsGetRequest(ctx context.Context, authHeader string, params datatransferobjects.AdvertisementParams) ([]datatransferobjects.BlockAdvertisementDTO, error)
	BlockAdvertisementGetRequest(ctx context.Context, authHeader string, id int) (datatransferobjects.BlockAdvertisementDTO, error)
	BlockAdvertisementPostRequest(ctx context.Context, authHeader string, BlockAdvertisement *datatransferobjects.BlockAdvertisementDTO) error
	BlockAdvertisementPutRequest(ctx context.Context, authHeader string, BlockAdvertisement *datatransferobjects.BlockAdvertisementDTO) error
	BlockAdvertisementDeleteRequest(ctx context.Context, authHeader string, id int) error
}

type LineAdvertisementRequests interface {
	CalculateLineAdvertisementCost(ctx context.Context, token string, params datatransferobjects.AdvertisementParams,
		lineAdv *datatransferobjects.LineAdvertisementDTO) (datatransferobjects.LineAdvertisementDTO, error)
	LineAdvertisementsGetRequest(ctx context.Context, authHeader string, params datatransferobjects.AdvertisementParams) ([]datatransferobjects.LineAdvertisementDTO, error)
	LineAdvertisementGetRequest(ctx context.Context, authHeader string, id int) (datatransferobjects.LineAdvertisementDTO, error)
	LineAdvertisementPostRequest(ctx context.Context, authHeader string, LineAdvertisement *datatransferobjects.LineAdvertisementDTO) error
	LineAdvertisementPutRequest(ctx context.Context, authHeader string, LineAdvertisement *datatransferobjects.LineAdvertisementDTO) error
	LineAdvertisementDeleteRequest(ctx context.Context, authHeader string, id int) error
}

type TagRequests interface {
	TagsGetRequest(ctx context.Context, authHeader string) ([]datatransferobjects.TagDTO, error)
	TagGetRequest(ctx context.Context, authHeader string, name string) (datatransferobjects.TagDTO, error)
	TagPostRequest(ctx context.Context, authHeader string, tag *datatransferobjects.TagDTO) error
	TagPutRequest(ctx context.Context, authHeader string, tag *datatransferobjects.TagDTO) error
	TagDeleteRequest(ctx context.Context, authHeader string, name string) error
}

type ExtraChargeRequests interface {
	ExtraChargesGetRequest(ctx context.Context, authHeader string) ([]datatransferobjects.ExtraChargeDTO, error)
	ExtraChargeGetRequest(ctx context.Context, authHeader string, name string) (datatransferobjects.ExtraChargeDTO, error)
	ExtraChargePostRequest(ctx context.Context, authHeader string, ExtraCharge *datatransferobjects.ExtraChargeDTO) error
	ExtraChargePutRequest(ctx context.Context, authHeader string, ExtraCharge *datatransferobjects.ExtraChargeDTO) error
	ExtraChargeDeleteRequest(ctx context.Context, authHeader string, name string) error
}

type CostRateRequests interface {
	CostRatesGetRequest(ctx context.Context, authHeader string) ([]datatransferobjects.CostRateDTO, error)
	CostRateGetRequest(ctx context.Context, authHeader string, name string) (datatransferobjects.CostRateDTO, error)
	CostRatePostRequest(ctx context.Context, authHeader string, costRate *datatransferobjects.CostRateDTO) error
	CostRatePutRequest(ctx context.Context, authHeader string, costRate *datatransferobjects.CostRateDTO) error
	CostRateDeleteRequest(ctx context.Context, authHeader string, name string) error
}

type UserRequests interface {
	UsersGetRequest(ctx context.Context, authHeader string) ([]datatransferobjects.UserDTO, error)
	UserGetRequest(ctx context.Context, authHeader string, name string) (datatransferobjects.UserDTO, error)
	UserPostRequest(ctx context.Context, authHeader string, user *datatransferobjects.UserDTO) (datatransferobjects.UserDTO, error)
	UserPutRequest(ctx context.Context, authHeader string, user *datatransferobjects.UserDTO) (datatransferobjects.UserDTO, error)
	UserDeleteRequest(ctx context.Context, authHeader string, name string) error
	AuthenticateUser(ctx context.Context, header string) ([]byte, error)
	CheckToken(ctx context.Context, head string) ([]byte, error)
}

type SubScribeRequests interface {
	SubscribeGetRequest(ctx context.Context, headers string, id string, ping bool) (datatransferobjects.SubscribeParams, error)
	SubscribersGetRequest(ctx context.Context, headers string) ([]datatransferobjects.SubscribeParams, error)

	SubscribeActiveRequest(ctx context.Context, headers string, addr string) (<-chan []byte, error)

	SubscribePostRequest(context.Context, string, datatransferobjects.SubscribeParams) (string, error)
	SubscribeDeleteRequest(context.Context, string, string) error
}

type FileRequests interface {
	FileGetRequest(ctx context.Context, headers string, name string) (string, error)
	FileGetFormatedRequest(ctx context.Context, headers string, name, size string) (datatransferobjects.File, error)
	FilesGetRequest(ctx context.Context, headers string, params map[string]string) ([]datatransferobjects.File, error)
	FilePostRequest(ctx context.Context, headers string, params map[string]string, name string, f io.ReadSeekCloser) (string, error)
	FilesPostRequest(ctx context.Context, headers string, params map[string]string, name datatransferobjects.Files) ([]datatransferobjects.File, error)
	FileDeleteRequest(ctx context.Context, headers string, name string) error
}

type FrontConverter interface {
	NewResponceMessage(int, string) converter.ResponceMessage
	FileToFront(file *datatransferobjects.File) converter.FileFront
}

type Base64EncodeDecoder interface {
	FromBase64URLString(source string) ([]byte, error)
}
