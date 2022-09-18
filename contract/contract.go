package contract

type ContainerInfo struct {
	ContainerName string
	ContainerId   string
}

var CInfo *ContainerInfo

func GetContainerInfo() *ContainerInfo {
	if CInfo == nil {
		CInfo = &ContainerInfo{}
	}
	return CInfo
}
