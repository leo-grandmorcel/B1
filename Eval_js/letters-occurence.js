function lettersOccurence(str) {
    var obj = {}
    for (let i = 0; i < str.length; i++) {
        if (obj[str[i]]) {
            obj[str[i]] += 1
        } else {
            obj[str[i]] = 1
        }
    }
    return obj
}

module.exports = lettersOccurence();