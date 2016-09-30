# GoIRS
Information Retrieval System written in Go

## About this project
This project was written for a course of Information Retrieval Systems at university. After a few months,
I decided to release it as an open source project.

## About the software
GoIRS is a library, but it has some test programs (goirs-* directories). They were only intended to work
with the Corpus we were provided, but the library should work with any text file (except some parts that
work only with spanish-written texts, as are the stopper and the stemmer). All of this may be solved
eventually.

Special thanks to [snowball stemmer library](https://github.com/kljensen/snowball) creator (saved me a lot of work).

## TODO list
- Translate comments (at the moment in spanish).
- Allow other programmers to write their own filter functions and attach them. (Provide an interface)
- Create an easy way to add process new documents (maybe feed them to a channel).
