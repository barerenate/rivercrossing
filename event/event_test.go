package event

import (
	"testing"
)

func TestTest(t *testing.T) {
	wanted := "[|rev|korn|  vvv\\_høne_pers_/ _________________/øøø]"
	got := Test(items)
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}
