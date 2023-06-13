import { styles } from "./pimp-my-style.data.js"

var click = 0

export const pimp = () => {
    var obj = Array.from(document.getElementsByClassName("button"))[0]
    if (obj.classList.contains("unpimp")) {
        obj.classList.remove(styles[click - 1])
        click -= 1
    } else {
        obj.classList.add(styles[click])
        click += 1
    }
    if (click == styles.length || (click == 0 && obj.classList.contains("unpimp"))) {
        obj.classList.toggle("unpimp")
    }

}