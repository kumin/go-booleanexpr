# go-booleanexpr
## What is go-booleanexr
This library provides method to calculating a boolean expression in Golang.
Input is a expression string and the list of Elements.
Every element is object which implement element interface.

### Motivation

In real life sometime you have to check list of conditions.
For example, check if a customer belong to a Corhot.

### Example

You have list of element A, B, C
+ A.Check() = True
+ B.Check() = False
+ C.Check() = True

The expression: (A | B) & C

The result is True

