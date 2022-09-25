package modules

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Desc struct {
	Description string `json:"description"`
}

type Job struct {
	Title      string `json:"title"`
	WorkFrom   string `json:"work_from"`
	Department string `json:"department"`
}

type Aggr struct {
	Jobs []Job
	Desc
}

func fetch(url string, d interface{}) error {
	var err error
	var resp *http.Response
	var b []byte

	resp, err = http.Get(url)
	b, err = io.ReadAll(resp.Body)

	err = json.Unmarshal(b, d)

	return err
}

type Cache struct {
	Aggr
	IsUsed bool
}

func (c *Cache) Aggregate() (Aggr, error) {
	defer CaculateTime(time.Now())

	c.IsUsed = !c.IsUsed

	var err error
	if c.IsUsed {
		descChan := make(chan Desc)
		descErrChan := make(chan error)

		go func(descChan chan Desc, descErrChan chan error) {
			var d Desc
			err := fetch("https://workspace-rho.vercel.app/api/description", &d)

			descChan <- d
			descErrChan <- err
		}(descChan, descErrChan)

		jobsChan := make(chan []Job)
		jobsErrChan := make(chan error)
		go func(jobsChan chan []Job, jobsErrChan chan error) {
			var j []Job
			err := fetch("https://workspace-rho.vercel.app/api/jobs", &j)

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
			Desc: desc,
			Jobs: jobs,
		}
	}

	return c.Aggr, err
}

func (c *Cache) AggregateSec() (Aggr, error) {
	defer CaculateTime(time.Now())

	c.IsUsed = !c.IsUsed

	var err error
	if c.IsUsed {

		var d Desc
		err = fetch("https://workspace-rho.vercel.app/api/description", &d)

		var j []Job
		err = fetch("https://workspace-rho.vercel.app/api/jobs", &j)

		c.Aggr = Aggr{
			Desc: d,
			Jobs: j,
		}
	}

	return c.Aggr, err
}

func Aggregator() {
	var c Cache
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

func CaculateTime(start time.Time) {
	fmt.Println("dari calculate ", start)
	fmt.Printf("took %v\n", time.Since(start))
}
