#!/bin/bash

trap "exit 1" ERR

echo "========> building sqlite for testing <========"
go install github.com/mattn/go-sqlite3

echo "========> Testing advertisements package <========"
go test github.com/VladimirRytov/advsrv/internal/advertisements
echo "========> Testing logging package <========"
go test github.com/VladimirRytov/advsrv/internal/logging
echo "========> Testing mapper package <========"
go test github.com/VladimirRytov/advsrv/internal/mapper
echo "========> Testing encodedecoder package <========"
go test github.com/VladimirRytov/advsrv/internal/encodedecoder
echo "========> Testing filestorage package <========"
go test github.com/VladimirRytov/advsrv/internal/filestorage
echo "========> Testing costRateCalculatorHandler package <========"
go test github.com/VladimirRytov/advsrv/internal/handlers/costcalculationhandler
echo "========> Testing broadcaster package <========"
go test github.com/VladimirRytov/advsrv/internal/handlers/broadcaster
echo "========> Testing broadcaster package <========"
go test github.com/VladimirRytov/advsrv/internal/validator
echo "========> Testing authorizator package <========"
go test github.com/VladimirRytov/advsrv/internal/authorizator

trap "rm internal/datastorage/sql/orm/testing.sqlite; exit 1" ERR

echo "========> Testing orm package <========"
echo "========> Testing tools <========"
go test -run '^(TestFetchTagsName|TestFetchExtraChargeName|TestReleaseDates)$' github.com/VladimirRytov/advsrv/internal/datastorage/sql/orm
echo "========> Testing mapper <========"
go test -run '^TestConvert' github.com/VladimirRytov/advsrv/internal/datastorage/sql/orm
go test -run 'ToModel$' github.com/VladimirRytov/advsrv/internal/datastorage/sql/orm
go test -run 'ToDto$' github.com/VladimirRytov/advsrv/internal/datastorage/sql/orm
echo "========> Testing connection <========"
go test -run '^TestConnectTo' github.com/VladimirRytov/advsrv/internal/datastorage/sql/orm
echo "========> Testing create requests <========"
go test -timeout 30s -run '^TestCreate' github.com/VladimirRytov/advsrv/internal/datastorage/sql/orm
echo "========> Testing get requests <========"
go test -run '^(TestClientByID|TestOrdersByID|TestLineAdvertisementByID|TestBlockAdvertisementByID|TestTagByName|TestExtraChargeByName|TestCostRateByName)$' github.com/VladimirRytov/advsrv/internal/datastorage/sql/orm
echo "========> Testing search requests <========"
go test -timeout 30s -run \
  '^(TestAllClients|TestAllOrders|TestOrdersByClientName|TestAllLineAdvertisements|TestLineAdvertisementsByOrderID|TestAllBlockAdvertisements|TestBlockAdvertisementsByOrderID|TestBlockAdvertisementBetweenReleaseDates|TestBlockAdvertisementFromReleaseDates|TestLineAdvertisementBetweenReleaseDates|TestLineAdvertisementFromReleaseDates|TestAllTags|TestAllExtraChargess|TestAllCostRates)$' \
  github.com/VladimirRytov/advsrv/internal/datastorage/sql/orm
echo "========> Testing update requests <========"
go test -timeout 30s -run '^TestUpdate' github.com/VladimirRytov/advsrv/internal/datastorage/sql/orm
echo "========> Testing remove requests <========"
go test -run '^TestRemove' github.com/VladimirRytov/advsrv/internal/datastorage/sql/orm
