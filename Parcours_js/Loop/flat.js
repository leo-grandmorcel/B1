function flat(array, int = 1) {
    if (int === Infinity) {
        int = 10
    }
    while (int > 0) {
        array = [].concat.apply([], array)
        int -= 1
    }
    return array
}
