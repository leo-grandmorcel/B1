export const getArchitects = () => {
    return [Array.from(document.getElementsByTagName("a")), Array.from(document.querySelectorAll("body :not(a)"))]
}
export const getClassical = () => {
    console.log(document.getElementsByClassName("classical"))
    return [Array.from(document.getElementsByClassName("classical")), Array.from(document.querySelectorAll("body :not(.classical)"))]
}
export const getActive = () => {
    return [Array.from(document.querySelectorAll(".active")), Array.from(document.querySelectorAll("body :not(.active)"))]
}

export const getBonannoPisano = () => {
    return [document.getElementById("BonannoPisano"), Array.from(document.querySelectorAll('a:not(#BonannoPisano)'))]
}