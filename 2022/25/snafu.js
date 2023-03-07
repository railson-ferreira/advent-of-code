class Snafu{
    value;
    constructor(snafu) {
        if(typeof snafu !== "string"){
            throw "Invalid Snafu number"
        }
        [...snafu].forEach(char=> {
            if (char !== "2" && char !== "1" && char !== "0" && char !== "-" && char !== "=") {
                throw "Invalid Snafu number"
            }
        })
        this.value = snafu
    }

    toString() {
        return this.value;
    }

    static parse(integer){
        if(integer !== Math.trunc(integer)){
            throw "Snafu cannot parse from a non integer value"
        }
        let places = []
        let remain = integer;
        while (remain>0){
            places.push(remain%5)
            remain = Math.trunc(remain/5)
        }
        let sumNext = false;
        places = places.map(place=>{
            const newPlace = sumNext?place+1:place
            if(newPlace>2){
                sumNext = true;
                return newPlace-5;
            }
            sumNext = false;
            return newPlace;
        })
        if(sumNext){
            places.push(1)
        }

        return new Snafu(places.map(place=>{
            switch (place){
                case -2:
                    return "="
                case -1:
                    return "-"
                case 0:
                    return "0"
                case 1:
                    return "1"
                case 2:
                    return "2"
                default:
                    throw "Invalid"
            }
        }).reverse().join(""))
    }

    toNumber(){
        let sum = 0;
        let multiplier = 1;
        [...this.value].map(place=>{
            switch (place){
                case "=":
                    return -2
                case "-":
                    return -1
                case "0":
                    return 0
                case "1":
                    return 1
                case "2":
                    return 2
                default:
                    throw "Invalid"
            }
        }).reverse().forEach(place=>{
            sum +=place*multiplier
            multiplier *= 5
        })
        return sum;
    }
}
module.exports = Snafu