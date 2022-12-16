# Initial Setup

## Hardware

We recommend the following hardware specifications:

- 16G RAM
- 4vCPUs
- 500GB Disk space

## Setup Instructions

We recommend using Ubuntu 18.04 or 20.04. The following setup instructions are assuming you are using one of these images and the setup may be different if not.

### Install Dependencies

Update packages:

```bash
sudo apt update
```

Install tools:

```bash
sudo apt install git build-essential wget jq -y
```

### Install Go

Download Go:

```bash
wget https://dl.google.com/go/go1.19.4.linux-amd64.tar.gz
```

Verify data integrity:

```bash
sha256sum go1.19.4.linux-amd64.tar.gz
```

Verify SHA-256 hash:

```bash
9df122d6baf6f2275270306b92af3b09d7973fb1259257e284dba33c0db14f1b
```

Unpack Go download:

```bash
sudo tar -C /usr/local -xzf go1.19.4.linux-amd64.tar.gz
```

Set up environment:

```bash
echo '
export GOPATH=$HOME/go
export GOROOT=/usr/local/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:/usr/local/go/bin:$GOBIN' >> ~/.profile
```

Source profile file:

```bash
. ~/.profile
```
