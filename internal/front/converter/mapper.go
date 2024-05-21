package converter

import (
	"bytes"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/encodedecoder"
)

type FrontConverter struct {
	b64 B64Enc
}

func NewFrontConverter(b64 B64Enc) *FrontConverter {
	return &FrontConverter{
		b64: b64,
	}
}

// func ToClient(c *datatransferobjects.ClientDTO) Client {
// 	var orders []Order
// 	client := Client{
// 		Name:                  c.Name,
// 		Phones:                c.Phones,
// 		Email:                 c.Email,
// 		AdditionalInformation: &c.AdditionalInformation,
// 	}

// 	if len(c.Orders) > 0 {
// 		orders = make([]Order, 0, len(c.Orders))
// 		for i := range c.Orders {
// 			orders = append(orders, ToOrder(&c.Orders[i]))
// 		}
// 		client.Orders = &orders
// 	}

// 	return client
// }

// func ToOrder(o *datatransferobjects.OrderDTO) Order {
// 	var (
// 		lines  []LineAdvertisement
// 		blocks []BlockAdvertisement
// 	)
// 	order := Order{
// 		ID:            o.ID,
// 		ClientName:    o.ClientName,
// 		Cost:          o.Cost,
// 		PaymentType:   &o.PaymentType,
// 		CreatedDate:   o.CreatedDate,
// 		PaymentStatus: o.PaymentStatus,
// 	}

// 	if len(o.LineAdvertisements) > 0 {
// 		lines = make([]LineAdvertisement, 0, len(o.LineAdvertisements))
// 		for i := range o.LineAdvertisements {
// 			lines = append(lines, ToLineAdvertisement(&o.LineAdvertisements[i]))
// 		}
// 		order.LineAdvertisements = &lines
// 	}

// 	if len(o.BlockAdvertisements) > 0 {
// 		blocks = make([]BlockAdvertisement, 0, len(o.BlockAdvertisements))
// 		for i := range o.LineAdvertisements {
// 			blocks = append(blocks, ToBlockAdvertisement(&o.BlockAdvertisements[i]))
// 		}
// 		order.BlockAdvertisements = &blocks
// 	}
// 	return order
// }

// func ToLineAdvertisement(l *datatransferobjects.LineAdvertisementDTO) LineAdvertisement {
// 	return LineAdvertisement{
// 		Advertisement: Advertisement{
// 			ID:           l.ID,
// 			OrderID:      l.OrderID,
// 			ReleaseCount: l.ReleaseCount,
// 			Cost:         l.Cost,
// 			Text:         &l.Text,
// 			Tags:         &l.Tags,
// 			ExtraCharges: &l.ExtraCharges,
// 			ReleaseDates: &l.ReleaseDates,
// 		},
// 	}
// }

// func ToBlockAdvertisement(b *datatransferobjects.BlockAdvertisementDTO) BlockAdvertisementFront {
// 	return BlockAdvertisementFront{
// 		AdvertisementFront: AdvertisementFront{
// 			ID:           b.ID,
// 			OrderID:      b.OrderID,
// 			ReleaseCount: b.ReleaseCount,
// 			Cost:         b.Cost,
// 			Text:         &b.Text,
// 			Tags:         &b.Tags,
// 			ExtraCharges: &b.ExtraCharges,
// 			ReleaseDates: &b.ReleaseDates,
// 		},
// 		Size:     b.Size,
// 		FileName: b.FileName,
// 	}
// }

func (fc *FrontConverter) NewResponceMessage(code int, message string) ResponceMessage {
	return ResponceMessage{
		Code:    code,
		Message: message,
	}
}

func (fc *FrontConverter) UserToFront(user *datatransferobjects.UserDTO) UserFront {
	perm := user.Permissions
	return UserFront{
		Name:        user.Name,
		Password:    "",
		Permissions: perm,
	}
}

func (fc *FrontConverter) UserFrontToDTO(user []byte) (datatransferobjects.UserDTO, error) {
	var userDTO datatransferobjects.UserDTO
	r := bytes.NewReader(user)
	err := encodedecoder.FromJSON(&userDTO, r)
	return userDTO, err
}

func (fc *FrontConverter) FileToFront(file *datatransferobjects.File) FileFront {
	var fileF FileFront
	fileF.Name = file.Name

	return FileFront{
		Name: file.Name,
		Size: file.Size,
		Data: fc.b64.ToBase64String(file.Data),
	}
}

func (fc *FrontConverter) DataBaseParamsToFront(name string, params []byte) ([]byte, error) {
	var dbParams NetworkDataBaseDSN
	r := bytes.NewReader(params)
	err := encodedecoder.FromJSON(&dbParams, r)
	if err != nil {
		return nil, err
	}
	p := DataBaseParams{
		Name:             name,
		ConnectionParams: dbParams,
	}
	var b bytes.Buffer
	err = encodedecoder.ToJSON(&b, p, false)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), err
}
