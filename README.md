# Garçon
**Garçon** is a simple static HTTP server. Intended as a way to learn Go, useful when doing frontend development.

## Installation

    $ go install github.com/alexei/garcon

## Usage

    $ garcon

### Arguments
`bind` specifies the address to bind to. Defaults to `127.0.0.1:8080`:

    $ garcon -bind 127.0.0.1:9000

The `path` argument specifies the root directory (defaults to current directory):

    $ garcon -path ~/www/

You can also specify a location prefix using the `prefix` argument:

    $ garcon -prefix /static/

`log` indicates the location to log to (defaults to stdout):

    $ garcon -log /var/log/garcon/access.log

## License

**Garçon** is licensed under the terms of the 3-clause BSD license.
