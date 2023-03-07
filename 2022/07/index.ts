import * as fs from "fs";
import {File, Folder, FileSystem} from "./file-system";

// Advent of Code 2022 - day 7


fs.readFile('input.txt', 'utf8', (err, data) => {
    if (err) {
        console.error(err);
        return;
    }
    challenge1(data)
    console.log("=======================")
    challenge2(data)
});


function challenge1(data: string) {
    const lines = data.split("\n").filter(item => item)
    const rootFileSystem: FileSystem = {};
    const rootFolder: Folder = {name: "/", fileSystem: rootFileSystem, totalSize: () => getFileSystemSize(rootFileSystem)}
    let currentFolder: Folder = rootFolder
    for(let x = 0 ; x< lines.length; x++){
        const line = lines[x];
        if(line.includes("$")){
            const [_,command,parameter] = line.split(" ");
            if(command == "ls"){
                x++
                let nextLine = lines[x]
                let data =""
                while (nextLine && !nextLine.includes("$")){
                    data += nextLine;
                    data +="\n"
                    x++;
                    nextLine = lines[x]
                }
                x--;
                currentFolder = handleCommand(command, data, currentFolder) ?? currentFolder
            }else {
                currentFolder = handleCommand(command, parameter, currentFolder) ?? currentFolder
            }
        }else{
            console.log(`something went wrong -> ${line}`);
        }
    }
    removeCircularPaths(rootFolder)
    const specificArray:{folderName: string, folderSize: number}[] = []
    addSpecificFoldersToArray1(rootFolder, specificArray);
    console.log(specificArray.map(item=>item.folderSize).reduce((a,b)=>a+b));
}

///////////////////// CHALlENGE 2 /////////////////

function challenge2(data: string ) {
    const lines = data.split("\n").filter(item => item)
    const rootFileSystem: FileSystem = {};
    const rootFolder: Folder = {name: "/", fileSystem: rootFileSystem, totalSize: () => getFileSystemSize(rootFileSystem)}
    let currentFolder: Folder = rootFolder
    for(let x = 0 ; x< lines.length; x++){
        const line = lines[x];
        if(line.includes("$")){
            const [_,command,parameter] = line.split(" ");
            if(command == "ls"){
                x++
                let nextLine = lines[x]
                let data =""
                while (nextLine && !nextLine.includes("$")){
                    data += nextLine;
                    data +="\n"
                    x++;
                    nextLine = lines[x]
                }
                x--;
                currentFolder = handleCommand(command, data, currentFolder) ?? currentFolder
            }else {
                currentFolder = handleCommand(command, parameter, currentFolder) ?? currentFolder
            }
        }else{
            console.log(`something went wrong -> ${line}`);
        }
    }
    removeCircularPaths(rootFolder)
    const specificArray:{folderName: string, folderSize: number}[] = []
    const totalSpace = 70000000;
    const totalUsedSpace =rootFolder.totalSize();
    const availableSpace = totalSpace-totalUsedSpace
    const spaceRequired = 30000000;
    const needToFree = spaceRequired - availableSpace;
    addSpecificFoldersToArray2(rootFolder, specificArray,needToFree);

    specificArray.sort((a,b)=>a.folderSize-b.folderSize)

    console.log(specificArray[0]);
    // console.log(specificArray.map(item=>item.folderSize));
    debugger
}


function handleCommand(command: string, data: string, folder: Folder): Folder | null {
    switch (command) {
        case "cd":
            if (data == "/") {
                let currentFolder = folder;
                while (currentFolder.name != "/") {
                    currentFolder = currentFolder.fileSystem[".."] as Folder
                }
                return currentFolder;
            } else if (data == "..") {
                return folder.fileSystem[".."] as Folder;
            } else {
                let subFolder: Folder | null = folder.fileSystem[data] as Folder | null;
                if (!subFolder) {
                    subFolder = folder.fileSystem[data] = newFolder(data, folder)
                }
                return subFolder
            }
        case "ls":
            const lines = data.split("\n").filter(item => item);
            lines.forEach(line => {
                if (!line) return;
                if (line.startsWith("dir")) {
                    const folderName = line.split(" ")[1];
                    if (!folder.fileSystem[folderName]) {
                        folder.fileSystem[folderName] = newFolder(folderName, folder)
                    }
                } else {
                    const [fileSize, fileName] = line.split(" ");
                    if (!folder.fileSystem[fileName]) {
                        folder.fileSystem[fileName] = {name: fileName, size: Number(fileSize)}
                    }
                }
            })
            break;
    }
    return null;
}


function newFolder(name: string, parentFolder: Folder): Folder {
    const fileSystem: FileSystem = {"..": parentFolder};
    function totalSize(): number{
        return getFileSystemSize(fileSystem);
    }
    return {name: name, fileSystem, totalSize}
}
function getFileSystemSize(fileSystem: FileSystem): number{
    let sizeSum = 0;
    Object.entries(fileSystem).forEach(([key,value])=>{
        if(key !== ".."){
            sizeSum += getSize(value)
        }
    })
    return sizeSum
}

function getSize(folderOrFile: Folder|File): number{
    if((folderOrFile as any).size !== undefined){
        return (folderOrFile as File).size;
    }else{
        let sizeSum = 0;
        Object.entries((folderOrFile as Folder).fileSystem).forEach(([key,value])=>{
            if(key !== ".."){
                sizeSum += getSize(value)
            }
        })
        return sizeSum
    }
}


function removeCircularPaths(folder: Folder):void{
    const fileSystem = folder.fileSystem;
    delete fileSystem[".."];
    Object.values(fileSystem).forEach(folderOrFile=>{
        if((folderOrFile as any).fileSystem){
            removeCircularPaths((folderOrFile as Folder))
        }
    })
}

function addSpecificFoldersToArray1(folderOrFile: Folder|File, array: Array<{folderName: string, folderSize: number}>): void{
    if((folderOrFile as any).size !== undefined){
        return;
    }
    const folder = folderOrFile as Folder;
    const totalSize = folder.totalSize();
    if(totalSize <= 100000){
        array.push({folderName:folder.name,folderSize: totalSize})
    }
    Object.values(folder.fileSystem).forEach(folderOrFile=>{
        addSpecificFoldersToArray1(folderOrFile, array);
    })
}

function addSpecificFoldersToArray2(folderOrFile: Folder|File, array: Array<{folderName: string, folderSize: number}>, minRequired: number): void{
    if((folderOrFile as any).size !== undefined){
        return;
    }
    const folder = folderOrFile as Folder;
    const totalSize = folder.totalSize();
    if(totalSize >= minRequired){
        array.push({folderName:folder.name,folderSize: totalSize})
    }
    Object.values(folder.fileSystem).forEach(folderOrFile=>{
        addSpecificFoldersToArray2(folderOrFile, array,minRequired);
    })
}