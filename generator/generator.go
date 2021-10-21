package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/kelseyhightower/envconfig"
)

func NewGenerator(influx InfluxClient) *Generator {
	var config generatorConfig
	err := envconfig.Process("GENERATOR", &config)
	if err != nil {
		panic(err)
	}

	source := rand.NewSource(time.Now().UnixNano())

	return &Generator{
		random:    rand.New(source),
		influx:    influx,
		frequency: config.Frequency,
		min:       config.Min,
		max:       config.Max,
	}
}

type generatorConfig struct {
	Frequency time.Duration `default:"2s"`
	Min       int           `default:"0"`
	Max       int           `default:"10"`
}

type Generator struct {
	influx    InfluxClient
	frequency time.Duration
	min       int
	max       int
	random    *rand.Rand
}

func (g *Generator) Start() {
	for {
		r := g.random.Intn(g.max-g.min) + g.min
		log.Printf("Generating... %d", r)

		err := g.influx.Write(r)
		if err != nil {
			log.Println(err)
		}
		<-time.After(g.frequency)

	}
}

// I don't care about concurrent read/write in this simple toy program..

func (g *Generator) Set(min, max *int) error {
	newMin := g.min
	if min != nil {
		newMin = *min
	}

	newMax := g.max
	if max != nil {
		newMax = *max
	}

	if newMin >= newMax {
		return fmt.Errorf("Min has to be smaller than Max: %d, %d", newMin, newMax)
	}

	g.min = newMin
	g.max = newMax
	log.Printf("Generator settings: <%d,%d)", g.min, g.max)
	return nil
}
