// Package collection contains different StepFn functions to do different pipeline actions.
package collection

import (
	"github.com/JonPeel/s3sync/pipeline"
	"github.com/JonPeel/s3sync/storage"
)

// LoadObjectMeta accepts an input object and downloads its metadata.
var LoadObjectMeta pipeline.StepFn = func(group *pipeline.Group, stepNum int, input <-chan *storage.Object, output chan<- *storage.Object, errChan chan<- error) {
	for obj := range input {
		err := group.Source.GetObjectMeta(obj)
		if err != nil {
			errChan <- &pipeline.ObjectError{Object: obj, Err: err}
		} else {
			output <- obj
		}
	}
}


// LoadObjectMeta accepts an input object and downloads its metadata.
var LoadObjectLocalMeta pipeline.StepFn = func(group *pipeline.Group, stepNum int, input <-chan *storage.Object, output chan<- *storage.Object, errChan chan<- error) {
	for obj := range input {
		err := group.Source.GetObjectLocalMeta(obj)
		if err != nil {
			errChan <- &pipeline.ObjectError{Object: obj, Err: err}
		} else {
			output <- obj
		}
	}
}

// LoadObjectData accepts an input object and downloads its content and metadata.
var LoadObjectData pipeline.StepFn = func(group *pipeline.Group, stepNum int, input <-chan *storage.Object, output chan<- *storage.Object, errChan chan<- error) {
	for obj := range input {
		err := group.Source.GetObjectContent(obj)
		if err != nil {
			errChan <- &pipeline.ObjectError{Object: obj, Err: err}
		} else {
			output <- obj
		}
	}
}
