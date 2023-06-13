function split(str, delimiter = ",", removeEmptyItems) {
    var byp = []
    if (!delimiter || delimiter.length === 0) {
        for (var k = 0; k < str.length; k++) {
            byp.push(str[k])
        }
        return byp
    }
    if (!str || str.length === 0) return [''];
    var result = [];
    var j = 0;
    var lastStart = 0;
    for (var i = 0; i <= str.length;) {
        if (i == str.length || str.substr(i, delimiter.length) == delimiter) {
            if (!removeEmptyItems || lastStart != i) {
                result[j++] = str.substr(lastStart, i - lastStart);
            }
            lastStart = i + delimiter.length;
            i += delimiter.length;
        } else i++;
    }
    return result;
}

function join(array, split = ",") {
    var res = ""
    for (let i = 0; i < array.length; i++) {
        res += array[i]
        res += split
    }
    return res.slice(0, -split.length)
}
