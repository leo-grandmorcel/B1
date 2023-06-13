const regURL = /(http|ftp|https):\/\/([\w_-]+(?:(?:[\w_-]+)+))([\w.,@?^=%&:\/\[\]~+#-]*[\w@?^=%&\/~+#-\{])/g
const regQuery = /((\?|&)[a-z]{1,}=[\w.,@?^%:\/\[\]~+#-]{1,}){3,}/

function getURL(url) {
    return url.match(regURL)
}

function greedyQuery(url) {
    var array = getURL(url)
    var res = []
    for (let i = 0; i < array.length; i++) {
        var temp = array[i].match(/=/g) || []
        if (temp.length >= 3) {
            res.push(array[i])
        }
    }
    return res
}

function notSoGreedy(url) {
    var array = getURL(url)
    var res = []
    for (let i = 0; i < array.length; i++) {
        var temp = array[i].match(/=/g) || []
        if (temp.length >= 2 && temp.length <= 3) {
            res.push(array[i])
        }
    }
    return res
}