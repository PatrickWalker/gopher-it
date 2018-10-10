# Gopher IT

## Overview

Inspired and forked from [go-102](https://github.com/timblair/go-102) with some additions 

[Go](https://golang.org/) is a small language, but one with a few differences
to many common languages.  For example, Go supports an OO-style of development,
but doesn't have classes, and it has built-in primitives to simplify concurrent
programming.

The content in this book is aimed at introducing these concepts to someone new
to Go (but not new to programming in general).  The content is split in to the
following sections:

* [Introduction](topics/the-basics/): built-in types, variable declaration,
  function and custom types.
    * Types
    * Conditional Logic
    * Structs
    * Functions
    * Errors
* [Object-oriented development](topics/object-orientation/)
    * [Methods](topics/object-orientation/methods.md)
    * [Interfaces](topics/object-orientation/interfaces.md)
    * [Embedding](topics/object-orientation/embedding.md)
    * [Composition](topics/object-orientation/composition.md)
* [Concurrency] (topics/concurrency)
* [Testing](topics/testing/)
* [User Input](topics/handling-input/) 
* [HTTP](topics/http/)

Each section explains the relevant concepts, walks through a number of examples
and code samples, and ends with an exercise to put the concepts in to practice.

## Pre-Requisites

Although you need have no real experience with Go itself, this is not a
"learning to code" guide.  Specifically, you are expected to have practical
experience with the principles behind object-oriented software development (in
any language).

## Exercises

There are templates available for each of the exercises.  If you've cloned this
repository under your `$GOPATH` then you can just edit and run the template
file in place.  Alternatively, follow the link to the template already set up
in the Go Playground, and work from there instead.

* The Basics: [source](topics/the-basics/exercises/basics/template/basics.go) /
  [playground](http://play.golang.org/p/ta6oFzjgwn)
* Object Orientation
  * Methods: [source](topics/object-orientation/exercises/methods/template/methods.go) /
    [playground](http://play.golang.org/p/jnBw-jtE3n)
  * Interfaces: [source](topics/object-orientation/exercises/methods/template/methods.go) /
    [playground](http://play.golang.org/p/rL5tT2VTJH)
  * Embedding: [source](topics/object-orientation/exercises/embedding/template/embedding.go) /
    [playground](http://play.golang.org/p/5qrrcfHdiZ)
* Concurrency
  * Goroutines: [source](topics/concurrency/exercises/goroutines/template/goroutines.go) /
    [playground](http://play.golang.org/p/EH_16WR5ND)
  * Channels: [source](topics/concurrency/exercises/channels/template/channels.go) /
    [playground](http://play.golang.org/p/H4F9aLKQVA)

Example solutions are also available, but don't look at those until you've had
a go yourself!

## Running a Workshop

This guide can be used as an individual study-aid, but it was designed to be
presented as a hands-on workshop.  In this format, the workshop should take
around 8 hours, with each section talked through by the presenter, before
the participants work through the relevant exercise (preferably in pairs).  

## Source and Contributing

The source for this material is [available on
GitHub](https://github.com/patrickwalker/gopher-it). Please raise any issues there, and
contributions (fixing typos, better formatting, improvements to code examples
etc) are more than welcome.

## Attribution

This is based and built upon [go-102](https://github.com/timblair/go-102) 

## License

<a rel="license" href="http://creativecommons.org/licenses/by-sa/4.0/"><img
alt="Creative Commons License" style="border-width:0"
src="https://i.creativecommons.org/l/by-sa/4.0/88x31.png" /></a>

Except where otherwise noted, the content of this repository is licensed under
a <a rel="license"
href="http://creativecommons.org/licenses/by-sa/4.0/">Creative Commons
Attribution-ShareAlike 4.0 International License</a>.
