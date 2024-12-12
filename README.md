[![Travis CI](https://img.shields.io/travis/theosiemensrhodes/wordsegmentation/master.svg?style=flat-square)](https://travis-ci.org/theosiemensrhodes/wordsegmentation)
[![Software License](https://img.shields.io/badge/License-MIT-orange.svg?style=flat-square)](https://github.com/theosiemensrhodes/wordsegmentation/LICENSE.md)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/theosiemensrhodes/wordsegmentation)

# Word segmentation
Word segmentation is the process of dividing a phrase without spaces back into its constituent parts. For example, consider a phrase like "thisisatest". Humans can immediately identify that the correct phrase should be "this is a test".

## Source and credits
This is a modified fork of https://github.com/YafimK/wordsegmentation - with the following main modifications:
1. Exposing the max length parameter allow users to improve performance as a tradeoff for decreasing the maximum length of a word that can be recognized
2. Various other small improvements

This package is heavily inspired by the Python module [grantjenks/wordsegment](https://github.com/grantjenks/wordsegment).

Copyright (c) 2015 by Grant Jenks under the Apache 2 license

The package is based on code from the chapter [Natural Language Corpus Data](http://norvig.com/ngrams/) by Peter Norvig from the book [Beautiful Data](http://oreilly.com/catalog/9780596157111/) (Segaran and Hammerbacher, 2009).

Copyright (c) 2008-2009 by Peter Norvig

## Getting started
You can grab this package with the following command:
```
go get gopkg.in/theosiemensrhodes/wordsegmentation.v0
```

## Usage
If you wanna use the default English corpus:
```go
package main

import (
    "fmt"

    "github.com/theosiemensrhodes/wordsegmentation"
    "github.com/theosiemensrhodes/wordsegmentation/corpus"
)

func main() {
    // Grab the default English corpus that will be created thanks to TSV files
    englishCorpus := corpus.NewEnglishCorpus()
    fmt.Println(wordsegmentation.Segment(englishCorpus, "thisisatest"))
}
```

## Unigrams and bigrams
> Information: an **n-gram** is a contiguous sequence of n items from a given sequence of text or speech.

This package ships with an English corpus by default that is ready to use. [Data files](https://github.com/theosiemensrhodes/wordsegmentation/tree/master/data) are derived from the [Google web trillion word corpus](http://googleresearch.blogspot.com/2006/08/all-our-n-gram-are-belong-to-you.html), as described by Thorsten Brants and Alex Franz, and [distributed](https://catalog.ldc.upenn.edu/LDC2006T13) by the Linguistic Data Consortium. This module contains only a subset of that data. The unigram data includes only the most common 333,000 words. Similarly, bigram data includes only the most common 250,000 phrases. Every word and phrase is lowercased with punctuation removed.

## Using a custom corpus
If you want to use a custom corpus, you will need to implement the `Corpus` interface to give to the `Segment` method.

The interface is as follow:
```go
// The corpus interface that lets access bigrams,
// unigrams, the total number of words from the corpus
// and a function to clean a string.
type Corpus interface {
    Bigrams() *models.Bigrams
    Unigrams() *models.Unigrams
    Total() float64
    Clean(string) string
}
```

Take a look at the [English corpus source code](corpus/english.go) to help you start!

## Documentation
The documentation of this package can be found on [GoDoc](https://godoc.org/github.com/theosiemensrhodes/wordsegmentation). Here is a list of links for the different modules:
- [`corpus`](https://godoc.org/github.com/theosiemensrhodes/wordsegmentation/corpus) - the default English corpus
- [`helpers`](https://godoc.org/github.com/theosiemensrhodes/wordsegmentation/helpers) - little functions to get the length of a string, remove special characters of a string, get the minimum between 2 given integers
- [`models`](https://godoc.org/github.com/theosiemensrhodes/wordsegmentation/models) - the various objects used (Unigrams, Bigrams, Arrangement, Candidate, Possibility)
- [`parsers`](https://godoc.org/github.com/theosiemensrhodes/wordsegmentation/parsers) - parsers to read tab-separated files into Unigrams and Bigrams
- [`segment`](https://godoc.org/github.com/theosiemensrhodes/wordsegmentation) - the 'main' package
