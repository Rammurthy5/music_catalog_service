# This is designed using Go's Gin framework to serve & listen at 8080. 

 For simplicity sake, in-memory DB is used


## Endpoint GET   /albums 
   Fetches all available albums

## Endpoint 2: GET    /albums/<id> 
     fetches only single album if available

## Endpoint 3: POST   /albums 
   adds new albums


### Sample Request Payload 
`{
    "id": "4",
    "title": "The Modern Sound of Betty Carter",
    "artist": "Betty Carter",
    "price": 49.99
}`


### Sample Response Payload

`
[
    {
        "id": "1",
        "title": "Blue Train",
        "artist": "John Coltrane",
        "price": 56.99
    },
    {
        "id": "2",
        "title": "Jeru",
        "artist": "Gerry Mulligan",
        "price": 17.99
    },
    {
        "id": "3",
        "title": "Sarah Vaughan and Clifford Brown",
        "artist": "Sarah Vaughan",
        "price": 39.99
    }
]
`

## cURLs

`
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
`



