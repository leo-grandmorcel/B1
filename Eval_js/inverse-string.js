function inverseString(str) {
    str = str.split(" ")
    res = ""
    for (let i = str.length - 1; i >= 0; i--) {
        res += str[i] + " "
    }
    return res.slice(0, -1)
}

module.exports = inverseString();