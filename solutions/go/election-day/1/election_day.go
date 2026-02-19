package electionday
import "fmt"

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
    var p *int
    p = &initialVotes
	return p
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	if (counter != nil) {
        return *counter
    }
    return 0
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	if (counter != nil) {
        *counter += increment
    }
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
    var p *ElectionResult
    p = &ElectionResult{
        Name: candidateName,
        Votes: votes,
    }
    return p
    
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	var p *ElectionResult
    p = result
    return fmt.Sprintf("%s (%x)", p.Name, p.Votes)
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
    results[candidate] -= 1
}
