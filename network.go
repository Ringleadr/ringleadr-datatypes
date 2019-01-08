package Datatypes

type Network struct {
	Name string `json:"name" binding:"required"`
}
