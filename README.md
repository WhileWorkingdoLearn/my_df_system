This is a Decentralized File System.

Purpose of this Project is to Build a DFS from Scratch (purely for educational reasons).

1. network infrastructure
    Peer-to-peer (P2P) communication: Each node should be able to communicate with others. For this reason the implementation of a custom application layer protocol will be provided. 
    Node discovery: How do the nodes find each other in the network? 
    An implementation could run via a distributed hash table (DHT).

2. file storage and distribution

    Chunking: Large files should be chunked into smaller blocks. How do you decide how big a block should be?
    Replication: To prevent data loss, files should be stored multiple times. How do you decide where to store the replicas?
    Encryption: Who should have access to the stored files? Does client-side encryption make sense?

3. consistency and file management

    Metadata management: Where and how do you store information about the files (e.g. which node owns which file)?
    Consistency mechanisms: If nodes fail, how is it ensured that the files are not lost or corrupted?
    Versioning: Should file versioning be implemented?

4. access and API

    User authentication: Who is authorized to upload, download or delete files?
    REST or gRPC API: Which API technology do you want to use? Should users interact via HTTP or via a special protocol?
    Indexing/Search: How can a user quickly find a file on the network?


