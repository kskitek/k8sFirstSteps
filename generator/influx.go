package main

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/kelseyhightower/envconfig"
)

func NewInfluxClient() InfluxClient {
	var config influxConfig
	err := envconfig.Process("INFLUX", &config)
	if err != nil {
		panic(err)
	}

	return InfluxClient{
		config: config,
	}
}

// usually such defaults are harmful but in this case will simplufy setup

type influxConfig struct {
	Url          string `default:"http://localhost:8086"`
	Organization string `default:"illuminati"`
	Bucket       string `default:"leakybucket"`
	Token        string `default:"supersecrettoken"`
}

type InfluxClient struct {
	config influxConfig
}

func (i InfluxClient) Write(v int) error {
	client := influxdb2.NewClient(i.config.Url, i.config.Token)
	writeAPI := client.WriteAPIBlocking(i.config.Organization, i.config.Bucket)
	p := influxdb2.NewPoint("rand", nil, map[string]interface{}{"value": float64(v)}, time.Now())

	return writeAPI.WritePoint(context.Background(), p)
}
