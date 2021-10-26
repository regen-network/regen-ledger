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

The prerequisites listed above for developers are also required for node operators. The following instructions will install the necessary prerequisites on a Linux machine.

:::tip NOTE
These commands are included in the [quickstart script](./running-a-validator.md#quickstart).
:::

Install tools:

```bash
sudo apt install git build-essential wget jq -y
```

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

You're now ready to set up a [full node](./running-a-full-node.md) or a [validator node](./running-a-validator.md).