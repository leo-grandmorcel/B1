const regPrice = /((USD)|(\$))\d*.\d*/g

function groupPrice(string) {
    var res = []
    var array = string.match(regPrice) || []
    for (let i = 0; i < array.length; i++) {
        var temp = array[i].match(/\d{1,}/g)
        var tab = []
        tab.push(array[i])
        tab.push(temp[0])
        tab.push(temp[1])
        res.push(tab)
    }
    return res
}