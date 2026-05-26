package main

import (
	"flag"
)

var flagShort = flag.Bool("short", false, "run the example with a smaller and insecure ring degree.")

func obliviousRiding() {
	_ = "STUB: not implemented"

	// This example simulates a situation where an anonymous rider
	// wants to find the closest available rider within a given area.
	// The application is inspired by the paper https://infoscience.epfl.ch/record/228219
	//
	//	A. Pham, I. Dacosta, G. Endignoux, J. Troncoso-Pastoriza,
	//	K. Huguenin, and J.-P. Hubaux. ORide: A Privacy-Preserving
	//	yet Accountable Ride-Hailing Service. In Proceedings of the
	//	26th USENIX Security Symposium, Vancouver, BC, Canada, August 2017.
	//
	// Each area is represented as a rectangular grid where each driver
	// anonymously signs in (i.e. the server only knows the driver is located
	// in the area).
	//
	// First, the rider generates an ephemeral key pair (riderSk, riderPk), which she
	// uses to encrypt her coordinates. She then sends the tuple (riderPk, enc(coordinates))
	// to the server handling the area she is in.
	//
	// Once the public key and the encrypted rider coordinates of the rider
	// have been received by the server, the rider's public key is transferred
	// to all the drivers within the area, with a randomized different index
	// for each of them, that indicates in which coefficient each driver must
	// encode her coordinates.
	//
	// Each driver encodes her coordinates in the designated coefficient and
	// uses the received public key to encrypt her encoded coordinates.
	// She then sends back the encrypted coordinates to the server.
	//
	// Once the encrypted coordinates of the drivers have been received, the server
	// homomorphically computes the squared distance: (x0 - x1)^2 + (y0 - y1)^2 between
	// the rider and each of the drivers, and sends back the encrypted result to the rider.
	//
	// The rider decrypts the result and chooses the closest driver.
	return
}

// Number of drivers in the area
//max is N

// Parameters (128 bit security) with plaintext modulus 65929217
// Creating encryption parameters from a default params with logN=14, logQP=438 with a plaintext modulus T=65929217

// Rider's keygen

// max values = floor(sqrt(plaintext modulus))
// binary mask upper-bound for the uniform sampling

// Rider coordinates [x, y, x, y, ....., x, y]

// driversData coordinates [0, 0, ..., x, y, ..., 0, 0]

func distance(a, b, c, d uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func main() {
	flag.Parse()
	obliviousRiding()
}
