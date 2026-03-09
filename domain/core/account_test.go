package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccountRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		req     CreateAccountRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "success",
			req: CreateAccountRequest{
				DocumentNumber: "123456",
			},
			wantErr: false,
		},
		{
			name: "should fail when empty document number",
			req: CreateAccountRequest{
				DocumentNumber: "",
			},
			wantErr: true,
			errMsg:  "invalid document number",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			if tt.wantErr {
				assert.ErrorContains(t, err, tt.errMsg)
			}
		})
	}

}
