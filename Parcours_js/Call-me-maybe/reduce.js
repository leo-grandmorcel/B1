function fold(array, func, acc = 0) {
    for (let i = 0; i < array.length; i++) {
        acc = func(acc, array[i])
    }
    return acc
}
function foldRight(array, func, acc = 0) {
    for (let i = array.length - 1; i >= 0; i--) {
        acc = func(acc, array[i])
    }
    return acc
}

function reduce(array, func, acc) {
    if (typeof array[0] === 'number') {
        acc = 0
    } else {
        acc = ''
    }
    return fold(array, func, acc)
}
function reduceRight(array, func, acc) {
    if (typeof array[0] === 'number') {
        acc = 0
    } else {
        acc = ''
    }
    return foldRight(array, func, acc)
}