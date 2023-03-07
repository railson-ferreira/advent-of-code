const fs = require('fs');
const rocks = require('./rocks');

// Advent of Code 2022 - day 17

fs.readFile('input.txt', 'utf8', (err, data) => {
    if (err) {
        console.error(err);
        return;
    }
    challenge1(data)
    console.log("=======================")
    challenge2(data)
    // challenge2A(data)
});

function challenge1(data) {
    const jets = [...data.replace("\n", "")]

    const rockShapesNumber = rocks.length;

    const chamber = [["-","-","-","-","-","-","-"]] // chamber's ground

    let jetIndex = 0;
    let stoppedRocksCount = 0;
    let highestRockIndex = 0;
    const flags = []
    for (let i = 0; i < 2022; i++) {
        const rockIndex = i % rockShapesNumber;
        const rock = rocks[rockIndex];
        let rockPosition = {x: 2, y: highestRockIndex + 4};

        let moveTurn = "jet";

        while (true) {
            const jet = jets[jetIndex%jets.length];
            if (moveTurn === "jet") {
                const newRockPosition = getNewRockPosition(rock, rockPosition, jet)
                if (!isRockIntoSomething(rock, newRockPosition, chamber)) {
                    rockPosition = newRockPosition;
                }
                jetIndex++;
                moveTurn = "down";
            } else {
                const newRockPosition = getNewRockPosition(rock, rockPosition, null)
                if (isRockIntoSomething(rock, newRockPosition, chamber)) {
                    addRockToTheChamber(rock, rockPosition, chamber)
                    stoppedRocksCount++;
                    break;
                } else {
                    rockPosition = newRockPosition;
                }
                moveTurn = "jet";
            }
        }
        highestRockIndex = getNewHighestRockIndex(chamber, highestRockIndex)

        if(i === 2021){
            console.log()
            showChamber(chamber, rock, rockPosition);
        }
    }
    console.log({jetIndex, stoppedRocksCount, highestRockIndex})
}


///////////////////// CHALLENGE 2 /////////////////

function challenge2A(data) {
    const jets = [...data.replace("\n", "")]

    const rockShapesNumber = rocks.length;

    const chamber = [["-","-","-","-","-","-","-"]] // chamber's ground

    let jetIndex = 0;
    let stoppedRocksCount = 0;
    let highestRockIndex = 0;
    const flags = []
    for (let i = 0; i < 2022; i++) {
        const rockIndex = i % rockShapesNumber;
        const rock = rocks[rockIndex];
        let rockPosition = {x: 2, y: highestRockIndex + 4};

        let moveTurn = "jet";

        while (true) {
            const jet = jets[jetIndex%jets.length];
            if (moveTurn === "jet") {
                const newRockPosition = getNewRockPosition(rock, rockPosition, jet)
                if (!isRockIntoSomething(rock, newRockPosition, chamber)) {
                    rockPosition = newRockPosition;
                }
                jetIndex++;
                moveTurn = "down";
            } else {
                const newRockPosition = getNewRockPosition(rock, rockPosition, null)
                if (isRockIntoSomething(rock, newRockPosition, chamber)) {
                    addRockToTheChamber(rock, rockPosition, chamber)
                    stoppedRocksCount++;
                    break;
                } else {
                    rockPosition = newRockPosition;
                }
                moveTurn = "jet";
            }
        }
        highestRockIndex = getNewHighestRockIndex(chamber, highestRockIndex)

        if(i === 2021){
            var string = getChamberString(chamber.reverse(), rock, rockPosition);
            debugger
        }
    }
    console.log({jetIndex, stoppedRocksCount, highestRockIndex})
}

