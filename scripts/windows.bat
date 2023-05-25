@echo off

@echo Building Linux edition
@set GOOS=linux
@set GOARCH=amd64
@go build -o builds/linux/EnsureLocal

@echo Building Mac edition - X64
@set GOOS=darwin
@set GOARCH=amd64
@go build -o builds/mac-x64/EnsureLocal

@echo Building Mac edition - Apple Silicon M1 ARM64
@set GOOS=darwin
@set GOARCH=arm64
@go build -o builds/mac-arm/EnsureLocal

@echo Building Windows edition
@set GOOS=windows
@set GOARCH=amd64
@go build -o builds/windows/EnsureLocal.exe
