const fs = require('fs');

// Advent of Code 2022 - day 13

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
    const structuredInput = []
    data.split("\n").filter(item=>item).forEach(line=>{
        structuredInput.push(JSON.parse(line))
    })
    const correctOrderIndexes = []
    for(let x = 1 ; x<structuredInput.length;x=x+2){
        const leftItem = structuredInput[x-1];
        const rightItem = structuredInput[x];
        if(isInCorrectOrder(leftItem,rightItem)){
            correctOrderIndexes.push((x+1)/2);
        }
    }
    console.log(correctOrderIndexes.reduce((a,b)=>a+b))
}

///////////////////// CHALLENGE 2 /////////////////

function challenge2(data) {
    const structuredInput = []
    data.split("\n").filter(item=>item).forEach(line=>{
        structuredInput.push(JSON.parse(line))
    })
    const dividerPackageA = [[2]]
    const dividerPackageB = [[6]]
    const orderedInputWithDiverPackages = [...structuredInput,dividerPackageA,dividerPackageB]
    orderedInputWithDiverPackages.sort((a,b)=>isInCorrectOrder(a,b)?-1:1);
    const packageAIndex = orderedInputWithDiverPackages.findIndex(signalPackage=>signalPackage === dividerPackageA)
    const packageBIndex = orderedInputWithDiverPackages.findIndex(signalPackage=>signalPackage === dividerPackageB)
    const result = (packageAIndex+1)*(packageBIndex+1)
    console.log(result)
}

// Utils Functions

function isInCorrectOrder(leftItem, rightItem){
    const leftJson = JSON.stringify(leftItem);
    const rightJson = JSON.stringify(rightItem);

    if(typeof leftItem === "number" && typeof rightItem === "number"){
        if(leftItem === rightItem){
            return null
        }
        return leftItem < rightItem
    }else{
        const leftArray = typeof leftItem === "number" ?[leftItem]:leftItem
        const rightArray = typeof rightItem === "number" ?[rightItem]:rightItem
        const maxLength = Math.max(leftArray.length,rightArray.length);
        for(let i = 0 ; i <maxLength ; i++){
            const leftSubItem = leftArray[i];
            const rightSubItem = rightArray[i];
            if(leftSubItem === undefined){
                return true;
            }
            if(rightSubItem === undefined){
                return false;
            }
            const result = isInCorrectOrder(leftSubItem,rightSubItem)
            if(result !== null){
                return result;
            }
        }
    }
    return null
}