function challenge2(data) {
    const jets = [...data.replace("\n", "")]

    const rockShapesNumber = rocks.length;

    const chamber = [["-","-","-","-","-","-","-"]] // chamber's ground

    let jetIndex = 0;
    let stoppedRocksCount = 0;
    let highestRockIndex = 0;
    let lastMemory = {highest: 0, stoppedRocks: 0, jetIndex: 0}
    let skippedRocks = 0;
    let skippedHeight = 0;
    for (let i = 0; i < 1000000000000; i++) {
        const rockIndex = i % rockShapesNumber;
        const rock = rocks[rockIndex];
        let rockPosition = {x: 2, y: highestRockIndex + 4 - skippedHeight};

        let moveTurn = "jet";

        while (true) {
            // console.log()
            // showChamber(chamber, rock, rockPosition);
            const jet = jets[jetIndex%jets.length];
            if (moveTurn === "jet") {
                const newRockPosition = getNewRockPosition(rock, rockPosition, jet)
                if (!isRockIntoSomething(rock, newRockPosition, chamber, skippedRocks)) {
                    rockPosition = newRockPosition;
                }
                jetIndex++;
                moveTurn = "down";
            } else {
                const newRockPosition = getNewRockPosition(rock, rockPosition, null)
                if (isRockIntoSomething(rock, newRockPosition, chamber, skippedRocks)) {
                    addRockToTheChamber(rock, rockPosition, chamber,skippedRocks)
                    stoppedRocksCount++;
                    break;
                } else {
                    rockPosition = newRockPosition;
                }
                moveTurn = "jet";
            }
        }
        highestRockIndex = getNewHighestRockIndex(chamber, highestRockIndex, skippedHeight)
        if(highestRockIndex >= 119){// 👈😁 Found analysing by text editor the result data from the previous challenge
            const keyHeight = 2642;// 👈😁 Found analysing by text editor the result data from the previous challenge
            if((highestRockIndex-119) % keyHeight === 0 && (highestRockIndex-119) / keyHeight <=5){
                const heightDifference = highestRockIndex-lastMemory.highest;
                const stoppedRocksDifference = stoppedRocksCount-lastMemory.stoppedRocks;
                const jetIndexDiff = jetIndex-lastMemory.jetIndex;
                console.log(`highest rock is ${highestRockIndex} with ${stoppedRocksCount} rocks, diff ${stoppedRocksDifference} : ${heightDifference}`)
                console.log(`jet diff ${jetIndexDiff}`)
                lastMemory.highest = highestRockIndex;
                lastMemory.stoppedRocks = stoppedRocksCount;
                lastMemory.jetIndex = jetIndex;
                if(heightDifference === keyHeight){
                    while (1000000000000-stoppedRocksCount>stoppedRocksDifference){
                        jetIndex += jetIndexDiff;
                        i += stoppedRocksDifference;
                        stoppedRocksCount += stoppedRocksDifference;
                        // chamber.length += stoppedRocksDifference;
                        const chamberLength = chamber.length;
                        // chamber.forEach((item,index)=>{
                        //     if(index<chamberLength-100){
                        //     }
                        // })
                        highestRockIndex +=heightDifference
                        if(stoppedRocksCount%1000000000 === 0)
                        console.log(`stopped: ${stoppedRocksCount}`)
                    }
                    skippedRocks = stoppedRocksCount-lastMemory.stoppedRocks
                    skippedHeight = highestRockIndex-lastMemory.highest
                    console.log(`skipped rocks ${stoppedRocksCount-lastMemory.stoppedRocks}. height diff ${highestRockIndex -lastMemory.highest}`)
                    console.log(`remains to calculate ${1000000000000-stoppedRocksCount}`)
                    debugger;
                }
            }
        }
    }
    console.log({jetIndex, stoppedRocksCount, highestRockIndex})
}

// Utils Functions

function showChamber(chamber, rock, rockPosition){
    const rockHeight = rock.length;
    for(let y = chamber.length-1;y>=0;y--){
        let line = "|"
        if(y === 0){
            line = "+"
        }
        const array = [...chamber[y]].map((item,x)=>{
            const rockX = x-rockPosition.x;
            const rockY = rockPosition.y - y + rockHeight - 1;
            if(rock[rockY]?.[rockX] === "#"){
                return "@";
            }else{
                return item;
            }
        });
        line += array.join("")

        if(y === 0){
            line += "+"
        }else{
            line += "|"
        }
        console.log(line)
    }
}

function getNewHighestRockIndex(chamber, oldHighestRockIndex = 0, offsetChange = 0) {
    if(offsetChange){
        debugger
    }
    let highest = oldHighestRockIndex;
    for (let i = highest - offsetChange; i < chamber.length; i++) {
        if (chamber[i].some((item) => item !== ".")) {
            highest = i+offsetChange;
        }
    }
    return highest;
}

function getNewRockPosition(rock, rockPosition, jet) {
    const {x: currentX, y: currentY} = rockPosition;
    switch (jet){
        case ">":
            return {x: currentX+1, y: currentY}
        case "<":
            return {x: currentX-1, y: currentY}
        default:
            return {
                x: currentX,
                y: currentY - 1
            }
    }
}

function isRockIntoSomething(rock, rockPosition, chamber) {
    const rockFilledSpaces = []
    for(let i = rock.length-1; i >=0;i--){ // starts from the bottom
        const invertedIndex = rock.length-1-i;
        const line = rock[i];
        for(let j = 0; j<line.length;j++){
            if(rock[i][j]!=="."){
                rockFilledSpaces.push({
                    x: rockPosition.x+j,
                    y: rockPosition.y+invertedIndex
                })
            }
        }
    }
    for(let i = 0; i < rockFilledSpaces.length ; i++){
     const position = rockFilledSpaces[i];
     const {x,y} = position;
     while (chamber.length-1<y){
         chamber.push(generateChamberEmptyLine())
     }
     if(chamber[y][x]!=="."){
         return true;
     }
    }
    return false;
}

function addRockToTheChamber(rock, rockPosition, chamber) {
    for(let i = rock.length-1; i >=0;i--){ // starts from the bottom
        const invertedIndex = rock.length-1-i;
        const line = rock[i];
        for(let j = 0; j<line.length;j++){
            if(rock[i][j]!=="."){
                const x= rockPosition.x+j;
                const y= rockPosition.y+invertedIndex;
                while (chamber.length-1<y){
                    chamber.push(generateChamberEmptyLine())
                }
                chamber[y][x]="#";
            }
        }
    }
}

function generateChamberEmptyLine(){
    return [".",".",".",".",".",".","."]
}



function getChamberString(chamber, rock, rockPosition){
    const rockHeight = rock.length;
    let body = ""
    for(let y = chamber.length-1;y>=0;y--){
        let line = "|"
        if(y === 0){
            line = "+"
        }
        const array = [...chamber[y]].map((item,x)=>{
            const rockX = x-rockPosition.x;
            const rockY = rockPosition.y - y + rockHeight - 1;
            if(rock[rockY]?.[rockX] === "#"){
                return "@";
            }else{
                return item;
            }
        });
        line += array.join("")

        if(y === 0){
            line += "+"
        }else{
            line += "|"
        }
        body+= line;
        body+= "\n";
    }
    return body
}