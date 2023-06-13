function manipDate(firstdate, seconddate) {
    if ((isNaN(firstdate) || firstdate == 0) || (isNaN(seconddate) || seconddate == 0)) {
        return "impossible"
    }
    return {
        daysBetween: Math.abs((seconddate - firstdate) / 86400000),
        isFirstLeap: ((firstdate.getUTCFullYear() % 4 == 0) && (firstdate.getUTCFullYear() % 100 != 0)) || (firstdate.getUTCFullYear() % 400 == 0),
        isSecondLeap: ((seconddate.getUTCFullYear() % 4 == 0) && (seconddate.getUTCFullYear() % 100 != 0)) || (seconddate.getUTCFullYear() % 400 == 0)
    }
}
module.exports = manipDate();