# Prototype

When it is easier to copy an existing object to fully initialise a new one

## Motivation
- Complicated objects (e.g., cars) aren't designed from scratch
    - They reiterate existing designs
- An existing (partially or fully constructed) design is a Prototype
- We make a copy of the protype and customise it
    - Requires 'deep copy' support
- We make the cloning convenient (e.g., via a Factory)

A **Prototype** is a partially or fully initialised object that you copy (clone) and make use of.

## Summary

- To implement a prototype, partially construct an object and store it somewhere
- Deep copy the prototype
- Customise the resulting instance
- A prototype factory provides a convenient API for using prototypes