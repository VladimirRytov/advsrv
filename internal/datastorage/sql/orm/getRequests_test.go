package orm

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/VladimirRytov/advsrv/internal/datastorage"
)

func TestClientByID(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf(err.Error(), v)
		}

		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		result, err := database.ClientByName(contex, clientDtoForTest.Name)
		if err != nil {
			t.Fatalf(err.Error(), v)
		}

		if clientDtoForTest.Name != result.Name || clientDtoForTest.Phones != result.Phones ||
			clientDtoForTest.Email != result.Email || clientDtoForTest.AdditionalInformation != result.AdditionalInformation {
			t.Fatalf("want name = %v,got client = %v", clientDtoForTest, result)
		}

		result, err = database.ClientByName(contex, "v")
		if err == nil || err != datastorage.ErrNotFound {
			t.Fatalf("want error %v", result)
		}
		database.Close()
	}
}

func TestOrdersByID(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		wantID := 1
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		result, err := database.OrderByID(contex, wantID)
		if err != nil {
			t.Fatalf(err.Error())
		}

		if wantID != result.ID {
			t.Fatalf("want id = %v,got order = %v", wantID, result)
		}

		wantID = 2
		result, err = database.OrderByID(contex, wantID)
		if err == nil || err != datastorage.ErrNotFound {
			t.Fatalf("want error,got %v ", result)
		}
		database.Close()
	}
}

func TestLineAdvertisementByID(t *testing.T) {
	CreateLogger()
	wantID := 1
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}

		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		result, err := database.LineAdvertisementByID(contex, wantID)
		if err != nil {
			t.Fatalf(err.Error())
		}
		if !reflect.DeepEqual(lineDtoForTest, result) {
			t.Fatalf("want id = %v,got LineAdvertisement = %v", lineDtoForTest, result)
		}

		_, err = database.LineAdvertisementByID(contex, 200)
		if err == nil || err != datastorage.ErrNotFound {
			t.Fatalf("want error")
		}
		database.Close()
	}
}

func TestBlockAdvertisementByID(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}

		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		wantID := 1
		result, err := database.BlockAdvertisementByID(contex, wantID)
		if err != nil {
			t.Fatalf(err.Error())
		}
		if !reflect.DeepEqual(result, blockDtoForTest) {
			t.Fatalf("database %s \nwant id = %v,got LineAdvertisement = %v", v, blockDtoForTest, result)
		}

		_, err = database.BlockAdvertisementByID(contex, 200)
		if err == nil || err != datastorage.ErrNotFound {
			t.Fatalf("want error")
		}
		database.Close()
	}
}

func TestTagByName(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}

		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		for _, v := range tagsDtoForTest {
			result, err := database.TagByName(contex, v.TagName)
			if err != nil {
				t.Fatalf(err.Error())
			}
			if result.TagName != v.TagName && result.TagCost != v.TagCost {
				t.Fatalf("want tag = %v,got tag = %v", v, result)
			}
			result, err = database.TagByName(contex, "xcg")
			if err == nil || err != datastorage.ErrNotFound {
				t.Fatalf("want error")
			}
		}
		database.Close()
	}
}

func TestExtraChargeByName(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}

		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		for _, v := range extraChargesDtoForTest {
			result, err := database.ExtraChargeByName(contex, v.ChargeName)
			if err != nil {
				t.Fatalf(err.Error())
			}
			if result.ChargeName != v.ChargeName && result.Multiplier != v.Multiplier {
				t.Fatalf("want tag = %v,got tag = %v", v, result)
			}
			result, err = database.ExtraChargeByName(contex, "xcg")
			if err == nil || err != datastorage.ErrNotFound {
				t.Fatalf("want error")
			}
		}
		database.Close()
	}
}

func TestCostRateByName(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}

		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		result, err := database.CostRateByName(contex, costRateDto.Name)
		if err != nil {
			t.Fatalf(err.Error())
		}
		if result.Name != costRateDto.Name {
			t.Fatalf("want costRate = %v,got costRate = %v", v, result)
		}
		result, err = database.CostRateByName(contex, "zxc")
		if err == nil || err != datastorage.ErrNotFound {
			t.Fatalf("want error %v,got err %v", datastorage.ErrNotFound, err)
		}
		database.Close()
	}
}
