function pyramid(str, int) {
    var res = ""
    var space = ""
    for (var i = 0; i < str.length; i++) {
        space += " "
    }
    for (var i = 0; i < int; i++) {
        for (var j = 1; j < int - i; j++) {
            res += space;
        }
        for (var k = 1; k <= (2 * i + 1); k++) {
            res += str;
        }
        res += "\n"
    }
    return res.slice(0, -1)
}