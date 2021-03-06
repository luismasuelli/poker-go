package high

import (
	"github.com/luismasuelli/poker-go/engine/games/cards"
	. "github.com/luismasuelli/poker-go/engine/games/cards/french"
	"testing"
)

func testHandPower(t *testing.T, expectedPower uint64, cards ...cards.Card) {
	_, power := Power(cards, nil)
	if power != expectedPower {
		t.Errorf("Testing hand: %v\nexpected power: %#064b\n     got power: %#064b\n", cards, expectedPower, power)
	}
}

func TestHandPowers(t *testing.T) {
	// Flush-straight.
	testHandPower(t, 0b1000000000000000000000000000001000000000000, CT, CA, CK, CJ, CQ)
	testHandPower(t, 0b1000000000000000000000000000000000000001000, H5, HA, H4, H2, H3)
	testHandPower(t, 0b1000000000000000000000000000000000001000000, S7, S4, S6, S5, S8)
	// 4 of a kind.
	testHandPower(t, 0b0111000000000000010000000000000000001000000, SA, CA, HA, DA, S8)
	testHandPower(t, 0b0111000000000000000000000000011000000000000, S2, C2, H2, D2, SA)
	testHandPower(t, 0b0111000000000000000000000010000010000000000, SQ, S5, C5, H5, D5)
	// Full house.
	testHandPower(t, 0b0110000000000000000000000010000010000000000, SQ, S5, C5, HQ, D5)
	testHandPower(t, 0b0110000000000000001000000000001000000000000, SA, SK, CK, HA, DK)
	testHandPower(t, 0b0110000000000000000000000000100000001000000, H3, S3, C3, H8, D8)
	// Flush.
	testHandPower(t, 0b0101000000000000000000000000000000011110100, S7, S9, S6, S4, S8)
	testHandPower(t, 0b0101000000000000000000000000001000101010010, HA, H6, H3, H8, HT)
	testHandPower(t, 0b0101000000000000000000000000001000100100110, DT, DA, D3, D4, D7)
	// Straights.
	testHandPower(t, 0b0100000000000000000000000000001000000000000, CT, SA, CK, CJ, CQ)
	testHandPower(t, 0b0100000000000000000000000000000000000001000, H5, HA, D4, H2, H3)
	testHandPower(t, 0b0100000000000000000000000000000000001000000, S7, S4, S6, C5, S8)
	// 3 of a kind.
	testHandPower(t, 0b0011000000000000000001000000000011000000000, CT, ST, DT, CJ, CQ)
	testHandPower(t, 0b0011000000000000010000000000000000000100010, DA, SA, CA, C3, H7)
	testHandPower(t, 0b0011000000000000000000000001000000010010000, C4, D4, S4, H6, S9)
	// Double Pair.
	testHandPower(t, 0b0010000000000000000001010000000010000000000, CT, ST, D8, C8, CQ)
	testHandPower(t, 0b0010000000000000010000001000000000100000000, DA, SA, HT, C7, H7)
	testHandPower(t, 0b0010000000000000000000000001010000010000000, C2, D2, S4, H4, S9)
	// Pair.
	testHandPower(t, 0b0001000000000000000001000000000010001000100, CT, ST, D4, C8, CQ)
	testHandPower(t, 0b0001000000000000010000000000000000100100010, DA, SA, HT, C3, H7)
	testHandPower(t, 0b0001000000000000000000000000010000010000110, C2, D2, S3, H4, S9)
	// Bust / High Card.
	testHandPower(t, 0b0000000000000000000000000000000000011110100, S7, S9, C6, H4, S8)
	testHandPower(t, 0b0000000000000000000000000000001000101010010, HA, D6, D3, H8, HT)
	testHandPower(t, 0b0000000000000000000000000000001000100100110, DT, CA, S3, H4, D7)
}
