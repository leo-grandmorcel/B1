function adder(array, x = 0) {
    return array.reduce((previousValue, currentValue) => previousValue + currentValue, x)
}

function sumOrMul(array, x = 0) {
    return array.reduce((previousValue, currentValue) => (currentValue % 2 == 0 ? previousValue * currentValue : previousValue + currentValue), x)
}

function funcExec(array, x) {
    return array.reduce((previous, current) => current(previous), x)
}