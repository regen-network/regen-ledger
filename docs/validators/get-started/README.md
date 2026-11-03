# Initial Setup

## Hardware

We recommend the following hardware specifications:

- 16G RAM
- 4vCPUs
- 500GB Disk space

## Setup Instructions

We recommend using Ubuntu 22.04 or 24.04. The following setup instructions are assuming you are using one of these images and the setup may be different if not.

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
wget https://dl.google.com/go/go1.26.10.linux-amd64.tar.gz
```

Verify data integrity:

```bash
sha256sum go1.26.10.linux-amd64.tar.gz
```

Verify SHA-256 hash:

```bash
ccfe7ae822d2c2f41f4fe7c956199967a4dac22d0ee0d2b2bfdb440aab3b20c5
```

Unpack Go download:

```bash
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.26.10.linux-amd64.tar.gz
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
