package date

const (
    baseYear = 1601
)

func isLeapYear(year int) bool {
    if year%4 == 0 {
        if year%400 == 0 {
            return true
        }
        if year%100 == 0 {
            return false
        }
        return true
    }
    return false
}

func offsetFromYear(year int) int {
    daycount := 0
    year -= baseYear
    if year < 0 {
        year += 1
        daycount = year*365 + year/4 + year/400 - year/100 - 366
    } else {
        daycount = year*365 + year/4 + year/400 - year/100
    }
    return daycount
}

var daysPerMonth = []int{
    0,
    31,  // February
    58,  // March
    90,  // April
    120, // May
    151, // June
    181, // July
    212, // August
    243, // September
    273, // October
    304, // November
    334, // December
}

func offsetFromDate(day int, month int, year int) int {
    daycount := offsetFromYear(year)
    if month > January {
        daycount += daysPerMonth[month-1]
        if month > February && isLeapYear(year) {
            daycount += 1
        }
    }
    return daycount + (day - 1)
}

func yearFromOffset(daycount int) int {
    year := daycount/365 + baseYear
    for daycount < offsetFromYear(year) {
        year -= 1
    }
    return year
}

func monthFromOffset(daycount int) int {
    daycount -= offsetFromYear(yearFromOffset(daycount))
    for month, mo := range daysPerMonth {
        if mo > daycount {
            return month
        }
    }
    return 12
}

func dayFromOffset(daycount int) int {
    year := yearFromOffset(daycount)
    is_leap := isLeapYear(year)
    daycount -= offsetFromYear(year)
    for month, mo := range daysPerMonth {
        if mo > daycount {
            daycount -= daysPerMonth[month-1]
            if month > January && is_leap {
                daycount -= 1
            }
            return daycount + 1
        }
    }
    daycount -= daysPerMonth[len(daysPerMonth)-1]
    if is_leap {
        daycount -= 1
    }
    return daycount + 1
}

func weekFromOffset(daycount int) int {
    days := daycount - yearFromOffset(daycount)
    return days/7 + 1
}

func weekdayFromOffset(daycount int) int {
    weekday := daycount % 7
    if weekday < 0 {
        weekday = 7 + weekday
    }
    return weekday
}

func advance(daycount int, n int, period int) int {
    switch period {
    case Year:
    case Month:
    case Week:
        daycount += n*7
    case Day:
        daycount += n
    }
    return daycoun
}
