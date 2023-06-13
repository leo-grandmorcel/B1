function isPositive(number) {
    return number > 0
}

function abs(number) {
    if (number == 0) {
        return number
    }
    if (isPositive(number)) {
        return number
    } else {
        return number *= -1
    }
}