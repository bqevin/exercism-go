package proteintranslation

import (
    "errors"
)

var (
	ErrStop        = errors.New("stop codon encountered")
	ErrInvalidBase = errors.New("invalid base")
)

func FromRNA(rna string) ([]string, error) {
	m := Proteins()
	r := make([]string, 0)
	l := len(rna)
	for i := 0; i < l; i++ {
		skip := i + 3
		if skip > l {
			skip = l
		}
        codon := rna[i:skip]
		if len(codon) != 3 {
			return r, ErrInvalidBase
		}
		if _, ok := m[codon]; !ok {
			return r, ErrInvalidBase
		}
		if m[codon] == "STOP" {
			return r, nil
		}
		r = append(r, m[codon])
		i = skip - 1
	}
	return r, nil
}

func FromCodon(codon string) (string, error) {
	m := Proteins()
	if _, ok := m[codon]; !ok {
		return "", ErrInvalidBase
	}
	if m[codon] == "STOP" {
		return "", ErrStop
	}
	return m[codon], nil
}

func Proteins() map[string]string {
	return map[string]string{
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
}
