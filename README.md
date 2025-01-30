# Initiator

A command-line tool that helps you initialize and scaffold new projects quickly. Initiator provides templates and configurations for various project types, saving you time and ensuring consistency across your projects.

## Features

- Multiple project templates (Go, Node.js with Typescript)
- Interactive project setup

## Installation

```bash
go install github.com/moabdelazem/initiator@latest
```

## Usage

```bash
initiator create [project-name] [--dir=/home/example]
```

Example:

```bash
initiator create my-awesome-project
```

```bash
initiator create my-awesome-project -d ~/personal
```

## Templates

Currently supported templates:

- go-api: RESTful API template with Go Echo
- go-plain: Plain Go Project
- node-express: Express.js web application

## Contributing

Contributions are always welcome!
