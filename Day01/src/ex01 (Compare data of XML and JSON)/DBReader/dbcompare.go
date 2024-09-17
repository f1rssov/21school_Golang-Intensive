package DBReader

import (
	"fmt"
	"reflect"
)

func cakeNameExists(old recipes, NewCakeName string) bool{
	for _, name := range old.Cake{
		if(name.Name == NewCakeName){
			return true
		}
	}
	return false
}

func getIndexCake(old recipes, new recipes, i int) int{
	for g, name := range old.Cake{
		if(name.Name == new.Cake[i].Name){
			return g
		}
	}
	return 0
}

func checkTime(old recipes, new recipes, i int) bool{
	j:=0
	for g, name := range old.Cake{
		if(name.Name == new.Cake[i].Name){
			j = g
			break
		}
	}
	return old.Cake[j].Time == new.Cake[i].Time
}

func checkExistIngr(old recipes, new recipes, i int, j int, needCake int) bool{
	for _, ingr := range old.Cake[needCake].Ingredients{
		if(ingr.IngName == new.Cake[i].Ingredients[j].IngName){
			return true
		}
	}
	return false
}

func getIndexIngr(old recipes, new recipes, i int, j int, needCake int) int{
	for g, name := range old.Cake[needCake].Ingredients{
		if(name.IngName == new.Cake[i].Ingredients[j].IngName){
			return g
		}
	}
	return 0
}

func checkIngrCount(old recipes, new recipes, i int, j int, needCake int, needIngr int) bool{
	return old.Cake[needCake].Ingredients[needIngr].IngCount == new.Cake[i].Ingredients[j].IngCount
}

func checkUnit(old recipes, new recipes, i int, j int, needCake int, needIng int) int{
	if old.Cake[needCake].Ingredients[needIng].IngUnit != "" && new.Cake[i].Ingredients[j].IngUnit == ""{
		return 1 //removed
	}else if(old.Cake[needCake].Ingredients[needIng].IngUnit == "" && new.Cake[i].Ingredients[j].IngUnit != ""){
		return 2 //changed
	}else if old.Cake[needCake].Ingredients[needIng].IngUnit != new.Cake[i].Ingredients[j].IngUnit{
		return 3 //changed
	}
	return 4 //okay
}

func dbCompare(old recipes, new recipes){
	if(reflect.DeepEqual(old, new)){
		fmt.Printf("EQUAL")
		return
	}
	//added/removed cake
	for i:=0; i < len(new.Cake); i++{
		if !cakeNameExists(old, new.Cake[i].Name){
			fmt.Printf("ADDED cake \"%s\"\n", new.Cake[i].Name)
		}
	}
	for i:=0; i<len(old.Cake); i++{
		if !cakeNameExists(new, old.Cake[i].Name){
			fmt.Printf("REMOVED cake \"%s\"\n", old.Cake[i].Name)
		}
	}

	//changed cooking time
	for i:=0; i < len(new.Cake); i++{
		if cakeNameExists(old, new.Cake[i].Name){
			if !checkTime(old, new, i){
				fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n",
							new.Cake[i].Name, new.Cake[i].Time, old.Cake[getIndexCake(old, new, i)].Time)
			}
		}
	}

	//added/removed ingredient
	for i:=0; i < len(new.Cake); i++{
		if cakeNameExists(old, new.Cake[i].Name){
			needCake := getIndexCake(old, new, i)
			for j:=0; j<len(new.Cake[i].Ingredients);j++{
				if!checkExistIngr(old, new, i, j, needCake){
					fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n",
								new.Cake[i].Ingredients[j].IngName, old.Cake[needCake].Name)
				}
			}
		}
	}
	for i:=0; i < len(old.Cake); i++{
		if cakeNameExists(new, old.Cake[i].Name){
			needCake := getIndexCake(new, old, i)
			for j:=0; j<len(old.Cake[i].Ingredients);j++{
				if!checkExistIngr(new, old, i, j, needCake){
					fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n",
								old.Cake[i].Ingredients[j].IngName, old.Cake[needCake].Name)
				}
			}
		}
	}

	//changed ingredient count
	for i:=0; i < len(new.Cake); i++{
		if cakeNameExists(old, new.Cake[i].Name){
			needCake := getIndexCake(old, new, i)
			for j:=0; j<len(new.Cake[i].Ingredients);j++{
				if checkExistIngr(old, new, i, j, needCake){
					needIngr := getIndexIngr(old, new, i, j, needCake)
					if checkUnit(old, new, i, j, needCake, needIngr) == 4{
						if !checkIngrCount(old, new, i, j, needCake, needIngr){
							fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n",
										old.Cake[needCake].Ingredients[needIngr].IngName, old.Cake[needCake].Name, new.Cake[i].Ingredients[j].IngCount, old.Cake[needCake].Ingredients[needIngr].IngCount)
						}
					}
				}
			}
		}
	}

	//change/removed unit
	for i:=0; i < len(new.Cake); i++{
		if cakeNameExists(old, new.Cake[i].Name){
			needCake := getIndexCake(old, new, i)
			for j:=0; j<len(new.Cake[i].Ingredients);j++{
				if checkExistIngr(old, new, i, j, needCake){
					needIngr := getIndexIngr(old, new, i, j, needCake)
					flag := checkUnit(old, new, i, j, needCake, needIngr)
					switch flag{
					case 1:
						fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake  \"%s\"\n",
						old.Cake[needCake].Ingredients[needIngr].IngUnit, old.Cake[needCake].Ingredients[needIngr].IngName, old.Cake[needCake].Name)
					case 2:
						fmt.Printf("CHANGED unit for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n",
						old.Cake[needCake].Ingredients[needIngr].IngName, old.Cake[needCake].Name, new.Cake[i].Ingredients[j].IngUnit, old.Cake[needCake].Ingredients[needIngr].IngUnit)
					case 3:
						fmt.Printf("CHANGED unit for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n",
						old.Cake[needCake].Ingredients[needIngr].IngName, old.Cake[needCake].Name, new.Cake[i].Ingredients[j].IngUnit, old.Cake[needCake].Ingredients[needIngr].IngUnit)
					default:
						continue
					}
				}
			}
		}
	}
}