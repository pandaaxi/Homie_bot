package database

import "time"

type VPS struct {
	ID              int
	Remark          string
	IP              string
	Region          string
	Provider        string
	Price           float64
	TotalData       int64  // Total data package (in MB/GB as defined)
	ResetTimes      int
	AvailableData   int64
	ResetDate       string // Could be converted to time.Time if preferred
	CalculationType string // "in", "out", or "in+out"
}

type UsageLog struct {
	ID         int
	VPSID      int
	Timestamp  time.Time
	InUsage    int64
	OutUsage   int64
	TotalUsage int64
}

type Alert struct {
	ID        int
	VPSID     int
	Threshold float64
	AlertType string // "in", "out", or "in+out"
}
