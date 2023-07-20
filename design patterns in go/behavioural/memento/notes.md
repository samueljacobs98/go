# Memento

## Motivation

- An object or system goes through changes
    - E.g., a bank account gets deposits and withdrawals
- There are different ways of navigating those changes
- One way is to record every change (Command) and teach a command to 'undo' itself
    - Also part of CQRS = Command Query Responsibility Segregation
- Another is to simply save snapshots of the system.

The Memento is a token representing the system state. It lets us roll back to the state when the token was generated. May or may not directly expose information.

## Memento vs Flyweight

- Both aptterns provide a 'token' clients can hold on to
- Memento is used _only_ to be fed back into the system
    - No public/mutable state
    - No methods
- A flyweight is similar to an ordinary reference to object
    - Can mutate state
    - Can provide additional funcitonality (fields/methods)

## Summary

- Mementos are used to roll back states arbitrarily
- A memento is simply a token/handle with (typically) no methods of its own
- A memento is not require to expose directly the state(s) to which it reverts the system
- Can be used to implement undo/redo
