package high

import (
	"github.com/luismasuelli/poker-go/engine/games/cards"
	. "github.com/luismasuelli/poker-go/engine/games/cards/french"
	"testing"
)

func testHandPower(t *testing.T, expectedPower uint64, expectedBest uint32, cards ...cards.Card) {
	best, power := Power(cards, nil)
	if power != expectedPower || best != expectedBest {
		t.Errorf(
			"Testing hand: %v\nexpected power: %#064b\n     got power: %#064b\nexpected best: %#07b\n     got best: %#07b\n",
			cards, expectedPower, power, expectedBest, best,
		)
	}
}

func TestHandPowers(t *testing.T) {
	// Flush-straight.
	testHandPower(t, 0b1000000000000000000000000000001000000000000, 0b1111100, DA, HA, CT, CA, CK, CJ, CQ)
	testHandPower(t, 0b1000000000000000000000000000000000000001000, 0b111110, C3, H5, HA, H4, H2, H3, D2)
	testHandPower(t, 0b1000000000000000000000000000000000001000000, 0b11111, S7, S4, S6, S5, S8, C5, D5)
	// 4 of a kind.
	testHandPower(t, 0b0111000000000000010000000000000000001000000, 0b0111101, D8, C8, SA, CA, HA, DA, S8)
	testHandPower(t, 0b0111000000000000000000000000011000000000000, 0b111110, D4, S2, C2, H2, D2, SA, S4)
	testHandPower(t, 0b0111000000000000000000000010000010000000000, 0b11111, SQ, S5, C5, H5, D5, D8, C8)
	// Full house.
	testHandPower(t, 0b0110000000000000000000000010000010000000000, 0b111011, SQ, S5, SA, C5, HQ, D5, D3)
	testHandPower(t, 0b0110000000000000001000000000001000000000000, 0b1011110, SQ, SA, SK, CK, HA, HQ, DK)
	testHandPower(t, 0b0110000000000000000000000000100000001000000, 0b1111010, D2, H3, H2, S3, C3, H8, D8)
	// Flush.
	testHandPower(t, 0b0101000000000000000000000000000000011110100, 0b1111100, CA, DA, S7, S9, S6, S4, S8)
	testHandPower(t, 0b0101000000000000000000000000001000101010010, 0b1111010, CA, HA, D4, H6, H3, H8, HT)
	testHandPower(t, 0b0101000000000000000000000000001000100100110, 0b101111, DT, DA, D3, D4, S4, D7, S7)
	// Straights.
	testHandPower(t, 0b0100000000000000000000000000001000000000000, 0b111011, CT, SA, HA, CK, CJ, CQ, D9)
	testHandPower(t, 0b0100000000000000000000000000000000000001000, 0b11111, H5, HA, D4, H2, H3, D5, C2)
	testHandPower(t, 0b0100000000000000000000000000000000001000000, 0b101111, S7, S4, S6, C5, D5, S8, C5)
	// 3 of a kind.
	testHandPower(t, 0b0011000000000000000001000000000011000000000, 0b1111010, D3, CT, D4, ST, DT, CJ, CQ)
	testHandPower(t, 0b0011000000000000010000000000000000000100100, 0b1010111, DA, SA, CA, C2, C4, C3, H7)
	testHandPower(t, 0b0011000000000000000000000001000000010010000, 0b1010111, C4, D4, S4, H5, H6, H2, S9)
	// Double Pair.
	testHandPower(t, 0b0010000000000000000001010000000010000000000, 0b1101011, CT, ST, C2, D8, H7, C8, CQ)
	testHandPower(t, 0b0010000000000000010000001000000000100000000, 0b1110110, H3, DA, SA, D3, HT, C7, H7)
	testHandPower(t, 0b0010000000000000000000000001010000010000000, 0b1111100, C3, S8, C2, D2, S4, H4, S9)
	// Pair.
	testHandPower(t, 0b0001000000000000000001000000000010001000100, 0b1111010, D2, CT, D3, ST, D4, C8, CQ)
	testHandPower(t, 0b0001000000000000010000000000000000100100100, 0b1110011, DA, SA, D2, D3, HT, C4, H7)
	testHandPower(t, 0b0001000000000000000000000000010000011100000, 0b1110011, C2, D2, S3, H4, S8, S7, S9)
	// Bust / High Card.
	testHandPower(t, 0b0000000000000000000000000000000000011110100, 0b0011111, S7, S9, C6, H4, S8, H2, H3)
	testHandPower(t, 0b0000000000000000000000000000001000101010100, 0b1101011, HA, D6, D2, D4, D3, H8, HT)
	testHandPower(t, 0b0000000000000000000000000000001000100111000, 0b1100111, DT, CA, S5, H4, D2, D6, D7)
}
