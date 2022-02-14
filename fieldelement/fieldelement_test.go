package fieldelement

import (
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
				Num:   7,
				Prime: 13,
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
			if diff := cmp.Diff(got, tt.want); diff != "" {
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
				Num:   7,
				Prime: 13,
			},
			want: true,
		},
		{
			name: "Eq returns false",
			args: FieldElement{
				Num:   6,
				Prime: 13,
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
			others, err := NewFieldElement(tt.args.Num, tt.args.Prime)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(fe.Eq(others), tt.want); diff != "" {
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
					Num:   44,
					Prime: 57,
				},
				fieldelement2: &FieldElement{
					Num:   33,
					Prime: 57,
				},
			},
			want: &FieldElement{
				Num:   20,
				Prime: 57,
			},
			wantErr: false,
		},
		{
			name: "9 + (-29)",
			args: args{
				fieldelement1: &FieldElement{
					Num:   9,
					Prime: 57,
				},
				fieldelement2: &FieldElement{
					Num:   -29,
					Prime: 57,
				},
			},
			want: &FieldElement{
				Num:   37,
				Prime: 57,
			},
			wantErr: false,
		},
		{
			name: "error diffrent primes",
			args: args{
				fieldelement1: &FieldElement{
					Num:   5,
					Prime: 13,
				},
				fieldelement2: &FieldElement{
					Num:   33,
					Prime: 57,
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
			if diff := cmp.Diff(got, tt.want); diff != "" {
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
					Num:   9,
					Prime: 57,
				},
				fieldelement2: &FieldElement{
					Num:   29,
					Prime: 57,
				},
			},
			want: &FieldElement{
				Num:   37,
				Prime: 57,
			},
			wantErr: false,
		},
		{
			name: "error diffrent primes",
			args: args{
				fieldelement1: &FieldElement{
					Num:   5,
					Prime: 13,
				},
				fieldelement2: &FieldElement{
					Num:   33,
					Prime: 57,
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
			if diff := cmp.Diff(got, tt.want); diff != "" {
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
					Num:   95,
					Prime: 97,
				},
				fieldelement2: &FieldElement{
					Num:   45,
					Prime: 97,
				},
				fieldelement3: &FieldElement{
					Num:   31,
					Prime: 97,
				},
			},
			want: &FieldElement{
				Num:   23,
				Prime: 97,
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
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("(-got +want)\n%v", diff)
			}
		})
	}
}
