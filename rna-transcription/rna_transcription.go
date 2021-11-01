package strand

const (
	adenine = 'A'
	cytosine = 'C'
	guanine = 'G'
	thymine = 'T'
	uracil = 'U'
)

func ToRNA(dna string) string {
	var rna string
	for _,ch := range dna {
		rna += string(translate(ch))
	}
	return rna
}

func translate(nucleotide rune) rune {
	switch nucleotide {
	case guanine:
		return cytosine
	case cytosine:
		return guanine
	case thymine:
		return adenine
	case adenine:
		return uracil
	}
	return ' '
}