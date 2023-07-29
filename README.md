# harvest

A framework for distributed web scraping

## Setup

This project assumes that Go 1.16 is used. It also requires several additional Go tools that are not included in the project packages. For code linting and formatting you will need [golangci-lint](https://github.com/golangci/golangci-linthttps://github.com/golangci/golangci-lint), and for generating swagger docs you will need [swag](https://github.com/swaggo/swag)

```bash
# Download golangci-lint for Mac using Brew
brew install golangci-lint
brew upgrade golangci-lint

# Download golangci-lint for Linux or Windows with curl
# (binary will be $(go env GOPATH)/bin/golangci-lint)
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.38.0

# Download golangci-lint with go get (this is not guaranteed to work)
go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.38.0

# Check that it installed successfully
golangci-lint --version

# Download swag with go get
go get -u github.com/swaggo/swag/cmd/swag

# (Optional) Add an alias for the project taskfile
# (use the command below to save it for one session, or add the command to
# ~/.bashrc to save for future sessions)
alias task=./taskfile

# Install Go modules
task install
```

## Development

To run this project in development mode use the below command.

```bash
task dev
```

You should then be able to view the swagger documentation at [http://localhost:5000/api/doc/index.html](http://localhost:5000/api/doc/index.html).

## Testing

To run tests, linting, and code formatting, the below commands are used.

```bash
task test
task lint
task format
```

## Contributing

#### Adding Packages

If you need to add any new packages, simply use `go get` and it will automatically be saved in `go.mod` and `go.sum`. If you end up with unused packages, use `go mod tidy` to remove any packages not referenced in your code.

#### Pipeline

When you push a commit, the pipeline will first run tests and linting. If both steps pass, it will build a docker image using the Dockerfile and publish it to the repository's container registry. Find the most recent docker image for your feature branch by looking for the image tagged with `FEATURE_BRANCH-latest`. You can also look up the image for a specific commit by using that commit's SHA.

These images can then be deployed to a test environment. Only images made from the master branch (`master-latest`) should be used in production.

## Usage

The primary way this project should be used is via the docker images created from its pipeline. It can also be run in a production mode using the taskfile as shown below.

```bash
# (optional) Set the port via the PORT environment variable
export PORT=5000

task prod
```
