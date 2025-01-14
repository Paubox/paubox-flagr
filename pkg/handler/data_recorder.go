package handler

import (
	"sync"

	"github.com/paubox/paubox-flagr/pkg/config"
	"github.com/paubox/paubox-flagr/swagger_gen/models"
)

var (
	singletonDataRecorder     DataRecorder
	singletonDataRecorderOnce sync.Once
)

// DataRecorder can record and produce the evaluation result
type DataRecorder interface {
	AsyncRecord(models.EvalResult)
	NewDataRecordFrame(models.EvalResult) DataRecordFrame
}

// GetDataRecorder gets the data recorder
func GetDataRecorder() DataRecorder {
	singletonDataRecorderOnce.Do(func() {
		recorderType := config.Config.RecorderType
		switch recorderType {
		case "kafka":
			singletonDataRecorder = NewKafkaRecorder()
		case "kinesis":
			singletonDataRecorder = NewKinesisRecorder()
		case "pubsub":
			singletonDataRecorder = NewPubsubRecorder()
		default:
			panic("recorderType not supported")
		}
	})

	return singletonDataRecorder
}
