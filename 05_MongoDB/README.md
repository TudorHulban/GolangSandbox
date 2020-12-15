### Docker Spin
```
docker run --name c-mongo --restart=always -d -p 27017:27017 mongo mongod
```

### Installation
```bash
# mx linux
sudo apt install mongodb-org
sudo mongod --dbpath mongodata/
```

### Operations
```bash
# enter shell
mongo
db.version()
```
Find records:
```bash
db.persons.find({Name:"mary"})
```
Delete records:
```bash
db.persons.deleteMany({Name:"mary"})
```
### Resources
```html
http://learningprogramming.net/golang/golang-and-mongodb/create-new-document-in-golang-and-mongodb/
https://kb.objectrocket.com/mongo-db/how-to-construct-mongodb-queries-from-a-string-using-golang-551
```
