// A program to ingest a list of words and return the longest word that
// exists that is a concatenation of other words in the list
// When a list is completely comprised of unique words this
// this program will let you know that no word is comprised
// of other words

package main

import (
    "log"
    "os"
    "io/ioutil"
//    "hash"
    "strings"
    "sort"
)


/* A function which returns true or false depending on whether
    there was an error or not, if there is an error msg will be
    logged to the console. */
func printError(err error, msg string) bool {
    if err != nil {
        log.Printf("%s: %s", msg, err)
        return true
    }
    return false
}

/* Sort functions for strings */
type ByLength []string
func (s ByLength) Len() int {
    return len(s)
}
func (s ByLength) Less(i, j int) bool {
    return len(s[i]) > len(s[j])
}
func (s ByLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
    
/* A function to take an input list of words and output an
    array of strings in the file delimited by newline as well
    as a map of those strings to boolean values for O(1) checks
    for existence in the file

    The returned array will be sorted in descending length
*/
func wordsFromFile(path string) (words []string, lookupTable map[string]bool, err error) {
    log.Println("Attempting to read the file: %s", path)
    data, err := ioutil.ReadFile(path)
    if printError(err, "Failed to read the file") {
        return nil, nil, err
    }
    
    tmp := strings.Split(string(data), "\n")
    words = []string{}
    for _, word := range tmp {
        if len(word) != 0 {
            words = append(words, word)
        }
    }
    
    lookupTable = make(map[string]bool)
    sort.Sort(ByLength(words))
    for idx, word := range words {
        words[idx] = strings.TrimSpace(word)
        lookupTable[words[idx]] = true
    }   
    return words, lookupTable, nil
}

/* A recursive function to determine if this word is composed of words in the
   dictionary of words, if it is it will return true and a list of all of the
   words used in its composition */
func recComposedOfSub (word string, lookupTable map[string]bool) (bool, []string)  {
    tmp := ""
    for idx, runeVal := range word {
        tmp += string(runeVal)
        if _, ok := lookupTable[tmp]; ok {
            subWords := []string{tmp}
            if word == tmp {
                // Matched and no more input
                return true, subWords
            }

            if ok, found := recComposedOfSub(word[idx + 1:], lookupTable); ok {
                // All remaining input matched
                subWords = append(subWords, found...)
                return true, subWords
            }
        }
    }
    return false, nil
}
/* A wrapper to the recursive function in order to create a base case */
func composedOfSub (word string, lookupTable map[string]bool) (bool, []string)  {
    delete(lookupTable, word)
    ok, res := recComposedOfSub(word, lookupTable)
    lookupTable[word] = true
    return ok, res
}
    
/* The main function */
func main() {
    if len(os.Args) < 2 {
        log.Fatal("Unable to read input file")
        return
    }   

    inputFileName := os.Args[1]
    words, table, err := wordsFromFile(inputFileName)
    if err != nil {
        return
    }
    for _, word := range words {
        if ok, composition := composedOfSub(word, table); ok {
            log.Println("Longest Word Found", word)
            log.Println("Composed of:")
            for _, subWord := range composition {
                log.Println(subWord)
            }
            return
        }
    }
    log.Println("All of the words in the input file are unique")
}