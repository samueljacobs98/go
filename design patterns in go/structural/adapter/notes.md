# Adapter

Getting the interface you want from the interface you have. An adapter is a construct which adapts an existing interface X to conform to the required interface Y.

## Summary

 - Implementing an Adapter is easy
 - Determine the API you have and the API you need
 - Create a component which aggregates (has a pointer to, ...) the adaptee
 - Intermediate representations can pile up: use caching and other optimisations