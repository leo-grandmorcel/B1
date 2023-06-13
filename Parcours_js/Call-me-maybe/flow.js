function flow(array, a = 0) {
    for (let i = 0; i < array.length; i++) {
        a = array[i](a)
    }
    return a
}