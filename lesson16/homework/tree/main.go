package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

func ListDir(dirPath string, deep int) error {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}
	if deep == 1 {
		fmt.Printf("|---%s\n", filepath.Base(dirPath))
	}
	// windows的目录分隔符是 \
	// linux的目录分隔符是 /
	sep := string(os.PathSeparator)
	for _, fi := range dir {
		//如果是目录，继续调用ListDir进行遍历
		if fi.IsDir() {
			fmt.Printf("|")
			for i := 0; i < deep; i++ {
				fmt.Printf("    |")
			}
			fmt.Printf("----%s\n", fi.Name())
			ListDir(dirPath+sep+fi.Name(), deep+1)
			continue
		}
		fmt.Printf("|")
		for i := 0; i < deep; i++ {
			fmt.Printf("    |")
		}
		fmt.Printf("----%s\n", fi.Name())
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "tree"
	app.Usage = "list all files"
	app.Action = func(c *cli.Context) error {
		var dir string = "."
		if c.NArg() > 0 {
			dir = c.Args().Get(0)
		}
		ListDir(dir, 1)
		return nil
	}
	app.Run(os.Args)
}
