build: main.go
	@echo "Building..."
	go build -o bin/midnight main.go

clean:
	@echo "Cleaning..."
	rm -rf bin

# disallow any parallelism (-j) for Make. This is necessary since some
# commands during the build process create temporary files that collide
# under parallel conditions.
.NOTPARALLEL:

# indicates that the target is not a file and should always be executed.
.PHONY: all build clean