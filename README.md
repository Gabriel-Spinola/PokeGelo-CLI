# PokeGelo-CLI
PokeGelo-CLI is a command-line interface (CLI) tool designed to simplify sending HTTP requests, making it easier and more intuitive compared to using curl. With PokeGelo-CLI, you can send HTTP requests by providing a JSON file containing the request and its body. This tool is especially useful for testing APIs, debugging, and automating HTTP requests. Inspired by [Quaezo](https://github.com/zahash/quaeso/tree/main)

## Features
- Send HTTP requests using a JSON file.
- Simplified syntax compared to curl.
- Supports various HTTP methods: GET, POST, PUT, PATCH, DELETE, etc.
- Unit testing implementation ensures the reliability of the tool.
- Easy-to-use command: `poke net send -f <path to the file>`
- Docker support for seamless containerization (coming soon).

## Installation
To install PokeGelo-CLI, follow these steps:

Clone the Repository:

```sh
git clone https://github.com/your-username/PokeGelo-CLI.git
cd PokeGelo-CLI
```
Build the Executable:

```sh
go build -o poke
```
Move the Executable to a Directory in Your $PATH:

```sh
sudo mv poke /usr/local/bin
```
Now, you can use poke from the command line.

## Usage
To send an HTTP request, create a JSON file specifying the request method, URL, headers, and body. For example:

```json
{
  "method": "POST",
  "url": "https://api.example.com/resource",
  "headers": {
    "Authorization": "Bearer YOUR_ACCESS_TOKEN"
  },
  "body": {
    "key": "value"
  }
}
```

Then, use the following command to send the request:

```sh
poke net send -f /path/to/your/request.json
```

This will send the specified HTTP request using the information provided in the JSON file.

```sh
poke net send -w -f /path/to/your/request.json
```
This will send the specified HTTP request using the information provided in the JSON file and write the response to the output folder.

## Unit Testing
PokeGelo-CLI includes comprehensive unit tests to ensure its functionality and reliability. You can run the tests using the following command:

```sh
cd poke/tests
go test
```

## Future Plans
- Docker Support: Dockerize PokeGelo-CLI for easy deployment and management.
- Interactive Mode: Implement an interactive mode for building requests directly from the command line.
- Better request bulding and other utilities.

## Contribution Guidelines
Contributions to PokeGelo-CLI are welcome! If you have any ideas for improvements, bug fixes, or new features, feel free to open an issue or submit a pull request.

## License
PokeGelo-CLI is licensed under the GPL 3.0 License. Feel free to use, modify, and distribute this tool as needed.

Thank you for using PokeGelo-CLI! If you have any questions or feedback, please don't hesitate to reach out. Happy coding!

### TODO:
 - [ ] Build project
 - [X] API requests like postman
 - [X] Write first test
 - [ ] Write possible request
 - [ ] Document
 - [ ] Dockerize
