export const build = (x) => {
    var start = 1
    var middle = 1

    let interval = setInterval(function () {
        const newBrick = document.createElement("div")
        if (start == x) {
            clearInterval(interval)
        }
        newBrick.setAttribute("id", "brick-" + start)
        if (middle == 2) {
            newBrick.dataset.foundation = true
            middle = 0
        } else {
            middle += 1
        }
        document.body.append(newBrick)
        start += 1
    }, 100)

}

export const repair = (...ids) => {
    ids.forEach(id => {
        var element = document.getElementById(id)
        if (element.hasAttribute("foundation", true)) {
            element.dataset.repaired = 'in progress'
        } else {
            element.dataset.repaired = true
        }
    })
}

export const destroy = () => {
    Array.from(document.getElementsByTagName("div")).slice(-1)[0].remove()
}