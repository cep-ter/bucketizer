package xxhash

import (
	"fmt"
	"github.com/cespare/xxhash"
	"github.com/ehsangolshani/bucketizer"
	"github.com/pkg/errors"
)

type XXHASHBucketizer struct {
	Buckets      []bucketizer.Bucket
	weightSum    uint64
	bucketRanges []uint64
	seed         []byte
}

func (b XXHASHBucketizer) BucketBytes(value []byte) (int, error) {
	a := append(value, b.seed...)
	hashDigest := xxhash.Sum64(a)
	reminder := hashDigest % b.weightSum
	n := len(b.bucketRanges)
	for i := 0; i < n-1; i++ {
		if reminder >= b.bucketRanges[i] && reminder < b.bucketRanges[i+1] {
			return i, nil
		}
	}
	return 0, errors.New("invalid reminder value")
}

func (b XXHASHBucketizer) BucketString(value string) (int, error) {
	valueInBytes := []byte(value)
	return b.BucketBytes(valueInBytes)
}

func (b XXHASHBucketizer) BucketInt(value int) (int, error) {
	valueInBytes := []byte(fmt.Sprintf("%d", value))
	return b.BucketBytes(valueInBytes)
}

func (b XXHASHBucketizer) BucketInt8(value int8) (int, error) {
	valueInBytes := []byte(fmt.Sprintf("%d", value))
	return b.BucketBytes(valueInBytes)
}

func (b XXHASHBucketizer) BucketInt16(value int16) (int, error) {
	valueInBytes := []byte(fmt.Sprintf("%d", value))
	return b.BucketBytes(valueInBytes)
}

func (b XXHASHBucketizer) BucketInt32(value int32) (int, error) {
	valueInBytes := []byte(fmt.Sprintf("%d", value))
	return b.BucketBytes(valueInBytes)
}

func (b XXHASHBucketizer) BucketInt64(value int64) (int, error) {
	valueInBytes := []byte(fmt.Sprintf("%d", value))
	return b.BucketBytes(valueInBytes)
}

func (b XXHASHBucketizer) BucketFloat64(value float64) (int, error) {
	valueInBytes := []byte(fmt.Sprintf("%g", value))
	return b.BucketBytes(valueInBytes)
}

func (b XXHASHBucketizer) BucketFloat32(value float32) (int, error) {
	valueInBytes := []byte(fmt.Sprintf("%g", value))
	return b.BucketBytes(valueInBytes)
}

func NewXXHASHBucketizer(seed string, buckets ...bucketizer.Bucket) XXHASHBucketizer {
	var sumOfWeights uint64 = 0
	var bucketRanges []uint64
	bucketRanges = append(bucketRanges, 0)
	for i, bucket := range buckets {
		sumOfWeights += uint64(bucket.Weight)
		startFromIndex := bucketRanges[i]
		bucketRanges = append(bucketRanges, startFromIndex+uint64(bucket.Weight))
	}
	return XXHASHBucketizer{
		Buckets:      buckets,
		weightSum:    sumOfWeights,
		bucketRanges: bucketRanges,
		seed:         []byte(seed),
	}
}
