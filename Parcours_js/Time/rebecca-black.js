const days = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday']

function isFriday(date) {
    return (date.getDay() == 5)
}

function isWeekend(date) {
    return (date.getDay() == 0 || date.getDay() == 6)
}
function isLeapYear(date) {
    var year = date.getFullYear()
    return ((year % 4 == 0) && (year % 100 != 0)) || (year % 400 == 0)
}
function isLastDayOfMonth(date) {
    var mois = date.getMonth()
    date.setHours(36)
    var newmois = date.getMonth()
    return mois != newmois
}