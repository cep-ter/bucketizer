package bucketizer

type Bucketizer interface {
	BucketString(value string) (int, error)
	BucketBytes(value []byte) (int, error)
	BucketInt(value int) (int, error)
	BucketInt8(value int8) (int, error)
	BucketInt16(value int16) (int, error)
	BucketInt32(value int32) (int, error)
	BucketInt64(value int64) (int, error)
	BucketFloat64(value float64) (int, error)
	BucketFloat32(value float32) (int, error)
	BucketInterface(value interface{}) (int, error)
}
