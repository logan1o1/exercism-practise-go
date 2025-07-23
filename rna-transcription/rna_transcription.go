package strand

import "strings"

func ToRNA(dna string) string {
	var rna []string

	for _, d := range dna {
		switch d {
		case 'G':
			rna = append(rna, "C")
		case 'C':
			rna = append(rna, "G")
		case 'T':
			rna = append(rna, "A")
		case 'A':
			rna = append(rna, "U")
		default:
			return ""
		}
	}
	return strings.Join(rna, "")
}
