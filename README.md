[![Software License](https://img.shields.io/badge/License-MIT-orange.svg?style=flat-square)](https://github.com/theosiemensrhodes/wordsegmentation/LICENSE.md)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/theosiemensrhodes/wordsegmentation)

# Word segmentation
Word segmentation is the process of dividing a phrase without spaces back into its constituent parts. For example, consider a phrase like "thisisatest". Humans can immediately identify that the correct phrase should be "this is a test".

## Source and credits
This is a modified fork of https://github.com/YafimK/wordsegmentation - with the following main modifications:
1. Overhauling the segmentation logic to used array based dynamic programming instead of memoization leading to a
2. Exposing the Segmentor's maximum word length parameter for user's to override. A shorter maximum word length is computationally cheaper at the cost of not being able to recognize words over the maximum word length's value.
3. Various other performance improvements in the parsing and searching of bigrams and unigrams

This package is inspired by the Python module [grantjenks/wordsegment](https://github.com/grantjenks/wordsegment).

Copyright (c) 2015 by Grant Jenks under the Apache 2 license

The package is based on from the chapter [Natural Language Corpus Data](http://norvig.com/ngrams/) by Peter Norvig from the book [Beautiful Data](http://oreilly.com/catalog/9780596157111/) (Segaran and Hammerbacher, 2009).

Copyright (c) 2008-2009 by Peter Norvig

## Getting started
You can grab this package with the following command:
```
go get gopkg.in/theosiemensrhodes/wordsegmentation
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
> Information: an **n-gram** is a contiguous sequence of n words from a given sequence of text or speech.

This package ships with an English corpus by default that is ready to use. [Data files](https://github.com/theosiemensrhodes/wordsegmentation/tree/master/data) are derived from intersecting the [Google Books Ngram Corpus - Version 2012](https://storage.googleapis.com/books/ngrams/books/datasetsv3.html) [(License)](http://creativecommons.org/licenses/by/3.0/) with a combination of [SCOWL - Spell Checker Oriented Word Lists](http://wordlist.aspell.net/) [(License)](http://wordlist.aspell.net/scowl-readme/) and a [wordlist scraped from Merriam Webster](link tbd). This module contains only a subset of that data, as it is not reasonable to load it all into memory. The full process in which this corupus was imported and created can be found in [this repo](link tbd)

## Using a custom corpus
If you want to use a custom corpus, you will need to implement the `Corpus` interface to give to the `NewSegmentor` constructor.

The interface is as follow:
```go
// The corpus interface that lets access bigrams,
// unigrams, the total number of words from the corpus,
// a function to clean a string and a function to get
// the maximum word length in the dictionary.
type Corpus interface {
    Bigrams() *models.Bigrams
    Unigrams() *models.Unigrams
    Total() float64
    Clean(string) string
    MaxWordLength() int
}
```

Take a look at the [English corpus source code](corpus/english.go) to help you start!

## Documentation
The documentation of this package can be found on [GoDoc](https://godoc.org/github.com/theosiemensrhodes/wordsegmentation). Here is a list of links for the different modules:
- [`corpus`](https://godoc.org/github.com/theosiemensrhodes/wordsegmentation/corpus) - the default English corpus
- [`models`](https://godoc.org/github.com/theosiemensrhodes/wordsegmentation/models) - the various objects used (Unigrams, Bigrams)
- [`parsers`](https://godoc.org/github.com/theosiemensrhodes/wordsegmentation/parsers) - parsers to read tab-separated files into Unigrams and Bigrams
- [`segment`](https://godoc.org/github.com/theosiemensrhodes/wordsegmentation) - the 'main' package
