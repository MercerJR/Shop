package entity

import "time"

type Order struct {
	CId string
	GID string
	SID string
	Number string
	Time time.Time
}
