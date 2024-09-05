package DBReader

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"os"
)

type recipes struct{
	Cake []struct{
		Name string `json:"name" xml:"name"`
		Time string `json:"time" xml:"stovetime"`
		Ingredients []struct{
			IngName string `json:"ingredient_name" xml:"itemname"`
			IngCount string `json:"ingredient_count" xml:"itemcount"`
			IngUnit string `json:"ingredient_unit,omitempty" xml:"itemunit"`
		} `json:"ingredients" xml:"ingredients>item"`
	} `json:"cake" xml:"cake"`
}

type JSON struct{
}
type XML struct{
}

func (j *JSON) recipy(dat []byte) recipes{
	var data recipes
	err := json.Unmarshal(dat, &data)
	if err != nil{
		log.Fatalln(err)
	}
	return data
}
func (x *XML) recipy(dat []byte) recipes{
	var data recipes
	err := xml.Unmarshal(dat, &data)
	if err != nil{
		log.Fatalln(err)
	}
	return data
}

type DBReader interface{
	recipy(dat []byte)recipes
}

func dbRead(reader DBReader, filename string) recipes{
	dat, err := os.ReadFile(filename)
	if err != nil{
		log.Fatalln("Can't read the file")
	}
	recipe := reader.recipy(dat)
	return recipe
}