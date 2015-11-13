package crypto

// A collection of functions useful for implementing cryptography
import (
	"math/big"
	"math/rand"
	"time"
	)

// Function to find a prime integer of the specified bitlength and return it as a
//  big Int
func BigPrime (bitLength int64) *big.Int {
	// Create a new random number generator
	random 		:= rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	// Set the minimum and maximum size for our large integer
	targetSize 	:= new(big.Int).Exp(big.NewInt(2),big.NewInt(bitLength + 1),nil)
	minSize 	:= new(big.Int).Exp(big.NewInt(2),big.NewInt(bitLength),nil)
	// New big.Int to hold our random values and initialize it randomly.
	bigInt 		:= big.NewInt(0).Rand(random, targetSize)
	// ProbablyPrime will tell us if our guess is a prime number with 
	//  error rate approximately 1 in 5 billion when called with 100
	//  we can cheaply choose new random numbers until it returns true
	for !bigInt.ProbablyPrime(100) {
		// Randomize bigInt and make sure it is in our target range before
		//  checking for primality
		bigInt.Rand(random, targetSize)
		for bigInt.Cmp(minSize) <= 0 {
			bigInt.Rand(random, targetSize)
		}
	}
	// Just return the number
	return bigInt
}

// Function to find a^-1, mod b
//  Just calls the extended Euclidean algorithm
func ModInverse (a, b *big.Int) *big.Int {
	var inv *big.Int
	// Call the extended Euclidean with the larger of our two arguments in
	//  the first position. Only save the output corresponding to the inverse of a
	if a.Cmp(b) < 0 {
		_, _, inv = eEuclid(b, a)
	} else {
		_, inv, _ = eEuclid(a, b)
	}
	// If the eGCD result is negative, we can add
	//  our modulus b to get a positive inverse
	if inv.Sign() < 0 {
		inv.Add(inv, b)
	}
	return inv
}

// Just a wrapper for func eEuclid --
//  Calls eEuclid withe the arguments reversed if the b is larger
//  than a
func EEuclid (a, b *big.Int) (d, x, y *big.Int) {
	if (a.Cmp(b) <= 0) {
		 d, y, x = eEuclid(b, a)
	} else {
		d, x, y = eEuclid(a, b)
	}
	return
}

// Recursive implementation of the extended Euclidean algorithm
//  on big Ints
func eEuclid (a, b *big.Int) (d, x, y *big.Int) {
	// When b is zero, return a, 1, 0
	if b.Sign() == 0 {
		return new(big.Int).Set(a), big.NewInt(1), big.NewInt(0)
	}
	// Call eEuclid recursively on b, a % b
	d, xp, yp := eEuclid(new(big.Int).Set(b), big.NewInt(0).Mod(a,b))
	// Return calculated return values
	return d, yp, xp.Sub(xp, new(big.Int).Mul(new(big.Int).Div(a,b), yp))
}