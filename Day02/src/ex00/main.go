package main

import (
	"flag"
	"fmt"
	"os"

	my "github.com/f1rssov/finding"
)


func main(){
	Flags := my.NewFlags()
	Flags.Parse()
	if len(os.Args) >= 2 && len(os.Args) <=7{
		args := flag.Args()
		if (len(args) < 1){
			fmt.Println("ERROR: -ext (Check args)")
			return
		}
		mainPath:=args[0]
		my.Find(mainPath, *Flags)
	}else{
		fmt.Println("Usage:\n\tgo run <main.go> /<dir_name> ")
	}
}