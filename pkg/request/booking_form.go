package request

import "time"

type Booking struct {
	CheckIn time.Time	`form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time 	`form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}
