InterPoller
-----------

I library to poll a host system's network interfaces and detect changes. 

# Features

After intialization return a event channel which will notify callers when changes to the Interfaces occur. 

Example:

`cmd/expose/main.go` will build an example usage writing all interfaces and their corresponding IP addresses to stdout.

