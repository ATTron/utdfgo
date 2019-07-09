package utdfgo

import (
	"errors"
	"io"
	"log"
)

var ErrBadFile = errors.New("Invalid Format: Packet Not In UTDF Format")

func check(err error) {
	if err != nil {
		if err != io.EOF {
			log.Fatal("Error running utdf package", err)
		}
	}
}

// calculate angle based functions
func calculateAngle(i int) float64 {
	angle := float64(i) / float64(fullCircle) * twoPi
	return angle
}

// calculate range and doppler functions
func calculateRD(i int) (float64, float64) {
	iHi := i / 65536
	iLo := i - iHi*65536
	return float64(iHi), float64(iLo)
}

// check if packet is in valid utdf format
func checkUTDF(utdf UTDF) error {
	if !utdf.isValid() {
		return ErrBadFile
	}
	return nil
}
