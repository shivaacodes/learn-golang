package main

import (
	imageprocessing "fan_out_fan_in/image_processing"
	"image"
	"strings"
)

type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

func loadImage(paths []string) []Job {
	var jobs []Job
	for _, p := range paths {
		job := Job{
			InputPath: p,
			OutPath:   strings.Replace(p, "images/", "images/output/", 1),
		}
		job.Image = imageprocessing.ReadImage(p)
		jobs = append(jobs, job)
	}
	return jobs
}

func resize(jobs *[]Job) <-chan Job {
	out := make(chan Job, len(*jobs))
	for _, job := range *jobs { // loop over jobs
		go func(job Job) { // create goroutines for each job
			job.Image = imageprocessing.Resize(job.Image) // in each goroutine, we resize the images into job.Image
			out <- job                                    // now place this job on channel out
		}(job) // pass job in the goroutine
	}
	return out
}

// fan in logic to collect the jobs back from teh channel
func collectJobs(input <-chan Job, imageCnt int) []Job {
	var resizedJobs []Job
	for i := 0; i < imageCnt; i++ {
		job := <-input
		resizedJobs = append(resizedJobs, job)
	}
	return resizedJobs
}

func saveImages(jobs *[]Job) {
	for _, job := range *jobs {
		imageprocessing.WriteImage(job.OutPath, job.Image)
	}
}

func main() {
	imagePaths := []string{
		"images/image1.jpeg",
		"images/image2.jpeg",
		"images/image3.jpeg",
		"images/image4.jpeg",
	}

	jobs := loadImage(imagePaths)
	// Fan out this function to multiple goroutines
	out := resize(&jobs)
	// Collect / fan-in
	resizedJobs := collectJobs(out, len(imagePaths))
	saveImages(&resizedJobs)
}
