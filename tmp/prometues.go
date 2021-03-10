// package main

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"strings"
// 	"time"

// 	prometheusapi "github.com/prometheus/client_golang/api"
// 	prometheusv1 "github.com/prometheus/client_golang/api/prometheus/v1"
// 	prommodel "github.com/prometheus/common/model"
// )

// // Metric represents a measured value in time.
// type Metric struct {
// 	Value float64
// 	TS    time.Time
// }

// // MetricSeries is a group of metrics identified by an ID and a context
// // information.
// type MetricSeries struct {
// 	ID      string
// 	Labels  map[string]string
// 	Metrics []Metric
// }

// // Query is the query that will be made to the datasource.
// type Query struct {
// 	Expr string `json:"expr,omitempty"`
// 	// Legend accepts `text.template` format.
// 	Legend       string `json:"legend,omitempty"`
// 	DatasourceID string `json:"datasourceID,omitempty"`
// }
// type Gatherer interface {
// 	// GatherSingle gathers one single metric at a point in time.
// 	GatherSingle(ctx context.Context, query Query, t time.Time) ([]MetricSeries, error)
// 	// GatherRange gathers multiple metrics based on a start and an end using a step duration
// 	// to know how many metrics needs to gather.
// 	// The returned metrics on the series should be ordered.
// 	GatherRange(ctx context.Context, query Query, start, end time.Time, step time.Duration) ([]MetricSeries, error)
// }

// // ConfigGatherer is the configuration of the Prometheus gatherer.
// type ConfigGatherer struct {
// 	// Client is the prometheus API client.
// 	Client prometheusv1.API
// 	// FilterSpecialLabels will return the metrics with the special labels filtered.
// 	// The special labels start with `__`, examples: `__name__`, `__scheme__`.
// 	FilterSpecialLabels bool
// }

// type gatherer struct {
// 	cli prometheusv1.API
// 	cfg ConfigGatherer
// }

// // NewGatherer returns a new metric gatherer for prometheus backends.
// func NewGatherer(cfg ConfigGatherer) Gatherer {
// 	return &gatherer{
// 		cli: cfg.Client,
// 		cfg: cfg,
// 	}
// }

// func (g *gatherer) GatherSingle(ctx context.Context, query Query, t time.Time) ([]MetricSeries, error) {
// 	// Get value from Prometheus.
// 	val, _, err := g.cli.Query(ctx, query.Expr, t)
// 	if err != nil {
// 		return []MetricSeries{}, err
// 	}

// 	// Translate prom values to domain.
// 	res, err := g.promToModel(val)
// 	if err != nil {
// 		return []MetricSeries{}, err
// 	}

// 	return res, nil
// }

// func (g *gatherer) GatherRange(ctx context.Context, query Query, start, end time.Time, step time.Duration) ([]MetricSeries, error) {
// 	// Get value from Prometheus.
// 	val, _, err := g.cli.QueryRange(ctx, query.Expr, prometheusv1.Range{
// 		Start: start,
// 		End:   end,
// 		Step:  step,
// 	})
// 	if err != nil {
// 		return []MetricSeries{}, err
// 	}

// 	// Translate prom values to domain.
// 	res, err := g.promToModel(val)
// 	if err != nil {
// 		return []MetricSeries{}, err
// 	}

// 	return res, nil
// }

// // promToModel converts a prometheus result metric to a domain model one.
// func (g *gatherer) promToModel(pm prommodel.Value) ([]MetricSeries, error) {
// 	res := []MetricSeries{}

// 	switch pm.Type() {
// 	case prommodel.ValScalar:
// 		scalar := pm.(*prommodel.Scalar)
// 		res = g.transformScalar(scalar)
// 	case prommodel.ValVector:
// 		vector := pm.(prommodel.Vector)
// 		res = g.transformVector(vector)
// 	case prommodel.ValMatrix:
// 		matrix := pm.(prommodel.Matrix)
// 		res = g.transformMatrix(matrix)
// 	default:
// 		return res, errors.New("prometheus value type not supported")
// 	}

