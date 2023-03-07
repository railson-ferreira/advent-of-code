const fs = require('fs');

// Advent of Code 2022 - day 6

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
    let count =0;
    while(count<data.length){
        const letter = data[count];
        if(count>2){
            const characters = []
            for(let x =0 ;x<4;x++){
                characters.push(data[count-x])
            }
            if(characters.length ===uniqueArray(characters).length){
                console.log(count+1);
                return;
            }
        }
        count++;
    }
}

///////////////////// CHALlENGE 2 /////////////////

function challenge2(data) {
    let count =0;
    while(count<data.length){
        const letter = data[count];
        if(count>2){
            const characters = []
            for(let x =0 ;x<14;x++){
                characters.push(data[count-x])
            }
            if(characters.length ===uniqueArray(characters).length){
                console.log(count+1);
                return;
            }
        }
        count++;
    }
}


function uniqueArray(array){
    return array.filter((item,index,array)=> array.indexOf(item)===index);
}
