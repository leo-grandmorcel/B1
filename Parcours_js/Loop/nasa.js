function nasa(n) {
    var res = ""
    for (let i = 1; i <= n; i++) {

        if (i % 3 == 0 && i % 5 == 0) {
            res += "NASA "
        } else if (i % 3 == 0) {
            res += "NA "
        } else if (i % 5 == 0) {
            res += "SA "
        } else {
            res += String(i)
            res += " "
        }
    }
    return res.slice(0, -1)
}
