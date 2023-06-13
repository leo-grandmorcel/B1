function slice(value, first = 0, last = value.length) {
    if (first < 0) {
        first = value.length + first
    }
    if (last < 0) {
        last = value.length + last
    }
    if (Array.isArray(value)) {
        var res = []
        for (let i = first; i < last; i++) {
            res.push(value[i])
        }
    } else {
        var res = ""
        for (let i = first; i < last; i++) {
            res += value[i]
        }
    }
    return res
}