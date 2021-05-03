:: OS: Windows
:: batch script to build the src using go-sdk for windows

:: Requirement
:: 1. Go sdk 1.16.x

:: Steps
:: run 'build -d' in cmd for getting all dependencies
:: run 'build -b' in cmd for Building
:: run 'build -i' in cmd for Installing using Go

@echo off

:: Check go is installed or not
go version /? >nul 2>&1
if errorlevel 1 (
    :: to build and install ghsh using go-sdk
    if %1% == "-i" goto install 

    :: this is only for building ghsh using go-sdk
    if %1% == "-b" goto build

    :: get all dependencies
    if %1% == "-d" goto dependencies
)

if not errorlevel 1 echo It seems like Go isn't installed.

:build
    echo Building ghsh using go..
    go build -o ../../ghsh.exe ../../main.go
    echo Completed
    pause
exit

:install
    echo Installing ghsh using go..
    go install -o ../../ghsh.exe ../../main.go
    echo Completed
    pause
exit

:dependencies
    echo Getting all dependencies..
    go get ../../...
    echo Completed
    pause
exit