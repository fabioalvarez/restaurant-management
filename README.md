# restaurant-management

## Init the project

`go mod init restaurant-management`

### Docker compose build

`docker-compose up --build -d`

### Docker exec container

`docker exec -it mongo bash`

### Mongo commands

- Start mongo shell: `mongosh`
- Show DB: `show dbs`
- Select DB: `use <database>` `use restaurant`
- Show Collections: `show collections`
- Query collection: `db.<collection>.find()` `db.menu.find()`
- Query with filter: `db.food.find({foodId: "64d2a30b920197bb8bb3e1fe"})`

### About the Context Library

`https://golangbot.com/context-timeout-cancellation/`

`https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go`

### About defer

`https://www.educative.io/answers/what-is-the-defer-keyword-in-golang`

### About validations

`https://medium.com/tunaiku-tech/go-validator-v10-c7a4f1be37df`

### Bind validations
`https://medium.com/@maktoobgar/how-to-validate-api-inputs-in-gin-f2af4a3ce43e`

### Init Mongo db

`https://dev.to/mikelogaciuk/docker-initialize-custom-users-and-databases-in-mongodb-3dkb`

`https://gist.github.com/gbzarelli/c15b607d62fc98ae436564bf8129ea8e`

`https://cj-hewett.medium.com/quick-mongodb-docker-setup-d1959c8fc8f2`

`https://www.mongodb.com/docs/upcoming/reference/method/db.getSiblingDB/`

### Init mongo db example - init.db

```
set -e

mongo <<EOF
db = db.getSiblingDB('restaurant')

db.createCollection('food')
db.createCollection('menu')

db = db.getSiblingDB('warehouse')

db.createUser({
  user: 'warehouse',
  pwd: '$WAREHOUSE_PASSWORD',
  roles: [{ role: 'readWrite', db: 'warehouse' }],
});
db.createCollection('documents')
db.createCollection('stocks')
db.createCollection('invoices')
db.createCollection('orders')

EOF
```

### Problems/Solutions
#### Connection problem

`https://stackoverflow.com/questions/53887738/server-selection-timeout-error-mongodb-go-driver-with-docker`