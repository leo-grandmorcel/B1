
function sameAmount(string, regex1, regex2) {
    regex1 = new RegExp(regex1, "g")
    regex2 = new RegExp(regex2, "g")
    var array1 = [...string.matchAll(regex1)] || []
    var array2 = [...string.matchAll(regex2)] || []
    console.log(regex1, regex2, array1, array2)
    return array1.length == array2.length
}
