export const generateLetters = () => {
    const Letters = ["A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"]
    function getRandomLetter() {
        return Letters[Math.floor(Math.random() * 26)];
    }
    var cpt = 1
    for (let i = 11; i <= 130; i++) {
        const Letter = document.createElement("div")
        Letter.textContent = getRandomLetter()

        if (i < 51) {
            Letter.style.fontWeight = 300
        } else if (i < 91) {
            Letter.style.fontWeight = 400
        } else {
            Letter.style.fontWeight = 600
        }
        Letter.style.fontSize = i + "px"
        document.body.append(Letter)
    }
}