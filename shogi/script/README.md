# Shogi Script

Minimal DSL for shogi.

## Syntax

* A script is sequence of lines of commands.
* Command line is constructed by space-separated command name and arguments.

Example:

```
# <- comment line

# space-separated command and arguments
command arg1 arg2 ...

# labelled argument
command label1:value1 arg2 label3:"value 2"

# multiline argument
command arg1:"""
foo
bar
""" arg2
```
