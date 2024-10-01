
# Challenge Lambda Function in Go

This project contains a simple AWS Lambda function written in Go.

## Prerequisites

- Go installed on your system (version 1.13+)
- AWS CLI configured with appropriate permissions

## Build & Package the Application

### 1. Download Dependencies

First, initialize the Go module and download the necessary dependencies:

```bash
go mod init challenge
go mod tidy
```

This will create the `go.mod` and `go.sum` files and download all the dependencies.

### 2. Build the Application

Create the distribution directory and build the application for a Linux environment with ARM64 architecture (required for AWS Lambda):

```bash
mkdir -p dist
env GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -tags lambda.norpc -ldflags="-s -w" -o ./dist/bootstrap main.go
```

- `GOOS=linux`: Ensures the binary is built for a Linux environment.
- `GOARCH=arm64`: Specifies the ARM64 architecture (you can change to `amd64` if needed).
- `CGO_ENABLED=0`: Disables CGO to ensure the build is compatible with Lambda's environment.
- `-tags lambda.norpc`: Optional flag to avoid including RPC code.
- `-ldflags="-s -w"`: Minimizes binary size by stripping debug information.

### 3. Change Execution Permissions

Make the output binary (`bootstrap`) executable:

```bash
chmod +x ./dist/bootstrap
```

### 4. Package the Application

AWS Lambda only accepts zipped files for uploading. Package the binary into a zip file:

```bash
cd ./dist && zip -r bootstrap.zip bootstrap ../templates && cd ..
```

The output should show something like:

```bash
updating: bootstrap (deflated 69%)
updating: ../templates/ (stored 0%)
updating: ../templates/email_template.html (deflated 69%)
updating: ../templates/email_template.txt (deflated 45%)
```

This will create a `bootstrap.zip` file in the `dist` directory, ready to be uploaded to AWS Lambda.

### 5. Upload to AWS Lambda



## Conclusion

Now your Go-based Lambda function is built, packaged, and ready for deployment to AWS Lambda.
