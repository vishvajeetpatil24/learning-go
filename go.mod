module topProject

go 1.17

replace github.com/myorg/project => ./project

replace bytedetective.net/workers => ./workers

require github.com/myorg/project v0.0.0-00010101000000-000000000000

require rsc.io/quote v1.5.2

require (
	bytedetective.net/workers v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.3.0
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)
