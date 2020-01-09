package collection

import (
	"github.com/JonPeel/s3sync/pipeline"
	"github.com/JonPeel/s3sync/storage"
)

// UploadObjectData read objects from input, put its content and meta to Target storage and send object to next pipeline steps.
var UploadObjectData pipeline.StepFn = func(group *pipeline.Group, stepNum int, input <-chan *storage.Object, output chan<- *storage.Object, errChan chan<- error) {
	for obj := range input {
		err := group.Target.PutObject(obj)
		if err != nil {
			errChan <- &pipeline.ObjectError{Object: obj, Err: err}
		} else {
			output <- obj
		}
	}
}
