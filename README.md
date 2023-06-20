# Setup

- clone to $GOPATH

```shell
git clone https://github.com/achntj/passman.git $GOPATH/passman
cd $GOPATH/passman
```

- create a pass.db file in the cloned directory (optional)

- add this function to your .zshrc or .bashrc file.

```shell
passman() {
    $GOPATH/passman/passman "$1";
}
```

- run-

```shell
passman add
```

or

```shell
passman get
```

# TODO

1. Add support for spaces in platform
2. Use OTPs for auth instead of a master key
3. Secure the pass.db file
4. Add grep search for platform / username
5. Maybe hide passwords while typing
