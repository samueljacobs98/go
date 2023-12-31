# Decorator

Adding behaviour without altering the type itself.

## Motivation

- Want to augment an object with additional functionality
- Do not want to rewrite or alter existing code (OCP)
- Want to keep new functionality separate (SRP)
- Need to be able to interact with existing structures
- Solution: embed the decorated object and provide additional functionality

The Decorator facilitates the addition of behaviours to individual objects through embedding.

## Summary

- A decorator embeds the decorated object
- Adds utility fields and methods to augment the object's features
- Often used to emulate multiple inheritance (may require extra work)