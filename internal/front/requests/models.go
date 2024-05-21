package requests

import (
	"context"
	"io"
	"time"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
)

type UserHandler interface {
	Authenticate(ctx context.Context, user *datatransferobjects.UserDTO) ([]byte, error)
	NewUser(context.Context, *datatransferobjects.UserDTO) (datatransferobjects.UserDTO, error)
	UserList(ctx context.Context) ([]datatransferobjects.UserDTO, error)
	UserByName(context.Context, string) (datatransferobjects.UserDTO, error)
	UpdateUser(context.Context, *datatransferobjects.UserDTO) (datatransferobjects.UserDTO, error)
	RemoveUser(context.Context, string) error
}

type BroadCaster interface {
	NewPassiveSub(addr, name string) (string, error)
	NewActiveSub(ctx context.Context, addr string) (<-chan []byte, error)
	RemoveSub(addr string)
	SubInfo(id string) (datatransferobjects.SubscribeParams, error)
	ListSubs() (int, []datatransferobjects.SubscribeParams)
	Ping(string) error
	ErrSubscriberExist() error
	ErrSubscriberNotExist() error
}

type FileHandler interface {
	Get(name string) (datatransferobjects.File, error)
	GetMiniature(name string, size int) (datatransferobjects.File, error)
	Set(name string, data io.ReadSeekCloser) (string, error)
	List() ([]datatransferobjects.File, error)
	ListWithMiniatures(int) ([]datatransferobjects.File, error)
	Path(name string) (string, error)
	Remove(file string) error
}

type B64Enc interface {
	ToBase64([]byte) []byte
	FromBase64([]byte) ([]byte, error)
}
type CostCalculator interface {
	CalculateOrderCost(ctx context.Context, adv datatransferobjects.OrderDTO, costRateName string) (datatransferobjects.OrderDTO, error)

	CalculateBlockAdvertisementCost(ctx context.Context, adv datatransferobjects.BlockAdvertisementDTO,
		costRateName string) (datatransferobjects.BlockAdvertisementDTO, error)

	CalculateLineAdvertisementCost(ctx context.Context, adv datatransferobjects.LineAdvertisementDTO,
		costRateName string) (datatransferobjects.LineAdvertisementDTO, error)
}

type Validator interface {
	Validate(token []byte) error
	FetchPayload(token []byte) (datatransferobjects.UserToken, error)
}

type Permissions interface {
	UsersPermission
	DatabasePermission
	RecordsPermission
}

type DatabasePermission interface {
	CanReadDatabase(user *datatransferobjects.UserToken) error
	CanWriteDatabase(user *datatransferobjects.UserToken) error
	CanCreateDatabase(user *datatransferobjects.UserToken) error
	CanDeleteDatabase(user *datatransferobjects.UserToken) error
}

type RecordsPermission interface {
	CanReadRecords(user *datatransferobjects.UserToken) error
	CanWriteRecords(user *datatransferobjects.UserToken) error
	CanCreateRecords(user *datatransferobjects.UserToken) error
	CanDeleteRecords(user *datatransferobjects.UserToken) error
}

type UsersPermission interface {
	CanReadUsers(user *datatransferobjects.UserToken) error
	CanWriteUsers(user *datatransferobjects.UserToken) error
	CanCreateUsers(user *datatransferobjects.UserToken) error
	CanDeleteUsers(user *datatransferobjects.UserToken) error
}

type Requests interface {
	Creator
	Getter
	Searcher
	Remover
	Updater
	Closer
}

type Creator interface {
	NewClient(context.Context, *datatransferobjects.ClientDTO) error
	NewAdvertisementsOrder(context.Context, *datatransferobjects.OrderDTO) error
	NewLineAdvertisement(context.Context, *datatransferobjects.LineAdvertisementDTO) error
	NewBlockAdvertisement(context.Context, *datatransferobjects.BlockAdvertisementDTO) error
	NewExtraCharge(context.Context, *datatransferobjects.ExtraChargeDTO) error
	NewTag(context.Context, *datatransferobjects.TagDTO) error
	NewCostRate(context.Context, *datatransferobjects.CostRateDTO) error
}

