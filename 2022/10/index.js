const fs = require('fs');
const assert = require("assert");

// Advent of Code 2022 - day 10

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
    const commands = [];
    data.split("\n").filter(item => item).forEach(line => {
        const commandRegex = /(addx|noop) ?(-?\d+)?/g
        const regexResult = commandRegex.exec(line);
        if(regexResult == null){
            debugger;
        }
        const [_, command, parameter] = regexResult
        const value = Number(parameter)
        commands.push({command, value})
    })

    let register = 1;
    let cycle = 0;
    let currentCommand = 0
    const log = []
    while(currentCommand<commands.length){
        cycle++;
        if(cycle >= 20){
            const test = (cycle-20)%40 === 0
            if(test){
                log.push({cycle, register});
            }
        }
        const commandIndex = Math.floor(currentCommand);
        const command = commands[commandIndex];
        if(!command){
            debugger
        }
        switch (command.command){
            case "addx":
                if(commandIndex !== currentCommand){
                    register += command.value;
                }
                currentCommand =currentCommand+0.5;
                break;
            case "noop":
                currentCommand++;
                break;
        }
    }
    console.log(log.map(item=>item.cycle*item.register).reduce((a,b)=>a+b));
}

///////////////////// CHALLENGE 2 /////////////////

function challenge2(data) {
    const commands = [];
    data.split("\n").filter(item => item).forEach(line => {
        const commandRegex = /(addx|noop) ?(-?\d+)?/g
        const regexResult = commandRegex.exec(line);
        if(regexResult == null){
            debugger;
        }
        const [_, command, parameter] = regexResult
        const value = Number(parameter)
        commands.push({command, value})
    })

    let register = 1;
    let cycle = 0;
    let currentCommand = 0
    const frame = getABlackFrame()
    while(currentCommand<commands.length){
        const line = Math.floor(cycle/40)
        const column = cycle%40
        const pixelPosition = column+1;
        if(pixelPosition>=register && pixelPosition <=register+2){
            frame[line][column] = "#"
        }
        cycle++;
        const commandIndex = Math.floor(currentCommand);
        const command = commands[commandIndex];
        if(!command){
            debugger
        }
        switch (command.command){
            case "addx":
                if(commandIndex !== currentCommand){
                    register += command.value;
                }
                currentCommand =currentCommand+0.5;
                break;
            case "noop":
                currentCommand++;
                break;
        }
    }
    console.log(renderFrame(frame));
    console.log("========================================")
}

// Utils Functions
function getABlackFrame(){
    return Array.from({length: 6}, ()=>Array.from({length: 40},()=>"."));
}

function renderFrame(frame){
    let render = ""
    frame.forEach(line=>{
        render+=line.join("");
        render+="\n"
    })
    return render;
}