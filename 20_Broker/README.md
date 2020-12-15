### Start Order
1. Broker
2. Consumer
3. Producer

### Testing
Broker needs to be restarted after each test. No persistance in modules.

### Notes
1. Suggesting LiteIDE for running tests. Broker was tested with one consumer and one producer.
2. Producer does not have REST interface.
3. For produce, check status of request does not work.
4. Suggested flow is frontend (missing) sends to API Gateway (missing) which passes to Producer. 
Producer checks and if request ready sends back by web sockets the links (composing links missing). 
5. Images stored on NFS share.
