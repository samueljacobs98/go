# Command

## Motivation

- Ordinary statements are perishable
    - Cannot undo field assignment
    - Cannot directly serialise a sequence of actions (calls)
- Want an object that represents an operation
    - person should change its _age_ to value _22_
    - _car_ should _explode()_
- Uses: GUI commands, multi-level undo/redo, macro recording and more!

The Command design pattern is an object which represents an instruction to perform a particular action. Contains all the information necessary for the action to be taken.

## Summary

- Encapsulate all details of an operation in a separate object
- Define functions for applying the command (either in the command itself, or elsewhere)
- Optionally define instructions for undoing the command
- Can create composite commands (a.k.a. macros)