# Golang-application-cli
A CLI in golang to interact with a single or multiple applications in Golang.

## Prerequisites

- Golang 1.16 or higher

To see the version of Golang installed on your machine, run the following command.

```bash
go version
```

If you don't have Golang installed on your machine, you can download it from the [official website](https://golang.org/dl/).

Or using apt or brew.

```bash
sudo apt search golang-go
sudo apt install golang-go
```

```bash
brew install go
```

## Installation

Clone the repository and run the following command at the root of the CLI.

```bash
go install
echo export GAC=$(pwd) >> ~/.zshrc
```

```bash
go build -o ./bin/go-app-cli
```

## Usage

### Create a new application

Create a new application using the following command.

```bash
./bin/go-app-cli init --name <YOUR_APP_NAME>
```

### List applications

List all the applications using the following command.

```bash
./bin/go-app-cli list
```

List a specific application using the following command.

```bash
./bin/go-app-cli list --name <YOUR_APP_NAME>
```

### Remove applications

Delete a specific application from the storage using the following command.

```bash
./bin/go-app-cli remove --name <YOUR_APP_NAME>
```

Delete completely the applications using the following command.

```bash
./bin/go-app-cli remove --remove-app --name <YOUR_APP_NAME>
```

### Update applications

Update the go version of a specific application using the following command.

```bash
./bin/go-app-cli update --name <YOUR_APP_NAME> --version <YOUR_GO_VERSION>
```

Update the go version of all the applications using the following command.

```bash
./bin/go-app-cli update --all --version <YOUR_GO_VERSION>
```

## Project Structure

```bash
.
├── bin
├── cmd
├── configs
├── storage
├── templates
├── .gitignore
├── go.mod
├── go.sum
├── main.go
└── README.md
```

1. **bin**: Contains the binary file of the CLI.
2. **cmd**: Contains the commands for the CLI.
3. **configs**: Contains the configuration files for application structure that the CLI generate.
4. **storage**: Contains the storage files for the CLI.
5. **templates**: Contains the templates for main.go that the CLI generate.
