var escapeStr = "\`\\\/\"\'";
var arr = [4, "2"];

const obj = {
    str: "t",
    num: 1,
    bool: true,
    undef: undefined
}
const nested = {
    arr: [4, undefined, "2"],
    obj
}
Object.freeze(obj)
Object.freeze(nested)
Object.freeze(nested.arr)
