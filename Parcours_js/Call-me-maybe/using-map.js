function citiesOnly(array) {
    return array.map(obj => obj['city'])
}
function temperaturesOnly(array) {
    return array.map(obj => obj['temperature'])
}
function statesOnly(array) {
    return array.map(obj => obj['state'])
}
function upperCasingStates(array) {
    return array.map(x => (x.split(" ")).map(y => capitalize(y)).join(" "))
}
function fahrenheitToCelsius(array) {
    return array.map(x => Math.floor((parseInt(x.slice(0, -2)) - 32) / 1.8) + "°C")
}

function trimTemp(array) {
    return array.map(obj => trim(obj))
}
function trim(obj) {
    obj['temperature'] = obj['temperature'].replace(/ /g, "")
    return obj
}

function tempForecasts(array) {
    var cities = upperCasingStates(citiesOnly(array))
    var temperatures = fahrenheitToCelsius(temperaturesOnly(trimTemp(array)))
    var states = upperCasingStates(statesOnly(array))
    var result = []
    for (let i = 0; i < array.length; i++) {
        result.push(temperatures[i] + "elsius in " + cities[i] + ", " + states[i])
    }
    return result
}
function capitalize(n) {
    return n[0].toUpperCase() + n.slice(1).toLowerCase()
}
