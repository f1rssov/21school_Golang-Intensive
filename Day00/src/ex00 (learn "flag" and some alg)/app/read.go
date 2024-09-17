package app
import(
	"bufio"
	"os"
	"fmt"
	conv "strconv"
)

//-100000 and 100000

type myData struct{
	Data []float64
	quantity int
}

func (d *myData) read(){
	input := bufio.NewScanner(os.Stdin)

	for input.Scan(){
		if(input.Text() == "q" || input.Text() == "Q"){
			break
		}
		if(input.Text() == ""){
			fmt.Println("Empty line. Nothing has been entered")
			continue
		}
		value, err := conv.Atoi(input.Text())
		if err != nil{
			fmt.Println("Invalid value entered")
			continue
		}
		if value >= -100000 && value <= 100000{
			d.Data = append(d.Data, float64(value))
		}else{
			fmt.Println("The entered value is greater than 100000 or less than -100000")
			continue
		}
	}
	d.quantity = len(d.Data)
}