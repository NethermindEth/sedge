package configs

import "testing"

func TestValidateLoggingFlag(t *testing.T) {
	tests := []struct {
		name    string
		flag    string
		wantErr bool
	}{
		{
			name:    "Valid logging flag, none",
			flag:    "none",
			wantErr: false,
		},
		{
			name:    "Valid logging flag, json",
			flag:    "json",
			wantErr: false,
		},
		{
			name:    "Invalid logging flag",
			flag:    "invalid",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateLoggingFlag(tt.flag); (err != nil) != tt.wantErr {
				t.Errorf("ValidateLoggingFlag() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
