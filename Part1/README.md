# Mutex and Channel basics

### What is an atomic operation?
> atomic operations are operations which cannot be stored midway. Meaning there can be no Contect switch during the operation; it cannot be interrupted.

### What is a semaphore?
> A semaphore is a variable used to keep control of access to common resources in concurrent programming.

### What is a mutex?
> *mutex, short for mutual exclusion, is also a tool for avoiding race conditions. Mutex is a locking mechanism and an object while semaphore is a signalling mechanism and an integer variable. 

### What is the difference between a mutex and a binary semaphore?
> Semaphore allows multiple threads to access a finite instance of resources, while Mutex allows multiple threads to access a single resource, but not simultaneously.

### What is a critical section?
> a critical section is a segment of code that accesses shared resources, and thus is prone to race conditions and hass to be executed as an atomic operation.

### What is the difference between race conditions and data races?
 > A race condition is when the timing of operations or ordering of events affects a programs corectness. Data races are when twow untimed threads access the same place in memory and we have undefined behaviour. 

### List some advantages of using message passing over lock-based synchronization primitives.
> upon using google, we found that message passing is generally belived to be easier to get right, and will bet against the programers capabilities. The programer cannot implement locks poorly if he/she is not granted access to the locks. Also, message passing, as the name suggests, allows for data transfer between threads. 

### List some advantages of using lock-based synchronization primitives over message passing.
> Often grants the programmer with morer freedom. It also appear to us that lock-based tools are more fit for when several threads have the need to access the same piece of information. 
