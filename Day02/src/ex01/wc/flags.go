package wc
import(
	"flag"
)

type Flags struct{
	Words *bool
	Lines *bool
	M *bool
}

func (f *Flags)Parse(){
	flag.Parse()
}

func NewFlags() *Flags{
	return &Flags{
		flag.Bool("w", false, "words"),
		flag.Bool("l", false, "lines"),
		flag.Bool("m", false, "characters"),
	}
}