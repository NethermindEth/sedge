package services

type VolumeType string

const (
	VolumeTypeBind   VolumeType = "bind"
	VolumeTypeVolume VolumeType = "volume"
)

type ContainerLogsMergedOptions struct {
	Follow     bool
	Since      string
	Until      string
	Timestamps bool
	Tail       string
}

type Mount struct {
	Type   VolumeType
	Source string
	Target string
}

type RunOptions struct {
	Network     string
	Args        []string
	Mounts      []Mount
	VolumesFrom []string
	AutoRemove  bool
}
