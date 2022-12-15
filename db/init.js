db = db.getSiblingDB('account');
db.createCollection('transactions');
db.createCollection('accounts');
db.accounts.insertOne({
    "_id": "52814c2d-657b-4e7b-be5c-9f28e59253f8",
    "document_number": "44639723024",
    "created_at": ISODate("2022-12-15T15:37:21.471-03:00")
});

db.accounts.insertOne({
    "_id": "74211f47-9c6f-4648-88a3-cb7d0614b5fe",
    "document_number": "67909198051",
    "created_at": ISODate("2022-12-15T15:37:21.471-03:00")
});

db.accounts.insertOne({
    "_id": "355daea3-bfdc-41d5-8ecf-c9bcd21f4dbf",
    "document_number": "51180817001",
    "created_at": ISODate("2022-12-15T15:37:21.471-03:00")
});

db.getCollection("transactions").createIndex({ "account_id": 1});
db.getCollection("transactions").createIndex({ "description": 1});
db.getCollection("transactions").createIndex({ "event_date": 1});
db.getCollection("transactions").createIndex({ "operation_type_id": 1});
db.getCollection("transactions").createIndex({ "amount": 1});