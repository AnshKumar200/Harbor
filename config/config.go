package config

const (
	DefaultImageStoreRootDir     = "/var/harbor/img"
	DefaultContainerStoreRootDir = "/var/harbor/cont"
	DefaultCGroupDir             = "/sys/fs/cgroup/"
	DefaultRegistry              = "https://registry-1.docker.io/"

	DefaultPidsLimit   = 32
	DefaultMemoryLimit = 64 * 1024 * 1024 // 64MB

	DefaultLogLevel = "debug"
)
