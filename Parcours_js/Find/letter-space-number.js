function letterSpaceNumber(string) {
    const reg = /[a-zA-Z]\s\d\b/g
    var array = string.match(reg) || []
    return array
}
