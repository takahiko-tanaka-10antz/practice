package model

import "time"

type User struct {
	Id       int64
	PlayerId string
	Ap       int
	Lp       int
	LoginAt  time.Time
}
