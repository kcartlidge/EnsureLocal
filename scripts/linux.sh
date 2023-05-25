echo Building Windows edition
env GOOS=windows GOARCH=amd64 go build -o builds/windows/EnsureLocal.exe

echo Building Mac edition - X64
env GOOS=darwin GOARCH=amd64 go build -o builds/mac-x64/EnsureLocal

echo Building Mac edition - Apple Silicon M1 ARM64
env GOOS=darwin GOARCH=arm64 go build -o builds/mac-arm/EnsureLocal

echo Building Linux edition
env GOOS=linux GOARCH=amd64 go build -o builds/linux/EnsureLocal
