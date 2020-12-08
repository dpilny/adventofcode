package halter

import (
	"testing"
)

func TestHalter(t *testing.T) {
	tests := []struct {
		name string
		code string
		want int
	}{
		{
			name: "AoC Example",
			code: "sample",
			want: 5,
		},
		{
			name: "AoC Task",
			code: "task",
			want: 1654,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr, err := Initialize(tt.code)
			if err != nil {
				t.Fatal("failed to initialize program code", err)
			}
			pr.run()
			if pr.Acc != tt.want {
				t.Errorf("Acc = %v, want %v", pr.Acc, tt.want)
			}
		})
	}
}