cp ../frontend/dist/ -Rv ./node_modules/filebrowser-frontend/
rice embed
go run cmd/filebrowser/main.go --file-scan-interval 60 --index-scan-interval 60 --scope $1 -p 8080