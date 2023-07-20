# Observer

## Motivation

- We need to be informed when certain things happen
    - Object's field changes
    - Object does something
    - Some external event occurs
- We want to listen to events and notified when they occur
- Two participants: _observable_ and _observer_

The Observer design pattern is where an _observer_ is an object that wishes to be informed about events happening in the system. The entity generating the events is an _observable_.

## Summary

- Observer is an intrusive approach
- Must provide a way of clients to subscribe
- Event data sent from observable to all observers
- Data represented as _interface{}_
- Unsubscription is possible