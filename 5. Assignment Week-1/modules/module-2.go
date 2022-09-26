package modules

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Description struct {
	Description string `json:"description"`
}

type Job struct {
	Title      string `json:"title"`
	WorkFrom   string `json:"work_from"`
	Department string `json:"department"`
}
type Aggr struct {
	Jobs []Job
	Description
}

func fetchData(url string, d interface{}) error {
	var err error
	var resp *http.Response
	var b []byte

	resp, err = http.Get(url)
	b, err = io.ReadAll(resp.Body)

	err = json.Unmarshal(b, d)

	return err
}

type Caching struct {
	Aggr
	IsUsed bool
}

func CalculateTime(start time.Time) {
	fmt.Println("dari calculate ", start)
	fmt.Printf("took %v\n", time.Since(start))
}

func (c *Caching) Aggregate() (Aggr, error) {
	defer CalculateTime(time.Now())

	c.IsUsed = !c.IsUsed
	var err error
	if c.IsUsed {
		descChan := make(chan Description)
		descErrChan := make(chan error)

		go func(descChandescChan chan Description, descErrChan chan error) {
			var d Description
			err := fetchData("https://workspace-rho.vercel.app/api/description", &d)

			descChan <- d
			descErrChan <- err
		}(descChan, descErrChan)

		jobsChan := make(chan []Job)
		jobsErrChan := make(chan error)
		go func(jobsChan chan []Job, jobsErrChan chan error) {
			var j []Job
			err := fetchData("https://workspace-rho.vercel.app/api/jobs", &j)

			jobsChan <- j
			jobsErrChan <- err
		}(jobsChan, jobsErrChan)

		desc := <-descChan
		descErr := <-descErrChan
		jobs := <-jobsChan
		jobsErr := <-jobsErrChan

		if descErr != nil {
			err = descErr
		}
		if jobsErr != nil {
			err = jobsErr
		}

		c.Aggr = Aggr{
			Description: desc,
			Jobs:        jobs,
		}
	}
	return c.Aggr, err
}

func (c *Caching) AggregateSec() (Aggr, error) {
	defer CalculateTime(time.Now())

	c.IsUsed = !c.IsUsed

	var err error
	if c.IsUsed {
		var d Description
		err = fetchData("https://workspace-rho.vercel.app/api/description", &d)

		var j []Job
		err = fetchData("https://workspace-rho.vercel.app/api/jobs", &j)

		c.Aggr = Aggr{
			Description: d,
			Jobs:        j,
		}
	}
	return c.Aggr, err
}

// For checking input time to terminal console (for debugging)
func Aggregator() {
	var c Caching
	aggr, err := c.AggregateSec()
	if err != nil {
		panic(err)
	}

	aggr, err = c.Aggregate()
	if err != nil {
		panic(err)
	}

	aggr, err = c.Aggregate()
	if err != nil {
		panic(err)
	}

	aggr, err = c.Aggregate()
	if err != nil {
		panic(err)
	}

	_ = aggr
}
