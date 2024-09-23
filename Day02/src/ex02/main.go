package main

import (
	"fmt"
	"os"
	"os/exec"
	"bufio"
)

func main(){
	comandString := os.Args[1:]
	var args []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		arg := scanner.Text()
		args = append(args, arg)
	}
	if scanner.Err() != nil{
		fmt.Println(scanner.Err())
		return
	}
	for _, ar := range args{
		comandString = append(comandString, ar)
	}

	cmd := exec.Command(comandString[0], comandString[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = "."
	if err:=cmd.Run();err!=nil{
		fmt.Println(err)
		return
	}
}