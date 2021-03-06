package common

import (
	"math/big"
	"math/rand"
)

// UnitStrMap is a map from unit string to value string
var UnitStrMap = map[string]string{
	"noether":    "0",
	"wei":        "1",
	"kwei":       "1000",
	"Kwei":       "1000",
	"babbage":    "1000",
	"femtoether": "1000",
	"mwei":       "1000000",
	"Mwei":       "1000000",
	"lovelace":   "1000000",
	"picoether":  "1000000",
	"gwei":       "1000000000",
	"Gwei":       "1000000000",
	"shannon":    "1000000000",
	"nanoether":  "1000000000",
	"nano":       "1000000000",
	"szabo":      "1000000000000",
	"microether": "1000000000000",
	"micro":      "1000000000000",
	"finney":     "1000000000000000",
	"milliether": "1000000000000000",
	"milli":      "1000000000000000",
	"ether":      "1000000000000000000",
	"kether":     "1000000000000000000000",
	"grand":      "1000000000000000000000",
	"mether":     "1000000000000000000000000",
	"gether":     "1000000000000000000000000000",
	"tether":     "1000000000000000000000000000000",
}

// UnitIntMap is a map from unit string to value as big int
var UnitIntMap = map[string]*big.Int{
	"noether":    big.NewInt(0),
	"wei":        big.NewInt(1),
	"kwei":       big.NewInt(1000),
	"Kwei":       big.NewInt(1000),
	"babbage":    big.NewInt(1000),
	"femtoether": big.NewInt(1000),
	"mwei":       big.NewInt(1000000),
	"Mwei":       big.NewInt(1000000),
	"lovelace":   big.NewInt(1000000),
	"picoether":  big.NewInt(1000000),
	"gwei":       big.NewInt(1000000000),
	"Gwei":       big.NewInt(1000000000),
	"shannon":    big.NewInt(1000000000),
	"nanoether":  big.NewInt(1000000000),
	"nano":       big.NewInt(1000000000),
	"szabo":      big.NewInt(1000000000000),
	"microether": big.NewInt(1000000000000),
	"micro":      big.NewInt(1000000000000),
	"finney":     big.NewInt(1000000000000000),
	"milliether": big.NewInt(1000000000000000),
	"milli":      big.NewInt(1000000000000000),
	"ether":      big.NewInt(1000000000000000000),
	"kether":     big.NewInt(-1), // Manual calculation is needed from here
	"grand":      big.NewInt(-1),
	"mether":     big.NewInt(-1),
	"gether":     big.NewInt(-1),
	"tether":     big.NewInt(-1),
}

// UnitFloatMap is a map from unit string to value as big float
var UnitFloatMap = map[string]*big.Float{
	"noether":    big.NewFloat(0),
	"wei":        big.NewFloat(1),
	"kwei":       big.NewFloat(1000),
	"Kwei":       big.NewFloat(1000),
	"babbage":    big.NewFloat(1000),
	"femtoether": big.NewFloat(1000),
	"mwei":       big.NewFloat(1000000),
	"Mwei":       big.NewFloat(1000000),
	"lovelace":   big.NewFloat(1000000),
	"picoether":  big.NewFloat(1000000),
	"gwei":       big.NewFloat(1000000000),
	"Gwei":       big.NewFloat(1000000000),
	"shannon":    big.NewFloat(1000000000),
	"nanoether":  big.NewFloat(1000000000),
	"nano":       big.NewFloat(1000000000),
	"szabo":      big.NewFloat(1000000000000),
	"microether": big.NewFloat(1000000000000),
	"micro":      big.NewFloat(1000000000000),
	"finney":     big.NewFloat(1000000000000000),
	"milliether": big.NewFloat(1000000000000000),
	"milli":      big.NewFloat(1000000000000000),
	"ether":      big.NewFloat(1000000000000000000),
	"kether":     big.NewFloat(1000000000000000000000),
	"grand":      big.NewFloat(1000000000000000000000),
	"mether":     big.NewFloat(1000000000000000000000000),
	"gether":     big.NewFloat(1000000000000000000000000000),
	"tether":     big.NewFloat(1000000000000000000000000000000),
}

// RandomUint64 generates random uint64
func RandomUint64() uint64 {
	return uint64(rand.Uint32())<<32 + uint64(rand.Uint32())
}
