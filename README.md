# Producer-Consumer problem 
There are two types of processes: "producer" and "consumer". Producers continuously produce a piece of information while consumers continuously consume. Communication occurs through a bounded buffer. Producer places a piece of information into the buffer. If there is no space, producer process blocks. Consumer takes a piece of information from the buffer. If a buffer is empty, consumer process blocks. No assumptions about the rate of production or consumption should be made. Access to the buffer is exclusive, only one process can produce or consume at a given time.

## Solution
Channel that is used as a buffer does not need any additional synchronisation. There are 4 basic cases to consider:
1. Channel is not full - producer puts item into the channel
2. Channel is full - producer process blocks
3. Channel is not empty - consumer takes item from the channel
4. Channels is empty - consumer process blocks

Producer-consumer problem defined by Dijkstra assumes that buffer has some capacity. That is why buffered channel is used as a buffer for the items. Sender can put value into the channel without having a corresponding receiver and vice versa. Capacity of the channel limits the number of such operations. When channel is full and there are senders and receivers waiting, it acts as a non-buffered channel and enforces synchronisation in communication.

## Wikipedia 
https://en.wikipedia.org/wiki/Producer%E2%80%93consumer_problem
