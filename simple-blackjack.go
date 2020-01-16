// Simple Blackjack for CSI 380
// This program play's a simple, single suit game of Blackjack
// against a computer dealer.

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Move a card from deck to hand
func drawCard(hand *[]string, deck *[]string) {
	i := rand.Intn(len(*deck)-1) + 1
	*hand = append(*hand, (*deck)[i])
	copy((*deck)[i:], (*deck)[i+1:])
	(*deck) = (*deck)[:len((*deck))-1]
}

// Calculate the score of the hand
func calculateScore(hand []string) int {
	score := 0
  ace := 0
	for i := 0; i < len(hand); i++ {
		j, err := strconv.Atoi(hand[i])
		if err == nil {
			score += j
		} else {
			switch hand[i] {
			case "A":
				ace += 1
			default:
				score += 10
			}
		}
	}
  if ace == 1 {
    if (score + 10) > 21 {
      score += 1
    } else {
      score += 10
    }
  }
	return score
}

// Print everyone's scores and hands
func printStatus(playerCards, dealerCards []string) {
  fmt.Println("")
	playerScore := calculateScore(playerCards)
	dealerScore := calculateScore(dealerCards)
	fmt.Println("Dealer: ", dealerCards, " = ", dealerScore)
	fmt.Println("Player: ", playerCards, " = ", playerScore)
	if playerScore < 22 {
		if dealerScore < 22 {
			if playerScore > dealerScore {
				fmt.Println("You Win!")
			} else {
				fmt.Println("You Lose!")
			}
		} else {
			fmt.Println("You Win!")
		}
	} else {
		fmt.Println("You Lose!")
	}
}

// Entry point and main game loop
func main() {
  rand.New(rand.NewSource(time.Now().UnixNano()))
	deck := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	playerCards := []string{}
	dealerCards := []string{}

	fmt.Println("Simple Blackjack!\n")

	//Deal Initial Cards
	drawCard(&dealerCards, &deck)
	drawCard(&playerCards, &deck)
	drawCard(&dealerCards, &deck)
	drawCard(&playerCards, &deck)

	//Player Draws Cards
	for {
		if calculateScore(playerCards) > 21 {
			fmt.Println("Player Bust!")
			break
		}
		fmt.Println("Your Hand: ", playerCards)
		fmt.Println("Would you like to [1] Draw? or [2] Hold?: ")
		var userInput string
		fmt.Scanln(&userInput)
		i, err := strconv.Atoi(userInput)
		_ = err
		if i == 1 || i == 2 {
			if i == 1 {
				drawCard(&playerCards, &deck)
			}
			if i == 2 {
				break
			}
		} else {
			fmt.Println("Please enter a valid value")
		}
	}

	//Dealer Draws Cards
  for {
    if calculateScore(dealerCards) > 21 {
        fmt.Println("Dealer Bust!")
        break;
    }
    if calculateScore(dealerCards) < calculateScore(playerCards) && calculateScore(dealerCards) <= 16 {
      drawCard(&dealerCards, &deck)
    } else {
      break;
    }
  }

  //Decide Winner
	printStatus(playerCards, dealerCards)
}