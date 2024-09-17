package DBReader

import (
	"log"
	"strings"
)

func checkFormat(filename string) string{ 
	if strings.HasSuffix(filename, ".json"){
		return "json"
	}else if strings.HasSuffix(filename, ".xml"){
		return "xml"
	}else {
		return ""
	}
}

func getRecipe(filename string) recipes{
	var recipe recipes
	file := checkFormat(filename)
	switch file {
	case "json":
		var j *JSON
		recipe = dbRead(j, filename)
	case "xml":
		var x *XML
		recipe = dbRead(x, filename)
	default:
		log.Fatalln("\nAvaliable format of file\n\t.xml\n\tor\n\t.json")
	}
	return recipe
}

func Start(){
	Flags := newFlag()
	Flags.parse()
	
	if *Flags.Old !="" && *Flags.New !="" {
		old_recipes:= getRecipe(*Flags.Old)
		new_recipes := getRecipe(*Flags.New)
		dbCompare(old_recipes, new_recipes)
	}else{
		log.Fatalln("\nUsage:\n\t--old old_file_name.xml --new new_file_name.json")
	}
}