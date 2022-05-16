# Initial Setup

## Hardware

We recommend the following hardware specifications:

- 16G RAM
- 4vCPUs
- 500GB Disk space

## Setup Instructions

We recommend using Ubuntu 18.04 or 20.04. The following setup instructions are assuming you are using one of these images and the setup may be different if not.

:::tip Note
The following commands are included in the [Quickstart Script](run-a-full-node.md#quickstart) and are therefore not required if you are using the script. The script also includes the steps outlined in [Run A Full Node](run-a-full-node.md).
:::

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
wget https://dl.google.com/go/go1.17.2.linux-amd64.tar.gz
```

Verify data integrity:

```bash
sha256sum go1.17.2.linux-amd64.tar.gz
```

Verify SHA-256 hash:

```bash
f242a9db6a0ad1846de7b6d94d507915d14062660616a61ef7c808a76e4f1676
```

Unpack Go download:

```bash
sudo tar -C /usr/local -xzf go1.17.2.linux-amd64.tar.gz
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
