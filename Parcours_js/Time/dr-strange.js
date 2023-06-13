function addWeek(date) {
    var days = ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday", "secondMonday", "secondTuesday", "secondWednesday", "secondThursday", "secondFriday", "secondSaturday", "secondSunday"]
    var start = new Date('0001-01-01')
    var difference = Math.ceil(date - start) / 86400000
    var index = difference - (Math.floor(difference / 14) * 14)
    return days[index]
}

function timeTravel({ date, hour, minute, second }) {
    date.setHours(hour)
    date.setMinutes(minute)
    date.setSeconds(second)
    return date
}