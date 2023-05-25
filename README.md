# EnsureLocal

**Ready to use**, despite only being a few commits. (No more were necessary.)

OneDrive is poor at making sure files are always available locally.
Even though the option exists, it often fails to work fully and will also remove local versions of files if it considers space is getting tight on your machine.
Other cloud services may not have the option at all.

*EnsureLocal* scans a nested folder structure.
The first byte of every file found is read (and ignored).
*This forces a local sync to occur*.

If a file cannot be read it tries again after 2 seconds and then after another 10 seconds, allowing cloud syncing to catch up if it is running slow.

Folders/files with names starting `.` are skipped.

You can stop the run at any time with Ctrl+C.
Restarting will whizz through previously accessed files as they are now local.

If run without a parameter, it will prompt for the folder.

[View the license (AGPL)](./LICENSE.txt).

## Pre-built executables

See the [builds](./builds) folder for executables for Mac, Windows, and Linux.

``` sh
cd builds
cd mac-arm      # Switch for your OS
./EnsureLocal   # Remove ./ on Windows
```

## Rebuilding locally

Clone the repo and run the relevant [scripts](./scripts) folder entry (scripts exist for running on *Mac*, *Linux*, and *Windows*).
You *may* need to apply `chmod +x` to the script file for Mac or Linux.

Scripts should be run from the top level folder (where this `README` lives).

``` sh
cd <repo>
./scripts/mac.sh
```

## Running straight from source

This shouldn't be necessary, but as it is a single source file it is simple.
Navigate into this folder in a terminal/command prompt to run it.

``` sh
cd <wherever>
go run main.go
```
