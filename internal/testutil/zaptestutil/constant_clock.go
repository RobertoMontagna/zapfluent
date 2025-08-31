package zaptestutil

import "time"

type ConstantClockForTest time.Time

func (c ConstantClockForTest) Now() time.Time {
	return time.Time(c)
}

func (c ConstantClockForTest) NewTicker(_ time.Duration) *time.Ticker {
	return &time.Ticker{}
}
