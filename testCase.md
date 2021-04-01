##1 wallet
1. parse key from JSON string✅
2. save key's JSON string data

##2 node init
1. create an account ✅
2. init config ✅
3. change config✅

##3 P2pNetwork
1. join a topic✅
2. subscribe topic✅
3. startup node✅
4. all work thread start success ✅
5. thread exit gracefully ✅
6. find peers for a public node  ✅
7. find peers for a private node ✅

        192.168.30.17
        192.168.30.214

8. find peers both in a public and private network ✅ 
9. send debug peer message ✅
10. sync online map when setup 
      - 10.1 one node setup -> user online
      - 10.2 two nodes setup
         + one node setup -> user online ->second node setup
         + one node setup ->second node setup -> user online
11. sync contact data when setup
12. check peers of all topics 
      - 12.1 one node setup
      - 12.2 more than one node setup
13. only one node setup
    - 13.1 boot node
    - 13.2 normal node
    
##4 message
1. immediate message
   - 1.1 filter by contact
2. unread message
3. online
   - 3.1 write to the local user table
   - 3.2 write to the online map
   - 3.3 publish to all peer node
   - 3.4 peer node add user address to the online map
   - 3.5 start the reading thread
   - 3.6 start writing thread

4. connect to localhost node
5. connect to the private network node
6. connect to the public network node
7. offline
8. online user data cmd line debug tools✅
9. ping pong status

##5 contact
1. add contact
2. update contact
3. remove contact
4. sync contact
5. connect to a private network node
6. connect to a public network node
7. connect to localhost node
8. sync contact from other peers
9. connect to a service without my contact

##6 refactor all proto message