TODO_BINARY=todoApp


## build_todo: builds the todo binary as a linux executable
build_todo:
	@echo "Building todo binary..."
	cd ../Todo && env GOOS=linux CGO_ENABLED=0 go build -o ${TODO_BINARY} ./cmd
	@echo "Done!"

