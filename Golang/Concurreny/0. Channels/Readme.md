
# Channels
- A channel is like a unix pipe which allows multiple readers and writers
- Channels are typed, so that a channel of type chan T can only be used to transfer messages of type T. 
- Channels are uni-directional , but have two ends which can be passed seperately as params.
- Channels by **default are unbuffered**
- **It's is used both as method of synchronization as well as communication tool.**
- It's a vehicle of transfering ownership of data, so that only one goroutine at a time is writing the data(avoids racing)
- **Don't communicate by sharing memeory, instead share memory by communicating** ~ Rob pike

## Visualizing Buffered vs Non-Buffered channels
------------------------------------------------------------------
- NOTE: Visualizing unbufferd is when both parties come together they create a block to exchange
    write -> | BUFFER | <-reader
    For buffered chan of even size the buffer is presisted as long as the chan is NOT CLOSED.

## Select block
---------------------------------------------------
- A select block is used for listening for ALL the channels simulentously which ever is ready is read. If more than 1 channels are ready then any one is choosen at random.
    example:
        for{
            select{
                case <-chan1:
                    fmt.Println("")
                case <-chan2:
                    //do somthing else
            }
        }
- In a select block, the default case is always  ready and will be choosen if no other case is
- DON'T use default inside a loop- the select will busy wait and waste CPU cycles.
## Channels and Synchroization
-----------------------------------------------------
0. A unbuffered channel syncs reader and writer.
1. Unbuffered Channels block unless ready to read or write
    - In case of buffered channels it blocks if no space to write or nothing to read.
    - We can read from a closed channel

2. A channel is ready to write iff 
    - it has buffer space or
    - atleast 1 reader is ready to read
3. A channel is ready to read iff 
    - it has data to read in it's buffer or 
    - atleast 1 writer is ready to write or 
    - it is closed**

## Closed channels
-----------------------------------------------------
1. A channel can only be closed once(else it will panic)
2. Only 1 goroutine can close a channel
3. Closing a channel is often a signal that work is DONE. and is done ALWAYS by the writer of the channel.


## Channel Axioms (VERY IMPORTANT.)
-------------------------------------------
A send to a nil channel blocks forever
A receive from a nil channel blocks forever
A close of nil channel panics.

A send to a closed channel panics (Once a channel is closed, you can't send further values on it else it panics. This is what you experience.)
A receive from a closed channel returns the zero value immediately



## Channel State reference
-------------------------------------------
| State                 | Receiver     | Sender       | Close          |
| :------------------  | :-------    | :-------   | :-------     |
| `Nil`                 | Gets Blocked | Gets Blocked | Panic          |
| `Empty`               | Gets Blocked | Write        | close          |  
| `Partially_Empty`     | Read         | Write        | Read           |
| `Full`                | Read         | Gets Blocked | Panic          |
| `Closed`              | Default Value| Panic        | Panic          |  
| `Receiver-only`       | OK           | Compile Error| Compiler Error |
| `Send-only`           | Compile Error| OK           |  OK            |

- **Select ignores a nil channel since it would always block**
- Reading a Buffered closed channel returns <defualt_value, !ok>
- Reading a Unbufferd closed channel returns DEADLOCK**

## Rendezvous and Buffering
-------------------------------------------------
- By default, channels are unbuffered (rendezvous model)
- Both reader and writer have to come togther to perform the task, this way both get sync together, which ever comes before has to wait for the other one.
    - example: A courier guy and you have to come together for the signature of packet to recieve at the door.

### Buffering
- Buffering allows the sender nd receiver without waiting ie they are indepenet.
- the sender deposits the item and returns immediately

Common uses of buffered channels:
-  avoid goroutines leaks
-  avoid rendezvous pauses (performances imporvments)
- **NOTE:** Don't buffer until it's needed: buffering may hide a race condition.
- Buffered channels are used to implement **counting seamphore pattern**.


**NOTE** We don't need to limit goroutines but we need to limit contention for shared resources.
**Amdal's law**


