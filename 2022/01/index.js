const fs = require('fs');

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
    let maxSum = null;
    const notesPerElf = data.split("\n\n");
    notesPerElf.forEach(elfNotes=>{
        const calories = elfNotes.split("\n").map(item=>Number(item));
        const sum = calories.reduce((a,b)=>a+b)
        if(maxSum==null){
            maxSum = sum;
        }else if(maxSum<sum){
            maxSum = sum;
        }
    })
    console.log(maxSum)
}

function challenge2(data){
    let topThree = [];
    const notesPerElf = data.split("\n\n");
    const sums = notesPerElf.map(elfNotes=>{
        const calories = elfNotes.split("\n").map(item=>Number(item));
        return calories.reduce((a,b)=>a+b)
    })
    sums.forEach(sum=>{
        if(topThree.length<3){
            topThree.push(sum)
            topThree.sort((a,b)=>a-b)
        }else if(topThree[0]<sum) {
            topThree[0]=sum
            topThree.sort((a,b)=>a-b)
        }
    })
    console.log(topThree)
    console.log(topThree.reduce((a,b)=>a+b))
}