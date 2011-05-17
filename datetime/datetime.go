package datetime

import (
    "time"
)

type DateTime struct {
    *Date
    time int
}

func Now() *DateTime {
    t := time.LocalTime()
    return NewDateTime(int(t.Year), t.Month, t.Day, t.Hour, t.Minute, t.Second)
}

func NewDateTime(args ...int) *DateTime {
    var hour, minute, second int

    year, month, day := 1970, 1, 1

    for i, v := range args {
        switch i {
        case 0:
            year = v
            break
        case 1:
            month = v
            break
        case 2:
            day = v
            break
        case 3:
            hour = v
            break
        case 4:
            minute = v
            break
        case 5:
            second = v
            break
        }
    }
    second += minute * 60 + hour * 3600
    return &DateTime{NewDate(year, month, day), second}
}

