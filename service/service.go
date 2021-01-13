package service

import (
	"ResourceServer/resource"
	"ResourceServer/resource/retry"
	"context"
	"errors"
	"golang.org/x/sync/semaphore"
	"log"
	"os"
	"time"
)

type Service interface {
	Load(ctx context.Context, link string) (string, error)
	LoadBatch(ctx context.Context, link []string) (map[string]string, error)
}

type ResourceService struct {
	filepath string
	sem      *semaphore.Weighted
}

// Load function, upload resource by given string and return
// name of uploaded resource. If there was an error
// retry function tries to upload anyway. Mat attempts 4.
func (r *ResourceService) Load(ctx context.Context, link string) (string, error) {
	var (
		filename string
		err      error
		options  = []retry.Option{
			retry.WithLinearFunc(),
			retry.WithMaxRetryTime(time.Second * 1),
			retry.WithRetryIntensity(time.Millisecond * 100),
			retry.WithMaxAttempts(4),
		}
	)
	err = retry.Go(func() error {
		filename, err = resource.Load(r.filepath, link)
		if err != nil {
			return err
		}

		return nil
	}, options...)

	if err != nil {
		return "", err
	}

	return filename, nil
}

// LoadBatch function loads an slice of urls with concurrency
func (r *ResourceService) LoadBatch(ctx context.Context, link []string) (map[string]string, error) {
	if len(link) == 0 {
		return nil, errors.New("empty links")
	}
	done := make(chan struct{}, 1)
	result := make(map[string]string)

	for _, re := range link {
		if err := r.sem.Acquire(ctx, 1); err != nil {
			log.Printf("failed to acquire semaphore")
			return nil, err
		}

		go func(next string) {
			defer r.sem.Release(1)
			filename, err := r.Load(ctx, next)
			if err != nil {
				result[next] = "none"
			} else {
				result[next] = filename
			}
			if len(result) >= len(link) {
				done <- struct{}{}
			}
		}(re)
	}

	<-done

	return result, nil
}

// Create new *ResourceService
func New(resourcepath string, parallel ...int) *ResourceService {
	l := 1
	if len(parallel) > 0 {
		l = parallel[0]
	}

	err := os.MkdirAll(resourcepath, 0777)
	if err != nil {
		log.Fatal(err)
	}

	return &ResourceService{
		filepath: resourcepath,
		sem:      semaphore.NewWeighted(int64(l)),
	}
}
