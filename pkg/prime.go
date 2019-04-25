package pkg

import (
	"context"
	"github.com/pkg/errors"
	"sort"
	"math"
)

type Primes struct {
	Known []uint64
}

func NewPrimes() *Primes {
	return &Primes{
		Known: []uint64{
			2, 3, 5, 7, 11, 13,
		},
	}
}

func (p *Primes) Next(ctx context.Context) (answer uint64, err error) {
	for answer == 0 {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			return
		default:
		}
	}

	return
}

// GenerateUpToIncluding generates prime numbers
// if the number provided is not a prime, stops after finding the next prime after the given number
// otherwise generates primes up to the number given
func (p *Primes) GenerateUpToIncluding(ctx context.Context, channel chan<- uint64, max uint64) (err error) {
	newPrimes := []uint64{}
	defer func() {
		p.Known = append(p.Known, newPrimes...)
	}()

	maxKnown := p.Known[len(p.Known) - 1]

	i := maxKnown + 2
	for {
		if maxKnown >= max {
			return
		}

		select {
		case <-ctx.Done():
			err = ctx.Err()
			return
		default:
		}

		isPrimeAnswer, err := p.IsPrime(ctx, i)
		if err != nil {
			return err
		}

		if isPrimeAnswer {
			newPrimes = append(newPrimes, i)
			maxKnown = i
			channel <- i
		}

		i += 2
	}
}

func (p *Primes) GenerateCountOf(ctx context.Context, channel chan<- uint64, maxCount uint64) (err error) {
	newPrimes := []uint64{}
	defer func() {
		p.Known = append(p.Known, newPrimes...)
	}()

	count := uint64(len(p.Known))
	i := p.Known[count - 1] + 2
	for count < maxCount {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			return
		default:
		}

		isPrimeAnswer, err := p.IsPrime(ctx, i)
		if err != nil {
			return err
		}

		if isPrimeAnswer {
			newPrimes = append(newPrimes, i)
			count++
			channel <- i
		}

		i += 2
	}

	return
}

func (p *Primes) IsPrime(ctx context.Context, x uint64) (answer bool, err error) {
	// since we know that Known is in order, we can use a binary search
	i := sort.Search(len(p.Known), func(i int) bool { return p.Known[i] >= x })
    if i < len(p.Known) && p.Known[i] == x {
		return true, nil
	}

	sqrt := math.Sqrt(float64(x))

	// we may not have enough prime numbers to do arbitrary calculations efficiently
	if p.Known[len(p.Known) - 1] < uint64(sqrt) + 1 {
		return false, errors.New("not enough known primes")
	}

	for _, kp := range p.Known {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			return
		default:
		}

		if kp > uint64(sqrt) + 1 {
			return true, nil
		}

		if x % kp == 0 {
			return false, nil
		}
	}

	return true, nil
}
