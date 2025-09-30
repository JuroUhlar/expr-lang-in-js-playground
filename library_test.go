package main

import "testing"

func TestValidateExpression(t *testing.T) {
	tests := []struct {
		expr      string
		wantValid bool
	}{
		// Basic fields
		{
			expr:      `bot == true`,
			wantValid: true,
		},
		{
			expr:      `country == "US"`,
			wantValid: true,
		},
		{
			expr:      `suspectScore > 50`,
			wantValid: true,
		},

		// Complex fields
		{
			expr:      `proxy.type == "residential"`,
			wantValid: true,
		},
		{
			expr:      `proxy.confidence > 0.8`,
			wantValid: true,
		},
		// Complex expressions
		{
			expr:      `proxy.type == "residential" && proxy.confidence > 0.9`,
			wantValid: true,
		},
		{
			expr:      `(vpn == true && proxy.type == "datacenter") || (proxy.confidence > 0.8 && country == "CN")`,
			wantValid: true,
		},
		// Invalid expressions
		{
			expr:      `unknownVariable == 1`,
			wantValid: false,
		},
		{
			expr:      `bot = true`,
			wantValid: false,
		},
		{
			expr:      `country.nonExistent == "x"`,
			wantValid: false,
		},
		{
			expr:      `proxy.unknownField == "value"`,
			wantValid: false,
		},
	}

	for _, tt := range tests {
		if got := isExpressionValid(tt.expr); got != tt.wantValid {
			t.Errorf("ValidateExpression(%q) = %v, want %v", tt.expr, got, tt.wantValid)
		}
	}
}
