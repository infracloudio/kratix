package writers

type StateStoreWriter interface {
	WriteObject(path string, objectName string, toWrite []byte) error
	RemoveObject(bucketName string, objectName string) error
}

const (
	S3  string = "s3"
	Git string = "git"
)
