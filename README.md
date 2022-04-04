# 1click
[![Go Report Card](https://goreportcard.com/badge/github.com/NethermindEth/1click)](https://goreportcard.com/report/github.com/NethermindEth/1click)

A tool to allow deploying validators with ease.

## Installation (Only UNIX systems)

### Using Go

If you have at least `go1.17.5` installed then this command will install the 1click executable along with the library and its dependencies:

```
go install github.com/NethermindEth/1click/cmd/1click@latest
```

The executable will be in `$GOBIN` (`$GOPATH/bin`) 

### Manual

Generate the executable manually (need Go installed):

```
git clone https://github.com/NethermindEth/1click.git
cd 1click
go build -o 1click cmd/1click/main.go
```

or if you have `make` installed:

```
git clone https://github.com/NethermindEth/1click.git
cd 1click
make compile
```

The executable will be in the `1click/build` folder

---
In case you want the binary in PATH (in case you don't have `$GOBIN` either in PATH), copy it to `/usr/local/bin`:

```
# Using go
sudo $GOPATH/bin/1click /usr/local/bin/
# Manual
sudo cp 1click/build/1click /usr/local/bin/
```

### Download the binary

> This is temporary until the first release

Download directly the binary and put it in `/usr/local/bin`:

```
sudo curl -LJ -o /usr/local/bin/1click https://github.com/NethermindEth/1click/raw/feature/main/build/1click
sudo chmod +x /usr/local/bin/1click
```