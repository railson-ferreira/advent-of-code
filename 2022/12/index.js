const fs = require('fs');

// Advent of Code 2022 - day 12

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
    const matrix = []
    data.split("\n").filter(item=>item).forEach(line=>{
        matrix.push([...line])
    })
    const initialPosition = getInitialPosition(matrix)
    const targetPosition = getTargetPosition(matrix)
    const distanceMap = getDistanceMatrix(matrix);
    let nextPositions = [initialPosition]
    let distance = 0;
    while (nextPositions.length>0){
        nextPositions = fillDistanceMap(nextPositions, distanceMap,distance);
        nextPositions = nextPositions.filter((item,index,array)=>{
            return array.findIndex(itemB=>itemB.x===item.x&&itemB.y===item.y) === index
        })
        distance++;
    }
    console.log(distanceMap[targetPosition.y][targetPosition.x].distance)
}

///////////////////// CHALLENGE 2 /////////////////

function challenge2(data) {
    const matrix = []
    data.split("\n").filter(item=>item).forEach(line=>{
        matrix.push([...line])
    })
    const initialPositions = getInitialPositions(matrix);
    let bestDistance = null;
    initialPositions.forEach(initialPosition=>{
        const targetPosition = getTargetPosition(matrix)
        const distanceMap = getDistanceMatrix(matrix);
        let nextPositions = [initialPosition]
        let distance = 0;
        while (nextPositions.length>0){
            nextPositions = fillDistanceMap(nextPositions, distanceMap,distance);
            nextPositions = nextPositions.filter((item,index,array)=>{
                return array.findIndex(itemB=>itemB.x===item.x&&itemB.y===item.y) === index
            })
            distance++;
        }
        const targetDistance = distanceMap[targetPosition.y][targetPosition.x].distance;
        if(!bestDistance&& targetDistance!==-1){
            bestDistance = targetDistance;
        }else if(targetDistance!==-1&&targetDistance<bestDistance){
            bestDistance = targetDistance;
        }
    })
    console.log(bestDistance)
}

// Utils Functions

function fillDistanceMap(positions,distanceMap, distance = 0){
    const nextPositions = []
    positions.forEach(position=>{
        const {x,y,elevation} = position;
        distanceMap[y][x].distance = distance;
        const possibleNextPositions = [
            {x:x  , y:y-1},
            {x:x-1, y:y  },
            {x:x  , y:y+1},
            {x:x+1, y:y  },
        ]
        for(let i = 0; i < possibleNextPositions.length; i++){
            const nextPosition = possibleNextPositions[i];
            const {x,y} = nextPosition
            const nextItem = distanceMap[y]?.[x];
            if(!nextItem)continue;
            if(nextItem.elevation>elevation+1)continue;
            if(nextItem.distance < 0){
                nextPositions.push({...nextPosition, elevation:nextItem.elevation});
            }
        }
    })
    return nextPositions;
}

function getInitialPosition(matrix){
    const initialPositionMark = "S";
    for(let y = 0; y<matrix.length ; y++){
        for(let x = 0; x < matrix[0].length; x++){
            if(matrix[y][x] === initialPositionMark){
                return {x,y,elevation:getElevationNumber(matrix[y][x])}
            }
        }
    }
}

function getInitialPositions(matrix){
    const initialPositionMarks = ["S","a"];
    const initialPositions = []
    for(let y = 0; y<matrix.length ; y++){
        for(let x = 0; x < matrix[0].length; x++){
            if(initialPositionMarks.includes(matrix[y][x])){
                initialPositions.push({x,y,elevation:getElevationNumber(matrix[y][x])})
            }
        }
    }
    return initialPositions;
}

function getTargetPosition(matrix){
    const targetMark = "E";
    for(let y = 0; y<matrix.length ; y++){
        for(let x = 0; x < matrix[0].length; x++){
            if(matrix[y][x] === targetMark){
                return {x,y}
            }
        }
    }
}

function getDistanceMatrix(matrix){
    const distanceMatrix = []
    matrix.forEach((line,y)=>{
        distanceMatrix[y]= []
        line.forEach((item,x)=>{
            distanceMatrix[y][x] = {distance: -1, elevation: getElevationNumber(item)};
        })
    })
    return distanceMatrix
}


function getElevationNumber(value){
    switch (value){
        case "S": return 0
        case "E": return 25
        default:
            return value.charCodeAt(0)-97
    }
}