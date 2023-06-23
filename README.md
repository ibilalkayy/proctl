# Proctl

Proctl, a project management software that utilizes the functionality of Monday.com in a command line interface, has several strong points that make it a valuable tool for managing projects. Some of these strengths include:

**Efficiency:** With a command-line interface, Proctl allows for quick and efficient task management without the need to navigate through a graphical user interface (GUI). This can save time and improve productivity for users who prefer a command line interface.

**Familiarity:** For users who are already comfortable working with command line interfaces, Proctl provides a familiar and intuitive way to manage tasks, making it easier to adopt as a new tool.

**Customization:** Proctl allows for customization and automation of tasks using scripting and other command-line tools, giving users more control over their workflow and allowing for greater flexibility in project management.

**Integration:** By using the functionality of Monday.com, Proctl can seamlessly integrate with other software tools and services, providing users with a comprehensive project management solution that can be tailored to their specific needs.

**Security:** Command line interfaces typically have fewer security vulnerabilities compared to GUI-based applications. Proctl provides an additional layer of security, ensuring that sensitive project information is kept secure and protected from unauthorized access.

## Overview

With Proctl, you can create workspaces, boards, projects and tasks, and arrange them according to your need. It is a simple and convenient way to manage your projects from the terminal.

## Usage

To run Proctl, simply run the following command:

`go run main.go`

**Note:** After signing up in the account, if you logout and login again, it will give you an error that the `proctl.Members` table does not exist.

Please make sure to create the Members table in MySQL database by copying this code.

    CREATE TABLE IF NOT EXISTS Members (
        id INT PRIMARY KEY AUTO_INCREMENT,
        emails VARCHAR(255) NOT NULL,
        passwords VARCHAR(255) NOT NULL,
        fullnames VARCHAR(255) NOT NULL,
        accountnames VARCHAR(255) NOT NULL,
        titles VARCHAR(255) NOT NULL,
        phones VARCHAR(255) NOT NULL,
        locations VARCHAR(255) NOT NULL,
        working_statuses VARCHAR(255) NOT NULL,
        is_active TINYINT NOT NULL,
        created_at DATETIME NOT NULL
    );

I will improve this code so that it is not used further.

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

- [Hashnode](https://hashnode.com/@ibilalkayy)

## License

- [Apache-2.0 license](https://raw.githubusercontent.com/ibilalkayy/proctl/master/LICENSE)
