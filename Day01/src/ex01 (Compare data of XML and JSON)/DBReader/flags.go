package DBReader
import(
	"flag"
)

type flags struct{
	Old *string
	New *string
}

func (f *flags) parse(){
	flag.Parse()
}

func newFlag() *flags{
	return &flags{
		flag.String("old", "", "old file"),
		flag.String("new", "", "new file"),
	}
}