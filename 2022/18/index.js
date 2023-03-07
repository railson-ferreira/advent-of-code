const fs = require('fs');

// Advent of Code 2022 - day 18

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
    const blocks = get3dBlocks(data);
    connectSides(blocks)
    let totalArea = 0;
    blocks.forEach(block => {
        if (block.sides.a === null)
            totalArea++;
        if (block.sides.b === null)
            totalArea++;
        if (block.sides.c === null)
            totalArea++;
        if (block.sides.d === null)
            totalArea++;
        if (block.sides.e === null)
            totalArea++;
        if (block.sides.f === null)
            totalArea++;
    })
    console.log(totalArea)
}


///////////////////// CHALLENGE 2 /////////////////

function challenge2(data) {
    const blocks = get3dBlocks(data);
    connectSides(blocks)
    const airBlocks = getAirBlocks(blocks);
    connectSides([...blocks, ...airBlocks])
    airBlocks.forEach(airBlock => {
        if (isTrapped(airBlock,blocks) === true) {
            airBlock.trapped = true;
        }
    })
    let totalArea = 0;
    blocks.forEach(block => {
        const {a, b, c, d, e, f} = block.sides;
        if (a === null || a.type === "air" && !a.trapped)
            totalArea++;
        if (b === null || b.type === "air" && !b.trapped)
            totalArea++;
        if (c === null || c.type === "air" && !c.trapped)
            totalArea++;
        if (d === null || d.type === "air" && !d.trapped)
            totalArea++;
        if (e === null || e.type === "air" && !e.trapped)
            totalArea++;
        if (f === null || f.type === "air" && !f.trapped)
            totalArea++;
    })
    console.log("wrong answer:",totalArea);
}

// Utils Functions


function get3dBlocks(data) {
    const blocks = []
    data.split("\n").filter(item => item).forEach((line, index) => {
        const regex = /(\d+),(\d+),(\d+)/g
        const result = regex.exec(line)
        const [_, x, y, z] = result
        blocks.push({
            type: "solid",
            id: index + 1,
            x: Number(x),
            y: Number(y),
            z: Number(z),
            sides: {
                a: null,
                b: null,
                c: null,
                d: null,
                e: null,
                f: null,
            }
        })

    })

    return blocks
}


function connectSides(blocks) {
    blocks.forEach(block => {
        const {x, y, z} = block;
        blocks.forEach(anotherBlock => {
            const {x: xb, y: yb, z: zb} = anotherBlock;
            if (x - xb === 1 && y === yb && z === zb) {
                block.sides.a = anotherBlock;
            }
            if (x - xb === -1 && y === yb && z === zb) {
                block.sides.b = anotherBlock;
            }
            if (x === xb && y - yb === 1 && z === zb) {
                block.sides.c = anotherBlock;
            }
            if (x === xb && y - yb === -1 && z === zb) {
                block.sides.d = anotherBlock;
            }
            if (x === xb && y === yb && z - zb === 1) {
                block.sides.e = anotherBlock;
            }
            if (x === xb && y === yb && z - zb === -1) {
                block.sides.f = anotherBlock;
            }
        })
    })
}

function getAirBlocks(connectedBlocks) {
    const airBlocks = []
    connectedBlocks.forEach(block => {
        if (block.sides.a === null) {
            const candidateAirBlock = {
                type: "air",
                x: block.x - 1,
                y: block.y,
                z: block.z,
            }
            if (airBlocks.findIndex(candidateBlock => {
                return candidateBlock.x === candidateAirBlock.x && candidateBlock.y === candidateAirBlock.y && candidateBlock.z === candidateAirBlock.z
            }) < 0)
                airBlocks.push(candidateAirBlock)
        }
        if (block.sides.b === null) {
            const candidateAirBlock = {
                type: "air",
                x: block.x + 1,
                y: block.y,
                z: block.z,
            }
            if (airBlocks.findIndex(candidateBlock => {
                return candidateBlock.x === candidateAirBlock.x && candidateBlock.y === candidateAirBlock.y && candidateBlock.z === candidateAirBlock.z
            }) < 0)
                airBlocks.push(candidateAirBlock)
        }
        if (block.sides.c === null) {
            const candidateAirBlock = {
                type: "air",
                x: block.x,
                y: block.y - 1,
                z: block.z,
            }
            if (airBlocks.findIndex(candidateBlock => {
                return candidateBlock.x === candidateAirBlock.x && candidateBlock.y === candidateAirBlock.y && candidateBlock.z === candidateAirBlock.z
            }) < 0)
                airBlocks.push(candidateAirBlock)
        }
        if (block.sides.d === null) {
            const candidateAirBlock = {
                type: "air",
                x: block.x,
                y: block.y + 1,
                z: block.z,
            }
            if (airBlocks.findIndex(candidateBlock => {
                return candidateBlock.x === candidateAirBlock.x && candidateBlock.y === candidateAirBlock.y && candidateBlock.z === candidateAirBlock.z
            }) < 0)
                airBlocks.push(candidateAirBlock)
        }
        if (block.sides.e === null) {
            const candidateAirBlock = {
                type: "air",
                x: block.x,
                y: block.y,
                z: block.z - 1,
            }
            if (airBlocks.findIndex(candidateBlock => {
                return candidateBlock.x === candidateAirBlock.x && candidateBlock.y === candidateAirBlock.y && candidateBlock.z === candidateAirBlock.z
            }) < 0)
                airBlocks.push(candidateAirBlock)
        }
        if (block.sides.f === null) {
            const candidateAirBlock = {
                type: "air",
                x: block.x,
                y: block.y,
                z: block.z + 1,
            }
            if (airBlocks.findIndex(candidateBlock => {
                return candidateBlock.x === candidateAirBlock.x && candidateBlock.y === candidateAirBlock.y && candidateBlock.z === candidateAirBlock.z
            }) < 0)
                airBlocks.push(candidateAirBlock)
        }
    })
    airBlocks.forEach(candidateBlock => {
        candidateBlock.sides = {
            a: null,
            b: null,
            c: null,
            d: null,
            e: null,
            f: null,
        }
    })
    return airBlocks
}

function isTrapped(airBlock) {
    airBlock.inProgress = true;
    return Object.values(airBlock.sides).map(sideBlock => {
        if (sideBlock === null) {
            return false;
        } else if (sideBlock.type === "solid") {
            return true;
        } else if (sideBlock.type === "air") {

            if (sideBlock.inProgress) {
                return true
            }
            return isTrapped(sideBlock)
        } else {
            throw "Invalid"
        }
    }).reduce((a, b) => a && b)
}
