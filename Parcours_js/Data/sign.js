function sign(number) {
    if (number === 0) {
        return 0
    }
    if (number > 0) {
        return 1
    } else {
        return -1
    }
}

function sameSign(n, p) {
    return (sign(n) === sign(p))
}