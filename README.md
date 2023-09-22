# About
Learning project for learning to use kafka with golang in highload examples.

# Case


## Architecture
1. Send POST request `/receipts` to master to begin generating 10m random transactions;
2. Sending generated transactions by batches to different partitions;
3. Each consumer receive messages from its individual partition in group;
4. Process transactions: 

    if transaction.Location == "LPP" || transaction.Amount > 100000 || transaction.TransactionType.String() == "Undefined" || transaction.TransactionType.String() == "RecurringPayment" 
    
    -->
    
     transaction marked as cancelled
5. Insert batches of processed transactions to DB.

 ![](/docs/architecture.jpg)

# Stack
* Golang 1.20.4
    * gin-gonic/gin 
    * go-gorm/gorm
    * sigmentio/kafka-go
    * ariga.io/atlas-provider-gorm (for versioned db migrations)
* Docker-Compose
* Kafka Broker
* PostgreSQL

# TODO
- [x] adapt to mapreduce in consumer
- [x] union general code in external library