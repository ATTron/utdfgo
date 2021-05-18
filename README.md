# UTDF-GO
[![Go Report Card](https://goreportcard.com/badge/github.com/attron/utdfgo)](https://goreportcard.com/report/github.com/attron/utdfgo)
UTDFGO is a UTDF Spacecraft Packet Decoder Written In Golang

# USAGE
------------------------------------------------------------------
```
    go get "github.com/ATTron/utdfgo"
```
## Using utdfgo in GO programs
```
    import "github.com/ATTron/utdfgo"

    . . .

    utdf, err := utdfgo.Run(filename)
    if err != nil {
      check(err)
    }
    for _, p := range utdf {
        // run functions on p
        p.GetYear()
    }
```
You can run the extraction functions with the returned utdf struct
```
