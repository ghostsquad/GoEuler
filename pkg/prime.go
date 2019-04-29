package pkg

import (
	"context"
	"github.com/pkg/errors"
	"math"
	"sort"
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

// GenerateUpToIncluding generates prime numbers
// if the number provided is not a prime, stops after finding the next prime after the given number
// otherwise generates primes up to the number given
func (p *Primes) GenerateUpToIncluding(ctx context.Context, max uint64) (<-chan uint64, <-chan error) {
	resultsChannel := make(chan uint64)
	errorChannel := make(chan error, 1)

	go func() {
		defer close(resultsChannel)
		defer close(errorChannel)

		maxKnown := p.Known[len(p.Known)-1]

		i := maxKnown + 2
		for {
			if maxKnown >= max {
				break
			}

			select {
			case <-ctx.Done():
				errorChannel <- ctx.Err()
				return
			default:
			}

			isPrimeAnswer, err := p.IsPrime(ctx, i)
			if err != nil {
				errorChannel <- err
				return
			}

			if isPrimeAnswer {
				p.Known = append(p.Known, i)
				maxKnown = i
				resultsChannel <- i
			}

			i += 2
		}

		errorChannel <- nil
	}()

	return resultsChannel, errorChannel
}

func (p *Primes) GenerateCountOf(ctx context.Context, maxCount uint64) (<-chan uint64, <-chan error) {
	resultsChannel := make(chan uint64)
	errorChannel := make(chan error, 1)

	go func() {
		defer close(resultsChannel)
		defer close(errorChannel)

		count := uint64(len(p.Known))
		i := p.Known[count-1] + 2
		for count < maxCount {
			select {
			case <-ctx.Done():
				errorChannel <- ctx.Err()
				return
			default:
			}

			isPrimeAnswer, err := p.IsPrime(ctx, i)
			if err != nil {
				errorChannel <- err
				return
			}

			if isPrimeAnswer {
				p.Known = append(p.Known, i)
				count++
				resultsChannel <- i
			}

			i += 2
		}

		errorChannel <- nil
	}()

	return resultsChannel, errorChannel
}

func (p *Primes) IsPrime(ctx context.Context, x uint64) (answer bool, err error) {
	// since we know that Known is in order, we can use a binary search
	i := sort.Search(len(p.Known), func(i int) bool { return p.Known[i] >= x })
	if i < len(p.Known) && p.Known[i] == x {
		return true, nil
	}

	sqrt := math.Sqrt(float64(x))

	// we may not have enough prime numbers to do arbitrary calculations efficiently
	if p.Known[len(p.Known)-1] < uint64(sqrt)+1 {
		return false, errors.New("not enough known primes")
	}

	for _, kp := range p.Known {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			return
		default:
		}

		if kp > uint64(sqrt)+1 {
			return true, nil
		}

		if x%kp == 0 {
			return false, nil
		}
	}

	return true, nil
}

func (p *Primes) Factorize(ctx context.Context, num uint64) (<-chan uint64, <-chan error) {
	resultsChannel := make(chan uint64)
	errorChannel := make(chan error, 1)

	go func() {
		defer close(resultsChannel)
		defer close(errorChannel)

		resultsChan, errsChan := p.GenerateUpToIncluding(ctx, uint64(num/2))
		for range resultsChan {
		}

		if err := <-errsChan; err != nil {
			errorChannel <- errors.Wrap(err, "prime number generation failed")
			return
		}

		result, err := p.IsPrime(ctx, num)
		if err != nil {
			errorChannel <- err
			return
		}

		if result {
			resultsChannel <- num
			errorChannel <- nil
			return
		}

		var tempNum float64 = float64(num)
		var newNum float64 = float64(num)
		primeIndex := 0
		prime := p.Known[primeIndex]

		for float64(prime) <= newNum {
			select {
			case <-ctx.Done():
				errorChannel <- ctx.Err()
				return
			default:
			}

			tempNum = newNum / float64(prime)

			if math.Mod(float64(tempNum), 1) == 0 {
				resultsChannel <- prime
				newNum = tempNum
			} else {
				primeIndex++
				prime = p.Known[primeIndex]
			}
		}

		errorChannel <- nil
	}()

	return resultsChannel, errorChannel
}
