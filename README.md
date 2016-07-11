InterPoller
-----------

I library to poll a host system's network interfaces and detect changes. (This is a bit of pet project to attempt to get better at writing clean library functionality in Go)

### Things I'm looking to learn well

* Bubbling errors up to the callers
* Informative logging
* (Less so, save for later project)Testing machine interfacing areas of code


# Features

After intialization return a event channel which will notify callers when changes to the Interfaces occur. 

## Internal data structures

Goal: Data structure to detect changes in network interfaces and addresses
* [ ] Build pipeline to create data structure representing state of network interfaces
* [ ] Hash of interfaces
* [ ] Hash of addresses

Example:

`cmd/expose/main.go` will build an example usage writing all interfaces and their corresponding IP addresses to stdout.


