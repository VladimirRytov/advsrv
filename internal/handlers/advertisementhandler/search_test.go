package advertisementhandler

import (
	"context"
	"testing"
	"time"
)

func TestAllClients(t *testing.T) {
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	clients, err := db.AllClients(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(clients) == 0 {
		t.Fatal("clients length = 0")
	}
}

func TestOrdersByClientID(t *testing.T) {
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	orders, err := db.OrdersByClientName(context.Background(), "Вася")
	if err != nil {
		t.Fatal(err)
	}

	if len(orders) == 0 {
		t.Fatal("orders length = 0")
	}
}

func TestBlockAdvertisementsByOrderID(t *testing.T) {
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	blocks, err := db.BlockAdvertisementsByOrderID(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}

	if len(blocks) == 0 {
		t.Fatal("blockAdvertisements length = 0")
	}
}

func TestBlockAdvertisementsBetweenReleaseDates(t *testing.T) {
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	blocks, err := db.BlockAdvertisementsBetweenReleaseDates(context.Background(), time.Date(2000, 12, 12, 12, 0, 0, 0, time.UTC), time.Now().AddDate(22, 2, 2))
	if err != nil {
		t.Fatal(err)
	}

	if len(blocks) == 0 {
		t.Fatal("blockAdvertisements length = 0")
	}
}

func TestBlockAdvertisementsActualReleaseDate(t *testing.T) {
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	blocks, err := db.BlockAdvertisementsActualReleaseDate(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if len(blocks) == 0 {
		t.Fatal("blockAdvertisements length = 0")
	}
}

func TestBlockAdvertisementsFromDate(t *testing.T) {
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	blocks, err := db.BlockAdvertisementsFromReleaseDate(context.Background(), time.Date(2000, 12, 12, 12, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatal(err)
	}

	if len(blocks) == 0 {
		t.Fatal("blockAdvertisements length = 0")
	}
}

func TestLineAdvertisementsByOrderID(t *testing.T) {
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	lines, err := db.LineAdvertisementsByOrderID(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}

	if len(lines) == 0 {
		t.Fatal("lineAdvertisements length = 0")
	}
}

func TestLineAdvertisementsBetweenReleaseDates(t *testing.T) {
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	lines, err := db.LineAdvertisementsBetweenReleaseDates(context.Background(), time.Date(2000, 12, 12, 12, 0, 0, 0, time.UTC), time.Now().AddDate(2, 2, 2))
	if err != nil {
		t.Fatal(err)
	}

	if len(lines) == 0 {
		t.Fatal("lineAdvertisements length = 0")
	}
}

func TestLineAdvertisementsActualReleaseDate(t *testing.T) {
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	lines, err := db.LineAdvertisementsActualReleaseDate(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if len(lines) == 0 {
		t.Fatal("lineAdvertisements length = 0")
	}
}

func TestLineAdvertisementsFromDate(t *testing.T) {
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	lines, err := db.LineAdvertisementsFromReleaseDate(context.Background(), time.Date(2000, 12, 12, 12, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatal(err)
	}

	if len(lines) == 0 {
		t.Fatal("lineAdvertisements length = 0")
	}
}

func TestAllTags(t *testing.T) {
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}
	tags, err := db.AllTags(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(tags) == 0 {
		t.Fatal("tags length = 0")
	}
}

func TestAllExtraCharges(t *testing.T) {
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}
	extraCharges, err := db.AllExtraCharges(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(extraCharges) == 0 {
		t.Fatal("extraCharges length = 0")
	}
}

func TestAllCostRates(t *testing.T) {
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}
	costRates, err := db.AllCostRates(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(costRates) == 0 {
		t.Fatal("costRates length = 0")
	}
}
