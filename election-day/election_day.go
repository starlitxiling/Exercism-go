package electionday

import (
	_ "fmt"
	"strconv"
)

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	// panic("Please implement the NewVoteCounter() function")
	var counter_pointer *int
	counter_pointer = &initialVotes
	return counter_pointer
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	// panic("Please implement the VoteCount() function")
	if counter == nil {
		return 0
	}
	return *counter
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	// panic("Please implement the IncrementVoteCount() function")
	*counter = *counter + increment
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	// panic("Please implement the NewElectionResult() function")

	// var result *ElectionResult							// also can pass test
	// result = &ElectionResult{Name: "peter", Votes: 20}
	// result.Name = candidateName
	// result.Votes = votes
	// return result

	result := &ElectionResult{
		Name:  candidateName,
		Votes: votes,
	}
	return result
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	// panic("Please implement the DisplayResult() function")

	// return fmt.Sprintf("%s (%d)", result.Name, result.Votes)		// also can pass test
	demo := result.Name + " (" + strconv.Itoa(result.Votes) + ")"
	return demo
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	// panic("Please implement the DecrementVotesOfCandidate() function")
	// result_pointer := &results
	// for i, j := range results {   // just copy iterater
	// 	if i == candidate {
	// 		j -= 1
	// 	}
	// }
	for i := range results { //correct answer
		results[i] -= 1
	}
	// if _, exists := results[candidate]; exists {   // also correct answer
	// 	results[candidate] -= 1
	// }
}
