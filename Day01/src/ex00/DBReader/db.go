package DBReader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
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

func (j *JSON) recipy(dat []byte) (string, error){
	var data recipes
	err := json.Unmarshal(dat, &data)
	if err != nil{
		log.Fatalln(err)
	}
	byt, err := xml.MarshalIndent(data, "", "    ")
	return string(byt), err
}
func (x *XML) recipy(dat []byte) (string, error){
	var data recipes
	err := xml.Unmarshal(dat, &data)
	if err != nil{
		log.Fatalln(err)
	}
	byt, err := json.MarshalIndent(data, "", "    ")
	return string(byt), err
}

type DBReader interface{
	recipy(dat []byte)(string, error)
}

func dbRead(dat []byte, isJson bool, isXml bool){
	var db DBReader
	var j *JSON
	var x *XML
	if isJson{
		db = j
	}
	if isXml{
		db = x
	}
	result, err := db.recipy(dat)
	if err!=nil{
		log.Fatalln(err)
	}else{
		fmt.Println(result)
	}
}