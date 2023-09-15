FROM golang:1.21 as dependencies

# Call all deps
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Build a new layer and compilate.
# Code changes won't require a new download, just a re-compile.
FROM dependencies as build
COPY . ./
# Build the object called main
RUN CGO_ENABLED=0 go build -o /main -ldflags="-w -s" .

# Move to a clean artifact, that just includes the exe.
FROM golang:1.21
COPY --from=build /main /main
CMD ["/main"]