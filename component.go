package Datatypes

type Component struct {
	Name           string            `json:"name"`
	Image          string            `json:"image" binding:"required"`
	Replicas       int               `json:"replicas"`
	Storage        []*storage        `json:"storage"`
	Ports          map[string]string `json:"ports"`
	Env            []string          `json:"env"`
	Status         string            `json:"status"`
	ScaleThreshold float64           `json:"scale_threshold"`
	ScaleMin       int               `json:"scale_min"`
	ScaleMax       int               `json:"scale_max"`
	NetworkInfo    []*NetworkInfo    `json:"network_info"`
}

type storage struct {
	Name      string `json:"name"`
	MountPath string `json:"mount_path"`
}

type NetworkInfo struct {
	Name string   `json:"name"`
	Addr []string `json:"addr"`
}
