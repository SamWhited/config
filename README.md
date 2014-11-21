# Go Config

Loads TOML config files and string blobs in Go. Config is loaded into `config.C map[string]interface{}`. If
`config.DefaultConfig` is set to a TOML blob, that is loaded on package init. If `config.ProjectName` is set, the
following files are loaded (after `config.DefaultConfig` if it's set) in order:

	1. `/etc/<ProjectName>/config`
	2. `~/.config/<ProjectName>/config`
	3. `./config`

Otherwise you can do this yourself later. Note that errors are not reported if config loading fails at any point during
init. They are simply ignored and the init process continues as if nothing has happened.

Note that `config.ProjectName` is considered a trusted variable (there's no checking for strings such as `../`. You
should NOT set the project name from user input.

To import and use this package in Go:

    import "github.com/SamWhited/config"
