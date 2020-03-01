package main

import "fmt"

type planetdata struct {
	name   string  // Name of planet
	tp     float64 // Period of orbit
	long   float64 // Longitude at the epoch
	peri   float64 // Longitude of the perihelion
	ecc    float64 // Eccentricity of the orbit
	axis   float64 // Semi-major axis of the orbit
	incl   float64 // Orbital inclination
	node   float64 // Longitude of the ascending node
	theta0 float64 // Angular diameter at 1 AU
	v0     float64 // Visual magnitude at 1 AU
}

func main() {
	planetDataCollection := make([]planetdata, 0)

	planetDataCollection = append(planetDataCollection, planetdata{
		name:   "Mercury",
		tp:     0.24085,
		long:   75.5671,
		peri:   77.612,
		ecc:    0.205627,
		axis:   0.387098,
		incl:   7.0051,
		node:   48.449,
		theta0: 6.74,
		v0:     -0.42})
	planetDataCollection = append(planetDataCollection, planetdata{
		name:   "Venus",
		tp:     0.615207,
		long:   272.30044,
		peri:   131.54,
		ecc:    0.006812,
		axis:   0.723329,
		incl:   3.3947,
		node:   76.769,
		theta0: 16.92,
		v0:     -4.4})
	planetDataCollection = append(planetDataCollection, planetdata{
		name:   "Earth",
		tp:     0.999996,
		long:   99.556772,
		peri:   103.2055,
		ecc:    0.016671,
		axis:   0.999985,
		incl:   -99.0,
		node:   -99.0,
		theta0: -99.0,
		v0:     -99.0})
	planetDataCollection = append(planetDataCollection, planetdata{
		name:   "Mars",
		tp:     1.880765,
		long:   109.09646,
		peri:   336.217,
		ecc:    0.093348,
		axis:   1.523689,
		incl:   1.8497,
		node:   49.632,
		theta0: 9.36,
		v0:     -1.52})
	planetDataCollection = append(planetDataCollection, planetdata{
		name:   "Jupiter",
		tp:     11.857911,
		long:   337.917132,
		peri:   14.6633,
		ecc:    0.048907,
		axis:   5.20278,
		incl:   1.3035,
		node:   100.595,
		theta0: 196.74,
		v0:     -9.4})
	planetDataCollection = append(planetDataCollection, planetdata{
		name:   "Saturn",
		tp:     29.310579,
		long:   172.398316,
		peri:   89.567,
		ecc:    0.053853,
		axis:   9.51134,
		incl:   2.4873,
		node:   113.752,
		theta0: 165.6,
		v0:     -8.88})
	planetDataCollection = append(planetDataCollection, planetdata{
		name:   "Uranus",
		tp:     84.039492,
		long:   356.135400,
		peri:   172.884833,
		ecc:    0.046321,
		axis:   19.21814,
		incl:   0.773059,
		node:   73.926961,
		theta0: 65.8,
		v0:     -7.19})
	planetDataCollection = append(planetDataCollection, planetdata{
		name:   "Neptune",
		tp:     165.845392,
		long:   326.895127,
		peri:   23.07,
		ecc:    0.010483,
		axis:   30.1985,
		incl:   1.7673,
		node:   131.879,
		theta0: 62.2,
		v0:     -6.87})

	// Display a specific field value for a specific record:
	fmt.Printf("Specific field ('name') value for a specific record (record 0): %v\n", planetDataCollection[0].name)

	// Find a specific record (e.g., "Jupiter") in a slice:
	for _, planetRecord := range planetDataCollection {
		if planetRecord.name == "Jupiter" {
			fmt.Printf("Found %v\n", planetRecord.name)
			fmt.Printf("\tAll elements: %v\n", planetRecord)
			fmt.Printf("\tThe orbital eccentricity of %v is %v\n", planetRecord.name, planetRecord.ecc)
		}
	}
}
