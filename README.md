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

    utdf := utdfgo.Run(filename)
    for _, p := range utdf {
        // run functions on p
    }
```
You can run the extraction functions with the returned utdf struct
## Using CLI
Download latest release:
Windows:
Use the utdf.exe
Linux + Mac:
```
    make install-linux-mac
```

# Functions

## Getting Started
#### Run(filename string)
returns an array of utdf packets from the file given

------------------------------------------------------------------
## Time Based Functions
#### GetYear()
returns the year
taken from byte 6
#### GetSeconds()
returns the seconds of year
taken from bytes 11:14
#### GetMicroseconds()
returns microseconds of seconds
taken from bytes 15:18
#### GetEpoch()
returns a UTC epoch
built from other time functions
#### GetTimestamps()
returns all time information for given utdf packet

------------------------------------------------------------------
## Angle Based Functions
#### GetAzimuth()
returns azimuth
taken from bytes 19:22
#### GetElevation()
returns elevation
taken from bytes 23:26

------------------------------------------------------------------
## Range + Doppler Based Functions
#### GetRangeDelay()
returns range delay hi and lo
#### GetRange()
returns calculated range of spacecraft
#### GetDopplerDelay()
returns doppler delay hi and lo
#### GetDoppler()
returns calculated doppler count

------------------------------------------------------------------
## Misc Functions
#### GetAGC()
returns spacecraft AGC
taken from bytes 39:40
#### GetTransmitFreq()
returns transmission frequency
taken from bytes 41:44
#### GetAntennaType()
returns antenna type
taken from byte 45
#### GetPADID()
returns antenna padid
taken from byte 46
#### GetRecieveAntennaType()
returns recieve antenna type
taken from byte 47
#### GetRecievePADID()
returns recieve antenna padid
taken from byte 48
#### GetSystemMode()
return system mode
taken from bytes 49:50
#### GetDataValidation()
return data validity
taken from byte 51
#### GetFrequencyBand()
return frequency band
taken from byte 52
#### GetTrackingInfo()
return tracking type and data
taken from bytes 53:54