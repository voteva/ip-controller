package entity

import "time"

type IpAccessEntity struct {
	IP            string    `db:"ip"`
	DatetimeFirst time.Time `db:"datetime_first"`
	DatetimeLast  time.Time `db:"datetime_last"`
}
