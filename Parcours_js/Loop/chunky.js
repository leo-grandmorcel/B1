function chunk(array, size) {
    var result = []
    for (let i = 0; i < array.length; i += size) {
        result.push(array.slice(i, i + size))
    }
    return result
}
