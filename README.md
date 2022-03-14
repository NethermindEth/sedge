# 1Click
[![Go Report Card](https://goreportcard.com/badge/github.com/NethermindEth/1Click)](https://goreportcard.com/report/github.com/NethermindEth/1Click)

A tool to allow deploying validators with ease.

## Installation (Only UNIX systems)

### Using Go

If you have at least `go1.17.5` installed then this command will install the 1Click executable along with the library and its dependencies:

```
go install github.com/2kodevs/domline/cmd/domline@latest
```

The executable will be in `$GOBIN` (`$GOPATH/bin`) 

### Manual

Generate the executable manually (need Go installed):

```
git clone https://github.com/NethermindEth/1Click.git
cd 1Click
make compile
```

The executable will be in the `1Click/build` folder

---
In case you want the binary in PATH (in case you don't have `$GOBIN` either in PATH), copy it to `/usr/local/bin`:

```
# Using go
sudo $GOPATH/bin/1Click /usr/local/bin/
# Manual
sudo cp 1Click/build/1Click /usr/local/bin/
```

### Download the binary

> This is temporary until the first release

Download directly the binary and put it in `/usr/local/bin`:

```
sudo curl -LJ -o /usr/local/bin/1Click https://github.com/NethermindEth/1Click/raw/feature/main/build/1Click
sudo chmod +x /usr/local/bin/1Click
```