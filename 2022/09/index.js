const fs = require('fs');
const assert = require("assert");

// Advent of Code 2022 - day 9

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
    const moves = [];
    data.split("\n").filter(item => item).forEach(line => {
        const parts = line.split(" ")
        moves.push({
            direction: parts[0],
            distance: Number(parts[1])
        })
    })
    const currentHeadPosition = {x: 0, y: 0}
    const currentTailPosition = {x: 0, y: 0}
    const tailPositionLog = [currentTailPosition];
    moves.forEach(move => {
        for(let i = 0; i < move.distance; i++){
            debugger
            switch (move.direction) {
                case "U":
                    currentHeadPosition.y--
                    break;
                case "L":
                    currentHeadPosition.x--
                    break;
                case "D":
                    currentHeadPosition.y++
                    break;
                case "R":
                    currentHeadPosition.x++
                    break;
            }
            const distanceX = Math.abs(currentHeadPosition.x-currentTailPosition.x)
            const directionX = currentHeadPosition.x-currentTailPosition.x<0?"west":"east"
            const distanceY = Math.abs(currentHeadPosition.y-currentTailPosition.y)
            const directionY = currentHeadPosition.y-currentTailPosition.y<0?"north":"south"
            if(distanceX >= 1 && distanceY >=1 && (distanceX !== distanceY||distanceX !== 1)){
                if(directionX==="west"){
                    currentTailPosition.x--
                }else{
                    currentTailPosition.x++
                }
                if(directionY==="north"){
                    currentTailPosition.y--
                }else{
                    currentTailPosition.y++
                }
            }else if(distanceX > 1){
                if(directionX==="west"){
                    currentTailPosition.x--
                }else{
                    currentTailPosition.x++
                }
            }else if(distanceY > 1){
                if(directionY==="north"){
                    currentTailPosition.y--
                }else{
                    currentTailPosition.y++
                }
            }
            tailPositionLog.push({x:currentTailPosition.x,y:currentTailPosition.y})
        }
    })
    const visitedPositionAtLeastOnce = tailPositionLog.filter((tailPosition, index, array)=>{
        return array.findIndex(tailPositionB=>tailPosition.x===tailPositionB.x&&tailPosition.y===tailPositionB.y) === index
    })
    console.log(visitedPositionAtLeastOnce.length);
}

///////////////////// CHALLENGE 2 /////////////////

function challenge2(data) {
    const moves = [];
    data.split("\n").filter(item => item).forEach(line => {
        const parts = line.split(" ")
        moves.push({
            direction: parts[0],
            distance: Number(parts[1])
        })
    })
    const knotsPositions = Array.from({length: 10}, ()=>({x: 0, y: 0}))
    const headKnotPosition = knotsPositions[0];
    const tailKnotPosition = knotsPositions[9];
    const tailPositionLog = [knotsPositions[9]];
    moves.forEach(move => {
        for(let i = 0; i < move.distance; i++){
            debugger
            switch (move.direction) {
                case "U":
                    headKnotPosition.y--
                    break;
                case "L":
                    headKnotPosition.x--
                    break;
                case "D":
                    headKnotPosition.y++
                    break;
                case "R":
                    headKnotPosition.x++
                    break;
            }
            for(let knotIndex = 1; knotIndex < knotsPositions.length; knotIndex++){
                const leadKnotPosition = knotsPositions[knotIndex-1];
                const trailingKnotPosition = knotsPositions[knotIndex];
                knotSmartMove(leadKnotPosition,trailingKnotPosition);
            }
            tailPositionLog.push({x:tailKnotPosition.x,y:tailKnotPosition.y})
        }
    })
    const visitedPositionAtLeastOnce = tailPositionLog.filter((tailPosition, index, array)=>{
        return array.findIndex(tailPositionB=>tailPosition.x===tailPositionB.x&&tailPosition.y===tailPositionB.y) === index
    })
    console.log(visitedPositionAtLeastOnce.length);
}


function knotSmartMove(leadKnotPosition, trailingKnotPosition){
    const distanceX = Math.abs(leadKnotPosition.x-trailingKnotPosition.x)
    const directionX = leadKnotPosition.x-trailingKnotPosition.x<0?"west":"east"
    const distanceY = Math.abs(leadKnotPosition.y-trailingKnotPosition.y)
    const directionY = leadKnotPosition.y-trailingKnotPosition.y<0?"north":"south"
    if(distanceX >= 1 && distanceY >=1 && (distanceX !== distanceY||distanceX !== 1)){
        if(directionX==="west"){
            trailingKnotPosition.x--
        }else{
            trailingKnotPosition.x++
        }
        if(directionY==="north"){
            trailingKnotPosition.y--
        }else{
            trailingKnotPosition.y++
        }
    }else if(distanceX > 1){
        if(directionX==="west"){
            trailingKnotPosition.x--
        }else{
            trailingKnotPosition.x++
        }
    }else if(distanceY > 1){
        if(directionY==="north"){
            trailingKnotPosition.y--
        }else{
            trailingKnotPosition.y++
        }
    }
}
