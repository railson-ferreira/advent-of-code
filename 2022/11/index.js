const fs = require('fs');
const parseData = require("./parse-data")

// Advent of Code 2022 - day 11

fs.readFile('input.txt', 'utf8', (err, data) => {
    if (err) {
        console.error(err);
        return;
    }
    challenge1(data)
    console.log("=======================")
    challenge2(data)
});


function challenge1(challengeData) {
    const monkeysData = parseData(challengeData);
    const roundsNumber = 20;
    let roundsCount = 0;
    while (roundsCount < roundsNumber) {
        roundsCount++;
        monkeysData.forEach(data => {
            //e.g.
            //Monkey 0:
            //   Monkey inspects an item with a worry level of 79.
            //     Worry level is multiplied by 19 to 1501.
            //     Monkey gets bored with item. Worry level is divided by 3 to 500.
            //     Current worry level is not divisible by 23.
            //     Item with worry level 500 is thrown to monkey 3.
            //   Monkey inspects an item with a worry level of 98.
            //     Worry level is multiplied by 19 to 1862.
            //     Monkey gets bored with item. Worry level is divided by 3 to 620.
            //     Current worry level is not divisible by 23.
            //     Item with worry level 620 is thrown to monkey 3.

            // console.log(`Monkey ${data.monkeyId}`)
            data.items.forEach((item, index) => {
                if (!data.inspectionTimes) {
                    data.inspectionTimes = 0
                }
                data.inspectionTimes++;
                const newWorryLevel = data.operation(item);
                const newestWorryLevel = Number(Math.floor(Number(newWorryLevel) / 3));
                const isDivisible = newestWorryLevel % data.testDivider === 0;
                const monkeyIdWhomThrowsTo = isDivisible ? data.monkeyIdIfTrue : data.monkeyIdIfFalse;
                monkeysData.find(monkeyData => monkeyData.monkeyId === monkeyIdWhomThrowsTo).items.push(newestWorryLevel)
                delete data.items[index];
                // console.log("  ",`Monkey inspects an item with a worry level of ${item}.`)
                // console.log("   ",`Worry level is changed from ${item} to ${newWorryLevel}.`)
                // console.log("   ",`Monkey gets bored with item. Worry level is divided by 3 to ${newestWorryLevel}.`)
                // console.log("   ",`Current worry level is ${isDivisible?"not ": ""} divisible by ${data.testDivider}.`)
                // console.log("   ",`Item with worry level ${newestWorryLevel} is thrown to monkey ${monkeyIdWhomThrowsTo}.`)
            })
        })
    }
    const twoMostActiveMonkeys = monkeysData.sort((a, b) => {
        return b.inspectionTimes - a.inspectionTimes
    }).slice(0, 2)
    // twoMostActiveMonkeys.forEach(monkeyData=>{
    //     console.log(`Monkey ${monkeyData.monkeyId} inspected items ${monkeyData.inspectionTimes} times`)
    // })
    console.log(twoMostActiveMonkeys.reduce((a, b) => a.inspectionTimes * b.inspectionTimes))
}

///////////////////// CHALLENGE 2 /////////////////

function challenge2(challengeData) {
    const monkeysData = parseData(challengeData);
    const roundsNumber = 10000;
    let roundsCount = 0;
    const monkeysDividersMmc = mmc(monkeysData.map(monkeyData => monkeyData.testDivider))
    while (roundsCount < roundsNumber) {
        roundsCount++;
        monkeysData.forEach(data => {
            data.items.forEach((item, index) => {
                const localRoundsCount = roundsCount;
                if (!data.inspectionTimes) {
                    data.inspectionTimes = 0
                }
                data.inspectionTimes++;
                let newWorryLevel = data.operation(item);
                // console.log(newWorryLevel%monkeysDividersMmc)
                if (newWorryLevel >= 2 * monkeysDividersMmc) {
                    const multiplier = Math.floor(newWorryLevel / monkeysDividersMmc)
                    newWorryLevel = newWorryLevel - monkeysDividersMmc * multiplier;
                }
                const isDivisible = newWorryLevel % data.testDivider === 0;
                const monkeyIdWhomThrowsTo = isDivisible ? data.monkeyIdIfTrue : data.monkeyIdIfFalse;
                monkeysData.find(monkeyData => monkeyData.monkeyId === monkeyIdWhomThrowsTo).items.push(newWorryLevel)
                delete data.items[index];
            })
            data.items = data.items.filter(item => item);
        })
        // if(roundsCount === 1 || roundsCount === 20 ||
        //     roundsCount % 1000 === 0){
        //     console.log(`== After round ${roundsCount} ==`);
        //     monkeysData.forEach(monkeyData=>{
        //         console.log(`Monkey ${monkeyData.monkeyId} inspected items ${monkeyData.inspectionTimes} times`)
        //     })
        // }
    }
    const twoMostActiveMonkeys = monkeysData.sort((a, b) => {
        return b.inspectionTimes - a.inspectionTimes
    }).slice(0, 2)
    console.log(twoMostActiveMonkeys.reduce((a, b) => a.inspectionTimes * b.inspectionTimes))
}

function divisibleAll(numbers, multiple) {

    for (let i = 0; i < numbers.length; i++) {

        if (multiple % numbers[i] !== 0) {

            return false;
        }
    }

    return true;
}

function mmc(numbers) {

    numbers.sort((a, b) => {

        if (a > b) {

            return -1;
        }

        if (a < b) {

            return 1;
        }

        return 0;
    });

    numbers = Array.from(new Set(numbers));

    var greather = numbers.shift();
    var i = 1;

    while (true) {

        let multiple = greather * i;

        if (divisibleAll(numbers, multiple)) {

            return multiple;
        }

        i++;
    }
}