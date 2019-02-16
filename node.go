package Datatypes

type Node struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	Active  bool   `json:"active"`
}

type NodeStatsEntry struct {
	Name  string      `json:"name" bson:"name"`
	Stats []NodeStats `json:"stats" bson:"stats"`
}

type NodeStats struct {
	TotalMem      uint64  `json:"total_mem" bson:"total_mem"`
	UsedMem       uint64  `json:"used_mem" bson:"used_mem"`
	Cpus          int     `json:"cpus" bson:"cpus"`
	CpuPercent    float64 `json:"cpu_percent" bson:"cpu_percent"`
	NumContainers int     `json:"num_containers" bson:"num_containers"`
	Timestamp     int64   `json:"timestamp" bson:"timestamp"`
}
