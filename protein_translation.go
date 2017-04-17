package protein

const testVersion = 1

var codons = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromCodon(input string) (expect string) {
	return codons[input]
}

func FromRNA(input string) []string {
	inputLen := len(input)
	output := make([]string, inputLen/3)
	for i := 0; i+3 <= inputLen; i = i + 3 {
		tProtein := FromCodon(input[i : i+3])
		if tProtein == "STOP" {
			return output[0 : i/3]
		}
		output[i/3] = tProtein
	}
	return output
}
