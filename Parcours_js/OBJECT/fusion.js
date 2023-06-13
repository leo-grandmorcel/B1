function fusion(obj1, obj2) {
    var newobj = {}
    var tmp1 = {}
    Object.keys(obj1).forEach(key => {
        if (Array.isArray(obj1[key])) {
            newobj[key] = obj1[key]
        } else if (typeof (obj1[key]) === "string") {
            newobj[key] = obj1[key]
        } else if (typeof (obj1[key]) === 'number') {
            newobj[key] = obj1[key]
        } else if (Object.prototype.toString.call(obj1[key]) === '[object Object]') {
            tmp1 = obj1[key]
        }
    })
    Object.keys(obj2).forEach(key => {
        if (Array.isArray(obj2[key])) {
            if (typeof newobj[key] != typeof obj2[key]) {
                newobj[key] = obj2[key]
            } else {
                obj2[key].forEach(elem => {
                    newobj[key].push(elem)
                })
            }
        } else if (typeof (obj2[key]) === "string") {
            if (typeof newobj[key] != typeof obj2[key]) {
                newobj[key] = obj2[key]
            } else {
                newobj[key] += ' ' + obj2[key]
            }
        } else if (typeof (obj2[key]) === 'number') {
            if (typeof newobj[key] != typeof obj2[key]) {
                newobj[key] = obj2[key]
            } else {
                newobj[key] += obj2[key]
            }
        } else if (Object.prototype.toString.call(obj2[key]) === '[object Object]') {
            newobj[key] = fusion(tmp1, obj2[key])
        }
    })
    return newobj
}

console.log(fusion({ a: 1 }, { a: { b: 1 } }))