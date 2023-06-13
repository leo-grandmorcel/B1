function filterKeys(obj, func) {
    var res = {}
    Object.keys(obj).forEach(key => {
        if (func(key)) {
            res[key] = obj[key]
        }
    })
    return res
}

function mapKeys(obj, func) {
    var res = {}
    Object.keys(obj).forEach(key => {
        res[func(key)] = obj[key]
    })
    return res
}

function reduceKeys(obj, func, acc = '') {
    Object.keys(obj).forEach(key => {
        if (acc == '') {
            acc = key
        } else {
            acc = func(acc, key)
        }

    })
    return acc
}