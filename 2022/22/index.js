const fs = require('fs');
const {Scenario,ScenarioPart2} = require("./scenario")

// Advent of Code 2022 - day 22

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
    let filteredScenarioData = "";
    let reachEndOfScenario = false
    let instructionsLine = null
    data.split("\n").forEach(line=>{
        if(instructionsLine!=null && line !== ""){
            throw "Invalid";
        }
        if(line === ""){
            reachEndOfScenario = true;
        }else if(!reachEndOfScenario){
            filteredScenarioData += line+"\n"
        }else{
            instructionsLine = line;
        }
    })
    if(instructionsLine==null){
        throw "Invalid";
    }
    const scenario = new Scenario(filteredScenarioData)
    const instructions = []
    instructionsLine.split(/[LR]/).forEach((numberStr,index)=> {
        if (index === 0 && instructionsLine.indexOf(numberStr) !== 0) {
            throw "Invalid"
        }
        instructions[index*2] = Number(numberStr);
    })
    instructionsLine.split(/\d+/).filter(item=>item).forEach((turn,index)=> {
        instructions[index*2+1] = turn;
    })
    instructions.forEach(instruction=>{
        scenario.runInstruction(instruction)
    })
    const {x,y} = scenario.currentPosition
    console.log(scenario.getTheFinalPassword());
}


///////////////////// CHALLENGE 2 /////////////////

function challenge2(data) {
    let filteredScenarioData = "";
    let reachEndOfScenario = false
    let instructionsLine = null
    data.split("\n").forEach(line=>{
        if(instructionsLine!=null && line !== ""){
            throw "Invalid";
        }
        if(line === ""){
            reachEndOfScenario = true;
        }else if(!reachEndOfScenario){
            filteredScenarioData += line+"\n"
        }else{
            instructionsLine = line;
        }
    })
    if(instructionsLine==null){
        throw "Invalid";
    }
    const scenario = new ScenarioPart2(filteredScenarioData)
    const instructions = []
    instructionsLine.split(/[LR]/).forEach((numberStr,index)=> {
        if (index === 0 && instructionsLine.indexOf(numberStr) !== 0) {
            throw "Invalid"
        }
        instructions[index*2] = Number(numberStr);
    })
    instructionsLine.split(/\d+/).filter(item=>item).forEach((turn,index)=> {
        instructions[index*2+1] = turn;
    })
    instructions.forEach(instruction=>{
        scenario.runInstruction(instruction)
    })
    console.log(scenario.getTheFinalPassword());
}
