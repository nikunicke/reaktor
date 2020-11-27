package warehouse

import (
	"fmt"
)

// Defined job types
const (
	JobTypeGetProducts                  = string("get_products")
	JobTypeGetAvailability              = string("get_availability")
	JobTypeMergeProductsAndAvailability = string("merge_products_and_availability")
)

// Job errors
const (
	ErrorInvalidJobType = Error("Invalid job type")
)

type Job struct {
	Type string
}

type JobExecutor struct {
	ProductService      ProductService
	AvailabilityService AvailabilityService
}

func (e *JobExecutor) ExecuteJob(job *Job) error {
	switch job.Type {
	case JobTypeGetProducts:
		return e.jobGetProducts(job)
	case JobTypeGetAvailability:
		return e.jobGetAvailability(job)
	case JobTypeMergeProductsAndAvailability:
		return e.jobMergeProductsAndAvailability(job)
	}
	return ErrorInvalidJobType
}

func (e *JobExecutor) jobGetProducts(job *Job) error {
	c1 := make(chan Products, 3)
	for _, category := range ProductCategories {
		fmt.Println("GET:", category)
		go func(ctg string) error {
			products, err := e.ProductService.GetProducts(ctg)
			if err != nil {
				return err
			}
			c1 <- products
			return nil
		}(category)
	}
	data := <-c1
	fmt.Printf("%v\n", data[0])
	return nil
}

func (e *JobExecutor) jobGetAvailability(job *Job) error {
	c1 := make(chan Availability)
	go func() error {
		availability, err := e.AvailabilityService.GetAvailability("availability/derp")
		if err != nil {
			return nil
		}
		c1 <- *availability
		return nil
	}()

	go func() error {
		availability, err := e.AvailabilityService.GetAvailability("availability/reps")
		if err != nil {
			return nil
		}
		c1 <- *availability
		return nil
	}()
	for i := 0; i < 2; i++ {
		select {
		case data := <-c1:
			fmt.Println(data.Code)
		}
	}
	return nil
}

func (e *JobExecutor) jobMergeProductsAndAvailability(job *Job) error {
	return nil
}
