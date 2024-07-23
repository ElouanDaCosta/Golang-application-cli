# Golang-microservice-cli
A CLI in golang to interact with a single or multiple applications in Golang.

## Installation

Clone the repository and run the following command to install the CLI.

```bash
go install
```

```bash
go build -o <APP_NAME>
```

## Usage

Create a new application using the following command.

```bash
./<APP_NAME> init --name <YOUR_APP_NAME>
```

List all the applications using the following command.

```bash
./<APP_NAME> list
```

List a specific application using the following command.

```bash
./<APP_NAME> list --name <YOUR_APP_NAME>
```