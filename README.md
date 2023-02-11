# Proctl

Proctl is a project management software that is using the functionality of Monday.com in a command line interface.

## Overview

With Proctl, you can create workspaces, boards, projects and tasks, and arrange them according to your need. It is a simple and convenient way to manage your projects from the terminal.

## Usage

To run Proctl, simply run the following command:

`go run main.go`

## Configuration

Proctl uses a configuration file to store your settings. By default, the configuration file is located in `$HOME/.proctl.yaml`. You can also specify a different config file using the `--config` flag.

## Environment Variables

Proctl also supports reading environment variables. Any environment variables that match the keys in the config file will be automatically loaded.

## Persistent Flags

Proctl supports persistent flags that can be defined and used globally throughout your application. For example, the `--config` flag can be used to specify a config file.

## Local Flags

Proctl also supports local flags that will only run when the action is called directly. For example, the `--toggle` flag can be used to turn a feature on or off.

## Contributing

If you would like to contribute to Proctl, then here is a [contribution guide](CONTRIBUTING.md) that you can check and raise an issue.

## Feedback

We’d love to hear your thoughts on this project. Feel free to drop us a note!

- [Twitter](https://twitter.com/ibilalkayy)

- [LinkedIn](https://www.linkedin.com/in/ibilalkayy/)

## License

- [Apache-2.0 license](https://raw.githubusercontent.com/ibilalkayy/proctl/master/LICENSE)
