# Todo-REST-GO

This is a simple REST API for the [https://github.com/leplasmo/todo-api](https://github.com/leplasmo/todo-api) module.

It supports CRUD operations:
- [GET] `/todo` or `/todo/` -- get all todos
- [GET] `/todo/1` -- get todo with id 1
- [POST] `/todo` -- create a new todo -- requires JSON payload
- [PATCH] `/todo/1?complete` -- set todo with id 1 as completed
- [DELETE] `/todo/1` -- remove todo with id 1 


## Usage

Supports flags:
- `-p` : Port to run on
- `-h` : Local IP to bind on (i.e. 127.0.0.1, localhost, 0.0.0.0)
- `-f` : File path/name to wite/read the todos from/to (JSON encoded)


```bash
# start the server
go run . -p 8080

# query todos
curl -s localhost:8080/todo | jq
#OUTPUT:
'
{
  "results": [
    {
      "Task": "My first todo",
      "Done": false,
      "CreatedAt": "2022-04-20T16:36:05.957792+02:00",
      "CompletedAt": "0001-01-01T00:00:00Z"
    },
    {
      "Task": "My second todo",
      "Done": false,
      "CreatedAt": "2022-04-20T16:36:13.352601+02:00",
      "CompletedAt": "0001-01-01T00:00:00Z"
    }
  ],
  "date": 1650465378,
  "total_results": 2
}
'

# query todo with id 2
curl -s localhost:8080/todo/2 | jq
#OUTPUT:
'
{
  "results": [
    {
      "Task": "My second todo",
      "Done": false,
      "CreatedAt": "2022-04-20T16:36:13.352601+02:00",
      "CompletedAt": "0001-01-01T00:00:00Z"
    }
  ],
  "date": 1650465521,
  "total_results": 1
}
'

# create new todo
curl -XPOST -d '{"Task":"This is a new TODO"}' localhost:8080/todo

# mark todo as completed
curl -XPATCH localhost:8080/todo/1?complete 

# delete a todo
curl -XDELETE localhost:8080/todo/1
```
