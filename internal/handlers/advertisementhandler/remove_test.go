package advertisementhandler

import (
	"context"
	"encoding/binary"
	"testing"
)

func TestRemoveClient(t *testing.T) {
	clientName := "Вася"
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	err = db.RemoveClientByName(context.Background(), clientName)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveOrder(t *testing.T) {
	orderId := 1
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	err = db.RemoveOrderByID(context.Background(), orderId)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveBlockAdvertisement(t *testing.T) {
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}
	blockAdvID := 1
	err = db.RemoveBlockAdvertisementByID(context.Background(), blockAdvID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveLineAdvertisement(t *testing.T) {
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}
	lineAdvID := 1
	err = db.RemoveLineAdvertisementByID(context.Background(), lineAdvID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveTag(t *testing.T) {
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	tagName := "tag A"
	err = db.RemoveTagByName(context.Background(), tagName)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveExtraCharge(t *testing.T) {
	db, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	chargeName := "charge A"
	err = db.RemoveExtraChargeByName(context.Background(), chargeName)
	if err != nil {
		t.Fatal(err)
	}
}
func BenchmarkBytesEnc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b := make([]byte, binary.MaxVarintLen64)
		binary.BigEndian.PutUint64(b, uint64(33365654))
	}
}
