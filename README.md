# MessageBird as a microservice
An OMG service for MessageBird, it is a cloud communications platform that connects enterprises to their global customers.

[![Open Microservice Guide](https://img.shields.io/badge/OMG-enabled-brightgreen.svg?style=for-the-badge)](https://microservice.guide)
<!-- [![Build Status](https://travis-ci.com/heaptracetechnology/microservice-intercom.svg?branch=master)](https://travis-ci.com/heaptracetechnology/microservice-intercom)
[![codecov](https://codecov.io/gh/heaptracetechnology/microservice-intercom/branch/master/graph/badge.svg)](https://codecov.io/gh/heaptracetechnology/microservice-intercom)
[![GolangCI](https://golangci.com/badges/github.com/golangci/golangci-web.svg)](https://golangci.com) -->


## [OMG](hhttps://microservice.guide) CLI

### OMG

* omg validate
```
omg validate
```
* omg build
```
omg build
```
### Test Service

* Test the service by following OMG commands

### CLI

##### Create and save user
```sh
$ omg run send -a from=<PHONE_NUMBER> -a to=<PHONE_NUMBER> -a message=<MESSAGE_TEXT_BODY> -e API_KEY=<API_KEY>
```
NOTE : Enter PHONE_NUMBER with country code example(INDIA) : "+910123456789"
## License
### [MIT](https://choosealicense.com/licenses/mit/)

## Docker
### Build
```
docker build -t microservice-messagebird .
```
### RUN
```
docker run -p 3000:3000 microservice-messagebird
```
