function RNA(ADN) {
    return ADN.replace(/G/g, "W").replace(/C/g, "G").replace(/W/g, "C").replace(/T/g, "W").replace(/A/g, "U").replace(/W/g, "A")
}
function DNA(ADN) {
    return ADN.replace(/C/g, "W").replace(/G/g, "C").replace(/W/g, "G").replace(/A/g, "W").replace(/U/g, "A").replace(/W/g, "T")
}
