##1 wallet
1. parse key from JSON string✅
2. save key's JSON string data

##2 node init
1. create an account ✅
2. init config ✅
3. change config✅
4. use command line to change config dynamically✅
##3 P2pNetwork
1. join a topic
   - 1.1 join success✅
   - 1.2 every topic has at least one peer node✅
2. subscribe topic✅
3. startup node✅
4. all work thread start success 
    - 4.1 debug topic thread✅
    - 4.2 online-offline topic thread✅
    - 4.3 message topic thread✅
    - 4.4 unread message topic thread✅
    - 4.5 contact operation topic thread✅
    - 4.6 websocket: dispatch thread for the client message✅
    - 4.7 websocket: websocket service thread✅
    - 4.8 contact http service thread✅
5. thread exit gracefully ✅
6. find peers for a public node  ✅
 ```       
    ninja.mac debug peers
```
7. find peers for a private node ✅
```
    192.168.30.17
    192.168.30.214
```
8. find peers both in a public and private network ✅ 
9. send debug peer message ✅
```
   ninja.mac debug push -m "hello"
```
10. sync online map when setup 
      - 10.1 one node setup -> user online✅
      - 10.2 two nodes setup
         + one node setup -> user online ->second node setup✅
         + one node setup ->second node setup -> user online✅

11. check peers of all topics 
      - 12.1 one node setup✅
      - 12.2 more than one node setup✅
```
     ninja.mac debug peers -t /0.1/Global/user/on_offline
     ninja.mac debug peers -t /0.1/Global/message/immediate
     ninja.mac debug peers -t /0.1/Global/message/unread
     ninja.mac debug peers -t /0.1/Global/contact/operate
```   
12. only one node setup, check the thread syncing and waiting
    - 13.1 boot node✅
    - 13.2 normal node✅
    
##4 message
1. the immediate message when both client is online
        
    setup 2 clients : 0 for account 0  2. 1 for account 1, open it in a new cmd window
```
    cd cli_lib/cmd
    go build .
    ./cmd 0
    ./cmd 1
```
    
- 1.1 both 2 clients on the same node
    - 1.1.1 message encrypt
    - 1.1.2 message decrypt 
    - 1.1.3 filter by contact
    
- 1.2 the 2 clients on different node
    - 1.2.1 message encrypt
    - 1.2.2 message decrypt
    - 1.2.3 filter by contact
    
- 1.3 the 2 clients on the different node and one node is in private network
  - 1.3.1 message encrypt
  - 1.3.2 message decrypt
  - 1.3.3 filter by contact
    
2. unread message
3. ping pong status✅
4. connect to localhost node
5. connect to the private network node
6. connect to the public network node

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

##6 online offline
```
ninja.mac debug ws -o
ninja.mac debug ws -l
```
1. online
    - 1.1 write to the local user table✅
    - 1.2 write to the online map✅
    - 1.3 publish to all peer node✅
    - 1.4 peer node add user address to the online map✅
    - 1.5 start the reading thread
    - 1.6 start writing thread

2. offline when client closed
    - 2.1 offline from local user table✅
    - 2.2 offline from online map✅
    - 2.3 publish to all peer node✅
    - 2.4 offline from all peer node✅
    - 2.5 reading thread exit
    - 2.6 writing thread exit

3. offline when node crash✅
3. online user data from cmd line debug tools✅

##7 refactor all proto message
##8 remove all warnings and typos