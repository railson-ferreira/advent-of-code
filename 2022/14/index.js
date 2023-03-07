const fs = require('fs');
//👇 do you want to see step by step?
const fps = null;// set a fps number or 'null' to render just the last frame

// Advent of Code 2022 - day 14

fs.readFile('input.txt', 'utf8', (err, data) => {
    if (err) {
        console.error(err);
        return;
    }
    challenge1(data)
    console.log("=======================")
    challenge2(data)
});

const SAND_INITIAL_COORDINATE = {x: 500, y: 0}


function challenge1(data) {
    const {plotting, corners} = structuredDataFromData(data)

    let sandCoordinate = {x: SAND_INITIAL_COORDINATE.x, y: SAND_INITIAL_COORDINATE.y}
    let frame = {number: 0}
    render(plotting, corners, sandCoordinate, () => {
        frame.number++
        let nextSandCoordinate;
        for (let i = 0; i < 3; i++) {
            switch (i) {
                case 0:
                    nextSandCoordinate = {x: sandCoordinate.x, y: sandCoordinate.y + 1};
                    break;
                case 1:
                    nextSandCoordinate = {x: sandCoordinate.x - 1, y: sandCoordinate.y + 1};
                    break;
                case 2:
                    nextSandCoordinate = {x: sandCoordinate.x + 1, y: sandCoordinate.y + 1};
                    break;
            }
            if (
                nextSandCoordinate.x < corners.beginning.x ||
                nextSandCoordinate.x > corners.ending.x ||
                nextSandCoordinate.y > corners.ending.y
            ) {
                sandCoordinate.x = -1;
                sandCoordinate.y = -1;
                return false;
            }
            const nextPlottingPoint = plotting[nextSandCoordinate.x][nextSandCoordinate.y];
            if (nextPlottingPoint === undefined || nextPlottingPoint === "+") {
                sandCoordinate.x = nextSandCoordinate.x;
                sandCoordinate.y = nextSandCoordinate.y;
                return true;
            } else if (i === 2) {
                plotting[sandCoordinate.x][sandCoordinate.y] = "o"
                sandCoordinate.x = SAND_INITIAL_COORDINATE.x;
                sandCoordinate.y = SAND_INITIAL_COORDINATE.y;
                return true;
            }
        }
        throw "Invalid, sand wasn't able to move to any direction"
    }, fps, () => {
        showPlotting(plotting, corners, sandCoordinate)
        console.log(countSand(plotting))
    })

}


///////////////////// CHALLENGE 2 /////////////////

function challenge2(data) {
    const {plotting, corners} = structuredDataFromData(data)

    for (let x = 0; x < 1000; x++) {
        if (!plotting[x]) {
            plotting[x] = [];
        }
        plotting[x][corners.ending.y + 2] = "#"
    }
    corners.ending.y = corners.ending.y + 2;

    let sandCoordinate = {x: SAND_INITIAL_COORDINATE.x, y: SAND_INITIAL_COORDINATE.y}
    let frame = {number: 0}
    render(plotting, corners, sandCoordinate, () => {
        frame.number++
        let nextSandCoordinate;
        for (let i = 0; i < 3; i++) {
            switch (i) {
                case 0:
                    nextSandCoordinate = {x: sandCoordinate.x, y: sandCoordinate.y + 1};
                    break;
                case 1:
                    nextSandCoordinate = {x: sandCoordinate.x - 1, y: sandCoordinate.y + 1};
                    break;
                case 2:
                    nextSandCoordinate = {x: sandCoordinate.x + 1, y: sandCoordinate.y + 1};
                    break;
            }
            if (nextSandCoordinate.x < corners.beginning.x) {
                corners.beginning.x = nextSandCoordinate.x;
            }
            if (nextSandCoordinate.x > corners.ending.x) {
                corners.ending.x = nextSandCoordinate.x;
            }
            if (
                nextSandCoordinate.y > corners.ending.y
            ) {
                sandCoordinate.x = -1;
                sandCoordinate.y = -1;
                return false;
            }
            const nextPlottingPoint = plotting[nextSandCoordinate.x][nextSandCoordinate.y];

            if (nextPlottingPoint === undefined || nextPlottingPoint === "+") {
                sandCoordinate.x = nextSandCoordinate.x;
                sandCoordinate.y = nextSandCoordinate.y;
                return true;
            } else if (i === 2) {
                if(plotting[sandCoordinate.x][sandCoordinate.y] === "o"){
                    sandCoordinate.x = -1;
                    sandCoordinate.y = -1;
                    return false
                }
                plotting[sandCoordinate.x][sandCoordinate.y] = "o"
                sandCoordinate.x = SAND_INITIAL_COORDINATE.x;
                sandCoordinate.y = SAND_INITIAL_COORDINATE.y;
                return true;
            }
        }
        throw "Invalid, sand wasn't able to move to any direction"
    }, fps, () => {
        showPlotting(plotting, corners, sandCoordinate)
        console.log(countSand(plotting))
    })

}

