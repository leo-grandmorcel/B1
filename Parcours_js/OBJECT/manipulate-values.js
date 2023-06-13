function filterValues(obj, func) {
    var res = {}
    Object.keys(obj).forEach(key => {
        if (func(obj[key], key, obj)) {
            res[key] = obj[key]
        }
    })
    return res
}
function mapValues(obj, func) {
    var res = {}
    Object.keys(obj).forEach(key => {
        res[key] = func(obj[key], key, obj)
    })
    return res
}
function reduceValues(obj, func, acc = 0) {
    Object.values(obj).forEach(value => {
        acc = func(acc, value)
    })
    return acc
}