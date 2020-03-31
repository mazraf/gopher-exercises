package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Queen, Suit: Club})
	fmt.Println(Card{Rank: King, Suit: Diamond})
	fmt.Println(Card{Suit: Joker})
	// Output:
	// Ace of Hearts
	// Two of Spades
	// Queen of Clubs
	// King of Diamonds
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 52 {
		t.Error("Wrong number of cards in the deck")
	}
}

func TestDetaultCards(t *testing.T) {
	cards := New(Sort(Less))
	exp := Card{Suit: Spade, Rank: Ace}
	if exp != cards[0] {
		t.Errorf("Expected %s  and got %s", exp, cards[0])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0

	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Errorf("Expected 3 got %d", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, card := range cards {
		if card.Rank == Two || card.Rank == Three {
			t.Error("Expected Two and Three to be filtered out.")
		}
	}
}

func TestDesk(t *testing.T) {
	cards := New(Deck((3)))
	if len(cards) != 13*4*3 {
		t.Errorf("Expected %d got %d", 13*4*3, len(cards))
	}
}

func TestShuffle(t *testing.T) {
	shuffleRand = rand.New(rand.NewSource(0))
	orig := New()
	first := orig[40]
	second := orig[35]

	cards := New(Shuffle)
	if cards[0] != first {
		t.Errorf("Expected %s got %s", first, cards[0])
	}
	if cards[1] != second {
		t.Errorf("Expected %s got %s", second, cards[1])
	}
}
