function getAcceleration(obj) {
    if (obj.f && obj.m) {
        return obj.f / obj.m
    } else if (obj.Δv && obj.Δt) {
        return obj.Δv / obj.Δt
    } else if (obj.d && obj.t) {
        return 2 * obj.d / obj.t ** 2
    } else {
        return "impossible"
    }
}