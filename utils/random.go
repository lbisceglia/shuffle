package utils

// import "sync"

// import (
// 	crand "crypto/rand"
// 	"encoding/binary"
// 	"errors"
// 	mrand "math/rand"
// 	"sync"
// )

// var once sync.Once

// // Randomize sets a pseudorandom seed for the global random number generator.
// func Randomize() {
// 	if seed, err := generateRandomSeed(); err == nil {
// 		SetRandomSeed(seed)
// 	} else {
// 		panic(err)
// 	}
// }

// // generateRandomSeed returns a securely generated random seed.
// // It will return an error if the system's secure random
// // number generator fails to function correctly, in which
// // case the caller should not continue.
// func generateRandomSeed() (int64, error) {
// 	b := make([]byte, 8)
// 	_, err := crand.Read(b)
// 	if err != nil {
// 		return 0, errors.New("cannot seed math/rand package with cryptographically secure random number")
// 	}
// 	return int64(binary.LittleEndian.Uint64(b)), nil
// }

// // SetRandomSeed sets a specific seed for the global random number generator.
// func SetRandomSeed(seed int64) {
// 	once.Do(func() {
// 		mrand.Seed(seed)
// 	})
// }

// // Testing Helpers

// Reset allows SetRandomSeed to be called with fresh values.
// func Reset() {
// 	once = sync.Once{}
// }
