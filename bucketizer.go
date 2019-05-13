package bucketizer

type Bucketizer interface {
	BucketString(value string) int
	BucketBytes(value []byte) int
	BucketInt(value int) int
	BucketInt8(value int8) int
	BucketInt16(value int16) int
	BucketInt32(value int32) int
	BucketInt64(value int64) int
	BucketFloat64(value float64) int
	BucketFloat32(value float32) int
	BucketInterface(value interface{}) int
}
