package validation

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/xinyi-chong/common-lib/consts"
	apperrors "github.com/xinyi-chong/common-lib/errors"
	"testing"
)

type RegisterParam struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
	Username string `json:"username" validate:"required"`
}

func TestValidationCheck(t *testing.T) {
	tests := []struct {
		name      string
		input     RegisterParam
		wantError error
	}{
		{
			name:      "All fields valid",
			input:     RegisterParam{Email: "test@example.com", Password: "password123", Username: "user1"},
			wantError: nil,
		},
		{
			name:      "Missing email",
			input:     RegisterParam{Password: "password123", Username: "user1"},
			wantError: apperrors.ErrXIsRequired,
		},
		{
			name:      "Invalid email",
			input:     RegisterParam{Email: "invalid-email", Password: "password123", Username: "user1"},
			wantError: apperrors.ErrInvalidX,
		},
		{
			name:      "Password too short",
			input:     RegisterParam{Email: "test@example.com", Password: "short", Username: "user1"},
			wantError: apperrors.ErrXMin,
		},
		{
			name:      "Missing username",
			input:     RegisterParam{Email: "test@example.com", Password: "password123"},
			wantError: apperrors.ErrXIsRequired,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Check(tt.input)
			t.Logf("err: %#v", err)

			if tt.wantError == nil {
				assert.NoError(t, err)
				return
			}
			assert.Error(t, err)
			var apperr *apperrors.Error
			ok := errors.As(err, &apperr)
			assert.True(t, ok, "error should be of type *apperrors.Error")

			var expected *apperrors.Error
			errors.As(tt.wantError, &expected)
			assert.Equal(t, expected.MessageKey, apperr.MessageKey)

			switch tt.name {
			case "Missing email", "Invalid email":
				t.Logf("tt.name: %#v", tt.name)
				assert.Equal(t, consts.Field("email"), apperr.TemplateData["Field"])
			case "Password too short":
				assert.Equal(t, consts.Field("password"), apperr.TemplateData["Field"])
				assert.Equal(t, "8", apperr.TemplateData["Value"])
			case "Missing username":
				assert.Equal(t, consts.Field("username"), apperr.TemplateData["Field"])
			}
		})
	}
}
