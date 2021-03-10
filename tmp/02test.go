package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

// Metric represents a measured value in time.
type Metric struct {
	Value float64
	TS    time.Time
}

// MetricSeries is a group of metrics identified by an ID and a context
// information.
type MetricSeries struct {
	ID      string
	Labels  map[string]string
	Metrics []Metric
}

// test time parse
func main() {
	// t1 := "2020-07-30 15:01:16"
	// t, err := time.Parse("2006-01-02 15:04:05", t1)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(t)
	g := &gatherer{
		path: "./data.txt",
	}
	g.GatherRange()
}

// ConfigGatherer is the configuration of the osw gatherer.
type ConfigGatherer struct {
	Path string
}

// type oswmetrics struct {
// 	name  string
// 	value []string
// }

func (c *ConfigGatherer) defaults() error {
	var err error

	c.Path = "./date.txt"
	return err
}

type gatherer struct {
	path string
	cfg  ConfigGatherer
}

func (g *gatherer) GatherRange() ([]MetricSeries, error) {
	res := []MetricSeries{}

	resp := g.GetMetric()

	// Build the metric series
	for name, result := range resp {
		metrics := []Metric{}
		for _, seria := range result {
			t, err := time.Parse("2006-01-02 15:04:05", seria[0])
			if err != nil {
				fmt.Println("time parse err")
			}

			v, _ := strconv.ParseFloat(seria[1], 10)
			metrics = append(metrics, Metric{
				TS:    t,
				Value: v,
			})
		}
		label := name
		//TODO(rochaporto): Allow alias based on a tag value
		//if serie.Tags != nil {
		//	if v1, ok := serie.Tags["version"]; ok {
		//		label = v1
		//	}
		//}
		res = append(res, MetricSeries{ID: label, Metrics: metrics})
	}
	fmt.Println(res[0].ID)
	return res, nil
}

// GetMetric is the gatherer to collect data
func (g *gatherer) GetMetric() map[string][][]string {
	f, err := os.Open(g.path)
	if err != nil {
		fmt.Errorf("file open err", err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	res := map[string][][]string{}
	var lim [][]string
	var proc [][]string
	for {
		line, err := buf.ReadString('\n')
		if err == io.EOF {
			fmt.Println("file read over!")
			break
		}
		if err != nil {
			fmt.Println("file read err.", err)
		}
		sl := strings.Split(line, "|")

		lim = append(lim, []string{strings.TrimSpace(sl[0]), strings.TrimSpace(sl[1])})

		proc = append(proc, []string{strings.TrimSpace(sl[0]), strings.TrimSpace(sl[2])})
		fmt.Println(proc)

	}
	res["lim"] = lim
	res["proc"] = proc
	return res
}
