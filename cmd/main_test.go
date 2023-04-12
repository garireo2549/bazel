package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name string
		i    int
		want string
	}{
		{
			name: "fizz",
			i:    3,
			want: fizz,
		},
		{
			name: "buzz",
			i:    5,
			want: buzz,
		},
		{
			name: "fizz buzz",
			i:    15,
			want: fbzz,
		},
		{
			name: "others",
			i:    2,
			want: "2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fizzBuzz(tt.i)
			if tt.want != got {
				t.Fatalf("unexpected error, want: %s, got: %s", tt.want, got)
			}
		})
	}
}
