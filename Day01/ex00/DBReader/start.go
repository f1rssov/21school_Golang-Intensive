package DBReader

import (
	"flag"
	"log"
	"os"
	"strings"
)


func Start(){
	Flags := newFlag()
	Flags.parse()

	var filename string
	
	if *Flags.F{
		if len(flag.Args()) < 1 && len(flag.Args()) > 1{
			log.Fatalln("\nUsage:\n\t-f file_name.xml\nor:\n\t-f file_name.json")
		}
		filename = flag.Args()[0]
		isJson := strings.HasSuffix(filename, ".json")
		isXml := strings.HasSuffix(filename, ".xml")
		if !isJson && !isXml {
			log.Fatalln("\nUsage:\n\t-f file_name.xml\nor:\n\t-f file_name.json")
		}
		dat, err := os.ReadFile(filename)
		if err != nil{
			log.Fatalln("ERROR: Can't read the file")
		}
		dbRead(dat, isJson, isXml)
	} else{
		log.Fatalln("\nUsage:\n\t-f file_name.xml\nor:\n\t-f file_name.json")
	}
}