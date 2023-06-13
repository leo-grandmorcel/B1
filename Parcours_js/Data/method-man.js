function words(n) {
    return n.split(" ")
}

function sentence(n) {
    return n.join(" ")
}

function yell(n) {
    return n.toUpperCase()
}

function whisper(n) {
    return "*" + n.toLowerCase() + "*"
}

function capitalize(n) {
    return n[0].toUpperCase() + n.slice(1).toLowerCase()
}