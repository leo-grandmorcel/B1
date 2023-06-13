function findExpression(target) {
    function find(current, history) {
        if (current == target) {
            return history;
        } else if (current > target) {
            return undefined;
        } else {
            return find(current + 4, `${history} ${add4}`) || find(current * 2, `${history} ${mul2}`);
        }
    }
    return find(1, "1");
}
