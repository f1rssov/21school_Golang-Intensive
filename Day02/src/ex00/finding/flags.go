package finding
import(
	"flag"
)

type Flags struct{
	Sl *bool
	Dir *bool
	File *bool
	Ext *string
}

func (f *Flags)Parse(){
	flag.Parse()
}

func NewFlags() *Flags{
	return &Flags{
		flag.Bool("sl", false, "Sym link"),
		flag.Bool("d", false, "dir"),
		flag.Bool("f", false, "file"),
		flag.String("ext", "", "ext"),
	}
}