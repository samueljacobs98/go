# Chain of Responsibility

A chain of components who all get a chance to process a command or a query, optionally having default processing implementation an ability to terminate the processing chain.

## Motivation

Examples:
- Unethical behaviour by an employee; who takes the blame?
    - Employee
    - Manager
    - CEO
- You click a grpahical element on a form
    - Button handles it, stops further processing
    - Underlying group box
    - Underlying window
- CCG computer game
    - Creature has attack and defense values
    - Those can be boosted by other cards

## Summary

- Chain of responsibility can be implemented as a linked list of pointers or a centralised construct
- Enlist objects in the chain, possibly controlling their order
- Control object removal from chain