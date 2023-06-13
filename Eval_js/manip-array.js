function manipArray(array) {
    if (array.length == 0) {
        return []
    } else {
        array = array.filter(x => typeof x == 'number')
        return [array.map(currentValue => currentValue * 2), array.map(currentValue => currentValue - 5), array.filter(x => x > 5)]
    }
}
module.exports = manipArray();