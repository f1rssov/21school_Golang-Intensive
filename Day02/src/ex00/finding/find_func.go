package finding

import (
	"fmt"
	"path/filepath"
	"os"
	"io/fs"
	"strings"
)

func Find(mypath string, flags Flags){

	err := filepath.Walk(mypath, func(path string, info os.FileInfo, err error) error {
			if os.IsPermission(err){
				return filepath.SkipDir
			} else if err !=nil{
				fmt.Println(err)
				return err
			}
			if *flags.Dir && info.IsDir(){
				fmt.Println("/"+path)
			}
			if *flags.File && *flags.Ext == "" && info.Mode().IsRegular(){
				fmt.Println("/"+path)
			}else if (*flags.File && *flags.Ext != "") && info.Mode().IsRegular(){
				if strings.HasSuffix(path, "." + *flags.Ext){
					fmt.Println("/"+path)
				}
			}
			if *flags.Sl{
				if info.Mode().Type() == fs.ModeSymlink {
					target, err := filepath.EvalSymlinks(path)
					if err != nil{
						fmt.Println("/"+path+" -> [broken]" )
					}else{
						fmt.Println("/"+path+" -> " + target)
					}
				}
			}

			if !*flags.Dir && !*flags.File && *flags.Ext == "" && !*flags.Sl{
				if info.Mode().Type() == fs.ModeSymlink {
					target, err := filepath.EvalSymlinks(path)
					if err != nil{
						fmt.Println("/"+path+" -> [broken]" )
					}else{
						fmt.Println("/"+path+" -> " + "/" + target)
					}
				}else{
					fmt.Println("/"+path)
				}
			}
		return nil
	})
	if err != nil{
		fmt.Println(err)
	}
}