// 	return res, nil
// }

// // transformScalar will get a prometheus Scalar and transform to a domain model
// // MetricSeries slice.
// func (g *gatherer) transformScalar(scalar *prommodel.Scalar) []MetricSeries {
// 	res := []MetricSeries{}

// 	m := Metric{
// 		TS:    scalar.Timestamp.Time(),
// 		Value: float64(scalar.Value),
// 	}
// 	res = append(res, MetricSeries{
// 		Metrics: []Metric{m},
// 	})

// 	return res
// }

// // transformVector will get a prometheus Vector and transform to a domain model
// // MetricSeries slice.
// // A Prometheus vector is an slice of metrics (group of labels) that have one
// // sample only (all samples from all metrics have the same timestamp)
// func (g *gatherer) transformVector(vector prommodel.Vector) []MetricSeries {
// 	res := []MetricSeries{}

// 	for _, sample := range vector {
// 		id := sample.Metric.String()
// 		labels := g.labelSetToMap(prommodel.LabelSet(sample.Metric))
// 		series := MetricSeries{
// 			ID:     id,
// 			Labels: g.sanitizeLabels(labels),
// 		}

// 		// Add the metric to the series.
// 		series.Metrics = append(series.Metrics, Metric{
// 			TS:    sample.Timestamp.Time(),
// 			Value: float64(sample.Value),
// 		})

// 		res = append(res, series)
// 	}

// 	return res
// }

// // transformMatrix will get a prometheus Matrix and transform to a domain model
// // MetricSeries slice.
// // A Prometheus Matrix is an slices of metrics (group of labels) that have multiple
// // samples (in a slice of samples).
// func (g *gatherer) transformMatrix(matrix prommodel.Matrix) []MetricSeries {
// 	res := []MetricSeries{}

// 	// Use a map to index the different series based on labels.
// 	for _, sampleStream := range matrix {
// 		id := sampleStream.Metric.String()
// 		labels := g.labelSetToMap(prommodel.LabelSet(sampleStream.Metric))
// 		series := MetricSeries{
// 			ID:     id,
// 			Labels: g.sanitizeLabels(labels),
// 		}

// 		// Add the metric to the series.
// 		for _, sample := range sampleStream.Values {
// 			series.Metrics = append(series.Metrics, Metric{
// 				TS:    sample.Timestamp.Time(),
// 				Value: float64(sample.Value),
// 			})
// 		}

// 		res = append(res, series)
// 	}

// 	return res
// }

// func (g *gatherer) labelSetToMap(ls prommodel.LabelSet) map[string]string {
// 	res := map[string]string{}
// 	for k, v := range ls {
// 		res[string(k)] = string(v)
// 	}

// 	return res
// }

// // sanitizeLabels will sanitize the map label values.
// // 	- Remove special labels if required (start with `__`).
// func (g *gatherer) sanitizeLabels(m map[string]string) map[string]string {

// 	// Filter if required.
// 	if !g.cfg.FilterSpecialLabels {
// 		return m
// 	}

// 	res := map[string]string{}
// 	for k, v := range m {
// 		if strings.HasPrefix(k, "__") {
// 			continue
// 		}
// 		res[k] = v
// 	}

// 	return res
// }

// func main() {

// 	// read data from prometheus
// 	cli, err := prometheusapi.NewClient(prometheusapi.Config{
// 		Address: "http://aol.fire-shahei-01.xyz:9090",
// 	})
// 	if err != nil {
// 		fmt.Println("read data error ")
// 	}
// 	g := NewGatherer(ConfigGatherer{
// 		Client: prometheusv1.NewAPI(cli),
// 	})
// 	q := &Query{
// 		Expr:         "go_goroutines",
// 		Legend:       "test",
// 		DatasourceID: "prometheus",
// 	}
// 	ctx := context.Background()
// 	met, err := g.GatherSingle(ctx, *q, time.Now().UTC())
// 	if err != nil {
// 		fmt.Println("----------", err)
// 	}
// 	for _, m := range met {
// 		fmt.Println(m.Metrics)
// 	}

// }
