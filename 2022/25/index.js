const fs = require('fs');
const Snafu = require('./snafu');

// Advent of Code 2022 - day 25

fs.readFile('input.txt', 'utf8', (err, data) => {
    if (err) {
        console.error(err);
        return;
    }
    challenge1(data)
    console.log("=======================")
    // challenge2(data)
});


function challenge1(data) {
    const snafus = []
    data.split("\n").filter(item=>item).forEach(snafu=>{
        snafus.push(new Snafu(snafu))
    })
    const decimalSum = snafus.map(snafu=>snafu.toNumber()).reduce((a,b)=>a+b)
    const snafu = Snafu.parse(decimalSum)
    console.log(`${snafu}`)
}


///////////////////// CHALLENGE 2 /////////////////

function challenge2(data) {
    // "Only 49 stars to go."
}
