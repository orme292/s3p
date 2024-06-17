package provider_v2

type Supports struct {
	BucketCreate          bool
	BucketDelete          bool
	ObjectDelete          bool
	ObjectUploadMultipart bool
}

func NewSupports(bucketCreate, bucketDelete, objectDelete, objectUploadMultipart bool) *Supports {
	return &Supports{
		BucketCreate:          bucketCreate,
		BucketDelete:          bucketDelete,
		ObjectDelete:          objectDelete,
		ObjectUploadMultipart: objectUploadMultipart,
	}
}
