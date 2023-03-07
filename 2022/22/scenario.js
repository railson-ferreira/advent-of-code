class Scenario {
    matrix = []
    currentPosition;
    facing = ">";


    constructor(filteredData) {
        filteredData.split("\n").filter(item => item).forEach((line, y) => {
            while (this.matrix.length - 1 < y) {
                this.matrix.push([])
            }
            [...line].forEach((tile, x) => {
                switch (tileType[tile]) {
                    case "void":
                    case "ground":
                    case "wall":
                        this.matrix[y][x] = tile;
                        break;
                    default:
                        throw "Invalid"

                }
                if (y === 0 && tileType[tile] === "ground" && !this.currentPosition) {
                    this.currentPosition = {x, y}
                }
            })
        })
    }

    runInstruction(instruction) {
        if (typeof instruction === "number") {
            for(let i=0;i<instruction;i++){
                this.moveForward()
                // console.log()
                // this.show();
                // debugger;
            }
        } else {
            switch (instruction) {
                case this.facing === ">" && "R":
                    this.facing = "v"
                    break
                case this.facing === "v" && "R":
                    this.facing = "<"
                    break
                case this.facing === "<" && "R":
                    this.facing = "^"
                    break
                case this.facing === "^" && "R":
                    this.facing = ">"
                    break
                case this.facing === ">" && "L":
                    this.facing = "^"
                    break
                case this.facing === "v" && "L":
                    this.facing = ">"
                    break
                case this.facing === "<" && "L":
                    this.facing = "v"
                    break
                case this.facing === "^" && "L":
                    this.facing = "<"
                    break
                default:
                    throw "Invalid"
            }
        }
    }

    moveForward(){
        let nextPosition = this.getNextPosition();
        let nextTile = this.matrix[nextPosition.y]?.[nextPosition.x];
        if(tileType[nextTile] === "void"){
            nextPosition = this.getTilePositionAfterVoidTile()
            nextTile = this.matrix[nextPosition.y]?.[nextPosition.x];
        }
        switch(tileType[nextTile]){
            case "ground":
                this.currentPosition = nextPosition;
                this.matrix[this.currentPosition.y][this.currentPosition.x] = this.facing
                return;
            case "wall":
                return;
        }
        throw "Invalid"
    }

    getNextPosition() {
        const x = this.currentPosition.x, y = this.currentPosition.y;
        switch (this.facing) {
            case ">":
                return {x: x + 1, y}
            case "v":
                return {x, y: y + 1}
            case "<":
                return {x: x - 1, y}
            case "^":
                return {x: x, y: y - 1}
        }
        throw "Invalid"
    }

    show() {
        this.matrix.forEach((columns, y) => {
            let line = "";
            columns.forEach((tile, x) => {
                if (x !== this.currentPosition.x || y !== this.currentPosition.y) {
                    line += tile;
                } else {
                    line += this.facing
                }
            })
            console.log(line)
        })
    }

    getTilePositionAfterVoidTile(){
        const x = this.currentPosition.x, y = this.currentPosition.y;
        let columns = null;
        let lines = null;
        switch (this.facing){
            case ">":
                columns = this.matrix[y];
                for(let x = 0 ; x< columns.length;x++){
                    if(tileType[this.matrix[y][x]] !== "void"){
                        return {x,y};
                    }
                }
                break;
            case "v":
                lines = this.matrix.map(columns=>columns[x]);
                for(let y = 0 ; y< lines.length;y++){
                    if(tileType[this.matrix[y][x]] !== "void"){
                        return {x,y};
                    }
                }
                break;
            case "<":
                columns = this.matrix[y];
                for(let x = columns.length-1 ; x>=0;x--){
                    if(tileType[this.matrix[y][x]] !== "void"){
                        return {x,y};
                    }
                }
                break;
            case "^":
                lines = this.matrix.map(columns=>columns[x]);
                for(let y = lines.length-1 ; y>=0;y--){
                    if(tileType[this.matrix[y][x]] !== "void"){
                        return {x,y};
                    }
                }
                break;
        }
        throw "Invalid"
    }

    getTheFinalPassword(){
        let facingNumber = NaN;
        switch (this.facing){
            case ">":
                facingNumber = 0
                break
            case "v":
                facingNumber = 1
                break
            case "<":
                facingNumber = 2
                break
            case "^":
                facingNumber = 3
                break
        }
        return 1000*(this.currentPosition.y+1)+4*(this.currentPosition.x+1)+facingNumber
    }
}

const tileType = {
    undefined: "void",
    " ": "void",
    ">": "ground",
    "v": "ground",
    "<": "ground",
    "^": "ground",
    ".": "ground",
    "#": "wall"
}


