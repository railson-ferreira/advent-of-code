const fs = require('fs');

// Advent of Code 2022 - day 23

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
    const matrix = [];
    const checkingOrder = [
        "N",
        "S",
        "W",
        "E"
    ]
    data.split("\n").filter(item=>item).forEach(line=>{
        matrix.push([...line]);
    })
    /// ⚠️ FIX MATRIX FOR EASIER USE
    for(let i = -1; i>=-100;i--){
        matrix.forEach(line=>{
            line[i] = "."
        })
    }
    const newXLength = matrix[0].length+100;
    for(let i = matrix[0].length; i<=newXLength;i++){
        matrix.forEach(line=>{
            line[i] = "."
        })
    }
    /// ⚠️ FIX MATRIX FOR EASIER USE
    let count = 0;
    while (count<10){
        const elvesProposes = []
        const {minX,maxX,minY,maxY} = getRanges(matrix)
        for(let x = minX;x<=maxX;x++){
            for(let y = minY;y<=maxY;y++){
                const tile = matrix[y]?.[x]
                if(tile === "."){
                    continue;
                }
                let isProposeMade = false
                for(let i = 0; i<checkingOrder.length;i++){
                    if(isProposeMade){
                        break;
                    }
                    if(matrix[y-1]?.[x-1] !== "#" && matrix[y-1]?.[x] !== "#" && matrix[y-1]?.[x+1] !== "#" && matrix[y]?.[x+1] !== "#" &&
                        matrix[y+1]?.[x+1] !== "#" && matrix[y+1]?.[x] !== "#" && matrix[y+1]?.[x-1] !== "#" && matrix[y]?.[x-1] !== "#"){
                        break;
                    }
                    const currentChecking = checkingOrder[i]
                    switch (currentChecking){
                        case "N":
                            if(matrix[y-1]?.[x-1] !== "#"&&matrix[y-1]?.[x] !== "#"&&matrix[y-1]?.[x+1] !== "#"){
                                elvesProposes.push({
                                    elf: `${x},${y}`,
                                    propose: `${x},${y-1}`
                                })
                                isProposeMade = true;
                            }
                            break;
                        case "S":
                            if(matrix[y+1]?.[x-1] !== "#"&&matrix[y+1]?.[x] !== "#"&&matrix[y+1]?.[x+1] !== "#"){
                                elvesProposes.push({
                                    elf: `${x},${y}`,
                                    propose: `${x},${y+1}`
                                })
                                isProposeMade = true;
                            }
                            break;
                        case "W":
                            if(matrix[y-1]?.[x-1] !== "#"&&matrix[y]?.[x-1] !== "#"&&matrix[y+1]?.[x-1] !== "#"){
                                elvesProposes.push({
                                    elf: `${x},${y}`,
                                    propose: `${x-1},${y}`
                                })
                                isProposeMade = true;
                            }
                            break;
                        case "E":
                            if(matrix[y-1]?.[x+1] !== "#"&&matrix[y]?.[x+1] !== "#"&&matrix[y+1]?.[x+1] !== "#"){
                                elvesProposes.push({
                                    elf: `${x},${y}`,
                                    propose: `${x+1},${y}`
                                })
                                isProposeMade = true;
                            }
                            break;
                    }
                }
            }
        }
       const notAllowedProposes = new Set(elvesProposes.map(item=>item.propose).filter((item,index,array)=>array.indexOf(item) !== index))

        elvesProposes.forEach(propose=>{
            if(!notAllowedProposes.has(propose.propose)){
                const [x,y] = propose.elf.split(",")
                const [nextX,nextY] = propose.propose.split(",")
                if(!matrix[y]) {
                    matrix[y] = []
                    for (let i = minX; i <= maxX; i++) {
                        matrix[y][i] = "."
                    }
                }
                matrix[y][x] = ".";
                if(!matrix[nextY]) {
                    matrix[nextY] = []
                    for (let i = minX; i <= maxX; i++) {
                        matrix[nextY][i] = "."
                    }
                }
                matrix[nextY][nextX] = "#";
            }
        })
        const aux = checkingOrder[0]
        checkingOrder[0] = checkingOrder[1]
        checkingOrder[1] = checkingOrder[2]
        checkingOrder[2] = checkingOrder[3]
        checkingOrder[3] = aux
        count++
    }
    console.log(getScore(matrix))
}


///////////////////// CHALLENGE 2 /////////////////

