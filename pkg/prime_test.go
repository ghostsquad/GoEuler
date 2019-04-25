package pkg

import (
	"context"
	"fmt"
	"math"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T) {
	isPrimeCases := []struct {
		num     uint64
		isPrime bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{10, false},
		{11, true},
		{12, false},
		{13, true},
		{14, false},
		{15, false},
		{16, false},
		{17, true},
		{18, false},
		{19, true},
		{20, false},
		{21, false},
		{22, false},
		{23, true},
		{29, true},
		{31, true},
		{37, true},
		{41, true},
		{43, true},
		{47, true},
		{53, true},
		{59, true},
		{61, true},
		{67, true},
		{71, true},
		{73, true},
		{79, true},
		{83, true},
		{89, true},
		{97, true},
		{101, true},
		{103, true},
		{107, true},
		{109, true},
		{113, true},
		{127, true},
		{131, true},
		{137, true},
		{139, true},
		{149, true},
		{151, true},
		{157, true},
		{163, true},
		{165, false},
		{167, true},
	}
	for _, tc := range isPrimeCases {
		t.Run(fmt.Sprintf("[ %d prime? %t ]", tc.num, tc.isPrime), func(t *testing.T) {
			kp := NewPrimes()
			ctx := context.TODO()

			result, err := kp.IsPrime(ctx, tc.num)

			assert.Nil(t, err)
			assert.Equal(t, tc.isPrime, result)
		})
	}

	t.Run("Given number larger than sqrt of known primes, returns err", func(t *testing.T) {
		kp := NewPrimes()
		ctx := context.TODO()
		_, err := kp.IsPrime(ctx, 701)

		assert.NotNil(t, err)
	})

	t.Run("When allowed to generate numbers up to sqrt then calculation succeeds", func(t *testing.T) {
		kp := NewPrimes()
		ctx := context.TODO()

		var num uint64 = 701

		_, err1 := kp.IsPrime(ctx, num)

		assert.NotNil(t, err1)

		sqrt := math.Sqrt(float64(num))
		sqrtInt := uint64(sqrt)

		channel := make(chan uint64)
		go func() {
			for {
				<-channel
			}
		}()

		err2 := kp.GenerateUpToIncluding(ctx, channel, sqrtInt)
		close(channel)

		assert.Nil(t, err2)

		result, err3 := kp.IsPrime(ctx, num)

		assert.Nil(t, err3)
		assert.True(t, result, fmt.Sprintf("%+v", kp.Known))

		largestKnownPrime := kp.Known[len(kp.Known) - 1]
		assert.True(t, largestKnownPrime > sqrtInt, fmt.Sprintf("%d should be > %d", largestKnownPrime, sqrtInt))
	})
}

func TestGenerateCountOf(t *testing.T) {
	kp := NewPrimes()
	ctx := context.TODO()

	channel := make(chan uint64)
	go func() {
		for {
			<-channel
		}
	}()
	assert.NotEqual(t, 20, len(kp.Known))

	err := kp.GenerateCountOf(ctx, channel, 20)

	assert.Nil(t, err)
	assert.Equal(t, 20, len(kp.Known))

	for _, n := range kp.Known {
		result, err := kp.IsPrime(ctx, n)

		assert.Nil(t, err)
		assert.True(t, result)
	}
}
