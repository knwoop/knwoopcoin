package fieldelement

import (
	"math/big"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewFieldElement(t *testing.T) {
	type args struct {
		Num   int64
		Prime int64
	}
	tests := []struct {
		name    string
		args    args
		want    *FieldElement
		wantErr bool
	}{
		{
			name: "success to new FieldElement",
			args: args{
				Num:   7,
				Prime: 13,
			},
			want: &FieldElement{
				Num:   big.NewInt(7),
				Prime: big.NewInt(13),
			},
		},
		{
			name: "error num out of range",
			args: args{
				Num:   14,
				Prime: 13,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error num and prime are same",
			args: args{
				Num:   13,
				Prime: 13,
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFieldElement(tt.args.Num, tt.args.Prime)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFieldElement(), error: %v, wantErr: %v", err, tt.wantErr)
			}

			ops := cmp.AllowUnexported(big.Int{})
			if diff := cmp.Diff(got, tt.want, ops); diff != "" {
				t.Errorf("(-got +want)\n%v", diff)
			}
		})
	}
}

func TestFieldElement_Eq(t *testing.T) {
	tests := []struct {
		name string
		args FieldElement
		want bool
	}{
		{
			name: "Eq returns true",
			args: FieldElement{
				Num:   big.NewInt(7),
				Prime: big.NewInt(13),
			},
			want: true,
		},
		{
			name: "Eq returns false",
			args: FieldElement{
				Num:   big.NewInt(6),
				Prime: big.NewInt(13),
			},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			fe, err := NewFieldElement(7, 13)
			if err != nil {
				t.Fatal(err)
			}
			others, err := NewFieldElement(tt.args.Num.Int64(), tt.args.Prime.Int64())
			if err != nil {
				t.Fatal(err)
			}

			ops := cmp.AllowUnexported(big.Int{})
			if diff := cmp.Diff(fe.Eq(others), tt.want, ops); diff != "" {
				t.Errorf("(-got +want)\n%v", diff)
			}
		})
	}
}

