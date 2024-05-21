package orm

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/VladimirRytov/advsrv/internal/datastorage"
)

func TestRemoveTagByName(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		testTags := tagsDtoForTest
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		for _, tag := range testTags {
			err = database.RemoveTagByName(contex, tag.TagName)
			if err != nil {
				t.Fatal(err)
			}
			_, err = database.TagByName(contex, tag.TagName)
			if !errors.Is(err, datastorage.ErrNotFound) {
				t.Fatal(err)
			}
		}
		database.Close()
	}
}

func TestRemoveExtraChargeByName(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		testExtraCharges := extraChargesDtoForTest
		for _, charge := range testExtraCharges {
			err = database.RemoveExtraChargeByName(contex, charge.ChargeName)
			if err != nil {
				t.Fatal(err)
			}
			_, err = database.ExtraChargeByName(contex, charge.ChargeName)
			if !errors.Is(err, datastorage.ErrNotFound) {
				t.Fatal(err)
			}
		}
		database.Close()
	}
}

func TestRemoveLineAdvertisementByID(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		err = database.RemoveLineAdvertisementByID(contex, 1)
		if err != nil {
			t.Fatal(err)
		}
		_, err = database.LineAdvertisementByID(contex, 1)
		if !errors.Is(err, datastorage.ErrNotFound) {
			t.Fatal(err)
		}
		database.Close()
	}
}

func TestRemoveBlockAdvertisementByID(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		err = database.RemoveBlockAdvertisementByID(contex, 1)
		if err != nil {
			t.Fatal(err)
		}
		_, err = database.BlockAdvertisementByID(contex, 1)
		if !errors.Is(err, datastorage.ErrNotFound) {
			t.Fatal(err)
		}
		database.Close()
	}
}

func TestRemoveOrderByID(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		orderTest := orderDtoForTest
		err = database.RemoveOrderByID(contex, orderTest.ID)
		if err != nil {
			t.Fatal(err)
		}
		_, err = database.OrderByID(contex, orderTest.ID)
		if !errors.Is(err, datastorage.ErrNotFound) {
			t.Fatal(err)
		}
		database.Close()
	}
}

func TestRemoveClientByName(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		clientTest := clientDtoForTest
		err = database.RemoveClientByName(contex, clientTest.Name)
		if err != nil {
			t.Fatal(err)
		}
		_, err = database.ClientByName(contex, clientTest.Name)
		if !errors.Is(err, datastorage.ErrNotFound) {
			t.Fatal(err)
		}
		database.Close()
	}
}

func TestRemoveCostRateByName(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		err = database.RemoveCostRateByName(contex, costRateDto.Name)
		if err != nil {
			t.Fatal(err)
		}
		_, err = database.CostRateByName(contex, costRateDto.Name)
		if !errors.Is(err, datastorage.ErrNotFound) {
			t.Fatal(err)
		}
		database.Close()
	}
}

func TestRemoveUserByName(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		err = database.RemoveUser(contex, v)
		if err != nil {
			t.Fatal(err)
		}
		_, err = database.UserByName(contex, v)
		if !errors.Is(err, datastorage.ErrNotFound) {
			t.Fatal(err)
		}
		database.Close()
	}
}
