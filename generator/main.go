package main

func main() {
	influx := NewInfluxClient()

	generator := NewGenerator(influx)
	go generator.Start()

	server := NewServer(generator)
	server.Start()
}
