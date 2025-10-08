package even_test

import (
	"context"
	"testing"

	"github.com/dcmcand/camp-quansight-2025-go-bof/pkg/even"
)

func TestIsEven(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		number  int
		want    bool
		wantErr bool
	}{
		{
			name:    "zero is even",
			number:  0,
			want:    true,
			wantErr: false,
		},
		{
			name:    "one is odd",
			number:  1,
			want:    false,
			wantErr: false,
		},
		{
			name:    "negatives throw an errror",
			number:  -1,
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := even.IsEven(context.Background(), tt.number)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("IsEven() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("IsEven() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
