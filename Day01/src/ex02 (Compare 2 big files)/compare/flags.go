package compare

import(
	"flag"
)

type flags struct{
	Old *string
	New *string
}

func (f *flags) Parse(){
	flag.Parse()
}

func NewFlag() *flags{
	return &flags{
		flag.String("old", "", "old file"),
		flag.String("new", "", "new file"),
	}
}