func TestFieldElement_Add(t *testing.T) {
	type args struct {
		fieldelement1 *FieldElement
		fieldelement2 *FieldElement
	}
	tests := []struct {
		name    string
		args    args
		want    *FieldElement
		wantErr bool
	}{
		{
			name: "44 + 33",
			args: args{
				fieldelement1: &FieldElement{
					Num:   big.NewInt(44),
					Prime: big.NewInt(57),
				},
				fieldelement2: &FieldElement{
					Num:   big.NewInt(33),
					Prime: big.NewInt(57),
				},
			},
			want: &FieldElement{
				Num:   big.NewInt(20),
				Prime: big.NewInt(57),
			},
			wantErr: false,
		},
		{
			name: "9 + (-29)",
			args: args{
				fieldelement1: &FieldElement{
					Num:   big.NewInt(9),
					Prime: big.NewInt(57),
				},
				fieldelement2: &FieldElement{
					Num:   big.NewInt(-29),
					Prime: big.NewInt(57),
				},
			},
			want: &FieldElement{
				Num:   big.NewInt(37),
				Prime: big.NewInt(57),
			},
			wantErr: false,
		},
		{
			name: "error diffrent primes",
			args: args{
				fieldelement1: &FieldElement{
					Num:   big.NewInt(5),
					Prime: big.NewInt(13),
				},
				fieldelement2: &FieldElement{
					Num:   big.NewInt(33),
					Prime: big.NewInt(57),
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.args.fieldelement1.Add(tt.args.fieldelement2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add(), error: %v, wantErr: %v", err, tt.wantErr)
			}
			ops := cmp.AllowUnexported(big.Int{})
			if diff := cmp.Diff(got, tt.want, ops); diff != "" {
				t.Errorf("(-got +want)\n%v", diff)
			}
		})
	}
}

func TestFieldElement_Sub(t *testing.T) {
	type args struct {
		fieldelement1 *FieldElement
		fieldelement2 *FieldElement
	}
	tests := []struct {
		name    string
		args    args
		want    *FieldElement
		wantErr bool
	}{
		{
			name: "9 - 29",
			args: args{
				fieldelement1: &FieldElement{
					Num:   big.NewInt(9),
					Prime: big.NewInt(57),
				},
				fieldelement2: &FieldElement{
					Num:   big.NewInt(29),
					Prime: big.NewInt(57),
				},
			},
			want: &FieldElement{
				Num:   big.NewInt(37),
				Prime: big.NewInt(57),
			},
			wantErr: false,
		},
		{
			name: "error diffrent primes",
			args: args{
				fieldelement1: &FieldElement{
					Num:   big.NewInt(5),
					Prime: big.NewInt(13),
				},
				fieldelement2: &FieldElement{
					Num:   big.NewInt(33),
					Prime: big.NewInt(57),
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.args.fieldelement1.Sub(tt.args.fieldelement2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sub(), error: %v, wantErr: %v", err, tt.wantErr)
			}

			ops := cmp.AllowUnexported(big.Int{})
			if diff := cmp.Diff(got, tt.want, ops); diff != "" {
				t.Errorf("(-got +want)\n%v", diff)
			}
		})
	}
}

func TestFieldElement_Mul(t *testing.T) {
	type args struct {
		fieldelement1 *FieldElement
		fieldelement2 *FieldElement
		fieldelement3 *FieldElement
	}
	tests := []struct {
		name    string
		args    args
		want    *FieldElement
		wantErr bool
	}{
		{
			name: "95 · 45 · 31",
			args: args{
				fieldelement1: &FieldElement{
					Num:   big.NewInt(95),
					Prime: big.NewInt(97),
				},
				fieldelement2: &FieldElement{
					Num:   big.NewInt(45),
					Prime: big.NewInt(97),
				},
				fieldelement3: &FieldElement{
					Num:   big.NewInt(31),
					Prime: big.NewInt(97),
				},
			},
			want: &FieldElement{
				Num:   big.NewInt(23),
				Prime: big.NewInt(97),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			fe3, err := tt.args.fieldelement1.Mul(tt.args.fieldelement2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mul(), error: %v, wantErr: %v", err, tt.wantErr)
			}
			got, err := fe3.Mul(tt.args.fieldelement3)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mul(), error: %v, wantErr: %v", err, tt.wantErr)
			}

			ops := cmp.AllowUnexported(big.Int{})
			if diff := cmp.Diff(got, tt.want, ops); diff != "" {
				t.Errorf("(-got +want)\n%v", diff)
			}
		})
	}
}

func TestFieldElement_Pow(t *testing.T) {
	type args struct {
		fieldelement1 *FieldElement
		exponent1     int64
		fieldelement2 *FieldElement
		exponent2     int64
	}
	tests := []struct {
		name    string
		args    args
		want    *FieldElement
		wantErr bool
	}{
		{
			name: "12**7 * 77**49",
			args: args{
				fieldelement1: &FieldElement{
					Num:   big.NewInt(12),
					Prime: big.NewInt(97),
				},
				exponent1: 7,
				fieldelement2: &FieldElement{
					Num:   big.NewInt(77),
					Prime: big.NewInt(97),
				},
				exponent2: 49,
			},
			want: &FieldElement{
				Num:   big.NewInt(64),
				Prime: big.NewInt(97),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			fe1, err := tt.args.fieldelement1.Pow(tt.args.exponent1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pow(), error: %v, wantErr: %v", err, tt.wantErr)
			}
			fe2, err := tt.args.fieldelement1.Pow(tt.args.exponent2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pow(), error: %v, wantErr: %v", err, tt.wantErr)
			}
			got, err := fe1.Mul(fe2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mul(), error: %v, wantErr: %v", err, tt.wantErr)
			}

			ops := cmp.AllowUnexported(big.Int{})
			if diff := cmp.Diff(got, tt.want, ops); diff != "" {
				t.Errorf("(-got +want)\n%v", diff)
			}
		})
	}
}
