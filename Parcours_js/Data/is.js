

is["num"] = function num(n) {
    return (typeof (n) === 'number')
}
is["nan"] = function nan(n) {
    return n !== n
}
is["str"] = function str(n) {
    return (typeof (n) === "string")
}
is["bool"] = function bool(n) {
    return (typeof (n) === 'boolean')
}

is['undef'] = function type(n) {
    return typeof (n) == 'undefined'
}
is["def"] = function def(n) {
    return (typeof (n) != 'undefined')
}
is["arr"] = function arr(n) {
    return Array.isArray(n)
}
is["obj"] = function obj(n) {
    return (Object.prototype.toString.call(n) === '[object Object]')
}

is["fun"] = function fun(n) {
    return (typeof (n) == 'function')
}

is["truthy"] = function truthy(n) {
    return Boolean(n) == true
}
is["falsy"] = function falsy(n) {
    return Boolean(n) == false
}





