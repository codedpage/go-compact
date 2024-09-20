Mutex vs Channels
We have solved the race condition problem using both mutexes and channels. So how do we decide what to use when. The answer lies in the problem you are trying to solve. If the problem you are trying to solve is a better fit for mutexes then go ahead and use mutex. Do not hesitate to use mutex if needed. If the problem seems to be a better fit for channels, then use it :).

Most Go newbies try to solve every concurrency problem using a channel as it is a cool feature of the language. This is wrong. The language gives us the option to either use Mutex or Channel and there is no wrong in choosing either.

###########################
In general use channels when Goroutines need to communicate with each other and mutexes when only one Goroutine should access the critical section of code.
###########################

In the case of the problem which we solved above, I would prefer to use mutex since this problem does not require any communication between the goroutines. Hence mutex would be a natural fit.

My advice would be to choose the tool for the problem and do not try to fit the problem for the tool :)

This brings us to an end of this tutorial. Have a great day.