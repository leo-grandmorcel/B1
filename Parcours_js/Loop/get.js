function get(src, path) {
    path = path.split(".")
    for (let i = 0; i < path.length; i++) {
        if (!(path[i] in src)) {
            return undefined
        }
        src = src[path[i]]
    }
    return src
}