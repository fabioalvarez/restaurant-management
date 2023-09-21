db = db.getSiblingDB('restaurant');

db.createCollection('food');
db.createCollection('menu');
db.createCollection('table');
db.createCollection('order');