// Utils Functions

function render(plotting, corners, sandCoordinate, afterFrame, fps = null, onFinish) {
    if (!fps) {
        while (afterFrame()) {
        }
        onFinish()
    } else {
        setTimeout(() => {
            if (afterFrame()) {
                render(plotting, corners, sandCoordinate, afterFrame, fps, onFinish)
            } else {
                console.log("✋✋✋");
                onFinish?.()
            }
        }, 1000 / fps);
        console.log();
        showPlotting(plotting, corners, sandCoordinate)
    }
}

function structuredDataFromData(data) {
    const sandCoordinates = {x: SAND_INITIAL_COORDINATE.x, y: SAND_INITIAL_COORDINATE.y};
    const plotting = []
    plotting[500] = ["+"]
    const corners = {beginning: sandCoordinates, ending: sandCoordinates}
    data.split("\n").filter(item => item).forEach(line => {
        const pairs = line.split(" -> ")
        updateCorners(corners, pairToCoordinates(pairs[0]))
        for (let i = 1; i < pairs.length; i++) {
            const previousPair = pairToCoordinates(pairs[i - 1]);
            const currentPair = pairToCoordinates(pairs[i]);
            updateCorners(corners, currentPair)
            if (previousPair.x === currentPair.x) {
                const x = previousPair.x;
                rangeLoop(previousPair.y, currentPair.y, (y) => {
                    if (!plotting[x]) {
                        plotting[x] = [];
                    }
                    plotting[x][y] = "#"
                })
            } else if (previousPair.y === currentPair.y) {
                const y = previousPair.y;
                rangeLoop(previousPair.x, currentPair.x, (x) => {
                    if (!plotting[x]) {
                        plotting[x] = [];
                    }
                    plotting[x][y] = "#"
                })
            } else {
                throw "Invalid coordinates, there wasn't found a straight line";
            }
        }
    })
    return {plotting, corners}
}

function pairToCoordinates(pair) {
    const [x, y] = pair.split(",")
    return {x: Number(x), y: Number(y)}
}

function rangeLoop(a, b, callback) {
    const distance = Math.abs(a - b);
    const min = Math.min(a, b)
    for (let i = min; i <= min + distance; i++) {
        callback(i);
    }
}

function updateCorners(corners, coordinates) {
    const {x, y} = coordinates;
    if (corners.beginning === null) {
        corners.beginning = coordinates
    } else {
        const {x: cornerX, y: cornerY} = corners.beginning;
        if (x < cornerX) {
            corners.beginning = {...corners.beginning, x};
        }
        if (y < cornerY) {
            corners.beginning = {...corners.beginning, y};
        }
    }
    if (corners.ending === null) {
        corners.ending = coordinates
    } else {
        const {x: cornerX, y: cornerY} = corners.ending;
        if (x > cornerX) {
            corners.ending = {...corners.ending, x};
        }
        if (y > cornerY) {
            corners.ending = {...corners.ending, y};
        }
    }
}

function showPlotting(plotting, corners, sandCoordinate) {
    const {beginning, ending} = corners
    const minX = Math.min(beginning.x, ending.x)
    const minY = Math.min(beginning.y, ending.y)
    const maxX = Math.max(beginning.x, ending.x)
    const maxY = Math.max(beginning.y, ending.y)

    let count = minX;
    const header = Array.from({length: maxX - minX + 1}, () => count++).map(item => item.toString().padStart(3, "0"));
    for (let i = 0; i < 3; i++) {
        if (maxY < 10) {
            console.log("  " + header.map(text => text[i]).join(" "))
        } else if (maxY < 100) {
            console.log("   " + header.map(text => text[i]).join(" "))
        } else {
            console.log("   " + header.map(text => text[i]).join(" "))
        }
    }
    for (let y = minY; y <= maxY; y++) {
        let line
        if (maxY < 10) {
            line = y + " ";
        } else if (maxY < 100) {
            line = y.toString().padStart(2, "0") + " "
        } else {
            line = y.toString().padStart(3, "0") + " "
        }
        for (let x = minX; x <= maxX; x++) {
            if (x === sandCoordinate.x && y === sandCoordinate.y) {
                line += "☆"
            } else {
                line += plotting[x]?.[y] || ".";
            }
            line += " ";
        }
        console.log(line)
    }
}


function countSand(plotting) {
    let sum = 0;
    plotting.forEach(items => {
        items.forEach(item => {
            if (item === "o") {
                sum++;
            }
        })
    })

    return sum;
}