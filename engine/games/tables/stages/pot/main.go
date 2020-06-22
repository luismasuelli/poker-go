package pot

import (
	"github.com/luismasuelli/poker-go/engine/games/tables/environment"
	"github.com/luismasuelli/poker-go/engine/games/tables/stages/resolution"
	"github.com/luismasuelli/poker-go/engine/games/tables/pots"
	"github.com/luismasuelli/poker-go/engine/games/messages/games"
	"github.com/luismasuelli/poker-go/engine/games/messages/tables"
	"github.com/luismasuelli/poker-go/engine/games/messages/tables/seats"
)

// Awards all the given pots to their winners. The
// given pots correspond to a single showdown instance
// (e.g. while Hold'Em has a single showdown, other
// modes may have more than one showdown, in a
// permanent or conditional way), so this function
// will be called once per showdown per hand cycle.
//
// For each showdown in certain hand (corresponding
// to certain table / game) this function is called
// exactly once (which could be twice in hi/lo modes).
func AwardPots(winners resolution.PotWinners, pots []*pots.Pot, gameID interface{}, tableID uint32, handID uint64,
	           messageFactory func(playerDisplay interface{}, handID uint64, potIndex int, prize uint64) interface{},
               broadcaster *environment.Broadcaster) {
    // Iterate over all the side pots (and the
    // main point) for a given showdown.
    for potIndex, pot := range pots {
    	// Iterate over all the showdown positions.
    	// The first one having at least one seat
    	// committed with this pot, earn[s] the pot
    	// since it is/are the winner[s].
		for _, tyingSeats := range winners {
			involvedWinners, amount, remainder := pot.Split(tyingSeats)
			if len(involvedWinners) != 0 {
				// Divide this pot, and break.
				for index, seat := range involvedWinners {
					// Get the seat ID, player display,
					// money to award (including the
					// remaining chip).
					display := seat.Player().Display()
					seatID := seat.SeatID()
					prize := amount
					if uint8(index) < remainder {
						prize += 1
					}
					// Add the chips to the seat.
					// On overflow: fuck you. How
					// do you overflow when having
					// uint64 as max limit?
					seat.AddStack(prize)
					// Notify with a game message about this
					// prize ($player received from $pot an
					// amount of $chips).
					broadcaster.Notify(games.GameMessage{
						gameID,
						tables.TableMessage{
							tableID,
							seats.SeatMessage{
								seatID,
								messageFactory(display, handID, potIndex, prize),
							},
						},
					})
				}
				break
			}
			// Otherwise, keep looking among the winners'
			// "levels" until at least one player, in the
			// same level, can claim the pot.
		}
	}
}
