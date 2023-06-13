function filter(array, func) {
    var res = []
    for (let i = 0; i < array.length; i++) {
        if (func(array[i], i, array)) {
            res.push(array[i])
        }
    }
    return res
}

function reject(array, func) {
    var res = []
    for (let i = 0; i < array.length; i++) {
        if (!func(array[i], i, array)) {
            res.push(array[i])
        }
    }
    return res
}

function partition(array, func) {
    return [filter(array, func), reject(array, func)]
}
