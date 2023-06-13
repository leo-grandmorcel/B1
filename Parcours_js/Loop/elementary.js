function multiply(n, p) {
    var res = 0
    for (let i = 0; i < Math.abs(p); i++) {
        res += n
    }
    if (n < 0 && p < 0) {
        return Math.abs(res)
    } else if (p < 0) {
        return -res
    } else {
        return res
    }
}

function divide(a, b) {
    var cp = 0
    var res = Math.abs(a)
    while (res >= Math.abs(b)) {
        res -= Math.abs(b)
        cp += 1
    }
    if (a < 0 && b < 0) {
        return cp
    } else if (a < 0 || b < 0) {
        return -cp
    } else {
        return cp
    }
}

function modulo(a, b) {
    var clone1 = divide(a, b)
    var clone2 = multiply(clone1, b)
    return a - clone2
}

console.log(modulo(-20, 5))