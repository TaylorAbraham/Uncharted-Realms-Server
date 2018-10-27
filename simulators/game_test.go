package simulators

import (
	"testing"

	"github.com/ryanabraham/urserver/helpers"
	"github.com/ryanabraham/urserver/models"
)

var one = helpers.FakeCard(1, 1, 1, nil)
var two = helpers.FakeCard(2, 2, 2, nil)
var oneDeck = helpers.DeckOf(one, 10)
var twoDeck = helpers.DeckOf(two, 10)
var emptyDeck = models.Deck{}

func TestGuaranteedWins(t *testing.T) {
	result := Play(oneDeck, emptyDeck)
	if result != 0 {
		t.Errorf("Game result incorrect, got: %d, expected %d.", result, 0)
	}
	result = Play(emptyDeck, oneDeck)
	if result != 1 {
		t.Errorf("Game result incorrect, got: %d, expected %d.", result, 1)
	}
}
func TestMoreCardsWins(t *testing.T) {
	oneMinusDeck := helpers.DeckOf(one, 9)
	result := Play(oneDeck, oneMinusDeck)
	if result != 0 {
		t.Errorf("Game result incorrect, got: %d, expected %d.", result, 0)
	}
	result = Play(oneMinusDeck, oneDeck)
	if result != 1 {
		t.Errorf("Game result incorrect, got: %d, expected %d.", result, 1)
	}
}
func TestStrongerCardsWins(t *testing.T) {
	result := Play(twoDeck, oneDeck)
	if result != 0 {
		t.Errorf("Game result incorrect, got: %d, expected %d.", result, 0)
	}
	result = Play(oneDeck, twoDeck)
	if result != 1 {
		t.Errorf("Game result incorrect, got: %d, expected %d.", result, 1)
	}
}