type Getter interface {
	ClientByName(context.Context, string) (datatransferobjects.ClientDTO, error)
	OrderByID(context.Context, int) (datatransferobjects.OrderDTO, error)
	LineAdvertisementByID(context.Context, int) (datatransferobjects.LineAdvertisementDTO, error)
	BlockAdvertisementByID(context.Context, int) (datatransferobjects.BlockAdvertisementDTO, error)
	TagByName(context.Context, string) (datatransferobjects.TagDTO, error)
	ExtraChargeByName(context.Context, string) (datatransferobjects.ExtraChargeDTO, error)
	CostRateByName(context.Context, string) (datatransferobjects.CostRateDTO, error)
}

type Searcher interface {
	OrdersByClientName(context.Context, string) ([]datatransferobjects.OrderDTO, error)

	BlockAdvertisementsByOrderID(context.Context, int) ([]datatransferobjects.BlockAdvertisementDTO, error)
	LineAdvertisementsByOrderID(context.Context, int) ([]datatransferobjects.LineAdvertisementDTO, error)

	BlockAdvertisementsBetweenReleaseDates(context.Context, time.Time, time.Time) ([]datatransferobjects.BlockAdvertisementDTO, error)
	LineAdvertisementsBetweenReleaseDates(context.Context, time.Time, time.Time) ([]datatransferobjects.LineAdvertisementDTO, error)

	BlockAdvertisementsActualReleaseDate(context.Context) ([]datatransferobjects.BlockAdvertisementDTO, error)
	LineAdvertisementsActualReleaseDate(context.Context) ([]datatransferobjects.LineAdvertisementDTO, error)

	BlockAdvertisementsFromReleaseDate(context.Context, time.Time) ([]datatransferobjects.BlockAdvertisementDTO, error)
	LineAdvertisementsFromReleaseDate(context.Context, time.Time) ([]datatransferobjects.LineAdvertisementDTO, error)

	AllTags(context.Context) ([]datatransferobjects.TagDTO, error)
	AllExtraCharges(context.Context) ([]datatransferobjects.ExtraChargeDTO, error)
	AllClients(context.Context) ([]datatransferobjects.ClientDTO, error)
	AllOrders(context.Context) ([]datatransferobjects.OrderDTO, error)
	AllLineAdvertisements(context.Context) ([]datatransferobjects.LineAdvertisementDTO, error)
	AllBlockAdvertisements(context.Context) ([]datatransferobjects.BlockAdvertisementDTO, error)
	AllCostRates(context.Context) ([]datatransferobjects.CostRateDTO, error)
}

type Remover interface {
	RemoveClientByName(context.Context, string) error
	RemoveOrderByID(context.Context, int) error
	RemoveLineAdvertisementByID(context.Context, int) error
	RemoveBlockAdvertisementByID(context.Context, int) error
	RemoveTagByName(context.Context, string) error
	RemoveExtraChargeByName(context.Context, string) error
	RemoveCostRateByName(context.Context, string) error
}

type Updater interface {
	UpdateClient(context.Context, *datatransferobjects.ClientDTO) error
	UpdateOrder(context.Context, *datatransferobjects.OrderDTO) error
	UpdateLineAdvertisement(context.Context, *datatransferobjects.LineAdvertisementDTO) error
	UpdateBlockAdvertisement(context.Context, *datatransferobjects.BlockAdvertisementDTO) error
	UpdateExtraCharge(context.Context, *datatransferobjects.ExtraChargeDTO) error
	UpdateTag(context.Context, *datatransferobjects.TagDTO) error
	UpdateCostRate(context.Context, *datatransferobjects.CostRateDTO) error
}
type Closer interface {
	Close() error
}
