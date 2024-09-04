package app

func CheckFlagsPrintCalc(){
	myFlags := NewFlags()
	myFlags.init()

	source := myData{}
	source.read()

	if *myFlags.Mean {
		source.meanAlg()
	}
	if *myFlags.Median {
		source.medianAlg()
	}
	if *myFlags.Mode {
		source.modeAlg()
	}
	if *myFlags.SD {
		source.sdAlg()
	}
	if !*myFlags.Mean && !*myFlags.Mode && !*myFlags.Median && !*myFlags.SD{
		source.meanAlg()
		source.medianAlg()
		source.modeAlg()
		source.sdAlg()
	}
}