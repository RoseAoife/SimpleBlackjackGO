// Simple Blackjack Tests for CSI 380 by David Kopec
// These tests ensure calculateScore() returns correct
// values for various hands.

package main

import "testing"

// Test K, J
func TestCalculateScoreKJ(t *testing.T) {
	hand := []string{"K", "J"}
	expected := 20
	actual := calculateScore(hand)
	if actual != expected {
		t.Errorf("Test failed: expected %v to be %d", hand, expected)
	}
}

// Test Q, 7, A
func TestCalculateScoreQ7A(t *testing.T) {
	hand := []string{"Q", "7", "A"}
	expected := 18
	actual := calculateScore(hand)
	if actual != expected {
		t.Errorf("Test failed: expected %v to be %d", hand, expected)
	}
}

// Test 2, 5, 8, Q, 9
func TestCalculateScore258Q9(t *testing.T) {
	hand := []string{"2", "5", "8", "Q", "9"}
	expected := 34
	actual := calculateScore(hand)
	if actual != expected {
		t.Errorf("Test failed: expected %v to be %d", hand, expected)
	}
}

// Test 2, 3, A
func TestCalculateScore23A(t *testing.T) {
	hand := []string{"2", "3", "A"}
	expected := 16
	actual := calculateScore(hand)
	if actual != expected {
		t.Errorf("Test failed: expected %v to be %d", hand, expected)
	}
}