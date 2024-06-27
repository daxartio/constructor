# Constructor

`constructor` is a Go code generator for creating constructors for your structs. This tool helps you generate constructor functions automatically, saving you time and ensuring consistency across your codebase.

## Installation

To install the constructor tool, use go get:

```
go get github.com/daxartio/constructor
```

## Usage

The `constructor` tool can be run from the command line with various options. Below are the available options and their descriptions:

```
Usage of constructor:
  -f string
        A format of the filename. (default "%s_constructor_gen.go")
  -h    Show help.
  -n    No prefix.
  -p string
        Package full path. (default ".")
  -s string
        Match structs name. (comma separated)
  -w    Write generated code.
```

### Options

- `-f string`: Specifies the format of the generated filename. By default, the filename format is `%s_constructor_gen.go`, where `%s` will be replaced with the name of the struct.
- `-h`: Displays help information for the constructor tool.
- `-n`: Disables the prefix in the generated constructor function names.
- `-p string`: Specifies the full path of the package where the structs are located. The default value is the current directory (.).
- `-s string`: Specifies the names of the structs to match, separated by commas.
- `-w:` Writes the generated code to the specified file. If this option is not set, the generated code will be printed to the standard output.

## Examples

Generate constructors for all structs in the current package and print the result to the standard output:

```
constructor -p .
```

Generate constructors for specific structs (User,Order) and write the generated code to files:

```
constructor -p . -s User,Order -w
```

Generate constructors without any prefix in the function names:

```
constructor -p . -n
```
