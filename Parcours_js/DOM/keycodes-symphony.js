export const compose = () => {
    document.addEventListener("keydown", (event) => {
        let asci = event.keyCode
        if (asci == 27) {
            document.querySelectorAll("div").forEach((div) => {
                div.remove()
            })
        } else if (asci == 8) {
            let divs = document.querySelectorAll("div")
            divs[divs.length - 1].remove()
        } else {
            let div = document.createElement("div")
            div.classList.add("note")
            div.style.background = "#50" + asci + asci
            document.body.append(div)
            div.innerHTML = event.key
        }
    })
}