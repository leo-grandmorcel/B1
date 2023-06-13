function filterShortStateName(array) {
    return array.filter(x => x.length < 7)
}

function filterStartVowel(array) {
    return array.filter(x => ["a", "e", "i", "o", "u", "A", "E", "I", "O", "U"].includes(x[0]))
}
function filter5Vowels(array) {
    return array.filter(x => x.match(/[aeiou]/gi).length >= 5)
}
function filter1DistinctVowel(array) {
    return array.filter(x => [...new Set(x.toLowerCase().match(/[aeiou]/g))].length <= 1)
}
function multiFilter(array) {
    return array.filter(x => x['capital'].length >= 8).filter(x => !(["a", "e", "i", "o", "u", "A", "E", "I", "O", "U"].includes(x["name"][0]))).filter(x => [... new Set(x['tag'].match(/[aeiou]/gi))].length >= 1).filter(x => x['region'] != 'South')
}
