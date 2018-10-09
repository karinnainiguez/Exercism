package strand

func ToRNA(dna string) string {
	var rna string

	for _, l := range dna {
		switch string(l) {
		case "G":
			rna += "C"
		case "C":
			rna += "G"
		case "T":
			rna += "A"
		case "A":
			rna += "U"
		}
	}

	return rna
}
