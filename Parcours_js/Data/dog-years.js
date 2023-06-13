function dogYears(planet, second) {
    var terre = 31557600;
    switch (planet) {
        case "earth":
            return parseFloat(((second / terre) * 7).toFixed(2))
        case "mercury":
            return parseFloat((second / (terre * 0.2408467) * 7).toFixed(2))
        case "venus":
            return parseFloat((second / (terre * 0.61519726) * 7).toFixed(2))
        case "mars":
            return parseFloat((second / (terre * 1.8808158) * 7).toFixed(2))
        case "jupiter":
            return parseFloat((second / (terre * 11.862615) * 7).toFixed(2))
        case "saturn":
            return parseFloat((second / (terre * 29.447498) * 7).toFixed(2))
        case "uranus":
            return parseFloat((second / (terre * 84.016846) * 7).toFixed(2))
        case "neptune":
            return parseFloat((second / (terre * 164.79132) * 7).toFixed(2))
    }
}