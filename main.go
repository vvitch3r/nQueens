package main

import (
	"fmt"
	"strings"
)

// Function to check if a queen can be placed at board[row][col]
func isSafe(board [][]string, row, col, n int) bool {
	// Check the current column
	for i := 0; i < row; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}

	// Check the upper left diagonal
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}

	// Check the upper right diagonal
	for i, j := row, col; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}

	return true
}

// Recursive function to solve the N-Queens problem
func solveNQueensUtil(board [][]string, row, n int) bool {
	// If all queens are placed, return true
	if row >= n {
		return true
	}

	// Try placing a queen in all columns one by one
	for col := 0; col < n; col++ {
		// Check if the queen can be placed at board[row][col]
		if isSafe(board, row, col, n) {
			// Place the queen
			board[row][col] = "Q"

			// Recur to place the rest of the queens
			if solveNQueensUtil(board, row+1, n) {
				return true
			}

			// If placing queen at board[row][col] doesn't lead to a solution,
			// then remove the queen (backtrack)
			board[row][col] = "."
		}
	}

	return false
}

// Function to solve the N-Queens problem
func solveNQueens(n int) [][]string {
	// Initialize the board with empty cells
	board := make([][]string, n)
	for i := range board {
		board[i] = make([]string, n)
		for j := range board[i] {
			board[i][j] = "."
		}
	}

	// Start solving from the first row
	if solveNQueensUtil(board, 0, n) {
		return board
	}

	// If no solution exists, return an empty board
	return [][]string{}
}

// Function to print the board
func printBoard(board [][]string) {
	for _, row := range board {
		fmt.Println(strings.Join(row, " "))
	}
}

func main() {
	// Define the size of the board
	n := 8

	// Solve the N-Queens problem
	board := solveNQueens(n)

	// Print the solution
	if len(board) == 0 {
		fmt.Println("No solution exists")
	} else {
		printBoard(board)
	}
}
