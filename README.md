# book-api-go-pg

This is a REST API for books written in Go that connects to a Postgres database. The app is provisioned for Heroku deployment using GB
dependency management. A running instance of this API can be found here: [http://rest-api-go-sql.herokuapp.com/api/v1/books](http://rest-api-go-sql.herokuapp.com/api/v1/books)



## Get All Books
Returns a complete list of every book in the database. The response is formatted as an array of JSON objects with each object representing a separate book. 
#### Url
`/api/v1/books`
#### Method
`GET`
#### Response
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



## Get One Book
Returns a JSON object that represents a book in the database. 
#### Url
`/api/v1/book/:id`
#### Method
`GET`
#### Response
```javascript
{
  "id":7,
  "title":"It",
  "author":"Stephen King",
  "publisher":"Viking Press"
}
```



## Get Book Count
Returns the total number of books in the database. 
#### Url
`/api/v1/count`
#### Method
`GET`
#### Response
```javascript
7
```



## Add a Book
Adds a new book to the database. 
#### Url
`/api/v1/books`
#### Method
`POST`
#### Parameters
| Key           | Value          |
| ------------- |:--------------:|
| Title         | text(required) |
| Author        | text(required) |
| Publisher     | text(required) |
#### Response
```javascript
{
  "success":"A book has been added to the database!"
}
```



## Update Book
Updates the information associated with a given book.
#### Url
`/api/v1/books/:id`
#### Method
`PUT`
#### Parameters
| Key           | Value          |
| ------------- |:--------------:|
| Title         | text(required) |
| Author        | text(required) |
| Publisher     | text(required) |
#### Response
```javascript
{
  "id":2,
  "title":"Harry Potter and The Prisoner of Azkaban",
  "author":"J.K. Rowling",
  "publisher":"Bloomsbury"
}
```



## Delete Book
Removes a book entry from the database. 
#### Url
`/api/v1/book/:id`
#### Method
`GET`
#### Response
```javascript
{
  "id":1,
  "title":"Green Eggs and Ham",
  "author":"Dr. Seuss",
  "publisher":"Random House"
}
```
