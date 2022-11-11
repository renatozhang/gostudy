package filter

func Replace(text string, replace string) (result string, isReplace bool) {
	result, isReplace = trie.Check(text, replace)
	return
}
