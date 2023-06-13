function triangle(str, int) {
    var res = ""
    for (let i = 1; i <= int; i++) {
        for (let y = 0; y < i; y++) {
            res += str
        }
        res += "\n"
    }
    return res.slice(0, -1)
}