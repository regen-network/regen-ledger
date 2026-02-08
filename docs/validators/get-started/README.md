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
wget https://dl.google.com/go/go1.23.8.linux-amd64.tar.gz
```

Verify data integrity:

```bash
sha256sum go1.23.8.linux-amd64.tar.gz
```

Verify SHA-256 hash:

```bash
45b87381172a58d62c977f27c4683c8681ef36580abecd14fd124d24ca306d3f
```

Unpack Go download:

```bash
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.23.8.linux-amd64.tar.gz
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
