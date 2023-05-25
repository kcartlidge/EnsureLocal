# EnsureLocal

OneDrive is very poor at making sure files are always available locally.
Even though the option is available it often fails to work fully.
Other cloud services may not even have the option at all.

*EnsureLocal* scans the folder structure starting where it is run from.
The first byte of every file found is read (and ignored).
*This ensures every file is locally synced*.

[View the license (AGPL)](./LICENSE.txt).

## Pre-built executables

See the [builds](./builds) folder for executables for Mac, Windows, and Linux.

``` sh
cd builds
cd mac-arm      # Switch for your OS
./EnsureLocal   # Remove ./ on Windows
```

## Rebuilding locally

Clone the repo and run the relevant [scripts](./scripts) folder entry.
You *may* need to apply `chmod +x` to the script file for Mac or Linux.

## Running straight from source

This shouldn't be necessary, but as it is a single source file it is simple.
Navigate into this folder in a terminal/command prompt to run it.

``` sh
cd <wherever>
go run main.go
```
