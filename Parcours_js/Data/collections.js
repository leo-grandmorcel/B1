function arrToSet(n) {
    return new Set(n)
}
function arrToStr(n) {
    return n.join("")
}
function setToArr(n) {
    return Array.from(n)
}
function setToStr(n) {
    n = Array.from(n)
    return arrToStr(n)
}
function strToArr(n) {
    return n.split("")
}
function strToSet(n) {
    n = strToArr(n)
    return arrToSet(n)
}
function mapToObj(n) {
    return Object.fromEntries(n)
}
function objToArr(n) {
    return Object.values(n)
}
function objToMap(n) {
    return new Map(Object.entries(n))
}
function arrToObj(n) {
    const obj = {}
    for (let i = 0; i < n.length; i++) {
        obj[i] = n[i]
    }
    return obj
}
function strToObj(n) {
    return arrToObj(strToArr(n))
}
function capitalize(n) {
    return n[0].toUpperCase() + n.slice(1).toLowerCase()
}

function superTypeOf(n) {
    if (n === null) {
        return "null"
    } else if (n === undefined) {
        return "undefined"
    }
    if (Object.getPrototypeOf(n) === Map.prototype) {
        return "Map"
    } else if (Object.getPrototypeOf(n) === Set.prototype) {
        return "Set"
    } else if (Object.getPrototypeOf(n) === Array.prototype) {
        return "Array"
    } else {
        return capitalize(typeof n)
    }
}
