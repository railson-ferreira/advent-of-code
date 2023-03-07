const fs = require('fs');

// Advent of Code 2022 - day 15

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
    const structuredData = getStructuredData(data);
    const visibleFreePositions = getVisibleFreePositionsAtY(structuredData, 2000000).sort((a,b)=>a.x-b.x);
    console.log(visibleFreePositions.length)
}


///////////////////// CHALLENGE 2 /////////////////

function challenge2(data) {
    const structuredData = getStructuredData(data);
    for(let y = 0; y <= 4000000;y++){
        const visibleRanges = []
        structuredData.forEach(sensorInfo=>{
            const range = sensorInfo.getVisibleRangeAtY(y)
            if(range.minX>range.maxX){
                return;
            }
            visibleRanges.push(range);
        })
         visibleRanges.sort((a,b)=>{
            return (a.minX+a.maxX)/2-(b.minX+b.maxX)/2 // median
        })
        let incrementalRange = visibleRanges[0];
        let foundPosition = null;
        for(let i = 1; i <visibleRanges.length; i++){
            if(visibleRanges[i].minX>incrementalRange.maxX+1){
                foundPosition = {x:incrementalRange.maxX+1,y:incrementalRange.y}
            }else{
                const {minX,maxX} = visibleRanges[i]
                if(minX<incrementalRange.minX){
                    incrementalRange.minX = minX;
                }
                if(maxX>incrementalRange.maxX){
                    incrementalRange.maxX = maxX;
                }
                if(foundPosition){
                    if(minX<=foundPosition.x && maxX>=foundPosition.x){
                        foundPosition = null
                    }
                }
            }
        }
        if(foundPosition){
            console.log(foundPosition.x*4000000+foundPosition.y)
            break;
        }
    }
}

// Utils Functions


function getStructuredData(data) {
    // const matrix = []
    // let maxXIndex = 0;
    const sensorInfos = []
    data.split("\n").filter(item => item).forEach(line => {
        const regex = /Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)/g
        const result = regex.exec(line)
        const [_, sensorXStr, sensorYStr, beaconXStr, beaconYStr] = result
        const
            sensorX = Number(sensorXStr), sensorY = Number(sensorYStr), beaconX = Number(beaconXStr),
            beaconY = Number(beaconYStr)
        sensorInfos.push(getSensorInfo({x: sensorX, y: sensorY}, {x: beaconX, y: beaconY}))

    })

    return sensorInfos
}

function getSensorInfo(sensorPosition, nearBeaconPosition) {
    const horizontalDist = Math.abs(sensorPosition.x - nearBeaconPosition.x)
    const verticalDist = Math.abs(sensorPosition.y - nearBeaconPosition.y)
    const manhattanDistance = horizontalDist + verticalDist;

    function isVisiblePosition(position) {
        const pointHorizontalDist = Math.abs(sensorPosition.x - position.x)
        const pointVerticalDist = Math.abs(sensorPosition.y - position.y)
        const pointManhattanDistance = pointHorizontalDist + pointVerticalDist;
        return pointManhattanDistance <= manhattanDistance;

    }

    function getVisibleRangeAtY(y) {
        const verticalDist = Math.abs(sensorPosition.y - y)
        return {minX: sensorPosition.x-(manhattanDistance-verticalDist), maxX:sensorPosition.x+(manhattanDistance-verticalDist),y, coords: `${sensorPosition.x},${sensorPosition.y}`, md: manhattanDistance};

    }

    return {sensorPosition, beaconPosition: nearBeaconPosition, isVisiblePosition, manhattanDistance,getVisibleRangeAtY}
}

function getVisibleFreePositionsAtY(structuredData, y) {
    const visibleCoordinates = new Set();
    structuredData.forEach(sensorInfo => {
        const minX = sensorInfo.sensorPosition.x - sensorInfo.manhattanDistance;
        const maxX = sensorInfo.sensorPosition.x + sensorInfo.manhattanDistance;
        for (let x = minX; x <= maxX; x++) {
            if (x !== sensorInfo.beaconPosition.x || y !== sensorInfo.beaconPosition.y) {
                if (sensorInfo.isVisiblePosition({x, y})) {
                    visibleCoordinates.add(`${x},${y}`)
                }
            }
        }
    })
    return [...visibleCoordinates].map(coordinates => {
        const [x, y] = coordinates.split(",");
        return {x, y}
    })
}