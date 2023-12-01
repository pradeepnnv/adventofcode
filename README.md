# adventofcodechallenge
Advent Of Code Challenge

### Setup cmd
`cmd` package contains the code to execute the program from command line.
```shell
cd adventofcode/
mkdir 2023/cmd -p
cd 2023/cmd
go mod init github.com/pradeepnnv/adventofcode/2023/cmd
```

All new modules should be added to the Go workspace at root.

```shell
cd adventofcode/
go work use 2023/cmd
```

### Init each day

```shell
mkdir day1
cd day1
go mod init github.com/pradeepnnv/adventofcode/2023/day1
cd ../../
go work use 2023/day1
```

## Using module in main

```shell
cd adventofcode/2023/cmd
go mod edit -replace  github.com/pradeepnnv/adventofcode/2023/day1=../day1
```