// Test file for functions

package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestFileRead(t *testing.T) {
    assert := assert.New(t)
    words, table, err := wordsFromFile("./non_existent.txt")
    assert.Nil(words, "words must not have loaded")
    assert.Nil(table, "table must not have loaded")
    assert.NotNil(err, "error needs to be defined")
    
    words, table, err = wordsFromFile("./test_words.txt")
    assert.Nil(err, "should have no errors")
    assert.Len(words, 3, "should only see 3 words")
    for _, word := range words {
        val, ok := table[word]
        assert.True(ok, "Word must be in table")
        assert.True(val, "True must have been set for the word")
    }
}

func TestComposedOfSub(t *testing.T) {
    assert := assert.New(t)
    
    words, table, err := wordsFromFile("./test_composition.txt")
    assert.NotNil(words, "Words need to exist")
    assert.Len(words, 6, "should only see 6 words")
    assert.NotNil(table, "Table needs to exist")
    assert.Nil(err, "should have no errors")
    
    assert.Equal(words[0], "anotherwordnotmaching")
    ok, list := composedOfSub(words[0], table)
    assert.False(ok, "'anotherwordnotmaching' isn't composed")
    assert.Nil(list)

    assert.Equal(words[1], "notmatchinganyother")
    ok, list = composedOfSub(words[1], table)
    assert.False(ok, "'notmatchinganyother' isn't composed")
    assert.Nil(list)

    assert.Equal(words[2], "anotherwordword")
    ok, list = composedOfSub(words[2], table)
    assert.True(ok, "'anotherwordword' should be composed")
    assert.Contains(list, "another", "'another' needs to be in list")
    assert.Contains(list, "word", "'word' needs to be in list")

    assert.Equal(words[3], "anotherword")
    ok, list = composedOfSub(words[3], table)
    assert.True(ok, "'anotherword' should be composed")
    assert.Contains(list, "another", "'another' needs to be in list")
    assert.Contains(list, "word", "'word' needs to be in list")

    assert.Equal(words[4], "another")
    ok, list = composedOfSub(words[4], table)
    assert.False(ok, "'another' isn't composed")
    assert.Nil(list)
    
    assert.Equal(words[5], "word")
    ok, list = composedOfSub(words[0], table)
    assert.False(ok, "'word' isn't composed")
    assert.Nil(list) 
}