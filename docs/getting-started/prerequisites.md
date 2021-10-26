# Prerequisites

## For Developers

### Git

For more information, see [Git](https://git-scm.com).

### Make

For more information, see [GNU Make](https://www.gnu.org/software/make/).

### Go

For more information, see [Go](https://golang.org/).

## For Node Operators

### Hardware

We recommend the following hardware specifications:

- 8GB RAM
- 4vCPUs
- 200GB Disk space

### Software

We recommend using Ubuntu 18.04 or 20.04.

The prerequisites listed above for developers are also required for node operators. The following instructions provide a quick start for installing prerequisites on a Linux machine. 

Install tools:
```
sudo apt install git build-essential wget jq -y
```

Download Go:
```
wget https://dl.google.com/go/go1.15.14.linux-amd64.tar.gz
```

Verify data integrity:
```
sha256sum go1.15.14.linux-amd64.tar.gz
```

Verify SHA-256 hash:
```
6f5410c113b803f437d7a1ee6f8f124100e536cc7361920f7e640fedf7add72d
```

Unpack Go download:
```
sudo tar -C /usr/local -xzf go1.15.14.linux-amd64.tar.gz
```

Set up environment:
```
echo '
export GOPATH=$HOME/go
export GOROOT=/usr/local/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:/usr/local/go/bin:$GOBIN' >> ~/.profile
```

Source profile file:
```
. ~/.profile
```

You're now ready to set up a [full node](./running-a-full-node.md) or a [validator node](./running-a-validator.md).