# Bridge

## Motivation

- Bridge prevents a 'Cartesian product' complexity explosion
- Example:
    - Common type ThreadScheduler
    - Can be preemptive or cooperative
    - Can run on Windows or Unix
    - End up with a 2x2 scenario: WindowsPTS, UnixPTS, WindowsCTS, UnixCTS

## Summary

- Decouple abstraction from implementation
- Both can exist as hierarchies