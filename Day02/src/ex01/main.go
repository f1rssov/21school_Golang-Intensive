package main

import (
	"fmt"
	"os"

	myWc "github.com/f1rrsov/wc"
)

func main(){
	Flags := myWc.NewFlags()
	Flags.Parse()

	if len(os.Args) >= 2{
		if(*Flags.M && *Flags.Words || *Flags.Words && *Flags.Lines || *Flags.M && *Flags.Lines){
			fmt.Println("Choose only 1 flag\n\t<-w> or <-l> or <-m>")
			return
		}
		myWc.GoRout(Flags, os.Args[1:])
	}else{
		fmt.Println("Usage\n\t go run main.go <-w> or <-l> or <-m>  <filename> .....")
	}
}