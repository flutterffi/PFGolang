package calc

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "small numbers", a: 2, b: 3, want: 5},
		{name: "with zero", a: 8, b: 0, want: 8},
		{name: "negative", a: -2, b: 5, want: 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Add(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("Add(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a       int
		b       int
		want    int
		wantErr bool
	}{
		{name: "exact division", a: 12, b: 3, want: 4, wantErr: false},
		{name: "truncated division", a: 10, b: 4, want: 2, wantErr: false},
		{name: "zero divisor", a: 10, b: 0, want: 0, wantErr: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Divide(tc.a, tc.b)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("Divide(%d, %d) expected an error", tc.a, tc.b)
				}
				return
			}

			if err != nil {
				t.Fatalf("Divide(%d, %d) returned unexpected error: %v", tc.a, tc.b, err)
			}

			if got != tc.want {
				t.Fatalf("Divide(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
