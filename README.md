# conway-app-go

[![Actions Status](https://github.com/otiai10/conway-app-go/workflows/Go/badge.svg)](https://github.com/otiai10/conway-app-go/actions)

```
go test ./... -v -cover
```

```
go run main.go
```

# More options

```
  -a string
        Expression string for alive cells (default "*")
  -d string
        Expression string for dead cells (default " ")
  -h int
        Height of your world (default 20)
  -s int
        Interval of alternation (milli sec) (default 200)
  -w int
        Width of your world (default 40)
```

# Examples

```
go run main.go -a "🤔" -d "  " -s 100 -w 40 -h 40
```

<img src="https://user-images.githubusercontent.com/931554/79042982-61bd6a80-7bfc-11ea-894d-d85464875e4f.png" width="50%" />