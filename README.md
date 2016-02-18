# book-api-go-pg

This is a REST API for books written in Go that connects to a Postgres database. The app is provisioned for Heroku deployment using GB
dependency management. 

##Get Books
Returns a complete list of every book in the database. The response is formatted as an array of JSON objects with each object representing
a separate book. 
###Url
`/books`
###Method
`GET`
###Response
```javascript
[{
  "id":1,
  "title":"Green Eggs and Ham",
  "author":"Dr. Seuss",
  "publisher":"Random House"
  },
{
  "id":2,
  "title":"Harry Potter and The Deathly Hallows",
  "author":"J.K. Rowling",
  "publisher":"Bloomsbury"
},
...
]
```
