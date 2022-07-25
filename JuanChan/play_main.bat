@echo off

cls
set startTimeTwo=%time%
echo Building and running Main
go build main.go
echo Started: %startTimeTwo%

echo Main is  currently running.
go run main.go
set finishTimeTwo=%time%
echo main.go has stopped
echo Ended: %finishTimeTwo%

