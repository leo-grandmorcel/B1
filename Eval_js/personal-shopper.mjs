import * as fs from 'node:fs'
var args = process.argv.slice(2);
console.log(add(args[0], args[1], args[2]))



function crea(file) {
    fs.open(file, 'wx', (err, fd) => {
        if (err) throw err;
        fs.close(fd, (err) => {
            if (err) throw err;
        })
    })
}

function del(file) {
    fs.unlink(file)
}
function add(file, key, value) {
    var data = ""
    fs.readFile(file, 'utf8', (err, data) => {
        console.log(JSON.parse(data));
    })
    if (key in data) {
        data[key] += value
    } else {
        data[key] = value
    }
    string = JSON.stringify(data)
    fs.writeFile(file, string,
        {
            encoding: "utf8",
            flag: "w",
            mode: 0o666
        }
    )
}

function red(file) {
    fs.readFile(file, 'utf8', (err, data) => {
        console.log(JSON.parse(data));
    })
}