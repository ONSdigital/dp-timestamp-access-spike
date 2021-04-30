dp-timestamp-access-spike
================
Investigation of using timestamp for accessing website content

The implementation of timestamp based access has been added to the `HelloHandler` in `api/hello.go`. This handler has been configured to handle all incoming requests. It maintains a map with each known URL as the key, and a value containing an array of versions for that URL.

There are 3 URL's that have been configured in the prototype code to represent 3 different scenarios: 

http://localhost:26000/publishedpage

A page with a single version that has a publish date in the past, so the content (`published page content`)should always be shown

http://localhost:26000/bulletin1

A page that has multiple versions. Two versions that have already been published, and a version that is due to be published. Out of the 3 versions, the middle version (`the latest bulletin1 published content`) should be shown to the user, as it's the most recent version that has been published.

http://localhost:26000/tobepublished

A page that has a single version with a publish time in the future. This scenario represents a new page due to be published, where no existing version has been published. In this case the content should not be shown, and a `content not found` response is returned.

### Getting started

* Run `make debug`

### Dependencies

* No further dependencies other than those defined in `go.mod`

### Configuration

| Environment variable         | Default   | Description
| ---------------------------- | --------- | -----------
| BIND_ADDR                    | :26000    | The host and port to bind to
| GRACEFUL_SHUTDOWN_TIMEOUT    | 5s        | The graceful shutdown timeout in seconds (`time.Duration` format)
| HEALTHCHECK_INTERVAL         | 30s       | Time between self-healthchecks (`time.Duration` format)
| HEALTHCHECK_CRITICAL_TIMEOUT | 90s       | Time to wait until an unhealthy dependent propagates its state to make this app unhealthy (`time.Duration` format)

### Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

### License

Copyright © 2021, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.

