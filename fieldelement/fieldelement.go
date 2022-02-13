package fieldelement

import (
	"errors"
	"fmt"
)

var (
	// ErrOutOfRangeNumber message
	ErrOutOfRangeNumber = errors.New("invalid Num in filed is out of range")
)

// FieldElement represents finite field
type FieldElement struct {
	Num   uint64
	Prime uint64
}

// NewFieldElement initializes struct
func NewFieldElement(num, prime uint64) (*FieldElement, error) {
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
