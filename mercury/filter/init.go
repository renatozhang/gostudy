package filter

import (
	"bufio"
	"io"
	"os"

	"github.com/renatozhang/gostudy/mercury/util"
)

var (
	trie *util.Trie
)

func Init(fileName string) (err error) {
	trie = util.NewTrie()
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		word, errRet := reader.ReadString('\n')
		if errRet == io.EOF {
			return
		}
		if errRet != nil {
			err = errRet
			return
		}
		err = trie.Add(word, nil)
		if err != nil {
			return
		}
	}
}
