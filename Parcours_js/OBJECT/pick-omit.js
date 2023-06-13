function pick(obj, arr) {
    var newobj = {}
    if (typeof arr == 'string') {
        arr = [arr]
    }
    Object.keys(obj).forEach(key => {
        if (arr.includes(key)) {
            newobj[key] = obj[key]
        }
    })
    return newobj
}

function omit(obj, arr) {
    var newobj = {}
    if (typeof arr == 'string') {
        arr = [arr]
    }
    Object.keys(obj).forEach(key => {
        if (!arr.includes(key)) {
            newobj[key] = obj[key]
        }
    })
    return newobj
}