function invert(obj) {
    return Object.fromEntries(Object.entries(obj).map(([k, v]) => [v, k]))
}