package costcalculationhandler

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"testing"

	"github.com/VladimirRytov/advsrv/internal/datastorage/sql/orm"
	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/handlers"
	"github.com/VladimirRytov/advsrv/internal/logging"
)

func CreateLogger() {
	logging.CreateLogger(".", 0, &slog.HandlerOptions{Level: slog.LevelInfo}, false, os.Stderr)
}

func TestCreateDBForTests(t *testing.T) {
	CreateLogger()

	_, err := CreateDBForTests()
	if err != nil {
		t.Fatal(err)
	}
}

func TestCalculateBlockAdvertisementCost(t *testing.T) {
	var err error
	CreateLogger()

	db, err := CreateDBForTests()
	if err != nil {
		t.Fatal(err)
	}
	calc := NewCostRateCalculator(db)

	block, err := calc.CalculateBlockAdvertisementCost(context.Background(), blockDtoForTest, "dsa")
	if err != nil {
		t.Fatal(err)
	}
	if block.Cost != 2664 {
		t.Fatal(block.Cost, blockDtoForTest)
	}
	db.Close()
}

func TestCalculateLineAdvertisementCostWordCount(t *testing.T) {
	var err error
	CreateLogger()

	db, err := CreateDBForTests()
	if err != nil {
		t.Fatal(err)
	}
	calc := NewCostRateCalculator(db)

	line, err := calc.CalculateLineAdvertisementCost(context.Background(), lineDtoForTest, "dsa")
	if err != nil {
		t.Fatal(err)
	}
	if line.Cost != 1208 {
		t.Fatalf("want cost = 1208, got cost = %d", line.Cost)
	}
	db.Close()
}

func TestCalculateLineAdvertisementCostSymbolCount(t *testing.T) {
	var err error
	CreateLogger()

	costRateDto.CalcForOneWord = false
	db, err := CreateDBForTests()
	if err != nil {
		t.Fatal(err)
	}
	calc := NewCostRateCalculator(db)

	line, err := calc.CalculateLineAdvertisementCost(context.Background(), lineDtoForTest, "dsa")
	if err != nil {
		t.Fatal(err)
	}
	if line.Cost != 1232 {
		t.Fatalf("want cost = 1232, got cost = %d", line.Cost)
	}
	db.Close()
}

func TestCalculateNewOrder(t *testing.T) {
	var err error
	CreateLogger()

	costRateDto.CalcForOneWord = false
	db, err := CreateDBForTests()
	if err != nil {
		t.Fatal(err)
	}
	calc := NewCostRateCalculator(db)
	if err != nil {
		t.Fatal(err)
	}
	newOrder := orderDtoForTest
	newOrder.LineAdvertisements = []datatransferobjects.LineAdvertisementDTO{lineDtoForTest}
	newOrder.LineAdvertisements[0].Cost = 100

	newOrder.BlockAdvertisements = []datatransferobjects.BlockAdvertisementDTO{blockDtoForTest}
	newOrder.BlockAdvertisements[0].Cost = 200
	newOrder.ID = 0
	order, err := calc.CalculateOrderCost(context.Background(), newOrder, "dsa")
	if err != nil {
		t.Fatal(err)
	}
	if order.Cost != 300 {
		t.Fatalf("want cost = 300, got cost = %d", order.Cost)
	}
	db.Close()
}

func TestCalculateExistedOrder(t *testing.T) {
	var err error
	CreateLogger()

	costRateDto.CalcForOneWord = false

	db, err := CreateDBForTests()
	if err != nil {
		t.Fatal(err)
	}
	calc := NewCostRateCalculator(db)
	if err != nil {
		t.Fatal(err)
	}
	order, err := calc.CalculateOrderCost(context.Background(), orderDtoForTest, "dsa")
	if err != nil {
		t.Fatal(err)
	}
	if order.Cost != 46 {
		t.Fatalf("want cost = 46, got cost = %d", order.Cost)
	}
	db.Close()
}

func CreateDBForTests() (handlers.DataBase, error) {
	context, cancel := context.WithCancel(context.Background())
	defer cancel()
	param := &datatransferobjects.LocalDSN{Name: ":memory:"}

	mar, err := json.Marshal(&param)
	if err != nil {
		return nil, err
	}
	db, err := orm.ConnectToSqlite(mar)
	if err != nil {
		return nil, err
	}
	db.InitializeDatabaseMode(false)
	for _, v := range tagsDtoForTest {
		db.NewTag(context, &v)
	}
	for _, v := range extraChargesDtoForTest {
		db.NewExtraCharge(context, &v)
	}
	_, err = db.NewClient(context, &clientDtoForTest)
	if err != nil {
		return nil, err
	}
	_, err = db.NewAdvertisementsOrder(context, &orderDtoForTest)
	if err != nil {
		return nil, err
	}
	_, err = db.NewBlockAdvertisement(context, &blockDtoForTest)
	if err != nil {
		return nil, err
	}
	_, err = db.NewLineAdvertisement(context, &lineDtoForTest)
	if err != nil {
		return nil, err
	}
	_, err = db.NewCostRate(context, &costRateDto)
	if err != nil {
		return nil, err
	}
	return db, nil
}
