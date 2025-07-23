package dna

import "errors"

type Histogram map[rune]int

type DNA string

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	for _, dn := range d {
		switch dn {
		case 'A', 'C', 'G', 'T':
			h[dn]++
		default:
			return nil, errors.New("invalid nucleotide")
		}
	}

	return h, nil
}
