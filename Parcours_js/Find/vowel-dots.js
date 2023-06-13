const vowels = /[aeiouAEIOU]/

function vowelDots(string) {
    var result = ""
    for (let i = 0; i < string.length; i++) {
        result += string[i]
        if (string[i].match(vowels)) {
            result += "."
        }
    }
    return result
}