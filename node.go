package Datatypes

type Node struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}
