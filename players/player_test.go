package players

import "testing"

func TestPlayers(t *testing.T) {
	var guestProbabilities = All()
	if len(guestProbabilities) < 1 {
		t.Errorf("Guest compilation incorrect, got: %d, want more than 1.", len(guestProbabilities))
	}
}
