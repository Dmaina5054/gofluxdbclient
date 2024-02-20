package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

// defining task type
const (
	TypeFluxdbFetch = "fluxdb:fetchrecords"
)

// task payload for any flux fetch related task
type FluxdbFetchPayload struct {
	//bucket name
	BucketName        string
	DestinationBucket string
}

// define function to start scheduler
func initScheduler() {
	//defining scheduler
	scheduler := asynq.NewScheduler(
		asynq.RedisClientOpt{Addr: ":6379"},
		&asynq.SchedulerOpts{
			Location: time.Local,
			LogLevel: asynq.DebugLevel,
		},
	)

	//define a function to schedule task with a given payload
	scheduleFluxdbFetchTask := func(bucketName, destinationBucket string) {
		payload, err := json.Marshal(FluxdbFetchPayload{BucketName: bucketName, DestinationBucket: destinationBucket})
		if err != nil {
			log.Fatal(err)
		}

		//schedule the task
		if _, err := scheduler.Register("*/5 * * * *", asynq.NewTask(TypeFluxdbFetch, payload)); err != nil {
			log.Fatal(err)
		}

	}

	//TODO: Spawn to container per region
	//schedule tasks with different payloads
	scheduleFluxdbFetchTask("MWKs", "MWKsDownsampled")
	scheduleFluxdbFetchTask("MWKn", "MWKnDownsampled")
	scheduleFluxdbFetchTask("KSNOnu", "KSNOnuDownsampled")
	scheduleFluxdbFetchTask("KWDOnu", "KWNOnuDownsampled")
	scheduleFluxdbFetchTask("STNOnu", "STNOnuDownsampled")
	scheduleFluxdbFetchTask("HTROnu", "HTROnuDownsampled")
	scheduleFluxdbFetchTask("LSMOnu", "LSMOnuDownsampled")
	scheduleFluxdbFetchTask("KRBSOnu", "KRBSOnuDownsampled")
	scheduleFluxdbFetchTask("UMJ3", "UMJ3Downsampled")

	// if no error
	// run scheduler with defined cron
	if err := scheduler.Run(); err != nil {
		log.Fatal(err)
	}

	log.Println("Scheduler running....")

}

func main() {

	initScheduler()

}
