function map(array, func) {
    var res = []
    for (let i = 0; i < array.length; i++) {
        res.push(func(array[i], i, array))
    }
    return res
}

function flatMap(array, func) {
    var res = []
    array = map(array, func)
    for (let map of array) {
        if (Array.isArray(map)) {
            for (let m of map) {
                res.push(m)
            }
        } else {
            res.push(map)
        }
    }
    return res
}