package xxhash

import (
	"github.com/cespare/xxhash"
	"github.com/ehsangolshani/bucketizer"
	"github.com/pkg/errors"
)

type xxhashBucketizer struct {
	Buckets      []bucketizer.Bucket
	weightSum    uint64
	bucketRanges []uint64
}

func (b xxhashBucketizer) BucketString(value string) (int, error) {
	hashDigest := xxhash.Sum64String(value)
	reminder := hashDigest % b.weightSum
	n := len(b.bucketRanges)
	for i := 0; i < n-1; i++ {
		if reminder >= b.bucketRanges[i] && reminder < b.bucketRanges[i+1] {
			return i, nil
		}
	}
	return 0, errors.New("invalid reminder value")
}

func (b xxhashBucketizer) BucketBytes(value []byte) (int, error) {
	hashDigest := xxhash.Sum64(value)
	reminder := hashDigest % b.weightSum
	n := len(b.bucketRanges)
	for i := 0; i < n-1; i++ {
		if reminder >= b.bucketRanges[i] && reminder >= b.bucketRanges[i+1] {
			return i, nil
		}
	}
	return 0, errors.New("invalid reminder value")
}

func (b xxhashBucketizer) BucketInt(value int) (int, error) {

	return 0, nil
}

func (b xxhashBucketizer) BucketInt8(value int8) (int, error) {

	return 0, nil
}

func (b xxhashBucketizer) BucketInt16(value int16) (int, error) {

	return 0, nil
}

func (b xxhashBucketizer) BucketInt32(value int32) (int, error) {

	return 0, nil
}

func (b xxhashBucketizer) BucketInt64(value int64) (int, error) {

	return 0, nil
}

func (b xxhashBucketizer) BucketFloat64(value float64) (int, error) {

	return 0, nil
}

func (b xxhashBucketizer) BucketFloat32(value float32) (int, error) {

	return 0, nil
}

func (b xxhashBucketizer) BucketInterface(value interface{}) (int, error) {

	return 0, nil
}

func NewXXHASHBucketizer(buckets ...bucketizer.Bucket) xxhashBucketizer {
	var sumOfWeights uint64 = 0
	var bucketRanges []uint64
	bucketRanges = append(bucketRanges, 0)
	for i, bucket := range buckets {
		sumOfWeights += uint64(bucket.Weight)
		startFromIndex := bucketRanges[i]
		bucketRanges = append(bucketRanges, startFromIndex+uint64(bucket.Weight))
	}
	return xxhashBucketizer{
		Buckets:      buckets,
		weightSum:    sumOfWeights,
		bucketRanges: bucketRanges,
	}
}
