package Datatypes

type Component struct {
	Name     string `json:"name"`
	Image    string `json:"image" binding:"required"`
	Replicas int    `json:"replicas"`
}
