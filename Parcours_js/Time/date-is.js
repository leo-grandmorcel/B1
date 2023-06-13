function isValid(date) {
    return !(isNaN(date) || date == 0)
}

function isAfter(date1, date2) {
    return date1 - date2 > 0
}

function isBefore(date1, date2) {
    return date2 - date1 > 0
}

function isFuture(date) {
    if (isValid(date)) {
        var current = Date.now()
        return current < date
    }
    return false
}

function isPast(date) {
    if (isValid(date)) {
        var current = Date.now()
        return current > date
    }
    return false
}

