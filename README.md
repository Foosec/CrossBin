# CrossBin

Packages MacOS, Linux and Windows binaries, into a single script that
is executable everywhere and executes the correct binary for the system.




# Building

## Requirements
- Golang

Running build.sh will cross compile for all 3 major OS's and output to dist/


# Usage

Supply it with pre compiled binaries for whichever platform you wish to run on

    ./CrossBin -w path_to_windows_binary -l path_to_linux_binary -m path_to_macos_binary -o output_file

Default output file is CrossBin.ps1





# Major contributors

## [@aJuvan](https://github.com/aJuvan) Helped with the development of the CrossShell script and added MacOs support. 
