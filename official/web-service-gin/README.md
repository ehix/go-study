# web-service-gin

`go run .` 

GET /albums
```
curl http://localhost:8080/albums \
    --header \
    "Content-Type: application/json" \
    --request "GET"
```

GET /albums/:id
```
curl http://localhost:8080/albums/2
```

POST /albums
```
curl http://localhost:8080/albums \
    --include --header \
    "Content-Type: application/json" \
    --request "POST" --data \
    '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```