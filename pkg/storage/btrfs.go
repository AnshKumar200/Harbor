package storage

import (
	"github.com/containerd/btrfs"
)

type COWFS interface {
	SubvolCreate(path string) error
	SubvolSnapshot(dst string, src string) error
}

type Btrfs struct {
}

func (b Btrfs) SubvolCreate(path string) error {
	return btrfs.SubvolCreate(path)
}

func (b Btrfs) SubvolSnapshot(dst string, src string) error {
	return btrfs.SubvolSnapshot(dst, src, false)
}
