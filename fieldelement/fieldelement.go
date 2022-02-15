package fieldelement

import (
	"errors"
	"fmt"
	"math/big"
)

var (
	// ErrOutOfRangeNumber message
	ErrOutOfRangeNumber = errors.New("invalid Num in filed is out of range")
	// ErrInvalidPrime is a message for invalid prime
	ErrInvalidPrime = errors.New("invalid prime to add two primes in different Fields")
)

// FieldElement represents finite field
type FieldElement struct {
	Num   *big.Int
	Prime *big.Int
}

// NewFieldElement initializes struct
func NewFieldElement(num, prime int64) (*FieldElement, error) {
	if num >= prime {
		return nil, ErrOutOfRangeNumber
	}
	return &FieldElement{
		Num:   big.NewInt(num),
		Prime: big.NewInt(prime),
	}, nil
}

func (fe *FieldElement) String() string {
	return fmt.Sprintf("FieldElement_%d_%d", fe.Prime, fe.Num)
}

func (fe *FieldElement) Eq(other *FieldElement) bool {
	return fe.Num.Cmp(other.Num) == 0 && fe.Prime.Cmp(other.Prime) == 0
}

func (fe *FieldElement) Add(other *FieldElement) (*FieldElement, error) {
	if fe.Prime.Cmp(other.Prime) != 0 {
		return nil, ErrInvalidPrime
	}
	num := fe.Num.Add(fe.Num, other.Num).Mod(fe.Num, fe.Prime)
	return NewFieldElement(num.Int64(), fe.Prime.Int64())
}

func (fe *FieldElement) Sub(other *FieldElement) (*FieldElement, error) {
	if fe.Prime.Cmp(other.Prime) != 0 {
		return nil, ErrInvalidPrime
	}
	num := fe.Num.Sub(fe.Num, other.Num).Mod(fe.Num, fe.Prime)
	return NewFieldElement(num.Int64(), fe.Prime.Int64())
}

func (fe *FieldElement) Mul(other *FieldElement) (*FieldElement, error) {
	if fe.Prime.Cmp(other.Prime) != 0 {
		return nil, ErrInvalidPrime
	}
	num := fe.Num.Mul(fe.Num, other.Num).Mod(fe.Num, fe.Prime)
	return NewFieldElement(num.Int64(), fe.Prime.Int64())
}

func (fe *FieldElement) Pow(exponent int64) (*FieldElement, error) {
	num := fe.Num.Exp(fe.Num, big.NewInt(exponent), fe.Prime)
	return NewFieldElement(num.Int64(), fe.Prime.Int64())
}
