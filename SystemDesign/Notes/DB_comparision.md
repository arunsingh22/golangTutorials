

# Availablity x DownTime(per year)
- 99%    	- 3.65days
- 99.9%  	- 8.77 hrs
- 99.99% 	- ~52 Min
- 99.999% 	- 5.26 Min**
- 99.9999   - 31.56 Seconds

# High level Design
- Mongo has WAL named syslog and MySql calls it binlog.
- 
- Oracle/MySql/Postgres do not support multi-master model which means when a write request it can only be written into the master/live DB node only. 

# Active Passive Architecture
- In a write heavy sys the master node becomes hot node
- If master node cluster goes down , the passive cluster would take some time to become primary so there will a lag observed and all writes will fail for the same time but reads will be successfull. 
- for all writes when both clusters are up also if the write gets forwarded to cluster 2 it would be sent to cluster 1 DB which will incure more latency. 

# Active-Active Architecture
- Casandra support multiple master nodes
- Conflict resolution is the most complex part in this
- Can handle far more traffic than active-passive arch





| DBMS        | Trade-off | Notes                       
|-------------|-----------|-----------------------------
| PostgreSQL  | CP        | Traditional RDBMS prioritizing strong consistency; 
                            high availability via read replicas but not partition tolerant by default.     
                            - Strong support to ACID properties
                            - ❌ **Multi-master**: No native support to multi-master, some 3rd party tools like Citus enable it
                            - ⚠️ **Sharding**: Limited native support (PG 16+ improving).
                            
| MySQL       | CP        | Standard MySQL favors consistency. Galera Cluster can support 
                            multi-master with trade-offs (can be configured as CA or CP).
                            - ❌ **Sharding**: No native support. Use **Vitess** (developed by youtube) which used by twitter and youtube or do a application-layer sharding.

| MongoDB|CP,tunable to AP| - Tunable consistency; can prioritize availability with eventual consistency
                            - ✅ Inbuilt support for sharding/ horizontal scaling
                            - ❌ **Multi-master**: No native support, one primary and others are replica nodes.

| Cassandra   | AP        | Prioritizes high availability and partition tolerance; consistency
                            is tunable (eventual or quorum).
                          - ✅ Inbuilt support for sharding/ horizontal scaling
                          - ✅ True multi-master (any node can accept writes); conflict resolution handled via timestamps
                          - Doesn't has any Leader , each node is treated as a master




Q: Why does RDBMS do not support sharding natively or struggle so much with sharding
# 1. Atomicity (A) – All-or-nothing transactions
Problem in sharding: In a distributed transaction that touches multiple shards (e.g., a money transfer  between two accounts in different shards), maintaining atomicity means either both commit or both roll back.

This requires two-phase commit (2PC) or distributed transactions, which are:
- Slow
- Error-prone
-Difficult to scale

- RDBMSs like MySQL/Postgres don’t natively support efficient 2PC across multiple nodes.
Example: A banking system transferring money between two users located in different shards must ensure both debiting and crediting succeed together — that's atomicity. Ensuring this across shards is non-trivial.

# 2. Isolation (Concurrent transactions don't interfere)
Problem in Sharding:
Isolation ensures that transactions appear to run one at a time even when they are concurrent.
This usually involves locking rows or tables or multi-version concurrency control (MVCC).

Why it's hard:
In a sharded system, ensuring isolation across shards means:
Coordinating locks across distributed systems.
Possibly deadlocks between transactions spanning different shards.
Strong isolation (like serializability) across shards is very expensive and degrades performance.

# 3. Consistency (C) Why Consistency Is Hard in Horizontal Scaling
1. Distributed Nature = Network Latency + Partition Risks.
  - When data is sharded, different pieces live on different machines.
  - To maintain consistent state across these, you need network calls and coordination. But networks are slow and can fail → introduces chances for inconsistency (e.g., partial writes).

**Example:**
You update a product's stock count and price—one in Shard A and one in Shard B.
If Shard B update fails mid-way (network glitch), your system now has an inconsistent view.

2.**Cross-Shard Referential Integrity Is Expensive**
## Relational databases enforce:
- Foreign keys
- Unique constraints
- Check constraints
These are trivial in a single-node system, but in a distributed system:

Foreign key order.user_id must validate user.id across shards—needs cross-node query.
Unique email check must scan all shards.
This violates the performance and isolation expected from scalable systems.

3. Global Transactions = Coordination Overhead
Maintaining a consistent view often involves:
Two-Phase Commit (2PC): To commit or abort transactions across shards.
This introduces:
- High latency
- Increased failure risk
- Blocking locks across nodes
Many RDBMS systems intentionally avoid supporting 2PC natively for this reason.

4. Eventual vs Strong Consistency Trade-offs
NoSQL systems like Cassandra or DynamoDB sacrifice consistency to gain scalability:

Use eventual consistency: updates will propagate over time, but not instantly.

Use quorum reads/writes to balance availability vs consistency.

Relational systems aim for strong consistency, making them:
Safer for financial or mission-critical apps
But harder to scale horizontally without major architectural changes

5. High Contention Workloads Break Down
In write-heavy apps with concurrent access:
Maintaining serializable isolation or even repeatable reads becomes hard when writes are spread across nodes.
Coordination overhead skyrockets (MVCC metadata sync, locking,