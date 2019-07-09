// Package utdfgo overview
// Spacecraft UTDF Packet Decoder
package utdfgo

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"time"
)

// speed of light consant
const sol = 299792.458

// full circle is 2^32
const fullCircle = 4294967296
const twoPi = 2 * math.Pi

// utdf packets are 75 bytes long
const bufferSize = 75

// UTDF data type creation
type UTDF []byte

// Run read in entire file then chunk into utdf packets
func Run(filename string) ([]UTDF, error) {
	var utdfArr []UTDF
	utdfFile, err := ioutil.ReadFile(filename)
	check(err)
	for i := 0; i < len(utdfFile); i += bufferSize {
		end := i + bufferSize
		if end > len(utdfFile) {
			end = len(utdfFile)
		}
		if err = checkUTDF(utdfFile[i:end]); err != nil {
			return nil, ErrBadFile
		}
		utdfArr = append(utdfArr, utdfFile[i:end])
	}
	return utdfArr, nil
}

// START TIME BASED FUNCTIONS

// GetYear gets last two digits of current year
// taken from byte 6
func (utdf UTDF) GetYear() int {
	year := int(utdf[5])
	if year < 70 {
		year += 2000
	} else {
		year += 1900
	}
	return year
}

// GetSeconds get seconds of year
// taken from bytes 11:14
func (utdf UTDF) GetSeconds() int {
	scY := (int(utdf[10]) << 24) + (int(utdf[11]) << 16) + (int(utdf[12]) << 8) + int(utdf[13])
	return scY
}

// GetMicroseconds get microseconds of seconds
// taken from bytes 15:18
func (utdf UTDF) GetMicroseconds() int {
	mcS := (int(utdf[14]) << 24) + (int(utdf[15]) << 16) + (int(utdf[16]) << 8) + int(utdf[17])
	return mcS
}

