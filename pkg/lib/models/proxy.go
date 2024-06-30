package models

type Status string

const (
	UP      Status = "up"
	DOWN    Status = "down"
	UNKNOWN Status = "not tested"
)

type Proxy struct {
	Ip     string
	Port   int64
	Speed  *string
	Status Status
}
