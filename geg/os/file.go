package main


import (
    "gitee.com/johng/gf/g/os/gfile"
    "fmt"
)

var dirpath1  = "/home/john/Workspace/temp/"
var dirpath2  = "/home/john/Workspace/temp/1"
var filepath1 = "/home/john/Workspace/temp/test.php"
var filepath2 = "/tmp/tmp.test"


type BinData struct{
    name string
    age  int
}

func info () {
    fmt.Println(gfile.Info(dirpath1))
}

func scanDir() {
    files := gfile.ScanDir(dirpath1)
    fmt.Println(files)
}

func getContents() {
    fmt.Printf("%s\n", gfile.GetContents(filepath1))
}

func putContents() {
    fmt.Println(gfile.PutContentsAppend(filepath2, []byte("123")))
}

func putBinContents() {
    data := []byte(BinData{"john", 31})
    fmt.Println(gfile.PutContents(filepath2, data))
}

func main() {
    //info()
    //getContents()
    //putContents()
    putBinContents()
    //scanDir()
}