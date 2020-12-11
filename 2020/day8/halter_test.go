package day8

import (
	"log"
	"testing"
)

func TestHalt(t *testing.T) {
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
			code, err := parseCode(tt.code)
			if err != nil {
				t.Fatalf("failed to parse instructions")
			}
			pr, err := Initialize(code)
			if err != nil {
				t.Fatal("failed to initialize program code", err)
			}
			err = pr.run()
			if err == nil {
				t.Errorf("Expected to fail run because of loop detection")
			}
			if pr.Acc != tt.want {
				t.Errorf("Acc = %v, want %v", pr.Acc, tt.want)
			}
		})
	}
}

func TestFixedHalt(t *testing.T) {
	tests := []struct {
		name string
		code string
		want int
	}{
		{
			name: "AoC Example",
			code: "sample",
			want: 8,
		},
		{
			name: "AoC Task",
			code: "task",
			want: 833,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, err := parseCode(tt.code)
			if err != nil {
				t.Fatalf("failed to parse instructions")
			}
			original := make([]Instruction, len(code))
			copy(original, code)
			offset := 0

			for {
				pr, err := Initialize(code)
				if err != nil {
					t.Fatal("failed to initialize program code", err)
				}

				err = pr.run()
				if err == nil {
					log.Printf("finished gracefully - acc is: %v\n", pr.Acc)
					if pr.Acc != tt.want {
						t.Errorf("Acc = %v, want %v", pr.Acc, tt.want)
					}
					break
				} else {
					code, offset = transformCode(original, offset)
				}

			}
		})
	}
}
