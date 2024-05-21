package advertisementhandler

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
)

func TestNewClient(t *testing.T) {
	client := datatransferobjects.ClientDTO{
		Name:                  "Вася",
		Phones:                "8-800-555-35-35",
		Email:                 "mail@mail.com",
		AdditionalInformation: "asdasd",
	}

	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}
	err = db.NewClient(context.Background(), &client)
	if err != nil {
		t.Fatal(err)
	}
}
func TestConvertationError(t *testing.T) {
	client := datatransferobjects.ClientDTO{
		Name:                  "В",
		Phones:                "8-800-555-35-35",
		Email:                 "mail@mail.com",
		AdditionalInformation: "asdasd",
	}

	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	err = db.NewClient(context.Background(), &client)
	if !errors.Is(err, ErrConvertation) {
		t.Fatal(err.Error())
	}
}

func TestNewOrder(t *testing.T) {
	order := datatransferobjects.OrderDTO{
		ID:            0,
		ClientName:    "Вася",
		Cost:          10,
		PaymentType:   "asd",
		CreatedDate:   time.Now(),
		PaymentStatus: true,
	}

	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	err = db.NewAdvertisementsOrder(context.Background(), &order)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewTag(t *testing.T) {
	tagDTO := datatransferobjects.TagDTO{TagName: "tag A", TagCost: 12}

	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	err = db.NewTag(context.Background(), &tagDTO)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewExtraCharge(t *testing.T) {
	extraChargeDTO := datatransferobjects.ExtraChargeDTO{ChargeName: "charge A", Multiplier: 2}
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	err = db.NewExtraCharge(context.Background(), &extraChargeDTO)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewBlockAdvertisement(t *testing.T) {
	blockDtoForTest := datatransferobjects.BlockAdvertisementDTO{
		Advertisement: datatransferobjects.Advertisement{
			ID:           0,
			OrderID:      1,
			ReleaseCount: 1,
			Cost:         23,
			Text:         "asd",
			Tags:         []string{"tag A", "tag B", "tag C"},
			ExtraCharges: []string{"charge A", "charge B", "charge C"},
			ReleaseDates: []time.Time{time.Now().AddDate(0, 0, 1)},
		},
		Size: 10,
	}
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	err = db.NewBlockAdvertisement(context.Background(), &blockDtoForTest)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewLineAdvertisement(t *testing.T) {
	lineDtoForTest := datatransferobjects.LineAdvertisementDTO{
		Advertisement: datatransferobjects.Advertisement{
			ID:           0,
			OrderID:      1,
			ReleaseCount: 1,
			Cost:         23,
			Text:         "asd",
			Tags:         []string{"tag A", "tag B", "tag C"},
			ExtraCharges: []string{"charge A", "charge B", "charge C"},
			ReleaseDates: []time.Time{time.Now().AddDate(0, 0, 1)},
		},
	}
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	err = db.NewLineAdvertisement(context.Background(), &lineDtoForTest)
	if err != nil {
		t.Fatal(err)
	}
}
