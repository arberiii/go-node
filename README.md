# peer
peer package facilitates usage of Listening UDP.
It was intended for Distributed Bloom Filter.
StartServer method takes two functions as arguments: handle is function
that handles messages, while periodicTask is placeholder if you have 
any task to do periodically(this was intended for Dist. BF) 
