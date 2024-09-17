package compare

import (
	"fmt"
	"os"
	"bufio"
)

func getMap(file string) map[string]bool{
	dic := make(map[string]bool)
	File, err := os.Open(file)
	if err!= nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer File.Close()
	scanner := bufio.NewScanner(File)
	for scanner.Scan(){
		line := scanner.Text()
		dic[line] = true
	}
	return dic
}

func readFile(old string, new string){
	OldDic:=getMap(old)
	NewDic := getMap(new)

	for line, b := range OldDic{
		if NewDic[line] != b{
			fmt.Println("REMOVED ", line)
		}
	}
	for line, b := range NewDic{
		if OldDic[line] != b{
			fmt.Println("ADDED ", line)
		}
	}
}

func Compare(){
	Flags := NewFlag()
	Flags.Parse()

	if *Flags.Old != "" && *Flags.New != ""{
		readFile(*Flags.Old, *Flags.New)
	}else{
		fmt.Println("Usage:\n\tgo run <main.go> --old <filename.txt> --new <filename.txt>")
		os.Exit(1)
	}
}