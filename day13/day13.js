const fs = require("fs")

const input = fs.readFileSync('input.txt', 'utf8')

const lines = input.split('\n')

const indices = []

const dividers = [[2], [6]]
const packets = [...dividers]

for (let i = 0; i < lines.length; i += 3) {
    const first = JSON.parse(lines[i])
    const second = JSON.parse(lines[i+1])

    if (compare(first,second) < 0) {
        indices.push((i/3)+1)
    }

    packets.push(first)
    packets.push(second)
}

const sum = indices.reduce((acc, i) => acc + i, 0)
console.log("Sum: " + sum)

packets.sort(compare)
const decoderKey =
    (packets.indexOf(dividers[0])+1) *
    (packets.indexOf(dividers[1])+1)
console.log("Decoder key: " + decoderKey)

function compare(a, b) {
    for (let i = 0; i < a.length; i++) {
        if (i >= b.length)
            return 1

        if (typeof a[i] === "number" && typeof b[i] === "number") {
            if (a[i] < b[i])
                return -1
            if (a[i] > b[i])
                return 1

            continue
        }

        if (typeof a[i] === "number") {
            const comp = compare([a[i]], b[i])
            if (comp === 0) continue
            else return comp
        }

        if (typeof b[i] === "number") {
            const comp = compare(a[i], [b[i]])
            if (comp === 0) continue
            else return comp
        }

        const comp = compare(a[i], b[i])
        if (comp !== 0) return comp
    }

    return a.length < b.length ? -1 : 0
}
