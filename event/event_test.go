package event

import (
	"testing"
)

func TestPutInBoat(t *testing.T) {
	wanted := "[|rev|korn|  vvv\\_høne_pers_/ _________________/øøø]"
	got := PutInBoat(høne)
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}
