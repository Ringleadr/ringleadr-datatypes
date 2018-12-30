package Datatypes

type Application struct {
	Name       string       `json:"name" binding:"required"`
	Copies     int          `json:"copies"`
	Components []*Component `json:"components" binding:"required"`
}
