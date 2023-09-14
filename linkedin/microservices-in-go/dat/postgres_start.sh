docker run -d --rm \
    --name local-pg \
    -e POSTGRES_PASSWORD=postgres \
    -p 5432:5432 \
    postgres

# Use `lsof -i :5432` to `kill` whatever process is bound to the port