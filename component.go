package Datatypes

type Component struct {
	Name     string      `json:"name"`
	Image    string      `json:"image" binding:"required"`
	Replicas int         `json:"replicas"`
	Storage  []*storeage `json:"storage"`
}

type storeage struct {
	Name      string `json:"name"`
	MountPath string `json:"mount_path"`
}
