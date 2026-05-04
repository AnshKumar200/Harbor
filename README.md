# Harbor

A container runtime built from scratch in Go using Linux namespaces, cgroups, and chroot.

## Prerequisites

Harbor runs only on Linux-based system with version 3.10 or higher of the Linux kernel.

Required packages:
- libcgroup-tools

Required configuration:
- A btrfs filesystem mounted under /var/harbor

## Install

```bash
git clone https://github.com/AnshKumar200/Harbor
cd Harbor
make
cd build
./Harbor --help
```

## Usage
```shell
./Harbor image list
./Harbor pull arm64v8/alpine
./H2arbor run arm64v8/alpine
```
