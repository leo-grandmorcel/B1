function dayOfTheYear(date) {
    var start = new Date('0001-01-01')
    start.setFullYear(date.getFullYear())
    date.setHours(24)
    return Math.ceil((date - start) / 86400000)
}