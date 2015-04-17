package stats

import "time"

type EpUri struct {
	Path     string
	IdMethod int
	Method   string
	Name     string
	Group    string

	DoStats StatTypes
}

type StatTypes struct {
	Frequency  bool
	Size       bool
	Longest    bool
	Shortest   bool
	StatusCode bool
}

type StatTime struct {
	Day   int
	Month int
	Year  int
	Hour  int
	Min   int
}

type StatMetric struct {
	DateTime    time.Time
	EPU         *EpUri
	Time        StatTime
	MetricName  string
	MetricUnit  string
	MetricValue float64
}

type EPURI_StatMetric struct {
	EP       EpUri
	Stats    StatMetric
	TimeSpan StatTime
}
