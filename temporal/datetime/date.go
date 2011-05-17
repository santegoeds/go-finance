package date

import (
    "time"
)

const (
    _   = iota
    January
    February
    March
    April
    May
    June
    July
    August
    September
    October
    November
    December
)

const (
    Monday = iota
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
    Sunday
)

const (
    Day = iota
    Week
    Month
    Year
)

type Date struct {
    dayCount int
}

func (d Date) Day() int {
    return dayFromOffset(d.dayCount)
}

func (d Date) Month() int {
    return monthFromOffset(d.dayCount)
}

func (d Date) Year() int {
    return yearFromOffset(d.dayCount)
}

func (d Date) Week() int {
    return weekFromOffset(d.dayCount)
}

func (d Date) WeekDay() int {
    return weekdayFromOffset(d.dayCount)
}

func (d Date) IsLeapYear() bool {
    return isLeapYear(d.dayCount)
}

func Today() *Date {
    t := time.LocalTime()
    return NewDate(int(t.Year), t.Month, t.Day)
}

func NewDate(year, month, day int) *Date {
    return &Date{offsetFromDate(year, month, day)}
}

