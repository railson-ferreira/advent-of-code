const fs = require('fs');
const priorityDictionary = require("./priority-dictionary");

// Advent of Code 2022 - day 3

fs.readFile('input.txt', 'utf8', (err, data) => {
    if (err) {
        console.error(err);
        return;
    }
    challenge1(data)
    console.log("=======================")
    challenge2(data)
});


function challenge1(data){
    const lines = data.split("\n");
    let prioritiesSum = 0;
    lines.forEach(line=>{
        const firstHalf = line.slice(0,line.length/2);
        const secondHalf = line.slice(line.length/2,line.length);
        [...firstHalf]
            .filter((itemType,index,array)=>array.indexOf(itemType)===index)
            .forEach(itemType=>{
            if(secondHalf.includes(itemType)){
                prioritiesSum += priorityDictionary[itemType];
            }
        })
    })
    console.log(prioritiesSum)
}

///////////////////// CHALlENGE 2 /////////////////

function challenge2(data){
    const linesNotEmpty = data.split("\n").filter(item=>item);

    let prioritiesSum = 0;
    for(let index =0;index < linesNotEmpty.length/3;index++){
        const firstElfRuckSack = linesNotEmpty[index*3]
        const secondElfRuckSack = linesNotEmpty[index*3 + 1]
        const thirdElfRuckSack = linesNotEmpty[index*3 + 2];
        [...firstElfRuckSack]
            .filter((itemType,index,array)=>array.indexOf(itemType)===index)
            .forEach(itemType=>{
            if(secondElfRuckSack.includes(itemType) && thirdElfRuckSack.includes(itemType)){
                prioritiesSum += priorityDictionary[itemType];
            }
        });
    }
    console.log(prioritiesSum)
}

