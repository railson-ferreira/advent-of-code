const fs = require('fs');

// Advent of Code 2022 - day 5

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
    const crateStacks = []

    let lineIndex = 0;
    let currentLine = lines[lineIndex];
    while(currentLine !== ""){
        lineIndex ++;
        currentLine = lines[lineIndex];
    }
    for(let index = lineIndex-1; index>=0;index--){
        currentLine = lines[index];
        if(index === lineIndex-1){
            const stackIds = currentLine.split(/\s+/)
            stackIds.forEach(stackId=>{
                if(!stackId){
                    return
                }
                const arrayValidIndex = Number(stackId)-1;
                crateStacks[arrayValidIndex] = []
            })
        }else{
            const nilCrate = "[!]";
            while(/\s{5}/.test(currentLine)){
                currentLine = currentLine.replaceAll(/\s{5}/g, ` ${nilCrate} `);
            }
            const cratesIds = currentLine.split(/\s+/).filter(item=>item);
            while(cratesIds.length<crateStacks.length){
                cratesIds.push(nilCrate);
            }
            cratesIds.forEach((crate,index)=>{
                if(crate && crate !== nilCrate){
                    crateStacks[index].push(crate);
                }
            })
        }
    }
    lineIndex++;
    currentLine = lines[lineIndex];
    // showCrateStacks(crateStacks)
    while(currentLine){
        // console.log(currentLine);
        const commandRegex = /move (\d+) from (\d+) to (\d+)/
        const [_,qtdStr, fromStr, toStr] = commandRegex.exec(currentLine)
        const qtd = Number(qtdStr)
        const from = Number(fromStr)
        const to = Number(toStr)
        for(let count = 0 ; count < qtd ; count ++){
            const crate = crateStacks[from-1].pop();
            if(crate){
                crateStacks[to-1].push(crate);
            }
        }
        lineIndex ++;
        // showCrateStacks(crateStacks)
        debugger
        currentLine = lines[lineIndex];
    }
    let result = "";
    crateStacks.forEach(stack=>{
        const crate = stack[stack.length-1];
        if(crate){
            result += crate;
        }else{
            result += " ";
        }
    })
    console.log(result.replaceAll("[","").replaceAll("]", ""))
}

///////////////////// CHALlENGE 2 /////////////////

function challenge2(data) {
    const lines = data.split("\n");
    let fullyContainedCount = 0;
    const crateStacks = []

    let lineIndex = 0;
    let currentLine = lines[lineIndex];
    while(currentLine !== ""){
        lineIndex ++;
        currentLine = lines[lineIndex];
    }
    for(let index = lineIndex-1; index>=0;index--){
        currentLine = lines[index];
        if(index === lineIndex-1){
            const stackIds = currentLine.split(/\s+/)
            stackIds.forEach(stackId=>{
                if(!stackId){
                    return
                }
                const arrayValidIndex = Number(stackId)-1;
                crateStacks[arrayValidIndex] = []
            })
        }else{
            const nilCrate = "[!]";
            while(/\s{5}/.test(currentLine)){
                currentLine = currentLine.replaceAll(/\s{5}/g, ` ${nilCrate} `);
            }
            const cratesIds = currentLine.split(/\s+/).filter(item=>item);
            while(cratesIds.length<crateStacks.length){
                cratesIds.push(nilCrate);
            }
            cratesIds.forEach((crate,index)=>{
                if(crate && crate !== nilCrate){
                    crateStacks[index].push(crate);
                }
            })
        }
    }
    lineIndex++;
    currentLine = lines[lineIndex];
    // showCrateStacks(crateStacks)
    while(currentLine){
        // console.log(currentLine);
        const commandRegex = /move (\d+) from (\d+) to (\d+)/
        const [_,qtdStr, fromStr, toStr] = commandRegex.exec(currentLine)
        const qtd = Number(qtdStr)
        const from = Number(fromStr)
        const to = Number(toStr)
        const cratesToMove = []
        for(let count = 0 ; count < qtd ; count ++){
            const crate = crateStacks[from-1].pop();
            if(crate){
                cratesToMove.push(crate);
            }
        }
        while (cratesToMove.length){
            crateStacks[to-1].push(cratesToMove.pop());
        }
        lineIndex ++;
        // showCrateStacks(crateStacks)
        debugger
        currentLine = lines[lineIndex];
    }
    let result = "";
    crateStacks.forEach(stack=>{
        const crate = stack[stack.length-1];
        if(crate){
            result += crate;
        }else{
            result += " ";
        }
    })
    console.log(result.replaceAll("[","").replaceAll("]", ""))
}



function showCrateStacks(crateStacks){
    console.log("(1)(2)(3)(4)(5)(6)(7)(8)(9)")
    let max = 0;
    crateStacks.forEach(stack=>{
        max = Math.max(stack.length,max);
    })
    for(let x =0 ; x<max ;x++){
        let line =""
        crateStacks.forEach(crateStack=>{
            const crate = crateStack[x];
            if(crate){
                line+=crate;
            }else{
                line+="   "
            }
        })
        console.log(line)
    }
}
