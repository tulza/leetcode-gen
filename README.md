# lcgen CLI

lcgen is a leetcode exercise generator that you can run on the cli

```bash
# initialise a leetcode folder using
lcgen init

# generate an exercise using
lcgen gen https://leetcode.com/problems/two-sum/description/
```

## Installation

### [Install on Windows](#installing-on-windows)

### [Install on Mac](#installing-on-macos)

---

<br/>

> [!NOTE]  
> You may need to install GoLang to compile this code.
> <br/>

### Installing on Windows

1. Compiling the code

```golang
go get .
go build
```

2. Create a folder `C:\lcgen` and put the lcgen.exe inside
3. Open environment variable settings <br/>
   `WIN + S` -> `Edit system environment settings` -> `Environment Variables`
4. In `User` or `System` variable find `Path`
5. Add `C:\lcgen` into the environment variable path
6. Run `lcgen` to see the result

### Installing on MacOS

1. Compiling the code

```golang
go get .
go build
```

2. Add to environment variable path using <br/>
   `mv lcgen /usr/local/bin`

3. Run `lcgen` to see the result

### Commands

```pwsh
lcgen
lcgen [-v | --version]
lcgen init
lcgen search [-d | --diff] [-n | --name] [-t | --topic]
lcgen gen [-n | --name] [-d | -daily] [URL]
```

### How does this work?

This cli program fetches infomation using [alfa leetcode api](https://alfa-leetcode-api.onrender.com/) to fetch user settings.
