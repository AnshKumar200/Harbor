# Harbor

A container runtime built from scratch in Go using Linux namespaces, cgroups, and chroot.

## Prerequisites

Harbor runs only on Linux-based system with version 3.10 or higher of the Linux kernel.

Required packages:
- libcgroup-tools

Required configuration:
- A btrfs filesystem mounted under /var/harbor (configurable)
- A cgroup filesystem mounted under /sys/fs/cgroup/ (configurable) if not already the case

## Install

```bash
git clone https://github.com/AnshKumar200/Harbor
cd Harbor
make
cd build
./Harbor --help
```

## Usage
When using Harbor, you need to be specific when pulling or running an image. For example, use `pull amd64/alpine` instead of just `pull alpine`.

```shell
# harbor requires root privileges
sudo su

./Harbor image list
./Harbor pull arm64v8/alpine
./Harbor run arm64v8/alpine:latest /bin/sh
./H2arbor run arm64v8/alpine
```
