---
License: MIT
LicenseFile: LICENSE
LicenseColor: yellow
---
# {{.Name}}

{{template "badge/travis" .}} {{template "badge/appveyor" .}} {{template "badge/goreport" .}} {{template "badge/godoc" .}} {{template "license/shields" .}}

{{pkgdoc}}

# {{toc 5}}

# Install
{{template "glide/install" .}}

# Cli

## Usage

#### $ {{exec "mutexer" "-help" | color "sh"}}

## Cli examples

```sh
# Create a mutexed version of demo/Tomate to gen_test/TomateSync
mutexer demo/Tomate:gen_test/TomateSync
# Specify the out pkg name
mutexer -p demo demo/Tomate:gen_test/TomateSync
# Use stdout
mutexer -p demo - demo/Tomate:gen_test/TomateSync
```
# API example

Following example demonstates a program using it to generate a sync version of a type.

#### > {{cat "demo/main.go" | color "go"}}

Following code is the generated implementation of `TomateSync` type.

#### > {{cat "demo/tomatessync.go" | color "go"}}


# Recipes

#### Release the project

```sh
gump patch -d # check
gump patch # bump
```

# History

[CHANGELOG](CHANGELOG.md)
