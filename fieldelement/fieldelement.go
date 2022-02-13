package fieldelement

import (
	"errors"
	"fmt"
)

var (
	// ErrOutOfRangeNumber message
	ErrOutOfRangeNumber = errors.New("invalid Num in filed is out of range")
	// ErrInvalidPrime is a message for invalid prime
	ErrInvalidPrime = errors.New("invalid prime to add two primes in different Fields")
)

// FieldElement represents finite field
type FieldElement struct {
	Num   int64
	Prime int64
}

// NewFieldElement initializes struct
func NewFieldElement(num, prime int64) (*FieldElement, error) {
	if num >= prime {
		return nil, ErrOutOfRangeNumber
	}
	return &FieldElement{
		Num:   num,
		Prime: prime,
	}, nil
}

func (fe *FieldElement) String() string {
	return fmt.Sprintf("FieldElement_%d_%d", fe.Prime, fe.Num)
}

func (fe *FieldElement) Eq(other *FieldElement) bool {
	return fe.Num == other.Num && fe.Prime == other.Prime
}

func (fe *FieldElement) Add(other *FieldElement) (*FieldElement, error) {
	if fe.Prime != other.Prime {
		return nil, ErrInvalidPrime
	}
	num := mod(fe.Num+other.Num, fe.Prime)
	return NewFieldElement(num, fe.Prime)
}

func (fe *FieldElement) Sub(other *FieldElement) (*FieldElement, error) {
	if fe.Prime != other.Prime {
		return nil, ErrInvalidPrime
	}
	num := mod(fe.Num-other.Num, fe.Prime)
	return NewFieldElement(num, fe.Prime)
}

func (fe *FieldElement) Mul(other *FieldElement) (*FieldElement, error) {
	if fe.Prime != other.Prime {
		return nil, ErrInvalidPrime
	}
	num := mod(fe.Num*other.Num, fe.Prime)
	return NewFieldElement(num, fe.Prime)
}

func mod(d, m int64) int64 {
	res := d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}
