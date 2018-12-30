package Datatypes

type Component struct {
	Name     string     `json:"name"`
	Image    string     `json:"image" binding:"required"`
	Replicas int        `json:"replicas"`
	Storage  []*storage `json:"storage"`
}

type storage struct {
	Name      string `json:"name"`
	MountPath string `json:"mount_path"`
}