// GetEpoch create UTC epoch
func (utdf UTDF) GetEpoch() int64 {
	ytd := time.Date(utdf.GetYear(), 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	epoch := int(ytd) + utdf.GetSeconds() + utdf.GetMicroseconds()/1000000
	return int64(epoch)
}

// GetTimestamps gets timestamp from utdf packet
func (utdf UTDF) GetTimestamps() string {
	return fmt.Sprintf("Year: %d, Epoch: %d, Seconds: %d, Microseconds: %d", utdf.GetYear(), utdf.GetEpoch(), utdf.GetSeconds(), utdf.GetMicroseconds())
}

// END TIME BASED FUNCTIONS

// START ANGLE + ELEVATION BASED FUNCTIONS

// GetAzimuth get angle_data[0] and calculate azimuth
// taken from bytes 19:22
func (utdf UTDF) GetAzimuth() float64 {
	azi := (int(utdf[18]) << 24) + (int(utdf[19]) << 16) + (int(utdf[20]) << 8) + int(utdf[21])
	aAngle := calculateAngle(azi)
	return aAngle
}

// GetElevation get angle_data[1] and calculate elevation
// taken from bytes 23-26
func (utdf UTDF) GetElevation() float64 {
	e := (int(utdf[22]) << 24) + (int(utdf[23]) << 16) + (int(utdf[24]) << 8) + int(utdf[25])
	eAngle := calculateAngle(e)
	return eAngle
}

// END ANGLE + ELEVATION BASED FUNCTIONS

// START RANGE + DOPPLER BASED FUNCTIONS

// GetRangeDelay calculate range delay hi low count
func (utdf UTDF) GetRangeDelay() (float64, float64) {
	rd := (int(utdf[26]) << 40) + (int(utdf[27]) << 32) + (int(utdf[28]) << 24) + (int(utdf[29]) << 16) + (int(utdf[30]) << 8) + int(utdf[31])
	rdHi, rdLo := calculateRD(rd)
	return float64(rdHi), float64(rdLo)
}

// GetRange calculate spacecraft range
func (utdf UTDF) GetRange() float64 {
	rh, rl := utdf.GetRangeDelay()
	r := (((rh + rl) - 0) * float64(sol) / 2000000)
	return r
}

// GetDopplerDelay get doppler hi low count
func (utdf UTDF) GetDopplerDelay() (float64, float64) {
	d := (int(utdf[32]) << 40) + (int(utdf[33]) << 32) + (int(utdf[34]) << 24) + (int(utdf[35]) << 16) + (int(utdf[36]) << 8) + int(utdf[37])
	dHi, dLo := calculateRD(d)
	return float64(dHi), float64(dLo)
}

// GetDoppler get spacecraft doppler count
func (utdf UTDF) GetDoppler() float64 {
	dh, dl := utdf.GetDopplerDelay()
	d := (((dh + dl) - 0) * float64(sol) / 2000000)
	return d
}

// END RANGE BASED FUNCTIONS

// START MISC FUNCTIONS

// GetAGC get AGC
// taken from bytes 39:40
func (utdf UTDF) GetAGC() float32 {
	agc := (int(utdf[38]) << 8) + int(utdf[39])
	return float32(agc)
}

// GetTransmitFreq get transmit frequency
// taken from bytes 41:44
func (utdf UTDF) GetTransmitFreq() string {
	tf := (int(utdf[40]) << 24) + (int(utdf[41]) << 16) + (int(utdf[42]) << 8) + int(utdf[43])
	hz := float64(tf) * 10
	return fmt.Sprintf("%1.10v", hz)
}

// GetAntennaType get antenna type
// taken from byte 45
func (utdf UTDF) GetAntennaType() byte {
	at := utdf[44]
	return at
}

// GetPADID get antenna padid
// taken from byte 46
func (utdf UTDF) GetPADID() int8 {
	pid := utdf[45]
	return int8(pid)
}

// GetRecieveAntennaType get receive antenna type
// taken from byte 47
func (utdf UTDF) GetRecieveAntennaType() byte {
	at := utdf[46]
	return at
}

// GetRecievePADID get receive antenna padid
// taken from byte 48
func (utdf UTDF) GetRecievePADID() int8 {
	pid := utdf[45]
	return int8(pid)
}

// GetSystemMode get system mode
// taken from bytes 49:50
func (utdf UTDF) GetSystemMode() int {
	m := (int(utdf[48]) << 8) + int(utdf[49])
	return m
}

// GetDataValidation get data validity
// taken from byte 51
func (utdf UTDF) GetDataValidation() byte {
	dv := utdf[50]
	return dv
}

// GetFrequencyBand get frequency band
// taken from byte 52
func (utdf UTDF) GetFrequencyBand() byte {
	fb := utdf[51]
	return fb
}

// GetTrackingInfo get tracking type and data
// taken from bytes 53:54
func (utdf UTDF) GetTrackingInfo() int {
	ti := (int(utdf[52]) << 8) + int(utdf[53])
	return ti
}

// END MISC FUNCTIONS

// GetSIC get spacecraft ID
func (utdf UTDF) GetSIC() uint64 {
	sic := utdf[6:8]
	sicE := hex.EncodeToString(sic)
	n, err := strconv.ParseUint(sicE, 16, 16)
	check(err)
	return n
}

// GetVID get vehicle ID
func (utdf UTDF) GetVID() uint64 {
	vid := utdf[9:10]
	vidE := hex.EncodeToString(vid)
	n, err := strconv.ParseUint(vidE, 16, 16)
	check(err)
	return n
}

func (utdf UTDF) isValid() bool {
	valid := true
	defer func() {
		if r := recover(); r != nil {
			valid = false
		}
	}()
	if utdf.GetYear() > time.Now().Year() {
		valid = false
	}
	return valid
}

// ToString currently just used for testing correct return values
func (utdf UTDF) ToString() string {
	return fmt.Sprintf("%s\t\t%1.15v\t%1.15v\t%1.15v", utdf.GetTimestamps(), utdf.GetAzimuth(), utdf.GetElevation(), utdf.GetRange())
}
