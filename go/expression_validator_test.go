package main

import "testing"

func TestValidateExpression(t *testing.T) {
	tests := []struct {
		expr      string
		wantValid bool
	}{
		{
			expr:      `products.botd.data.bot.result == "notDetected"`,
			wantValid: true,
		},
		{
			expr:      `products.identification.data.browserDetails.browserName == "Chrome" && products.vpn.data.result == false`,
			wantValid: true,
		},
		{
			expr:      `unknownVariable == 1`,
			wantValid: false,
		},
		{
			expr:      `products.botd.data.bot.result = "notDetected"`, // assignment, invalid syntax
			wantValid: false,
		},
		// Deep invalid property: on a struct field that does not exist
		{
			expr:      `products.identification.data.nonExistentField == true`,
			wantValid: false,
		},
		// Deep invalid nested property inside a known object
		{
			expr:      `products.identification.data.browserDetails.nonExistent == "x"`,
			wantValid: false,
		},
		// Invalid property on a known leaf type
		{
			expr:      `products.ipInfo.data.v4.asn.nonExistent == "x"`,
			wantValid: false,
		},
		// Complex logical expression with brackets and multiple conditions (valid)
		{
			expr:      `(products.botd.data.bot.result == "notDetected" && products.vpn.data.result == false) || (products.ipBlocklist.data.result == false && products.identification.data.browserDetails.browserName == "Chrome")`,
			wantValid: true,
		},
		// Complex expression mixing negation and comparison on int (valid)
		{
			expr:      `!products.tampering.data.result && (products.suspectScore.data.result < 10 || products.proxy.data.result == false)`,
			wantValid: true,
		},
		// Complex expression with a nonexistent deep field inside brackets (invalid)
		{
			expr:      `(products.botd.data.bot.result == "notDetected" && products.identification.data.nonExistent == false) || products.vpn.data.result == false`,
			wantValid: false,
		},
	}

	for _, tt := range tests {
		if got := ValidateExpression(tt.expr); got != tt.wantValid {
			t.Errorf("ValidateExpression(%q) = %v, want %v", tt.expr, got, tt.wantValid)
		}
	}
}
