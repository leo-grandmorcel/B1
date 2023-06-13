function bloodySunday(date) {
    var days = ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"]
    var start = new Date('0001-01-01')
    var difference = Math.ceil(date - start) / 86400000
    var index = difference - (Math.floor(difference / 6) * 6)
    return days[index]
}