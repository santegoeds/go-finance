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

var offsetsPerMonth = []int{
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
        daycount += offsetsPerMonth[month-1]
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
    for month, mo := range offsetsPerMonth {
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
    for month, mo := range offsetsPerMonth {
        if mo > daycount {
            daycount -= offsetsPerMonth[month-1]
            if month > January && is_leap {
                daycount -= 1
            }
            return daycount + 1
        }
    }
    daycount -= offsetsPerMonth[len(offsetsPerMonth)-1]
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

var daysPerMonth = []int{
    31,
    28,
    31,
    30,
    31,
    30,
    31,
    31,
    30,
    31,
    30,
    31,
}

func advance(daycount int, n int, period int) int {
    var day, month, year, maxday int
    switch period {
    case Day:
        daycount += n
    case Week:
        daycount += n*7
    case Month:
        day = dayFromOffset(daycount)
        month = monthFromOffset(daycount)
        year = yearFromOffset(daycount)
        switch {
        case month > 12:
            year += month/12
            month = month%12
            if month == 0 {
                month = 12
            }
        case month < 1:
            yrs := (month-12)/12
            year += yrs
            month -= yrs*12
        }
        maxday = offsetsPerMonth[month-1]
        if month > 1 && isLeapYear(year) {
            maxday += 1
        }
        if day > maxday {
            day = maxday
        }
        daycount = offsetFromDate(day, month, year)
    case Year:
        day = dayFromOffset(daycount)
        month = monthFromOffset(daycount)
        year = yearFromOffset(daycount) + n
        maxday = offsetsPerMonth[month-1]
        if month > 1 && isLeapYear(year) {
            maxday += 1
        }
        if day > maxday {
            day = maxday
        }
        daycount = offsetFromDate(day, month, year)
    }
    return daycount
}
