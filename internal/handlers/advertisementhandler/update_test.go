package advertisementhandler

import (
	"context"
	"testing"
	"time"

	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
)

func TestUpdateClient(t *testing.T) {
	newClient := datatransferobjects.ClientDTO{
		Name:                  "Вася",
		Phones:                "8-800-555-35-35",
		Email:                 "mail@mail.com",
		AdditionalInformation: "zzxczx",
	}

	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateClient(context.Background(), &newClient)
	if err != nil {
		t.Fatal(err)
	}

}

func TestUpdateOrder(t *testing.T) {
	newOrder := datatransferobjects.OrderDTO{
		ID:            1,
		ClientName:    "Вася",
		Cost:          10,
		PaymentType:   "иит",
		CreatedDate:   time.Now(),
		PaymentStatus: false,
	}

	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateOrder(context.Background(), &newOrder)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateBlockAdvertisement(t *testing.T) {
	newBlockDtoForTest := datatransferobjects.BlockAdvertisementDTO{
		Advertisement: datatransferobjects.Advertisement{
			ID:           1,
			OrderID:      1,
			ReleaseCount: 1,
			Cost:         23,
			Text:         "asd",
			Tags:         []string{"tag A", "tag B", "tag C"},
			ExtraCharges: []string{"charge A", "charge B", "charge C"},
			ReleaseDates: []time.Time{time.Now()},
		},
		Size: 210,
	}

	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}
	err = db.UpdateBlockAdvertisement(context.Background(), &newBlockDtoForTest)
	if err != nil {
		t.Fatal(err)
	}

}

func TestUpdateLienAdvertisement(t *testing.T) {
	newLineDtoForTest := datatransferobjects.LineAdvertisementDTO{
		Advertisement: datatransferobjects.Advertisement{
			ID:           1,
			OrderID:      1,
			ReleaseCount: 1,
			Cost:         55,
			Text:         "cxzx",
			Tags:         []string{"tag A", "tag B", "tag C"},
			ExtraCharges: []string{"charge A", "charge B", "charge C"},
			ReleaseDates: []time.Time{time.Now()},
		},
	}
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateLineAdvertisement(context.Background(), &newLineDtoForTest)
	if err != nil {
		t.Fatal(err)
	}

}

func TestUpdateTag(t *testing.T) {
	newTagDTO := datatransferobjects.TagDTO{TagName: "tag A", TagCost: 22}

	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateTag(context.Background(), &newTagDTO)
	if err != nil {
		t.Fatal(err)
	}

}

func TestUpdateExtraCharge(t *testing.T) {
	newExtraChargeDTO := datatransferobjects.ExtraChargeDTO{ChargeName: "charge A", Multiplier: 4}
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateExtraCharge(context.Background(), &newExtraChargeDTO)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateCostRate(t *testing.T) {
	newCostRateDTO := datatransferobjects.CostRateDTO{CalcForOneWord: true, Name: "стандарт", ForOneWordSymbol: 1, ForOnecm2: 2}
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateCostRate(context.Background(), &newCostRateDTO)
	if err != nil {
		t.Fatal(err)
	}
}
