package practice

/*
Problem:
Ref: https://leetcode.com/problems/word-ladder/description/

Given:
- beginWord and endWord
- a dictionary wordList
- A transformation sequence from word beginWord to word endWord using a dictionary wordList is a sequence of words beginWord -> s1 -> s2 -> ... -> sk such that:

Every adjacent pair of words differs by a single letter.
Every si for 1 <= i <= k is in wordList. Note that beginWord does not need to be in wordList.
sk == endWord

Return:
- Number of word in the shortest transformation sequence
- 0 if no such sequence exists

Constraints:
- length of word: [1, 10]
- length of wordList: [1, 5000]
- all words have the same length
- all words consist of lowercase English letters
- all words in wordList are unique
*/

/*
First approach:
Consider each transformation between 2 word is an edge of 2 nodes.
Use BFS to traverse until reaching end word
*/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	canDo := false
	for _, word := range wordList {
		if word == endWord {
			canDo = true
		}
	}

	if !canDo {
		return 0
	}

	wordLen := len(beginWord)
	visited := make(map[string]bool)
	queue := make([]string, 0)
	queue = append(queue, beginWord)

	sequence := 0
	for len(queue) > 0 {
		sequence++
		count := len(queue)
		for _, word := range queue {
			if !visited[word] {
				visited[word] = true
				for _, nextWord := range wordList {
					diff := 0
					for i := 0; i < wordLen; i++ {
						if word[i] != nextWord[i] {
							diff++
						}
					}

					if diff == 1 {
						if nextWord == endWord {
							return sequence
						}

						queue = append(queue, nextWord)
					}
				}
			}
		}

		queue = queue[count:]
	}

	return 0
}
