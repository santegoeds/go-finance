package date

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
    Day
    Week
    Month
    Year
)

type Temporal interface {
    Day() int
    Month() int
    Year() int
    Week() int
    WeekDay() int
    IsLeapYear() bool
    Advance(n int, period int)
}

type Date struct {
    dayCount int
}

func (d *Date) Day() int {
    return dayFromOffset(d.dayCount)
}

func (d *Date) Month() int {
    return monthFromOffset(d.dayCount)
}

func (d *Date) Year() int {
    return yearFromOffset(d.dayCount)
}

func (d *Date) Week() int {
    return weekFromOffset(d.dayCount)
}

func (d *Date) IsLeapYear() bool {
    return isLeapYear(d.dayCount)
}
