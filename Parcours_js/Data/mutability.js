let clone1 = Object.assign({}, person)
let clone2 = Object.assign({}, person)

var samePerson = person
person.age += 1;
person.country = "FR"
console.log(samePerson, person)