function challenge2(data) {
    const matrix = [];
    const checkingOrder = [
        "N",
        "S",
        "W",
        "E"
    ]
    data.split("\n").filter(item=>item).forEach(line=>{
        matrix.push([...line]);
    })
    /// ⚠️ FIX MATRIX FOR EASIER USE
    for(let i = -1; i>=-100;i--){
        matrix.forEach(line=>{
            line[i] = "."
        })
    }
    const newXLength = matrix[0].length+100;
    for(let i = matrix[0].length; i<=newXLength;i++){
        matrix.forEach(line=>{
            line[i] = "."
        })
    }
    /// ⚠️ FIX MATRIX FOR EASIER USE
    let count = 0;
    while (true){
        const elvesProposes = []
        const {minX,maxX,minY,maxY} = getRanges(matrix)
        for(let x = minX;x<=maxX;x++){
            for(let y = minY;y<=maxY;y++){
                const tile = matrix[y]?.[x]
                if(tile === "."){
                    continue;
                }
                let isProposeMade = false
                for(let i = 0; i<checkingOrder.length;i++){
                    if(isProposeMade){
                        break;
                    }
                    if(matrix[y-1]?.[x-1] !== "#" && matrix[y-1]?.[x] !== "#" && matrix[y-1]?.[x+1] !== "#" && matrix[y]?.[x+1] !== "#" &&
                        matrix[y+1]?.[x+1] !== "#" && matrix[y+1]?.[x] !== "#" && matrix[y+1]?.[x-1] !== "#" && matrix[y]?.[x-1] !== "#"){
                        break;
                    }
                    const currentChecking = checkingOrder[i]
                    switch (currentChecking){
                        case "N":
                            if(matrix[y-1]?.[x-1] !== "#"&&matrix[y-1]?.[x] !== "#"&&matrix[y-1]?.[x+1] !== "#"){
                                elvesProposes.push({
                                    elf: `${x},${y}`,
                                    propose: `${x},${y-1}`
                                })
                                isProposeMade = true;
                            }
                            break;
                        case "S":
                            if(matrix[y+1]?.[x-1] !== "#"&&matrix[y+1]?.[x] !== "#"&&matrix[y+1]?.[x+1] !== "#"){
                                elvesProposes.push({
                                    elf: `${x},${y}`,
                                    propose: `${x},${y+1}`
                                })
                                isProposeMade = true;
                            }
                            break;
                        case "W":
                            if(matrix[y-1]?.[x-1] !== "#"&&matrix[y]?.[x-1] !== "#"&&matrix[y+1]?.[x-1] !== "#"){
                                elvesProposes.push({
                                    elf: `${x},${y}`,
                                    propose: `${x-1},${y}`
                                })
                                isProposeMade = true;
                            }
                            break;
                        case "E":
                            if(matrix[y-1]?.[x+1] !== "#"&&matrix[y]?.[x+1] !== "#"&&matrix[y+1]?.[x+1] !== "#"){
                                elvesProposes.push({
                                    elf: `${x},${y}`,
                                    propose: `${x+1},${y}`
                                })
                                isProposeMade = true;
                            }
                            break;
                    }
                }
            }
        }
        const notAllowedProposes = new Set(elvesProposes.map(item=>item.propose).filter((item,index,array)=>array.indexOf(item) !== index))

        let moved = false;
        elvesProposes.forEach(propose=>{
            if(!notAllowedProposes.has(propose.propose)){
                const [x,y] = propose.elf.split(",")
                const [nextX,nextY] = propose.propose.split(",")
                if(!matrix[y]) {
                    matrix[y] = []
                    for (let i = minX; i <= maxX; i++) {
                        matrix[y][i] = "."
                    }
                }
                matrix[y][x] = ".";
                if(!matrix[nextY]) {
                    matrix[nextY] = []
                    for (let i = minX; i <= maxX; i++) {
                        matrix[nextY][i] = "."
                    }
                }
                matrix[nextY][nextX] = "#";
                moved = true;
            }
        })
        const aux = checkingOrder[0]
        checkingOrder[0] = checkingOrder[1]
        checkingOrder[1] = checkingOrder[2]
        checkingOrder[2] = checkingOrder[3]
        checkingOrder[3] = aux
        count++
        if(!moved){
            break;
        }
    }
    console.log(count)
}

// Utils Functions

function showMatrix(matrix){
    const {minX,maxX,minY,maxY} = getRanges(matrix)
    for(let y = minY;y<=maxY;y++) {
        let line = ""
        for (let x = minX; x <= maxX; x++) {
            line += matrix[y]?.[x];
        }
        console.log(line);
    }
}

function getScore(matrix){
    let minRecX = Infinity
    let minRecY = Infinity
    let maxRecX = -Infinity
    let maxRecY = -Infinity
    const {minX,maxX,minY,maxY} = getRanges(matrix)
    for(let x = minX;x<=maxX;x++){
        for(let y = minY;y<=maxY;y++){
            const tile = matrix[y]?.[x]
            if(tile === "#") {
                if (x < minRecX)
                    minRecX = x;
                if (x > maxRecX)
                    maxRecX = x;
                if (y < minRecY)
                    minRecY = y;
                if (y > maxRecY)
                    maxRecY = y;
            }
        }
    }
    let score = 0;
    for(let x = minRecX; x<=maxRecX; x++){
        for(let y = minRecY;y<=maxRecY;y++){
            const tile = matrix[y]?.[x];
            if(tile==="."){
                score++;
            }
        }
    }
    return score
}

function getRanges(matrix){
    let minX = 0;
    let minY = 0;
    let maxX = -Infinity;
    let maxY = matrix.length-1;
    while (matrix[minY-1]){
        minY--;
    }
    for(let y = minY; y<matrix.length;y++){
        while (matrix[y][minX-1]){
            minX--;
        }
        if(matrix[y].length-1>maxX){
            maxX = matrix[y].length-1
        }
    }
    return {minX,maxX,minY,maxY}
}