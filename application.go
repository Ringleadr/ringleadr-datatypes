package Datatypes

type Application struct {
	Name       string       `json:"name" binding:"required"`
	Copies     int          `json:"copies"`
	Components []*Component `json:"components" binding:"required"`
	Networks   []string     `json:"networks"`
	Messages   []string     `json:"messages"`
	Node       string       `json:"node"`
}
