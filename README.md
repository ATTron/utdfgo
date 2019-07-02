# UTDF-GO
UTDF-GO is a golang decoder used to extract information out of 
UTDF Packets

# USAGE
```
    go get "github.com/ATTron/utdf-go"
```
------------------------------------------------------------------
```
    import "github.com/ATTron/utdf-go"

    . . .

    utdf := run(filename)
    for _, p := range utdf {
        // run functions on p
    }
```
You can run the extraction functions with the returned utdf struct

# Functions

## Getting Started
#### run(filename string)
returns an array of utdf packets from the file given

------------------------------------------------------------------
## Time Based Functions
#### getYear()
returns the year
taken from byte 6
#### getSeconds()
returns the seconds of year
taken from bytes 11:14
#### getMicroseconds()
returns microseconds of seconds
taken from bytes 15:18
#### getEpoch()
returns a UTC epoch
built from other time functions
#### getTimestamps()
returns all time information for given utdf packet

------------------------------------------------------------------
## Angle Based Functions
#### getAzimuth()
returns azimuth
taken from bytes 19:22
#### getElevation()
returns elevation
taken from bytes 23:26

------------------------------------------------------------------
## Range + Doppler Based Functions
#### getRangeDelay()
returns range delay hi and lo
#### getRange()
returns calculated range of spacecraft
#### getDopplerDelay()
returns doppler delay hi and lo
#### getDoppler()
returns calculated doppler count

------------------------------------------------------------------
## Misc Functions
#### getAGC()
returns spacecraft AGC
taken from bytes 39:40
#### getTransmitFreq()
returns transmission frequency
taken from bytes 41:44
#### getAntennaType()
returns antenna type
taken from byte 45
#### getPADID()
returns antenna padid
taken from byte 46
#### getRecieveAntennaType()
returns recieve antenna type
taken from byte 47
#### getRecievePADID()
returns recieve antenna padid
taken from byte 48
#### getSystemMode()
return system mode
taken from bytes 49:50
#### getDataValidation()
return data validity
taken from byte 51
#### getFrequencyBand()
return frequency band
taken from byte 52
#### getTrackingInfo()
return tracking type and data
taken from bytes 53:54