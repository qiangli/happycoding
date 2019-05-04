package dynamic

import (
	"testing"
)

func TestWorkBreak(t *testing.T) {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	if !wordBreak(s, wordDict) {
		t.FailNow()
	}

	s = "applepenapple"
	wordDict = []string{"apple", "pen"}
	if !wordBreak(s, wordDict) {
		t.FailNow()
	}

	s = "catsandog"
	wordDict = []string{"cats", "dog", "sand", "and", "cat"}
	if wordBreak(s, wordDict) {
		t.FailNow()
	}

	s = "aaaaaaa"
	wordDict = []string{"aaaa", "aaa"}
	if !wordBreak(s, wordDict) {
		t.FailNow()
	}

	s = "bb"
	wordDict = []string{"a", "b", "bbb", "bbbb"}
	if !wordBreak(s, wordDict) {
		t.FailNow()
	}

	s = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	wordDict = []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	if wordBreak(s, wordDict) {
		t.FailNow()
	}
}
