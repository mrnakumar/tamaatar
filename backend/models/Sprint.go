package models

type Sprint struct {
	Name     string
	Duration uint8
	EndTime  int64
	HourOffset int8
	MinuteOffset uint8
}
