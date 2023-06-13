const reverse = obj => Object.fromEntries(Object.entries(obj).map(([k, v]) => [v, k]))

function extractObject(str) {
    return str.match(/{\S*}/g).map(x => reverse(JSON.parse(x)))
}

module.exports = extractObject();