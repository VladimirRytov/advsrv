package advertisementhandler

import (
	"context"
	"testing"
)

func TestGetClient(t *testing.T) {
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}
	clientName := "Вася"
	client, err := db.ClientByName(context.Background(), clientName)
	if err != nil {
		t.Fatal(err)
	}
	if client.Name != clientName {
		t.Fatalf("want Client name %s, got %s", clientName, client.Name)
	}
}

func TestGetOrder(t *testing.T) {
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	orderId := 1
	order, err := db.OrderByID(context.Background(), orderId)
	if err != nil {
		t.Fatal(err)
	}

	if order.ID != orderId {
		t.Fatalf("want OrderID %d, got %d", orderId, order.ID)
	}
}

func TestGetBlockAdvertisement(t *testing.T) {
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	blockAdvID := 1
	blockAdv, err := db.BlockAdvertisementByID(context.Background(), blockAdvID)
	if err != nil {
		t.Fatal(err)
	}

	if blockAdvID != blockAdv.ID {
		t.Fatalf("want BlockAdvertisementID %d, got %d", blockAdvID, blockAdv.ID)
	}
}

func TestGetLineAdvertisement(t *testing.T) {
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	lineAdvID := 1
	lineAdv, err := db.LineAdvertisementByID(context.Background(), lineAdvID)
	if err != nil {
		t.Fatal(err)
	}

	if lineAdvID != lineAdv.ID {
		t.Fatalf("want LineAdvertisementID %d, got %d", lineAdvID, lineAdv.ID)
	}
}

func TestGetTag(t *testing.T) {
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}
	tagName := "tag A"
	tag, err := db.TagByName(context.Background(), tagName)
	if err != nil {
		t.Fatal(err)
	}

	if tagName != tag.TagName {
		t.Fatalf("want tag name = %s, got %s", tagName, tag.TagName)
	}
}

func TestGetExtraCharge(t *testing.T) {
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}
	chargeName := "charge A"
	charge, err := db.ExtraChargeByName(context.Background(), chargeName)
	if err != nil {
		t.Fatal(err)
	}

	if chargeName != charge.ChargeName {
		t.Fatalf("want charge name = %s, got %s", chargeName, charge.ChargeName)
	}
}
