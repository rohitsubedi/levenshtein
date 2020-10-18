package test_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rohitsubedi/levenshtein"
)

func TestMinOperationToConvertOneStringToAnother1(t *testing.T) {
	destString := "DE-2341-00001"
	// For same strings
	minOperation := levenshtein.GetLevenshteinDistance([]rune(destString), []rune(destString), false)
	assert.Equal(t, 0, minOperation)
	// For same strings but different case(lower case)
	minOperation = levenshtein.GetLevenshteinDistance([]rune(strings.ToLower(destString)), []rune(destString), false)
	assert.Equal(t, 0, minOperation)
	// For string with 1 char missing
	minOperation = levenshtein.GetLevenshteinDistance([]rune("D-2341-00001"), []rune(destString), false)
	assert.Equal(t, 1, minOperation)
	minOperation = levenshtein.GetLevenshteinDistance([]rune("DE-231-00001"), []rune(destString), false)
	assert.Equal(t, 1, minOperation)
	minOperation = levenshtein.GetLevenshteinDistance([]rune("DE-2341-0001"), []rune(destString), false)
	assert.Equal(t, 1, minOperation)
	minOperation = levenshtein.GetLevenshteinDistance([]rune("DE-234100001"), []rune(destString), false)
	assert.Equal(t, 1, minOperation)
	// For string with 1 extra char
	minOperation = levenshtein.GetLevenshteinDistance([]rune("DEE-2341-00001"), []rune(destString), false)
	assert.Equal(t, 1, minOperation)
	minOperation = levenshtein.GetLevenshteinDistance([]rune("DE-23441-00001"), []rune(destString), false)
	assert.Equal(t, 1, minOperation)
	minOperation = levenshtein.GetLevenshteinDistance([]rune("DE-2341-000001"), []rune(destString), false)
	assert.Equal(t, 1, minOperation)
	minOperation = levenshtein.GetLevenshteinDistance([]rune("DE-2341--00001"), []rune(destString), false)
	assert.Equal(t, 1, minOperation)
	// For string with 2 char missing
	minOperation = levenshtein.GetLevenshteinDistance([]rune("-2341-00001"), []rune(destString), false)
	assert.Equal(t, 2, minOperation)
	minOperation = levenshtein.GetLevenshteinDistance([]rune("DE-21-00001"), []rune(destString), false)
	assert.Equal(t, 2, minOperation)
	minOperation = levenshtein.GetLevenshteinDistance([]rune("DE-2341-000"), []rune(destString), false)
	assert.Equal(t, 2, minOperation)
	minOperation = levenshtein.GetLevenshteinDistance([]rune("DE234100001"), []rune(destString), false)
	assert.Equal(t, 2, minOperation)
}

func TestMinOperationToConvertOneStringToAnother2(t *testing.T) {
	destString := "DE-2341-00001"
	// For string with 2 extra char
	minOperation := levenshtein.GetLevenshteinDistance([]rune("DEE--2341-00001"), []rune(destString), false)
	assert.Equal(t, 2, minOperation)
	minOperation = levenshtein.GetLevenshteinDistance([]rune("DE-2344-1-00001"), []rune(destString), false)
	assert.Equal(t, 2, minOperation)
	minOperation = levenshtein.GetLevenshteinDistance([]rune("DE-2341-0000199"), []rune(destString), false)
	assert.Equal(t, 2, minOperation)
	minOperation = levenshtein.GetLevenshteinDistance([]rune("DE--2341--00001"), []rune(destString), false)
	assert.Equal(t, 2, minOperation)
	// For string with more than 2 char missing
	minOperation = levenshtein.GetLevenshteinDistance([]rune("2341-00001"), []rune(destString), false)
	assert.Equal(t, 3, minOperation)
	minOperation = levenshtein.GetLevenshteinDistance([]rune("DE-34-001"), []rune(destString), false)
	assert.Equal(t, 4, minOperation)
	minOperation = levenshtein.GetLevenshteinDistance([]rune("DE34-001"), []rune(destString), false)
	assert.Equal(t, 5, minOperation)
	// For string with more than 2 extra char
	minOperation = levenshtein.GetLevenshteinDistance([]rune("DEE--23341-00001"), []rune(destString), false)
	assert.Equal(t, 3, minOperation)
	minOperation = levenshtein.GetLevenshteinDistance([]rune("DEE--23341-000001"), []rune(destString), false)
	assert.Equal(t, 4, minOperation)
	minOperation = levenshtein.GetLevenshteinDistance([]rune("DEE--233414-000001"), []rune(destString), false)
	assert.Equal(t, 5, minOperation)
	// Test case sensitive
	minOperation = levenshtein.GetLevenshteinDistance([]rune("ROHIT"), []rune("rohit"), true)
	assert.Equal(t, 5, minOperation)
	// Test case in sensitive
	minOperation = levenshtein.GetLevenshteinDistance([]rune("ROHIT"), []rune("rohit"), false)
	assert.Equal(t, 0, minOperation)
	// UTF-8 characters
	minOperation = levenshtein.GetLevenshteinDistance([]rune("rohit"), []rune("röhit"), false)
	assert.Equal(t, 1, minOperation)
	// Test empty string
	minOperation = levenshtein.GetLevenshteinDistance([]rune("ROHIT"), []rune(""), true)
	assert.Equal(t, 5, minOperation)
	// Test empty string
	minOperation = levenshtein.GetLevenshteinDistance([]rune("ö"), []rune(""), true)
	assert.Equal(t, 1, minOperation)
	// Test case in sensitive
	minOperation = levenshtein.GetLevenshteinDistance([]rune("ROHIT"), []rune("tihor"), false)
	assert.Equal(t, 4, minOperation)
	// UTF-8 characters same
	minOperation = levenshtein.GetLevenshteinDistance([]rune("röhit"), []rune("röhit"), false)
	assert.Equal(t, 0, minOperation)
}
