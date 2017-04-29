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
# Create a mutexed version os Tomate to MyTomate
mutexer -p demo tomate_gen.go Tomate:MyTomate
```
# API example

Following example demonstates a program using it to generate a mutexed version of a type.

#### > {{cat "demo/lib.go" | color "go"}}

Following code is the generated implementation of `Tomate` type.

#### > {{cat "demo/vegetuxed_gen.go" | color "go"}}


# Recipes

#### Release the project

```sh
gump patch -d # check
gump patch # bump
```

# History

[CHANGELOG](CHANGELOG.md)
