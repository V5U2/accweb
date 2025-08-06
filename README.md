# Assetto Corsa Competizione Server Web Interface (Enhanced Auth Fork)

This is a fork of [accweb](https://github.com/assetto-corsa-web/accweb) with enhanced authentication features. The original project was created by the Assetto Corsa Web team.

[![Go Report Card](https://goreportcard.com/badge/github.com/V5U2/accweb)](https://goreportcard.com/report/github.com/V5U2/accweb)

ACCWeb lets you manage your Assetto Corsa Competizione servers via a nice and simple web interface. You can start, stop and configure server instances and monitor their status. This fork adds flexible authentication options including OAuth support.

## What's New in This Fork

* Flexible authentication modes:
  * Standard password-based authentication (original behavior)
  * OAuth2 authentication support (new)
  * No authentication mode for trusted environments
* Easy configuration through environment variables
* Full compatibility with the original version

## Table of contents

1. [Features](#features)
2. [Changelog](#changelog)
3. [Installation](#installation-and-configuration)
4. [Docker](#docker)
5. [Backup](#backup)
6. [Contribute and support](#contribute-and-support)
7. [Build release](#build-release)
8. [Links](#links)
9. [License](#license)


## Features

* create and manage as many server instances as you like
* configure your instances in browser
* start/stop instances and monitor their status
* view server logs
* copy server configurations
* import/export server configuration files
* delete server configurations
* flexible authentication options:
    * standard mode: three different permission levels (admin, mod, read-only) using passwords
    * OAuth mode: authentication through providers like GitHub or Google
    * no-auth mode: for trusted environments or when using external authentication
* easy setup
    * no database required
    * simple configuration using environment variables
    * Docker support with all authentication options
    
## Changelog
<a name="changelog" />

See [CHANGELOG.md](CHANGELOG.md).

## Installation and configuration

accweb is installed by extracting the zip on your server, modifing the YAML configuration file to your needs and starting it in a terminal.

### Authentication Configuration

The application supports three authentication modes that can be configured in `config.yml`:

1. **Standard Authentication** (default)
   ```yaml
   auth:
     mode: standard
     admin_password: your_admin_password
     moderator_password: your_mod_password
     read_only_password: your_readonly_password
     timeout: 5h
   ```

2. **OAuth Authentication**
   ```yaml
   auth:
     mode: oauth
     timeout: 5h
     oauth:
       provider: "github"  # or "google"
       client_id: "your_client_id"
       client_secret: "your_client_secret"
       callback_url: "https://your-domain/api/auth/oauth/callback"
   ```

3. **No Authentication**
   ```yaml
   auth:
     mode: none
   ```

### Manual Installation

1. download the latest release from the release section on GitHub
2. extract the zip file on your server
3. edit the `config.yml` to match your needs and configure authentication
4. open a terminal
5. change directory to the accweb installation location
6. start accweb using `./accweb` on Linux and `accweb.exe` on Windows
7. leave the terminal open (or start in background using screens on Linux for example)
8. visit the server IP/domain and port you've configured, for example: http://example.com:8080

I recommend to setup an SSL certificate, but that's out of scope for this instructions. You can enable a certificate inside the `config.yml`.

**Note that you have to install [wine](https://www.winehq.org/) if you're on Linux.**

## Docker 

This fork is available on Docker Hub at `v5u2/accweb`. You can configure authentication using environment variables:

```bash
# Standard authentication (default)
docker run -e ACCWEB_AUTH_MODE=standard \
  -e ACCWEB_ADMIN_PASSWORD=admin \
  -e ACCWEB_MOD_PASSWORD=mod \
  -e ACCWEB_RO_PASSWORD=readonly \
  v5u2/accweb

# OAuth authentication
docker run -e ACCWEB_AUTH_MODE=oauth \
  -e ACCWEB_OAUTH_PROVIDER=github \
  -e ACCWEB_OAUTH_CLIENT_ID=your_client_id \
  -e ACCWEB_OAUTH_CLIENT_SECRET=your_client_secret \
  -e ACCWEB_OAUTH_CALLBACK_URL=https://your-domain/api/auth/oauth/callback \
  v5u2/accweb

# No authentication
docker run -e ACCWEB_AUTH_MODE=none v5u2/accweb
```

The original version is still available at [accweb/accweb](https://hub.docker.com/r/accweb/accweb) if you prefer the standard authentication only.

## Backup

To backup your files, copy and save the `config` directory as well as the `config.yml`. The `config` directory can later be placed inside the new accweb version directory and you can adjust the new `config.yml` based on your old configuration (don't overwrite it, there meight be breaking changes).

## Contribute and support

If you like to contribute, have questions or suggestions you can open tickets and pull requests on GitHub.

All Go code must have been run through go fmt. The frontend and backend changes must be (manually) tested on your system. If you have issues running it locally open a ticket.

To run the accweb locally is really simple, make sure that the attribute `dev` is set to true in your `config.yml` file.

### Frontend development environment

Our current frontend was built using Vue.js and can be found inside `public` directory.

To run the watcher use the following command.

```shell
make run-dev-frontend
```
Then when you edit any js file, the watcher will detect and rebuild the js package.

### Backend development environment

ACCweb backend is running over golang and can be found inside `internal` directory.

Use the following command to run the backend on your terminal.

```shell
make run-dev-backend
```
Keep in mind that you need to restart the command for see the changes that you made in the code working (or not :zany_face:) 

### Visual Studio Code - Remote container

There is a pre-built development environment setup for ACCWeb for Visual Studio Code and Remote Containers. Please, check here how to setup and use: https://code.visualstudio.com/docs/remote/containers

## Build release

To build a release, execute the `build_release.sh` script (on Linux) or follow the steps inside the script. You need to pass the build version as the first parameter. Example:
To build a release, execute the `build_release.sh` script (on Linux) or follow the steps inside the script. You need to pass the build version as the first parameter. Example:

```shell
./build/build_release.sh 1.2.3
```

This will create a directory `releases/accweb_1.2.3` containing the release build of accweb. This directory can be zipped, uploaded to GitHub and deployed on a server.

## Links

* [Docker Hub](https://cloud.docker.com/repository/docker/kugel/accweb/general)
* [Assetto Corsa Forums](https://www.assettocorsa.net/forum/index.php?threads/release-accweb-assetto-corsa-competizione-server-management-tool-via-web-interface.57572/)

## License

MIT
