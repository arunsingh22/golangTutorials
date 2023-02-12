
- set key value 
- get key 
- strlen key 
- key * // displays all the keys present 
- flushall // deletes everything from the redis
- setex key sec value 
- ttl key 


- setnx key value  // sets key: value iff key is not present , if key is already defined it wont replace 
- mset key value key2 value2 key3 value3 ....

- incr key 
- decr key 
- incrby key value 
- decr key value 
- append key "appendvalue"