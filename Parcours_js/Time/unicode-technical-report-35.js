const Months = ['January', 'February', 'March', 'April', 'May', 'June', 'July',
    'August', 'September', 'October', 'November', 'December']
const Days = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday']

function format(date, format) {
    let isPm = false
    var result = format
    if (format.match(/yyyy/)) {
        if (date.getFullYear() >= 0) {
            result = result.replace(/yyyy/, String(date.getFullYear()).padStart(4, '0'))
        } else {
            result = result.replace(/yyyy/, String(-(date.getFullYear())).padStart(4, '0'))
        }
    } else if (format.match(/y/)) {
        if (date.getFullYear() >= 0) {
            result = result.replace(/y/, String(date.getFullYear()))
        } else {
            result = result.replace(/y/, String(-(date.getFullYear())))
        }

    }
    if (format.match(/GGGG/)) {
        if (date.getFullYear() < 0) {
            result = result.replace(/GGGG/, "Before Christ")
        } else {
            result = result.replace(/GGGG/, "Anno Domini")
        }
    } else if (format.match(/G/)) {
        if (date.getFullYear() < 0) {
            result = result.replace(/G/, "BC")
        } else {
            result = result.replace(/G/, "AD")
        }
    }
    if (format.match(/MMMM/)) {
        result = result.replace(/MMMM/, Months[date.getMonth()])
    } else if (format.match(/MMM/)) {
        result = result.replace(/MMM/, Months[date.getMonth()].slice(0, 3))
    } else if (format.match(/MM/)) {
        result = result.replace(/MM/, String(date.getMonth() + 1).padStart(2, "0"))
    } else if (format.match(/M/)) {
        result = result.replace(/M/, String(date.getMonth() + 1))
    }
    if (format.match(/dd/)) {
        result = result.replace(/dd/, String(date.getDate()).padStart(2, "0"))
    } else if (format.match(/d/)) {
        result = result.replace(/d/, String(date.getDate()))
    }
    if (format.match(/EEEE/)) {
        result = result.replace(/EEEE/, Days[date.getDay()])
    } else if (format.match(/E/)) {
        result = result.replace(/E/, Days[date.getDay()].slice(0, 3))
    }
    if (format.match(/a/)) {
        if (date.getHours() >= 12) {
            result = result.replace(/a/, "PM")
            isPm = true
        } else {
            result = result.replace(/a/, "AM")
        }
    }
    if (format.match(/[Hh][Hh]/)) {
        if (isPm) {
            result = result.replace(/[Hh][Hh]/, String(date.getHours() - 12).padStart(2, "0"))
        } else {
            result = result.replace(/[Hh][Hh]/, String(date.getHours()).padStart(2, "0"))
        }
    } else if (format.match(/[Hh]/)) {
        if (isPm) {
            result = result.replace(/[Hh]/, String(date.getHours() - 12))
        } else {
            result = result.replace(/[Hh]/, String(date.getHours()))
        }

    }
    if (format.match(/mm/)) {
        result = result.replace(/mm/, String(date.getMinutes()).padStart(2, "0"))
    } else if (format.match(/m/)) {
        result = result.replace(/m/, String(date.getMinutes()))
    }
    if (format.match(/ss/)) {
        result = result.replace(/ss/, String(date.getSeconds()).padStart(2, "0"))
    } else if (format.match(/s/)) {
        result = result.replace(/s/, String(date.getSeconds()))
    }
    return result
}
console.log(format(new Date(-585, 4, 28), 'y'))