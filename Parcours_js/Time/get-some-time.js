function firstDayWeek(week, year) {
    var result = new Date(year)
    if (week === 1) {
        result.setHours(24)
        return datetostring(result)
    }
    let dayPlus = week * 7 * 24
    result.setHours(dayPlus - 123)
    for (let i = 0; i < 7; i++) {
        if (getWeekDay(result) == "Monday") {
            return datetostring(result)
        }
        result.setHours(-24)
    }
    return datetostring(result)
}

function datetostring(date) {
    var year = String(date.getFullYear()).padStart(4, '0')
    var month = String(date.getMonth() + 1).padStart(2, '0')
    var day = String(date.getDate() - 1).padStart(2, "0")
    return String(day + "-" + month + "-" + year)
}

function getWeekDay(date) {
    let days = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'];
    return days[date.getDay() - 1];
}
