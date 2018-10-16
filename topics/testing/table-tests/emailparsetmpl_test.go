package emailparse

import (
	"fmt"
	"testing"
)

func TestEmailParseTmpl(t *testing.T) {
	//this is our anonymous struct to represent the test
	tests := []struct {
		name string
		//we will probably need more than a name what else would you add

	}{
		//add your test cases here
	}

	//uses range to loop through your tests
	for _, test := range tests {
		//t.Run should be used here

		//printing the test so this compiles but remove it
		fmt.Println(test)
	}
}
