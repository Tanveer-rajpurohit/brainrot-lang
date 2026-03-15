# During development — just run directly
go run main.go run examples/hello.brt

# Build a single .exe / binary for your OS
go build -o brainrot main.go

# Run it like a real CLI tool
./brainrot run examples/hello.brt

# Build for Windows FROM Mac/Linux (cross compile — Go is amazing at this!)
GOOS=windows GOARCH=amd64 go build -o brainrot.exe main.go

# Build for Mac FROM Windows
GOOS=darwin GOARCH=amd64 go build -o brainrot main.go

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o brainrot main.go