class ScenarioPart2 extends Scenario{
    faceSize
    faces = [
        // ⚠️ manually setting faces TODO make it automatic for any data
        {i: 1, j: 0, ">": null, "v": null, "<": null, "^": null,},//0
        {i: 2, j: 0, ">": null, "v": null, "<": null, "^": null,},//1
        {i: 1, j: 1, ">": null, "v": null, "<": null, "^": null,},//2
        {i: 0, j: 2, ">": null, "v": null, "<": null, "^": null,},//3
        {i: 1, j: 2, ">": null, "v": null, "<": null, "^": null,},//4
        {i: 0, j: 3, ">": null, "v": null, "<": null, "^": null,},//5
    ]

    constructor(filteredData) {
        super(filteredData);
        this.faceSize = this._getFaceSize()

        // ⚠️ manually connect faces TODO make it automatic for any data
        this.faces[0][">"] = this.faces[1];this.faces[0]["v"] = this.faces[2];this.faces[0]["<"] = this.faces[3];this.faces[0]["^"] = this.faces[5];
        this.faces[1][">"] = this.faces[4];this.faces[1]["v"] = this.faces[2];this.faces[1]["<"] = this.faces[0];this.faces[1]["^"] = this.faces[5];
        this.faces[2][">"] = this.faces[1];this.faces[2]["v"] = this.faces[4];this.faces[2]["<"] = this.faces[3];this.faces[2]["^"] = this.faces[0];
        this.faces[3][">"] = this.faces[4];this.faces[3]["v"] = this.faces[5];this.faces[3]["<"] = this.faces[0];this.faces[3]["^"] = this.faces[2];
        this.faces[4][">"] = this.faces[1];this.faces[4]["v"] = this.faces[5];this.faces[4]["<"] = this.faces[3];this.faces[4]["^"] = this.faces[2];
        this.faces[5][">"] = this.faces[4];this.faces[5]["v"] = this.faces[1];this.faces[5]["<"] = this.faces[0];this.faces[5]["^"] = this.faces[3];
    }

    _getFaceSize(){
        let minLineWidth = Infinity;
        let minColumnHeight = Infinity;
        const lines = this.matrix;
        const columns = []
        lines.forEach((line)=>{
            let lineWidth = 0;
            line.forEach((tile,x)=>{
                if(!columns[x]){
                    columns[x] = []
                }
                columns[x].push(tile);
                if(tileType[tile] !== "void"){
                    lineWidth++;
                }

            })
            if(lineWidth<minLineWidth){
                minLineWidth = lineWidth;
            }
        })

        columns.forEach((column)=>{
            let columnHeight = 0;
            column.forEach((tile)=>{
                columns.push(tile);
                if(tileType[tile] !== "void"){
                    columnHeight++;
                }

            })
            if(columnHeight<minColumnHeight){
                minColumnHeight = columnHeight;
            }
        })

        return {width: minLineWidth, height: minColumnHeight}
    }

    getCurrentFaceCords(){
        return {
            i: Math.trunc(this.currentPosition.x/this.faceSize.width),
            j: Math.trunc(this.currentPosition.y/this.faceSize.height)
        }
    }
    getFaceCoordsOf(position){
        const fixNegativeX = position.x<0?-1:0
        const fixNegativeY = position.y<0?-1:0
        return {
            i: Math.trunc(position.x/this.faceSize.width)+fixNegativeX,
            j: Math.trunc(position.y/this.faceSize.height)+fixNegativeY
        }
    }

    moveForward(){
        let nextPosition = this.getNextPosition();
        let nextFacing = this.facing;
        let nextTile = this.matrix[nextPosition.y]?.[nextPosition.x];
        if(tileType[nextTile] === "void"){
            [nextPosition, nextFacing] = this.getTilePositionAndFacingAfterVoidTile()
            nextTile = this.matrix[nextPosition.y]?.[nextPosition.x];
        }
        switch(tileType[nextTile]){
            case "ground":
                this.currentPosition = nextPosition;
                this.facing = nextFacing;
                this.matrix[this.currentPosition.y][this.currentPosition.x] = this.facing
                return;
            case "wall":
                return;
        }
        throw "Invalid"
    }

