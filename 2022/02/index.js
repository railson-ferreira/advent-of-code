const fs = require('fs');

// Advent of code - day 2

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
    const rounds = data.split("\n");
    let totalScore = 0;
    rounds.forEach(round=>{
        const shapes = round.split(" ");
        totalScore += calculateScoreChallenge1(shapes[0],shapes[1])
    })
    console.log(totalScore)
}

function calculateScoreChallenge1(enemyChoice, myChoice){
    let score = 0;
    switch (myChoice){
        case "X":
            score +=1;
            switch (enemyChoice){
                case "A":
                    score +=3;
                    break;
                case "B":
                    score += 0;
                    break;
                case "C":
                    score +=6;
                    break;
            }
            break;
        case "Y":
            score +=2;
            switch (enemyChoice){
                case "A":
                    score +=6;
                    break;
                case "B":
                    score += 3;
                    break;
                case "C":
                    score +=0;
                    break;
            }
            break
        case "Z":
            score +=3;
            switch (enemyChoice){
                case "A":
                    score +=0;
                    break;
                case "B":
                    score += 6;
                    break;
                case "C":
                    score +=3;
                    break;
            }
            break
    }
    return score;
}

///////////////////// CHALlENGE 2 /////////////////

function challenge2(data){
    const rounds = data.split("\n");
    let totalScore = 0;
    rounds.forEach(round=>{
        const shapes = round.split(" ");
        totalScore += calculateScoreChallenge2(shapes[0],shapes[1])
    })
    console.log(totalScore)
}


function calculateScoreChallenge2(enemyChoice, roundResult){
    let myChoice;
    switch (enemyChoice) {
        case roundResult === "X"&&"A":
            myChoice = "Z";
            break;
        case roundResult === "Y" && "A":
            myChoice = "X";
            break;
        case roundResult === "Z" && "A":
            myChoice = "Y";
            break;
        case roundResult === "X" && "B":
            myChoice = "X";
            break;
        case roundResult === "Y" && "B":
            myChoice = "Y";
            break;
        case roundResult === "Z" && "B":
            myChoice = "Z";
            break;
        case roundResult === "X" && "C":
            myChoice = "Y";
            break;
        case roundResult === "Y" && "C":
            myChoice = "Z";
            break;
        case roundResult === "Z" && "C":
            myChoice = "X";
            break;
    }
    let score = 0;
    switch (myChoice){
        case "X":
            score +=1;
            switch (enemyChoice){
                case "A":
                    score +=3;
                    break;
                case "B":
                    score += 0;
                    break;
                case "C":
                    score +=6;
                    break;
            }
            break;
        case "Y":
            score +=2;
            switch (enemyChoice){
                case "A":
                    score +=6;
                    break;
                case "B":
                    score += 3;
                    break;
                case "C":
                    score +=0;
                    break;
            }
            break
        case "Z":
            score +=3;
            switch (enemyChoice){
                case "A":
                    score +=0;
                    break;
                case "B":
                    score += 6;
                    break;
                case "C":
                    score +=3;
                    break;
            }
            break
    }
    return score;
}
