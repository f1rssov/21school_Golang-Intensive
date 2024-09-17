package DBReader
import(
	"flag"
)

type flags struct{
	F *bool
}

func (f *flags) parse(){
	flag.Parse()
}

func newFlag() *flags{
	return &flags{
		flag.Bool("f", false, "read file"),
	}
}