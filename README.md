[![build on release](https://github.com/juanCortelezzi/Royallist/actions/workflows/build-on-release.yml/badge.svg)](https://github.com/juanCortelezzi/Royallist/actions/workflows/build-on-release.yml)

# Royallist

A glorified "ls" command.

Don't be surprised, it's just ls with icons and written in Go

## How to use?

Call royallist and watch you current directory files shine in the terminal. If you want to use
another directory, pass it as a first argument just like "ls" and on you go.
```bash
# To print local directory
royallist

# Or add as a parameter the desired directory
royallist ~/Documents

# Or even multiple
royallist ~/Documents ~/Pictures


# Recursive/tree view (only one directory)
royallist -r
royallist -r ~/Documents

# For help
royallist -h
royallist --help
```

For ease of use, I recommend aliasing Royallist to a simpler command:

Just add the following to your shell RC file

```bash
alias l="royallist"
```

## Install

Please note you need a NerdFont or similar.

In order to install, compile the source code, or just Follow release
instructions.
