function reverse(string) {
    if (Array.isArray(string)) {
        var res = []
        for (let i = string.length - 1; i >= 0; i--) {
            res.push(string[i])
        }
    } else {
        var res = ""
        for (let i = string.length - 1; i >= 0; i--) {
            res += string[i]
        }
    }
    return res
}