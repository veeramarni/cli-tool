# Cdecode CLI [![Build Status](https://travis-ci.org/cdecode/src-cli.svg)](https://travis-ci.org/cdecode/src-cli) [![Build status](https://ci.appveyor.com/api/projects/status/fwa1bkd198hyim8a?svg=true)](https://ci.appveyor.com/project/cdecode/src-cli) [![Go Report Card](https://goreportcard.com/badge/cdecode/src-cli)](https://goreportcard.com/report/cdecode/src-cli)

The Cdecode `cli` CLI provides access to [Cdecode](https://cdecode.com) via a command-line interface.

![image](https://user-images.githubusercontent.com/3173176/43567326-3db5f31c-95e6-11e8-9e74-4c04079c01b0.png)

It currently provides the ability to:

- **Execute search queries** from the command line and get nice colorized output back (or JSON, optionally).
- **Execute GraphQL queries** against a Cdecode instance, and get JSON results back (`src api`).
  - You can provide your API access token via an environment variable or file on disk.
  - You can easily convert a `src api` command into a curl command with `src api -get-curl`.
- **Manage repositories, users, and organizations** using the `src repos`, `src users`, and `src orgs` commands.

If there is something you'd like to see Cdecode be able to do from the CLI, let us know! :)

## Installation

### Mac OS:

```bash
curl -L https://github.com/cdecode/cli-tool/releases/download/latest/src_darwin_amd64 -o /usr/local/bin/src
chmod +x /usr/local/bin/src
```

### Linux:

```bash
curl -L https://github.com/cdecode/cli-tool/releases/download/latest/src_linux_amd64 -o /usr/local/bin/src
chmod +x /usr/local/bin/src
```

### Windows:

Note: Windows support is still rough around the edges, but is available. If you encounter issues, please let us know by filing an issue :)

Run in PowerShell as administrator:

```powershell
New-Item -ItemType Directory 'C:\Program Files\Cdecode'
Invoke-WebRequest https://github.com/cdecode/cli-tool/releases/download/latest/src_windows_amd64.exe -OutFile 'C:\Program Files\Cdecode\src.exe'
[Environment]::SetEnvironmentVariable('Path', [Environment]::GetEnvironmentVariable('Path', [EnvironmentVariableTarget]::Machine) + ';C:\Program Files\Cdecode', [EnvironmentVariableTarget]::Machine)
$env:Path += ';C:\Program Files\Cdecode'
```

Or manually:

- [Download the latest src_windows_amd64.exe](https://github.com/cdecode/cli-tool/releases/download/latest/src_windows_amd64.exe) and rename to `src.exe`.
- Place the file under e.g. `C:\Program Files\Cdecode\src.exe`
- Add that directory to your system path to access it from any command prompt

## Usage

Consult `src -h` and `src api -h` for usage information.

## Authentication

Some Cdecode instances will be configured to require authentication. You can do so via the environment:

```sh
SRC_ACCESS_TOKEN="secret" src ...
```

Or via the configuration file (`~/src-config.json`):

```sh
	{"accessToken": "secret"}
```

See `src -h` for more information on specifying access tokens.

To acquire the access token, visit your Cdecode instance (or https://cdecode.com), click your profile picture, and select **access tokens** in the left hand menu.

## Development

If you want to develop the CLI, you can install it with `go get`:

```
go get -u github.com/cdecode/src-cli/cmd/src
```

## Releasing

1.  Find the latest version (either via the releases tab on GitHub or via git tags) to determine which version you are releasing.
2.  `VERSION=9.9.9 ./release.sh` (replace `9.9.9` with the version you are releasing)
3.  Travis will automatically perform the release. Once it has finished, **confirm that the curl commands fetch the latest version above**.
