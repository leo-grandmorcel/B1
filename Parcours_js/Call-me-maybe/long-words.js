function longWords(array) {
    return array.every((currentValue) => currentValue.length >= 5)
}

function oneLongWord(array) {
    return array.some((currentValue) => currentValue.length >= 10)
}

function noLongWords(array) {
    return array.every((currentValue) => currentValue.length < 7)
}