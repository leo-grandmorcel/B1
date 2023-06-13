const reg = /[a-zA-Z]*tion\b/g

function ionOut(string) {
    var array = string.match(reg) || []
    for (let i = 0; i < array.length; i++) {
        array[i] = array[i].slice(0, -3)
    }
    return array
}