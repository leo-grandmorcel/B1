function matchCron(string, date) {
    var monbool = true
    var cron = string.split(" ")
    if (cron[0] != "*") {
        monbool = cron[0] == date.getMinutes()
    }
    if (cron[1] != "*") {
        monbool = cron[1] == date.getHours()
    }
    if (cron[2] != "*") {
        monbool = cron[2] == date.getDate()
    }
    if (cron[3] != "*") {
        monbool = cron[3] == date.getMonth() + 1
    }
    if (cron[4] != "*") {
        monbool = cron[4] == date.getDay()
    }
    return monbool
}
console.log(matchCron('* * * * *', new Date()))