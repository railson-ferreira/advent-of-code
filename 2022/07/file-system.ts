export type FileSystem = { [name: string]: File|Folder }

export type File = {
    name: string
    size: number
}

export type Folder = {
    name: string
    fileSystem: FileSystem
    totalSize: ()=>number
}