UTM Stripper Microservice (Go)
================================

Very simple microservice written in Go for removing URL variables from URLs.

To use the url utm stripper, run this command with a URL ENCODED url:
`./stripper http://domain.com/url/path?utm_campaign=fo%26utm_feed=bar`

----

Build binary with all dependencies:
`CGO_ENABLED=0 go build -a -ldflags '-s' stripper.go`

Verify with (should show “not a dynamic executable”):
`ldd stripper`

Building Dockerfile:
`docker build -t cloudspace/utm-stripper .`

Running Dockerfile:
`docker run -ti cloudspace/utm-stripper http://domain.com/url/path?utm_campaign=foo%26utm_feed=bar`

Finding docker images:
`docker images -a`

Finding docker containers:
`docker ps -a`

Removing a container:
`docker rm containerid`
