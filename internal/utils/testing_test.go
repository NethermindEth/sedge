package utils

import (
	"errors"
	"testing"
)

func TestCheckErr(t *testing.T) {
	testErr := errors.New("test error")

	tcs := []struct {
		name    string
		isErr   bool
		err     error
		wantErr bool
	}{
		{
			name:    "Test case 1, got error and don't expect error",
			isErr:   false,
			err:     testErr,
			wantErr: true,
		},
		{
			name:    "Test case 2, got error and expect error",
			isErr:   true,
			err:     testErr,
			wantErr: false,
		},
		{
			name:    "Test case 3, got no error and don't expect error",
			isErr:   false,
			err:     nil,
			wantErr: false,
		},
		{
			name:    "Test case 4, got no error and expect error",
			isErr:   true,
			err:     nil,
			wantErr: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			err := CheckErr("", tc.isErr, tc.err)
			if tc.wantErr && err == nil {
				t.Errorf("CheckErr(\"\", %v, %v) failed: expected error, got nil", tc.isErr, tc.err)
			}

			if !tc.wantErr && err != nil {
				t.Errorf("CheckErr(\"\", %v, %v) failed: expected nil, got error", tc.isErr, tc.err)
			}
		})
	}
}
