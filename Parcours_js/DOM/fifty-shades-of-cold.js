import { colors } from "./fifty-shades-of-cold.data.js"

export const generateClasses = () => {
    var style = document.createElement("style")
    colors.forEach(color => {
        var test = '.' + color + '{background:' + color + ';}'
        style.append(test)
    });
    document.head.append(style)
}

export const generateColdShades = () => {
    colors.forEach(color => {
        if (color.match(/(aqua|blue|turquoise|green|cyan|navy|purple)/g)) {
            var div = document.createElement("div")
            div.classList.add(color)
            div.textContent = color
            document.body.append(div)
        }
    })
}

export const choseShade = (shade) => {
    var elems = document.querySelectorAll('div')
    elems.forEach(elem => {
        let style = elem.className
        elem.classList.replace(style, shade)
    });
}