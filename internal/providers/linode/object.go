package linode

import (
	"fmt"
	"os"

	"github.com/orme292/objectify"
	"s3p/internal/provider"
)

type LinodeObject struct {
	job *provider.Job

	f *os.File

	key string

	bucket string
}

func NewLinodeObject(job *provider.Job) provider.Object {
	return &LinodeObject{
		job: job,
	}
}

func (o *LinodeObject) Destroy() error {
	return o.Post()
}

func (o *LinodeObject) Generate() error {

	o.key = o.job.Key
	o.bucket = o.job.App.Bucket.Name
	return nil

}

func (o *LinodeObject) Post() error {
	return o.f.Close()
}

func (o *LinodeObject) Pre() error {

	o.job.Metadata.Update()

	if !o.job.Metadata.IsExists || !o.job.Metadata.IsReadable {
		return fmt.Errorf("file no longer accessible")
	}

	var target string
	if o.job.Metadata.Mode == objectify.EntModeLink {
		target = o.job.Metadata.TargetFinal
	} else {
		target = o.job.Metadata.FullPath()
	}

	f, err := os.Open(target)
	if err != nil {
		fmt.Printf("Error opening file %s: %s\n", o.job.Metadata.FullPath(), err)
		return err
	}

	o.f = f

	return nil

}
