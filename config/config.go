package config

import "github.com/AnshKumar200/Harbor/pkg/logging"

const (
	DefaultImageStoreRootDir     = "/var/harbor/img"
	DefaultContainerStoreRootDir = "/var/harbor/cont"
	DefaultCGroupDir             = "/sys/fs/cgroup/"
	DefaultRegistry              = "https://registry-1.docker.io/"

	DefaultLogLevel = logging.Debug
)
