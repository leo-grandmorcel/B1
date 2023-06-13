function forEach(array, func) {
    var res = []
    for (let i = 0; i < array.length; i++) {
        res.push(func(array[i], i, array))
    }
    return res
}