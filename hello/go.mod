module example/hello

go 1.23.1

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.2
)

require (
	golang.org/x/exp v0.0.0-20241009180824-f66d83c29e7c // indirect
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)

replace example.com/greetings => ../greetings
