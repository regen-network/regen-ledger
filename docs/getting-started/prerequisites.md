# Prerequisites

We recommend using Ubuntu 18.04 or higher, the official images can be found at [Ubuntu.com](https://ubuntu.com/tutorials/install-ubuntu-desktop#1-overview).

## Software Requirements

### Git

```
sudo apt update
sudo apt install git
```

For more information, see [Git](https://git-scm.com/).

### Make (Build Essential)

Chances are you will need things like `gcc` to actually do the building so you might as well install `build-essential`. The `build-essential` package includes `gcc`, `make`, and other tools.

```
sudo apt update
sudo apt install build-essential
```

For more information, see [GNU Make](https://www.gnu.org/software/make/).

### Go

```
wget https://dl.google.com/go/go1.15.2.linux-amd64.tar.gz
tar -xvf go1.15.2.linux-amd64.tar.gz
```

```
sudo mv go /usr/local
```

Add go to local path

```
echo '' >> ~/.profile
echo 'export GOPATH=$HOME/go' >> ~/.profile
echo 'export GOROOT=/usr/local/go' >> ~/.profile
echo 'export GOBIN=$GOPATH/bin' >> ~/.profile
echo 'export PATH=$PATH:/usr/local/go/bin:$GOBIN' >> ~/.profile
```

```
. ~/.profile
```

Check to make sure you have the correct version installed (version 1.15 or higher):

```
go version
```

For more information, see [Go](https://golang.org/)

## Hardware Requirements

The following requirements are only necessary if you are running a validator node or a full node for Regen Mainnet (`regen-1`) or Regen Devnet (`regen-devnet-5`).

- 8GB RAM
- 4vCPUs
- 200GB Disk space
