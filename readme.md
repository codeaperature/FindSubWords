# Find Sub Words

Given a list of words, like word.list, the FindSubWords.go 
program's compiled executable will find the longest compound-word in the 
list. These compounds are a concatenation of other sub-words that exist
in the list. The program allow the user to input different data via the 
commmand-line by specifying a file name. Any programming language can be 
used, but Go is preferred.


## Stipulations
Stipulations are being made for ease of coding. Programming changes would be
required in order to bypass these restrictions:

1. The dictionary contains words that start with all letters (a-z)
2. There is a compound word to be found (composing of other dictionary words)
3. Compounds can be made of 2 OR MORE other dictionary words
4. The dictionary is sorted from the start
5. Dictionary is ASCII characters and all LOWER CASE
6. No --help parameter is required


## Compiling

In order to compile the program, Golang 1.4 is suggested although other 
versions may be compatible (TBD). Here is the tested version:

$ go version
go version go1.4 darwin/386
$



## Running
The go source file is FindSubWords.go and the sample data file is FindSubWords.txt. 
Here is the compile line & run line:

$ go run FindSubWords.go FindSubWords.txt
Using Dictionary:  FindSubWords.txt
FINAL Word:  antidisestablishmentarianisms
$

Swap out FindSubWords.txt for any data file ... such as the original word.list 
file. (See Stipulations and requirements listed above.)

## Comment
This code was an entrance to an interview posed by an unnamed company.


Enjoy, Stephan


