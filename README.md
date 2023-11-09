# Project Name: Card Deck and Utility for Managing It.

## Description
This project is a Go programming language-based program that allows you to create and manage a deck of cards. The core of the project consists of a 54-card deck, including standard suits and jokers. The program offers the following functionality:

- **Deck Creation:** Create a new deck that includes standard suits (Spades, Hearts, Diamonds, Clubs) and two jokers (Colorful and Black and White).

- **Deck Shuffling:** Thoroughly shuffle the deck to ensure randomness of the cards.

- **Card Dealing:** Deal a specified number of cards from the deck. For example, you can deal 6 random cards.

- **Deck Saving:** Save the generated or dealt deck on your computer in a text file for later use.

- **Deck Loading:** Load a saved deck from a file and continue working with it.

The project includes functions for creating, shuffling, dealing, saving, and loading decks, providing a simple and user-friendly interface for card management.

## Installation

To get started with this project, you can follow these steps:

   ```bash
   git clone https://github.com/Polovichok/cards.git
   cd cards
   go run main.go deck.go
  ```
 
## Testing with deck_test.go
```go
go test 
```

The project includes a test file named `deck_test.go` that contains test cases for validating the functionality of the card deck manipulation functions. Here's a brief description of the tests included in this file:
#### TestNewDeck

This test ensures that the `newDeck` function correctly creates a deck of cards. It checks the following conditions:
- The length of the deck should be 54.
- The first card should be "Ace of Spades."- The third-to-last card should be "Ten of Clubs."
- The last card should be "Black and White Joker."

```go
func TestNewDeck(t *testing.T) {
    // Test code and assertions...
}
```
#### TestSaveToDeckAndNewDeckFromFile
This test checks the functionality of saving a deck to a file and then loading it back. It verifies:
- A file named "_decktesting" is created and removed after the test.
- The saved deck has 54 cards.
- The loaded deck is consistent with the original deck
```go
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
    // Test code and assertions...
}
```

