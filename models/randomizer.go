package models

import (
	crand "crypto/rand"
	"encoding/binary"
	"errors"
	mrand "math/rand"
	"sync"
)

type Randomizer interface {
	Seed(seed int64)
	Randomize()
	Shuffle(s shoe)
}

type rng struct {
	once sync.Once
	r    *mrand.Rand
}

func NewRNG() *rng {
	seed, _ := GenerateRandomSeed()
	rng := &rng{
		once: sync.Once{},
		r:    mrand.New(mrand.NewSource(seed)),
	}
	return rng
}

// GenerateRandomSeed returns a securely generated random seed.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomSeed() (int64, error) {
	b := make([]byte, 8)
	_, err := crand.Read(b)
	if err != nil {
		return 0, errors.New("cannot seed math/rand package with cryptographically secure random number")
	}
	return int64(binary.LittleEndian.Uint64(b)), nil
}

// SetRandomSeed sets a specific seed for the global random number generator.
func (rng *rng) Seed(seed int64) {
	rng.once.Do(func() {
		rng.r.Seed(seed)
	})
}

// Randomize sets a pseudorandom, securlely generated seed for the random number generator.
func (rng *rng) Randomize() {
	if seed, err := GenerateRandomSeed(); err == nil {
		rng.Seed(seed)
	} else {
		panic(err)
	}
}

// Shuffle randomly shuffles a Shoe.
func (rng *rng) Shuffle(s shoe) {
	rng.r.Shuffle(len(s), func(i, j int) { s[j], s[i] = s[i], s[j] })
}
