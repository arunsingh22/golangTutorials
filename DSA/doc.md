

sync.NoCopy : It's a runtime primitive type which is used in all the sync packages 
- It prevents any struct form being copied, the runtime checks this automatically 
- sync.WaitGroup, sync.Mutex, sync.cond all have sync.NoCopy
- sync.Pool also uses it and therefore we cannot be copied.