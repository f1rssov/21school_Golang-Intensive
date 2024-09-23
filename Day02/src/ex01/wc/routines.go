package wc

import(
	"fmt"
	"sync"
	"os"
	"bufio"
)

func calc(flag *Flags, filename string, wg *sync.WaitGroup){
	defer wg.Done()
	file, err := os.Open(filename);
	if err !=nil{
		fmt.Println(err)
		return
	}else{
		defer file.Close()
	}
	switch true{
	case *flag.Words:
		scanner := bufio.NewScanner(file)
		count :=0
		for scanner.Scan(){
			for _, ch := range scanner.Text(){
				if ch == ' '{
					count++
				}
			}
			count++
		}
		fmt.Printf("%d\t%s\n", count, filename)
	case *flag.Lines:
		scanner := bufio.NewScanner(file)
		count :=0
		for scanner.Scan(){
			count++
		}
		fmt.Printf("%d\t%s\n", count, filename)
	case *flag.M:
		scanner := bufio.NewScanner(file)
		count :=0
		lines :=0
		for scanner.Scan(){
			lines++
			for _, ch:= range scanner.Text(){
				ch++
				count++
			}
		}
		count += lines-1
		fmt.Printf("%d\t%s\n", count, filename)
	default:
		scanner := bufio.NewScanner(file)
		count :=0
		for scanner.Scan(){
			for _, ch := range scanner.Text(){
				if ch == ' '{
					count++
				}
			}
			count++
		}
		fmt.Printf("%d\t%s\n", count, filename)
	}
}

func GoRout(flag *Flags, args []string){
	wg := new(sync.WaitGroup)
	if (*flag.Words || *flag.Lines || *flag.M){
		args = args[1:]
	}
	for _, arg := range args{
		wg.Add(1)
		go calc(flag, arg, wg)
	}
	wg.Wait()
}