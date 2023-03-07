const fs = require('fs');

// Advent of Code 2022 - day 4

fs.readFile('input.txt', 'utf8', (err, data) => {
    if (err) {
        console.error(err);
        return;
    }
    challenge1(data)
    console.log("=======================")
    challenge2(data)
});


function challenge1(data) {
    const lines = data.split("\n");
    let fullyContainedCount = 0;
    lines.filter(item => item).forEach(line => {
        const [firstRange, secondRange] = line
            .split(",")
            .map(range => {
                const [start, end] = range.split("-")
                return {start:Number(start), end:Number(end)}
            });
        if (
            firstRange.start >= secondRange.start && firstRange.end <= secondRange.end
            ||
            secondRange.start >= firstRange.start && secondRange.end <= firstRange.end
        ) {
            fullyContainedCount++;
        }

    })
    console.log(fullyContainedCount)
}

///////////////////// CHALlENGE 2 /////////////////

function challenge2(data) {
    const lines = data.split("\n");
    let fullyContainedCount = 0;
    lines.filter(item => item).forEach(line => {
        const [firstRange, secondRange] = line
            .split(",")
            .map(range => {
                const [start, end] = range.split("-")
                return {start:Number(start), end:Number(end)}
            });
        if (
            firstRange.start <= secondRange.end && firstRange.end >= secondRange.start
        ) {
            fullyContainedCount++;
        }

    })
    console.log(fullyContainedCount)
}

