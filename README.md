# TX EDIT

Webapp using text editor, sqlite, templates... 

## Getting Started

```bash
# Running
go mod download

go run main.go
```

## Docker Commands

```bash
# Create image
docker build --tag=txedit:1 .

# Running container
docker run --name=txedit --restart=always -d -p 8080:8080 txedit:1
```
