package Datatypes

type Node struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	Active  bool   `json:"active"`
}

type NodeStats struct {
	TotalMem   uint64  `json:"total_mem"`
	UsedMem    uint64  `json:"used_mem"`
	Cpus       int     `json:"cpus"`
	CpuPercent float64 `json:"cpu_percent"`
}
