const fs = require('fs');

// Advent of Code 2022 - day 21

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
    const rootMonkey = structuredData.find(monkey=>monkey.name === "root")
    rootMonkey.calculateJob()
    console.log(rootMonkey.job);
}


///////////////////// CHALLENGE 2 /////////////////

function challenge2(data) {
    const structuredData = getStructuredData(data, true);
    const rootMonkey = structuredData.find(monkey=>monkey.name === "root")
    rootMonkey.calculateJob()
    console.log(rootMonkey.job);
}

// Utils Functions

function getStructuredData(data, isPart2 = false){
    const monkeys  = []
    data.split("\n").filter(item=>item).forEach(monkey=>{
        const regex = /(\w+): ((\w+ . \w+)|(-?\d+))/
        const result = regex.exec(monkey);
        const [_,name,__,pendingInfo, readyInfo] = result;
        const separatedPendingInfo = pendingInfo?.split(" ") || {}
        const monkeyAName = separatedPendingInfo[0];
        const operation = separatedPendingInfo[1];
        const monkeyBName = separatedPendingInfo[2]
        const monkeyObj = {
            name,
            isReady: isPart2 && name === "humn"?true:!pendingInfo,
            job: isPart2 && name === "humn"?(n)=>n:Number(readyInfo),
            calculateJob(){
                const monkeyA = monkeys.find(monkeyFromList=>monkeyFromList.name === monkeyAName);
                const monkeyB = monkeys.find(monkeyFromList=>monkeyFromList.name === monkeyBName);
                if(isPart2){
                    switch (name){
                        case "root":
                            this.job = newRootMonkeyJobCalculation(monkeyA,operation,monkeyB)
                            break;
                        default:
                            this.job = calculateMonkeyJob(monkeyA,operation,monkeyB)
                            break;
                    }
                }else{
                    this.job = calculateMonkeyJob(monkeyA,operation,monkeyB)
                }
                this.isReady = !!this.job;
            }
        }
        monkeys.push(monkeyObj)
    })
    return monkeys;
}

function calculateMonkeyJob(monkeyA, operation, monkeyB){
    if(!monkeyA.isReady){
        monkeyA.calculateJob()
    }
    if(!monkeyB.isReady){
        monkeyB.calculateJob()
    }
    if(typeof monkeyA.job === "function" ||typeof monkeyB.job === "function"){
        return calculateMonkeyJobInDifferentWay(monkeyA,operation,monkeyB)
    }
    switch (operation){
        case "+":
            return monkeyA.job + monkeyB.job
        case "-":
            return monkeyA.job - monkeyB.job
        case "/":
            return monkeyA.job / monkeyB.job
        case "*":
            return monkeyA.job * monkeyB.job
    }
    throw "Invalid"
}

function calculateMonkeyJobInDifferentWay(readyMonkeyA,operation, readyMonkeyB){
    const hasMonkeyAFunctionAsJob = typeof readyMonkeyA.job === "function"
    const hasMonkeyBFunctionAsJob = typeof readyMonkeyB.job === "function"
    return (n)=>{
        const monkeyAJob = hasMonkeyAFunctionAsJob?readyMonkeyA.job(n):readyMonkeyA.job;
        const monkeyBJob = hasMonkeyBFunctionAsJob?readyMonkeyB.job(n):readyMonkeyB.job;
        switch (operation){
            case "+":
                return monkeyAJob + monkeyBJob
            case "-":
                return monkeyAJob - monkeyBJob
            case "/":
                return monkeyAJob / monkeyBJob
            case "*":
                return monkeyAJob * monkeyBJob
        }
    }
}


function newRootMonkeyJobCalculation(monkeyA, operation, monkeyB){
    if(!monkeyA.isReady){
        monkeyA.calculateJob()
    }
    if(!monkeyB.isReady){
        monkeyB.calculateJob()
    }

    function test(n){
        const monkeyAJob = typeof monkeyA.job === "function"?monkeyA.job(n):monkeyA.job;
        const monkeyBJob = typeof monkeyB.job === "function"?monkeyB.job(n):monkeyB.job;
        return monkeyAJob-monkeyBJob;
    }

    let n = 0;
    let increment = 100000000000;
    let result = test(n);
    while (result!==0){
        if(result<0 && increment>0){
            increment /= -10;
        }
        if(result>0 && increment<0){
            increment /= -10;
        }
        n += increment;
        result = test(n);
    }

    return n
}