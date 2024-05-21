package converter

import "time"

type ClientFront struct {
	Name                  string        `json:"name"`
	Phones                *string       `json:"phones"`
	Email                 *string       `json:"email"`
	AdditionalInformation *string       `json:"additionalInformation"`
	Orders                *[]OrderFront `json:"orders"`
}

type OrderFront struct {
	ID                  int                        `json:"id"`
	ClientName          *string                    `json:"clientName"`
	Cost                *int                       `json:"cost"`
	PaymentType         *string                    `json:"paymentType"`
	CreatedDate         *time.Time                 `json:"createdDate"`
	PaymentStatus       *bool                      `json:"paymentStatus"`
	LineAdvertisements  *[]LineAdvertisementFront  `json:"lineAdvertisements"`
	BlockAdvertisements *[]BlockAdvertisementFront `json:"blockAdvertisements"`
}

type AdvertisementFront struct {
	ID           int         `json:"id"`
	OrderID      int         `json:"orderID"`
	ReleaseCount int16       `json:"releaseCount"`
	Cost         int         `json:"cost"`
	Text         string      `json:"text"`
	Tags         []string    `json:"tags,omitempty"`
	ExtraCharges []string    `json:"extraCharges,omitempty"`
	ReleaseDates []time.Time `json:"releaseDates,omitempty"`
}

type BlockAdvertisementFront struct {
	AdvertisementFront
	Size     *int16  `json:"size"`
	FileName *string `json:"fileName"`
}

type LineAdvertisementFront struct {
	AdvertisementFront
}

type TagFront struct {
	Name string `json:"name"`
	Cost *int   `json:"cost"`
}

type ExtraChargeFront struct {
	ChargeName string `json:"chargeName"`
	Multiplier *int   `json:"multiplier"`
}

type CostRateFront struct {
	CalcForOneWord   *bool  `json:"calcForOneWord"`
	Name             string `json:"name"`
	ForOneWordSymbol *int   `json:"forOneSymbol"`
	ForOnecm2        *int   `json:"forOnecm2"`
}

type ResponceMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Token struct {
	AccessToken string `json:"accessToken"`
}

type UserFront struct {
	Name        string `json:"name"`
	Password    string `json:"password,omitempty"`
	Permissions int32  `json:"permissions,omitempty"`
}

type FileFront struct {
	Name string `json:"name"`
	Size int64  `json:"size,omitempty"`
	Data string `json:"data,omitempty"`
}

type DataBaseParams struct {
	Type             string `json:"type"`
	Name             string `json:"name"`
	ConnectionParams NetworkDataBaseDSN
}

type NetworkDataBaseDSN struct {
	Type     string `json:"type,omitempty"`
	Adress   string `json:"adress,omitempty"`
	DbName   string `json:"dbName,omitempty"`
	UserName string `json:"UserName"`
	Password bool   `json:"password,omitempty"`
	Port     uint   `json:"port,omitempty"`
}

type B64Enc interface {
	ToBase64([]byte) []byte
	ToBase64String([]byte) string
	FromBase64([]byte) ([]byte, error)
}
