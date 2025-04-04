# buf-fmt-stdin

[Buf protocol buffer formatter](https://buf.build/docs/reference/cli/buf/format/) wrapper for stdin text editor integration

## Context

Some code editors support stdin integration, allowing users to pipe code directly into the editor for formatting and linting. This tool provides a convenient way to integrate Buf with such editors.

## Installation

```sh
wget https://github.com/Fuabioo/buf-fmt-stdin/releases/latest/download/buf-fmt-stdin_$(uname -s)_$(uname -m).tar.gz
tar -xzf buf-fmt-stdin_$(uname -s)_$(uname -m).tar.gz
chmod +x buf-fmt-stdin
```

## Usage

Since this is a wrapper to buf-fmt fix command then you can use any flag
defined in this tool such as `--config`.

### Zed

With the [Zed editor](https://zed.dev/), you can use this tool by adding the [configuration](zed-configuration.json) to your settings json.

How [Zed formatting](https://zed.dev/docs/configuring-zed#formatter) works for anything other than a language server formatter is that you set up an external CLI tool, to which it will pipe the code. Currently there is no support for providing something like a filepath to the external tool or the like.

> For more information about Protocol Buffers language support, see [Proto - Zed](https://zed.dev/docs/languages/proto).

### Developing

If you have the go/bin directory in your PATH, you can simply run `go install` to install the tool from the source code in your local machine.
Once that is done, just pipe any code into the tool or
(if you already have one) use your configured text editor.

Execute the tool with the following command:

```sh
echo '


syntax = "proto3";

  package example;


  message TestMessage {
    string name =  1;
    int32 age = 2;
  }' | buf-fmt-stdin
```

This should throw the following output (the editor will know what to do with it):

```proto3

syntax = "proto3";

package example;

message TestMessage {
  string name = 1;
  int32 age = 2;
}

```
