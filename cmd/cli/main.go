package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/funayman/countdown-solver"
	flag "github.com/spf13/pflag"
)

var (
	needsHelp bool
	dictFile  string
	maxSize   int
	minLength int
)

func usage() {
	fmt.Printf("usage: %s [options] <letters>\n", os.Args[0])
	fmt.Println("options:")
	flag.PrintDefaults()
	fmt.Println()
	fmt.Println("examples:")
	fmt.Println("  $ ./countdown 'sraneimax'")
	fmt.Println("  $ ./countdown --max-size 10 'mhseoudma'")
	fmt.Println("  $ ./countdown --dict /usr/share/dict/words --min-length 4 'eoiadnnzt'")
}

func init() {
	flag.StringVarP(&dictFile, "dict", "d", "words.txt", "Location of dictionary file")
	flag.BoolVarP(&needsHelp, "help", "h", false, "Prints this message")
	flag.IntVarP(&minLength, "min-length", "l", 1, "Minimum length answers need to be")
	flag.IntVarP(&maxSize, "max-size", "s", 0, "Maximum size of word list")
	flag.Parse()
}

func main() {
	if needsHelp {
		usage()
		return
	}

	if len(flag.Args()) != 1 {
		usage()
		log.Fatalln("missing letters input")
	}

	letters := strings.ToLower(flag.Arg(0))

	game, err := countdown.New(dictFile)
	if err != nil {
		log.Fatalf("countdown.New(%q): %v", dictFile, err)
	}

	results := game.Solve(letters)

	if maxSize == 0 {
		maxSize = len(results)
	}

	for i := 0; i < len(results) && i < maxSize && minLength <= len(results[i]); i++ {
		fmt.Println(results[i])
	}
}
