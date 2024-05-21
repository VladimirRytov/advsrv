package datatransferobjects

import (
	"time"
)

type JsonStr struct {
	Clients      []ClientDTO      `json:"clients"`
	Tags         []TagDTO         `json:"tags"`
	ExtraCharges []ExtraChargeDTO `json:"extraCharges"`
	CostRate     []CostRateDTO    `json:"costRates"`
}

type ClientDTO struct {
	Name                  string     `json:"name"`
	Phones                string     `json:"phones"`
	Email                 string     `json:"email"`
	AdditionalInformation string     `json:"additionalInformation"`
	Orders                []OrderDTO `json:"orders,omitempty"`
}

type OrderDTO struct {
	ID                  int                     `json:"id"`
	ClientName          string                  `json:"clientName"`
	Cost                int                     `json:"cost"`
	PaymentType         string                  `json:"paymentType"`
	CreatedDate         time.Time               `json:"createdDate"`
	PaymentStatus       bool                    `json:"paymentStatus"`
	LineAdvertisements  []LineAdvertisementDTO  `json:"lineAdvertisements,omitempty"`
	BlockAdvertisements []BlockAdvertisementDTO `json:"blockAdvertisements,omitempty"`
}

type Advertisement struct {
	ID           int         `json:"id"`
	OrderID      int         `json:"orderID"`
	ReleaseCount int16       `json:"releaseCount"`
	Cost         int         `json:"cost"`
	Text         string      `json:"text"`
	Tags         []string    `json:"tags,omitempty"`
	ExtraCharges []string    `json:"extraCharges,omitempty"`
	ReleaseDates []time.Time `json:"releaseDates,omitempty"`
}

type BlockAdvertisementDTO struct {
	Advertisement
	Size     int16  `json:"size"`
	FileName string `json:"fileName"`
}

type LineAdvertisementDTO struct {
	Advertisement
}

type TagDTO struct {
	TagName string `json:"name"`
	TagCost int    `json:"cost"`
}

type ExtraChargeDTO struct {
	ChargeName string `json:"name"`
	Multiplier int    `json:"multiplier"`
}

type CostRateDTO struct {
	CalcForOneWord   bool   `json:"calcForOneWord"`
	Name             string `json:"name"`
	ForOneWordSymbol int    `json:"forOneWordSymbol"`
	ForOnecm2        int    `json:"forOnecm2"`
}

type UserDTO struct {
	Name        string `json:"name"`
	Password    []byte `json:"password,omitempty"`
	Permissions int32  `json:"permissions,omitempty"`
}
