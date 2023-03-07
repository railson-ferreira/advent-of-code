const monkeysOperations = require("./monkeys-operations")

module.exports = (data)=>{
    const monkeysData = []
    data.split("\n\n").filter(item=>item).map(item=>item.trim()).forEach(item=>{
        const monkeyDataRegex = /Monkey (\d+):\n.*Starting items: ((\d+,?.*)+)\n.*Operation: .*\n.*Test: divisible by (\d+)\n.*If true: throw to monkey (\d+)\n.*If false: throw to monkey (\d+)/g
        const regexResult = monkeyDataRegex.exec(item);
        const [_, monkeyIdStr, startingItemsStr, __,testDividerStr, monkeyIdIfTrueStr, monkeyIdIfFalseStr] = regexResult;
        const monkeyId = Number(monkeyIdStr);
        const monkeyData = {
            monkeyId,
            items: startingItemsStr.split(",").map(item=>Number(item)),
            operation: monkeysOperations[monkeyId],
            testDivider: Number(testDividerStr),
            monkeyIdIfTrue: Number(monkeyIdIfTrueStr),
            monkeyIdIfFalse: Number(monkeyIdIfFalseStr)
        }
        monkeysData.push(monkeyData)
    })
    return monkeysData;
}