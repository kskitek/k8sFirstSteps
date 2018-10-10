package influx

import (
	"time"

	influx "github.com/influxdata/influxdb/client/v2"
	"github.com/kskitek/k8sFirstSteps/value"
	log "github.com/sirupsen/logrus"
)

type influxSaver struct {
	value  int
	client influx.Client
	db     string
}

func New(address, user, passwd, db string, initialValue int) (value.Saver, error) {
	c, err := influx.NewHTTPClient(influx.HTTPConfig{
		Addr:     address,
		Username: user,
		Password: passwd,
		Timeout:  time.Second * 1,
	})
	if err != nil {
		return nil, err
	}

	saver := &influxSaver{
		value:  initialValue,
		client: c,
		db:     db,
	}
	return saver, nil
}

func (is *influxSaver) Save(value int) {
	p, err := intToPoint(value)
	if err != nil {
		log.Error(err)
		return
	}

	bp, err := pointToBatch(p, is.db)
	if err != nil {
		log.Error(err)
		return
	}

	if err := is.client.Write(bp); err != nil {
		log.Error(err)
	}
}

func intToPoint(value int) (*influx.Point, error) {
	tags := map[string]string{
		"userValue": "1",
	}
	fields := map[string]interface{}{
		"value": value,
	}
	return influx.NewPoint("mes1", tags, fields, time.Now())
}

func pointToBatch(p *influx.Point, db string) (influx.BatchPoints, error) {
	bp, err := influx.NewBatchPoints(influx.BatchPointsConfig{
		Database:        db,
		Precision:       "s",
		RetentionPolicy: "autogen",
	})
	if err != nil {
		return nil, err
	}

	bp.AddPoint(p)
	return bp, err
}
