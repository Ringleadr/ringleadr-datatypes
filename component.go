package Datatypes

type Component struct {
	Name     string            `json:"name"`
	Image    string            `json:"image" binding:"required"`
	Replicas int               `json:"replicas"`
	Storage  []*storage        `json:"storage"`
	Ports    map[string]string `json:"ports"`
	Env      []string          `json:"env"`
	Status   string            `json:"status"`
}

type storage struct {
	Name      string `json:"name"`
	MountPath string `json:"mount_path"`
}
