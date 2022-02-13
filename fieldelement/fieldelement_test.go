package fieldelement

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewFieldElement(t *testing.T) {
	type args struct {
		Num   uint64
		Prime uint64
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
