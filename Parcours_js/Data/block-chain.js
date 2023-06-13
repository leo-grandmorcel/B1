function blockChain(data, prev = { index: 0, hash: '0' }) {
    const obj = {}
    obj.index = prev.index + 1
    obj.hash = hashCode(obj.index + prev.hash + JSON.stringify(data))
    obj.data = data
    obj.prev = prev
    obj.chain = function c(n) {
        return blockChain(n, obj)
    }
    return obj
}