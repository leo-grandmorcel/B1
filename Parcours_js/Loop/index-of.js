function indexOf(array, value, index = 0) {
    for (let i = index; i < array.length; i++) {
        if (array[i] == value) {
            return i
        }
    }
    return -1
}

function lastIndexOf(array, value, index = array.length - 1) {
    for (let i = index; i >= 0; i--) {
        if (array[i] == value) {
            return i
        }
    }
    return -1
}

function includes(array, value, index = 0) {
    if (indexOf(array, value, index) != -1) {
        return true
    } else {
        return false
    }
}