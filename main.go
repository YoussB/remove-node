package main

import ( 
	"fmt"
	"os"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("Hello GO.")

    inputFolder := os.Args[1]

	readFolderDeleteNode(inputFolder)

}

func readFolderDeleteNode (name string) {
	
	files, err := ioutil.ReadDir(name)
    if err != nil {
        log.Fatal(err)
	}
	
    for _, f := range files {
			// fmt.Println(f.Name())
			if f.IsDir() {
				currLoc := name + "/" + f.Name()
				if f.Name() != "node_modules" && f.Name() != "bower_components" {
					readFolderDeleteNode(currLoc)
				}else if f.Name() == "node_modules" {
					fmt.Println(currLoc + " is a node.")
					os.RemoveAll(currLoc)
					fmt.Println(currLoc + " deleted.")
				} else if f.Name() == "bower_components" {
					fmt.Println(currLoc + " is a bower.")
					os.RemoveAll(currLoc)
					fmt.Println(currLoc + " deleted.")
				}
			}
    }
}