    getTilePositionAndFacingAfterVoidTile(){
        const x = this.currentPosition.x, y = this.currentPosition.y;
        let nextPosition = this.getNextPosition();
        let nextTile = this.matrix[nextPosition.y]?.[nextPosition.x];
        if(tileType[nextTile] !== "void"){
            console.log(nextTile)
            throw "Invalid"
        }
        const currentFaceCoords = this.getCurrentFaceCords()
        const nextFaceCoords = this.fixFace(this.getFaceCoordsOf(nextPosition))
        const currentFace = this.faces.find(face=>face.i === currentFaceCoords.i && face.j === currentFaceCoords.j);
        const nextFace = this.faces.find(face=>face.i === nextFaceCoords.i && face.j === nextFaceCoords.j);
        const forwardConnection = this.faceConnectionOf(currentFace,nextFace)
        const backwardConnection = this.faceConnectionOf(nextFace,currentFace)
        const currentRelativePosition = this.getRelativePosition(this.currentPosition,currentFaceCoords)

        const fc = forwardConnection;
        const bc = backwardConnection;
        const crp = currentRelativePosition;


        if( this.faceSize.width!== this.faceSize.height)
            throw "Invalid"
        const maxFaceIndex = this.faceSize.width-1
        let newFacing = null;
        switch (bc){
            case ">":
                newFacing = "<"
                break;
            case "v":
                newFacing = "^"
                break;
            case "<":
                newFacing = ">"
                break;
            case "^":
                newFacing = "v"
                break;
        }
        let newPosition = null
        if(fc === ">" && bc === ">")
            newPosition = this.getAbsolutePosition({x:crp.x,y:maxFaceIndex-crp.y},nextFaceCoords)
        if(fc === ">" && bc === "v")
            newPosition = this.getAbsolutePosition({x:crp.y,y:crp.x},nextFaceCoords)
        if(fc === ">" && bc === "<")
            newPosition = this.getAbsolutePosition({x:maxFaceIndex-crp.x,y:crp.y},nextFaceCoords)
        if(fc === ">" && bc === "^")
            newPosition = this.getAbsolutePosition({x:maxFaceIndex-crp.y,y:maxFaceIndex-crp.x},nextFaceCoords)

        if(fc === "v" && bc === ">")
            newPosition = this.getAbsolutePosition({x:crp.y,y:crp.x},nextFaceCoords)
        if(fc === "v" && bc === "v")
            newPosition = this.getAbsolutePosition({x:maxFaceIndex-crp.x,y:crp.y},nextFaceCoords)
        if(fc === "v" && bc === "<")
            newPosition = this.getAbsolutePosition({x:maxFaceIndex-crp.y,y:maxFaceIndex-crp.x},nextFaceCoords)
        if(fc === "v" && bc === "^")
            newPosition = this.getAbsolutePosition({x:crp.x,y:maxFaceIndex-crp.y},nextFaceCoords)

        if(fc === "<" && bc === ">")
            newPosition = this.getAbsolutePosition({x:maxFaceIndex-crp.x,y:crp.y},nextFaceCoords)
        if(fc === "<" && bc === "v")
            newPosition = this.getAbsolutePosition({x:maxFaceIndex-crp.y,y:maxFaceIndex-crp.x},nextFaceCoords)
        if(fc === "<" && bc === "<")
            newPosition = this.getAbsolutePosition({x:crp.x,y:maxFaceIndex-crp.y},nextFaceCoords)
        if(fc === "<" && bc === "^")
            newPosition = this.getAbsolutePosition({x:crp.y,y:crp.x},nextFaceCoords)

        if(fc === "^" && bc === ">")
            newPosition = this.getAbsolutePosition({x:maxFaceIndex-crp.y,y:maxFaceIndex-crp.x},nextFaceCoords)
        if(fc === "^" && bc === "v")
            newPosition = this.getAbsolutePosition({x:crp.x,y:maxFaceIndex-crp.y},nextFaceCoords)
        if(fc === "^" && bc === "<")
            newPosition = this.getAbsolutePosition({x:crp.y,y:crp.x},nextFaceCoords)
        if(fc === "^" && bc === "^")
            newPosition = this.getAbsolutePosition({x:maxFaceIndex-crp.x,y:crp.y},nextFaceCoords)
        if(newPosition !== null && newFacing !== null){
            return [newPosition, newFacing]
        }
        throw "Invalid";
    }

    faceConnectionOf(faceA,faceB){
        const connections = [">","v","<","^"];
        for(let i = 0; i<connections.length;i++){
            const connectedFace = faceA[connections[i]];
            if(connectedFace.i === faceB.i && connectedFace.j === faceB.j){
                return connections[i];
            }
        }
        throw "Invalid"
    }

    getRelativePosition(position, faceCoords){
        return {
            x: position.x-faceCoords.i*this.faceSize.width,
            y: position.y-faceCoords.j*this.faceSize.height
        }
    }
    getAbsolutePosition(relativePosition, faceCoords){
        return {
            x: relativePosition.x+faceCoords.i*this.faceSize.width,
            y: relativePosition.y+faceCoords.j*this.faceSize.height
        }
    }

    fixFace(face){
        if(this.faces.findIndex(faceA=>faceA.i === face.i && faceA.j === face.j)>=0){
            throw "Invalid! face is already correct"
        }
        let desiredFace = null;
        switch (this.facing){
            case ">":
                desiredFace = {i: face.i-1, j:face.j}
                break;
            case "v":
                desiredFace = {i: face.i, j:face.j-1}
                break;
            case "<":
                desiredFace = {i: face.i+1, j:face.j}
                break;
            case "^":
                desiredFace = {i: face.i, j:face.j+1}
                break;
        }
        const foundFace = this.faces.find(faceA=>faceA.i === desiredFace.i && faceA.j === desiredFace.j);
        if(foundFace){
            return {i: foundFace[this.facing].i,j: foundFace[this.facing].j}
        }
        throw "Invalid"
    }


}

module.exports = {Scenario,ScenarioPart2}