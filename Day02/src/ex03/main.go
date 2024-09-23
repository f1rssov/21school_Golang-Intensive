package main

import (
	"archive/tar"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Flags struct{
	Archive *string
}

func (f *Flags)Parse(){
	flag.Parse()
}

func NewFlags() *Flags{
	return &Flags{
		flag.String("a", "", "archive to"),
	}
}

func main(){
	flags := NewFlags()
	flags.Parse()
	//Для записи пути куда сохранять файл
	path := ""
	if *flags.Archive != ""{
		pathToSave := *flags.Archive
		file,err := os.Stat(pathToSave)
		if err!=nil{
			fmt.Println(err)
			return
		}
		if !file.IsDir(){
			fmt.Println("-a: No such directory")
		}
		if pathToSave[len(pathToSave)-1] != '/'{
			pathToSave += "/"
		}
		path = pathToSave
	}
	var wg sync.WaitGroup
	for _, arg := range flag.Args(){
		wg.Add(1)
		line := strings.Split(arg, "/")
		filen:=line[len(line)-1] //только название файла
		file :=arg //путь и название файла (чтобы узнать абсолютный путь)
		go func(){
			defer wg.Done()
			filepat := ""
			filename := strings.TrimSuffix(filen, ".log") + "_" + strconv.Itoa(int(time.Now().Unix())) + ".tar.gz"
			if path != ""{
				filepat = filepath.Join(path, filename)
			}
			if filepat == ""{
				fpath, _ := filepath.Abs(filepath.Dir(file))
				filepat = filepath.Join(fpath, filename)
			}
			newfile, err := os.Create(filepat)
			if err != nil{
				fmt.Println(err)
				return
			}
			defer newfile.Close()
			tarWriter := tar.NewWriter(newfile)
			defer tarWriter.Close()
			content, _ := os.ReadFile(file)
			header := &tar.Header{
				Name: filename,
				Size: int64(len(content)),
			}
			if err := tarWriter.WriteHeader(header); err!=nil{
				fmt.Println(err)
				return
			}
			if _, err := tarWriter.Write(content); err!=nil{
				fmt.Println(err)
				return
			}
		}()
	}
	wg.Wait()
}