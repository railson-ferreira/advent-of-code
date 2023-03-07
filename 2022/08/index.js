const fs = require('fs');
const assert = require("assert");

// Advent of Code 2022 - day 8

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
    data.split("\n").filter(item => item).forEach(line => {
        matrix.push([...line].map(item => Number(item)))
    })
    const visibleTreesFromTheLeftSide = visibleTreesHorizontally(matrix, "ltr");
    const visibleTreesFromTheRightSide = visibleTreesHorizontally(matrix, "rtl");
    const visibleTreesFromTheTopSide = visibleTreesVertically(matrix, "ttb");
    const visibleTreesFromTheBottomSide = visibleTreesVertically(matrix, "btt");

    const allVisibleTreeWithRedundancy = [
        ...visibleTreesFromTheLeftSide,
        ...visibleTreesFromTheRightSide,
        ...visibleTreesFromTheTopSide,
        ...visibleTreesFromTheBottomSide,
    ]

    const allVisibleTree = allVisibleTreeWithRedundancy.filter((tree, index, array)=>{
        return array.findIndex((treeB)=>treeB.line === tree.line && treeB.column === tree.column) === index
    })
    console.log(allVisibleTree.length);
}

///////////////////// CHALlENGE 2 /////////////////

function challenge2(data) {
    const matrix = [];
    data.split("\n").filter(item => item).forEach(line => {
        matrix.push([...line].map(item => Number(item)))
    })
    const columnsNumber = matrix.length;
    assert(matrix.every(line => line.length === columnsNumber), "Invalid Matrix lines number != columns number");

    let highestScenicScore = 0
    for (let line = 0; line < matrix.length; line++) {
        for (let column = 0; column < matrix[0].length; column++) {
            const treeHeight = matrix[line][column]
            let topScoreSum = 0;
            let leftScoreSum = 0;
            let bottomScoreSum = 0;
            let rightScoreSum = 0;
            for(let x = line-1 ; x >=0; x--){ // Looking up
                const anotherTreeHeight = matrix[x][column]
                if(anotherTreeHeight>=treeHeight){
                    topScoreSum++;
                    break;
                }
                topScoreSum++;
            }
            for(let x = column-1 ; x >=0; x--){ // Looking left
                const anotherTreeHeight = matrix[line][x]
                if(anotherTreeHeight>=treeHeight){
                    leftScoreSum++;
                    break;
                }
                leftScoreSum++;
            }
            for(let x = line+1 ; x <matrix.length; x++){ // Looking down
                const anotherTreeHeight = matrix[x][column]
                if(anotherTreeHeight>=treeHeight){
                    bottomScoreSum++;
                    break;
                }
                bottomScoreSum++;
            }
            for(let x = column+1 ; x <matrix[0].length; x++){ // Looking right
                const anotherTreeHeight = matrix[line][x]
                if(anotherTreeHeight>=treeHeight){
                    rightScoreSum++;
                    break;
                }
                rightScoreSum++;
            }
            const scenicScore = topScoreSum*leftScoreSum*bottomScoreSum*rightScoreSum
            if(scenicScore>highestScenicScore){
                highestScenicScore = scenicScore;
            }
        }
    }
    console.log(highestScenicScore)
}

// Utils Functions

function visibleTreesHorizontally(matrix, dir) {
    const columnsNumber = matrix.length;
    assert(matrix.every(line => line.length === columnsNumber), "Invalid Matrix lines number != columns number");
    assert(dir, "A direction was not specified as a second parameter")
    const trees = [];
    switch (dir) {
        case "ltr":
            for (let line = 0; line < matrix.length; line++) {
                let lineTrees = []
                for (let column = 0; column < matrix[0].length; column++) {
                    const treeHeight = matrix[line][column]
                    if (column === 0) {
                        lineTrees.push({line, column, height: treeHeight});
                    } else if (isEveryTreeSmallerThan(lineTrees, treeHeight)) {
                        lineTrees.push({line, column, height: treeHeight});
                    } else {
                    }
                }
                trees.push(...lineTrees)
            }
            break;
        case "rtl":
            for (let line = 0; line < matrix.length; line++) {
                const lineTrees = []
                for (let column = matrix[0].length - 1; column >= 0; column--) {
                    const treeHeight = matrix[line][column]
                    if (column === matrix[0].length - 1) {
                        lineTrees.push({line, column, height: treeHeight});
                    } else if (isEveryTreeSmallerThan(lineTrees, treeHeight)) {
                        lineTrees.push({line, column, height: treeHeight});
                    }
                }
                trees.push(...lineTrees)
            }
            break;
        default:
            throw "Invalid Direction"
    }
    return trees;
}

function visibleTreesVertically(matrix, dir) {
    const columnsNumber = matrix.length;
    assert(matrix.every(line => line.length === columnsNumber), "Invalid Matrix lines number != columns number");
    assert(dir, "A direction was not specified as a second parameter");
    const trees = [];

    switch (dir) {
        case "ttb":
            for (let column = 0; column < matrix.length; column++) {
                let columnTrees = []
                for (let line = 0; line < matrix[0].length; line++) {
                    const treeHeight = matrix[line][column]
                    if (line === 0) {
                        columnTrees.push({line, column, height: treeHeight});
                    } else if (isEveryTreeSmallerThan(columnTrees, treeHeight)) {
                        columnTrees.push({line, column, height: treeHeight});
                    }
                }
                trees.push(...columnTrees)
            }
            break;
        case "btt":
            for (let column = 0; column < matrix.length; column++) {
                const columnTrees = []
                for (let line = matrix.length - 1; line >= 0; line--) {
                    const treeHeight = matrix[line][column]
                    if (line === matrix.length - 1) {
                        columnTrees.push({line, column, height: treeHeight});
                    } else if (isEveryTreeSmallerThan(columnTrees, treeHeight)) {
                        columnTrees.push({line, column, height: treeHeight});
                    }
                }
                trees.push(...columnTrees)
            }
            break;
        default:
            throw "Invalid Direction"
    }
    return trees;
}

function isEveryTreeSmallerThan(trees, height) {
    const result = !trees.find(tree => tree.height >= height);
    return !trees.find(tree => tree.height >= height)
}