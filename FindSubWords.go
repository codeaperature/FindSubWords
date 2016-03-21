/*
Q: Given a list of words like
https://github.com/NodePrime/quiz/blob/master/word.list find the longest
compound-word in the list, which is also a concatenation of other sub-words
that exist in the list. The program should allow the user to input different
data. The finished solution shouldn't take more than one hour. Any
programming language can be used, but Go is preferred.

Fork this repo, add your solution and documentation on how to compile and
run your solution, and then issue a Pull Request.

Obviously, we are looking for a fresh solution, not based on others' code.

A. Stipulations
1. The dictionary contains words that start with all letters
2. There is a compound word to be found (composing of other dictionary words)
3. Compounds can be made of 2 OR MORE other dictionary words
4. The dictionary is sorted from the start
5. Dictionary is ASCII characters
...
100. Code documentation does not need to follow any standards
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type FinderType struct {
	Length int // could be a smaller type size if words are smaller than int
	Index  int
}

// Use this to sort the disctionary words by longest to shortest word
type LenSorter []FinderType

func (a LenSorter) Len() int           { return len(a) }
func (a LenSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a LenSorter) Less(i, j int) bool { return a[i].Length > a[j].Length }

// alfwords is the word dictionary starting wiht the letter associated
// with the first index and has second index related to the number of words
// starting with that letter (a = 0, z = 25)
// This list will be sorted longest to shortest
var alfwords [26][]FinderType

// allwords is the entire dictionary of words sorted longest to shortest
// but represented as only a length and index into the dictionary
var allwords []FinderType

// dictionary -- each dictionary words
var dictionary []string
var allinc int

//
// readAndOrganizeDictionary - reads a whole file into memory and returns
// a slice of its dictionary words. The rountine preps allwords &
// alfwords by sorting each
//
func readAndOrganizeDictionary(path string) error {

	// open file
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// use lastletter to flip to the next main index of alfwords
	var lastletter uint8

	// initialize the counters
	var alfinc int
	allinc = 0
	alfinc = 0

	// file read stuff
	scanner := bufio.NewScanner(file)
	// build the dictionary and organize our search criteria
	for scanner.Scan() {
		// build the dictionary
		dictionary = append(dictionary, scanner.Text())
		// check for flipping to next alfwords
		firstletter := dictionary[allinc][0] - 'a'
		if firstletter != lastletter {
			lastletter = firstletter
			alfinc = 0
		}
		// in order to sort our slices by dictionary word's len
		wordlen := len(dictionary[allinc])
		// add to our slices
		alfwords[firstletter] = append(alfwords[firstletter], FinderType{wordlen, allinc})
		allwords = append(allwords, FinderType{wordlen, allinc})
		allinc++
		alfinc++
	}

	// sort all the data to help quickly find our words
	sort.Sort(LenSorter(allwords))
	for i := 0; i < 26; i++ {
		sort.Sort(LenSorter(alfwords[i]))
	}
	return scanner.Err()
}

//
// recurse - simply named function searches for solutions for compound
// words by comparing words that start with the same 1st letter. If match is
// found, then partial word is stripped off to be handled back to a
// recursive search
//
func recurse(level uint8, word string) uint8 {

	// are we done? no word to find sub word
	if word == "" {
		// if we have several parts ...
		if level > 1 {
			//return winner
			return level
		} else {
			// same word overlayed ... try again
			return 0
		}
	}

	// find the test words's length
	wordlen := len(word)

	// strip off 1st letter of test word
	firstletter := word[0] - 'a'

	// length of array to use
	alflen := len(alfwords[firstletter])

	// define rval default return value
	var rval uint8
	rval = 0

	i := 0
	if level == 0 {
		// loop to at shorter word - first time - skip past words
		// that are as long as the search word
		for ; (i < alflen) && (alfwords[firstletter][i].Length >= wordlen); i++ {
		}
	} else {
		// loop to shorter or equal length word - 2nd+ time, last or final
		//  word can match length or be part of compound words
		for ; (i < alflen) && (alfwords[firstletter][i].Length > wordlen); i++ {
		}
	}

	// loop over remaining words that start with the letter in the remaining
	// longest word
	for ; (i < alflen) && (alfwords[firstletter][i].Length <= wordlen); i++ {
		if strings.Index(word, dictionary[alfwords[firstletter][i].Index]) == 0 {
			rval = recurse(level+1, word[alfwords[firstletter][i].Length:])
			// yes! ... we found compound word quit loop and return
			if rval > 0 {
				break
			}
		}
	}
	return rval
}

//
// main - simple main that calls two main routines that do all the work
//
func main() {
	// get dictionary and organized structures
	fmt.Println("Using Dictionary: ", os.Args[1])
	err := readAndOrganizeDictionary(os.Args[1])
	if err != nil {
		log.Fatalf("Problem getting / organizing dictionary: %s", err)
	}
	// loop over longest to shortest words to find 1st compound word
	for i := 0; i < allinc; i++ {
		if recurse(0, dictionary[allwords[i].Index]) > 0 {
			fmt.Println("FINAL Word: ", dictionary[allwords[i].Index])
			break
		}
	}
}
