What are buffered channels?
All the channels we discussed in the previous tutorial were basically unbuffered. As we discussed in the channels tutorial in detail, sends and receives to an unbuffered channel are blocking.

It is possible to create a channel with a buffer. Sends to a buffered channel are blocked only when the buffer is full. Similarly receives from a buffered channel are blocked only when the buffer is empty.

Buffered channels can be created by passing an additional capacity parameter to the make function which specifies the size of the buffer.

ch := make(chan type, capacity)  
capacity in the above syntax should be greater than 0 for a channel to have a buffer. The capacity for an unbuffered channel is 0 by default and hence we omitted the capacity parameter while creating channels in the previous tutorial.

Lets write some code and create a buffered channel.