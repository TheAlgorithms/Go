# Linked List
***
## What is it?
***
[Linked list](https://www.scaler.com/topics/linked-list/) is a data structure that is a linear chain of data elements whose order is not given by their phyisical placement in memory. This structure is built up of nodes which point to the element ahead or behind this particular node (depending on the type of Linked List).

# Types of Linked List implemented within this repository

## Singly Linked List
Singly Linked List is a linked list in which a node only points to the next element.
Some of the main applications of Singly Linked List are in the construction
of fundamental data structures such as Stacks and Queues.

## Doubly Linked List
Doubly Linked List is a linked list in which a node points to both the element ahead and behind the node.
Any application which requires us to keep track of forward and backward information uses doubly linked list.
For example, the feature of undo and redo are implemented using these doubly linked lists.

## Cyclic Linked List AKA Looped Linked List
Looped Linked Lists are singly or doubly-linked that chase their own tail:
A points to B, B points to C, C points to D, and D points to A. 
They are better suited for cyclic data such as train schedules.
These lists are missing the first and last items.
Therefore, it is necessary to introduce the concept of the current position.

This picture shows similar lists:
![Alt text](./Linked_Cyclic_List.jpg?raw=true)
