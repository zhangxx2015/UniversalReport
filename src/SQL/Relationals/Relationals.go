package Relationals

type Relationals int

const (
	_ Relationals = iota
	EQ
	NE
	GT
	LT
	GE
	LE
	LIKE
)

func (r Relationals) ToString() string {
	return []string{" ?? ", " = ", " <> ", " > ", " < ", " >= ", " <= ", " LIKE "}[r]
}
