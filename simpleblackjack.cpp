//
//  simpleblackjack.cpp
//  Simple Blackjack
//
//  Created by David Kopec on 8/29/16.
//  Ported from Chapter 6 of
//  Dart for Absolute Beginners by David Kopec
//  Copyright 2014
//

#include <string>
#include <vector>
#include <iostream>
#include <random>
#include <algorithm>

using namespace std;

// calculate the score of the cards Vector
int calculateScore(vector<string> &cards) {
    int score = 0;
    bool hasAce = false; // only one ace in Simple Blackjack
    
    // iterate through all cards
    for (string const& card: cards) {
        try {
            int cardValue = stoi(card);
            score += cardValue;
        } catch (std::invalid_argument& e) {
            // not a number card
            if (card == "A") { // it's an ace
                hasAce = true;
            } else { // it's a face card
                score += 10;
            }
        }
    }
    
    if (hasAce) {
        if ((score + 11) > 21) {  // don't let ace cause bust
            score += 1;
        } else {
            score += 11;
        }
    }
    
    return score;
}

// Print everyone's scores and decks
void printStatus(vector<string> &playerCards, vector<string> &dealerCards) {
    cout << "\n";  // blank line
    cout << "Player's total is " << calculateScore(playerCards) << ":\n";
    for (string const& card: playerCards) {
        cout << card << ", ";
    }
    cout << "\n";
    
    cout << "Dealer's total is " << calculateScore(dealerCards) << ":\n";
    for (string const& card: dealerCards) {
        cout << card << ", ";
    }
    cout << "\n";
}

// main game logic/loop
int main(int argc, const char * argv[]) {
    vector<string> deck = {"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"};
    vector<string> playerCards = vector<string>();
    vector<string> dealerCards = vector<string>();
    
    // shuffle randomly the deck
    random_device rng;
    mt19937 urng(rng());
    shuffle(deck.begin(), deck.end(), urng);
    
    cout << "Dealer draws first card.\n";
    dealerCards.push_back(deck.back());
    deck.pop_back();
    cout << "Player receives two cards.\n";
    playerCards.push_back(deck.back());
    deck.pop_back();
    playerCards.push_back(deck.back());
    deck.pop_back();
    printStatus(playerCards, dealerCards);
    
    while (true) { // player decision loop
        cout << "Do you want to (H)it, (S)tay, or (Q)uit?\n";
        string selection;
        cin >> selection;
        // transform to upper case
        transform(selection.begin(), selection.end(),selection.begin(), ::toupper);
        
        if (selection == "H") { // hit
            playerCards.push_back(deck.back());
            deck.pop_back();
            printStatus(playerCards, dealerCards);
            
            if (calculateScore(playerCards) > 21) {
                cout << "You busted! You lose!\n";
                return 0;
            }
        } else if (selection == "S") { // stay
            break; // stop offering to hit, leave loop
        } else if (selection == "Q") { // quit
            return 0;
        }
    }
    
    cout << "Dealer draws rest of cards.\n";
    // keep drawing cards till 17
    while (calculateScore(dealerCards) < 17) {
        dealerCards.push_back(deck.back());
        deck.pop_back();
    }
    printStatus(playerCards, dealerCards);
    
    if (calculateScore(dealerCards) > 21) { // dealer bust
        cout << "Dealer busts! You win!\n";
    } else if (calculateScore(dealerCards) > calculateScore(playerCards)) {
        cout << "Dealer wins!\n";
    } else if (calculateScore(dealerCards) < calculateScore(playerCards)) {
        cout << "You win!\n";
    } else { // must be a tie by default
        cout << "It's a tie!\n";
    }
    return 0;
}
