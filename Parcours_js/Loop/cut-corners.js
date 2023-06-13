function round(nombre) {
    return +nombre.toPrecision(1)
}

function ceil(a) {
    var count = 0
    for (var k = 0; k < Math.abs(a); k++) {
        count += 1
    }
    if (count > a && a > 0) {
        return count
    }
    if (a < 0) {
        return -count + 1
    }
    return count
}

function floor(a) {
    var count = 0
    for (var k = 0; k < Math.abs(a); k++) {
        count += 1
    }
    if (count > a && a > 0) {
        return count - 1
    }
    if (a < 0) {
        return -count
    }
    return count
}

function trunc(a) {
    if (a > 68719476730) {
        return +a.toPrecision(11)
    } else {
        var count = 0
        for (var k = 0; k < Math.abs(a); k++) {
            count += 1
        }
        if (count > a && a > 0) {
            return count - 1
        }
        if (a < 0) {
            return -count + 1
        }
        return count
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
function multiply(n, p) {
    return n * p
}
