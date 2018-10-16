package emailparse

import (
	"testing"
)

func TestEmailParse(t *testing.T) {

	tests := []struct {
		name           string
		email          string
		expectedDomain string
	}{
		{"success", "a@b.com", "b.com"},
		{"long domain", "a@bc.bcbcbcbcccbcj.ssoss.comcoc.com", "bc.bcbcbcbcccbcj.ssoss.comcoc.com"},
		//If you remove the following tests this will pass. That's because it only does happy/positive path testing
		{"empty string", "", ""},
		{"not an email", "broccoli", ""},
		{"two @ symbols", "a@b@c.com", ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			domain := getEmailDomain(test.email)
			if domain != test.expectedDomain {
				t.Errorf("Unexpected Domain. Got %v : Expected %v ", domain, test.expectedDomain)
			}
		})
	}
}
