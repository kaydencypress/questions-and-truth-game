# "Questions and Truth" Game
Inspired by the final game "Questions and Truth" in Season 2 Episode 12 of the Netflix show "The Devil's Game".

## Running the game
Download the binary for your operating system and architecture:
* questions-and-truth-windows-amd64
* questions-and-truth-mac-amd64
* questions-and-truth-linux-amd64
* questions-and-truth-windows-arm64
* questions-and-truth-windows-arm64
* questions-and-truth-windows-arm64

Then, run the executable to start the command line interface.

Or if you have Go installed, you can download the source code and build your own executable using *go build*

## Winning the game
The player asks questions to attempt to deduce the computer's hand of cards in as few rounds as possible. The player wins the game by correctly guessing the value of the cards and their order. 

## Rules
* The computer selects 8 cards from a standard deck of 52 playing cards (excluding jokers).
* The computer choses an order to arrange the cards.
* Cards within the same suit must be arranged from lowest value to highest value from left to right, but do not need to be adjacent to each other.
* Cards of different suits do not need to be arranged in ascending order relative to each other.
* Aces, Jacks, Queens, and Kings are face cards with values of 1, 11, 12, and 13, respectively.
* Each turn, the player will choose "Question" to ask a question about the computer's hand, or "Truth" to guess the computer's hand.
* If "Question" is chosen, the player selects from a pre-defined list of questions and receives this information about the computer's hand.
* If "Truth" is selected, the player submits their guess. The player enters the values of the cards from left to right. The suit of the cards is irrelevant when submitting a guess. 
* If the player's guess is correct, the game ends with the player's victory. If incorrect, the game continues with the next round.

## Available questions
1. Sum
    1. What is the sum of cards in three specified positions (1-8)?
    2. What is the total sum of cards of the specified suit?
    3. What is the total sum of face cards (Aces, Jacks, Queens, Kings)?
    4. What is the total sum of the number cards (2 - 10)?
2. Count
    1. What is the total count of face cards (Aces, Jacks, Queens, Kings)?
    2. What is the total count of the number cards (2 - 10)?
    3. What is the total count of cards with the specified value (1-13)?
3. Position
    1. Which positions contain cards of the specified suit?
    2. Which positions contain cards with the same value?
    3. Which positions contain cards with consecutive values?
    4. Which position(s) contains the card(s) with the highest value?
    5. Which position(s) contains the card(s) with the lowest value?

## Differences from the original game
This implementation is a 1-player game vs a computer instead of a 2-player game. The player chooses "Question" or "Truth" each round, instead of betting tokens vs their opponent for the opportunitity to play the round.

## Enhancements to consider
* For each possible hand, calculate the minimum number of rounds required to definitively deduce the answer.
    * Assign a difficulty rating to each hand accordingly, and allow users to select a difficulty rating.
    * Compare the minimum number of rounds vs the total rounds played at the conclusion of the game.
* Option to enforce a time limit per round.
* Option to require guessing the suit of each card in addition to its value.
* Add frontend UI.
* Add 2-player version of the game (player vs player and/or player vs computer).