# _MessageBird_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)
[![Build Status](https://travis-ci.com/heaptracetechnology/microservice-messagebird.svg?branch=master)](https://travis-ci.com/heaptracetechnology/microservice-messagebird)
[![codecov](https://codecov.io/gh/heaptracetechnology/microservice-messagebird/branch/master/graph/badge.svg)](https://codecov.io/gh/heaptracetechnology/microservice-messagebird)

An OMG service for MessageBird, it is a cloud communications platform that connects enterprises to their global customers.

## Direct usage in [Storyscript](https://storyscript.io/):

##### Send Message
```coffee
>>> messagebird send from:'senderNumber' to:'receiverNumber' message:'testMessage'
{"ID": "operationID","Type": "sms","Originator": "+917507704328","Body": "messageBody","CreatedDatetime": "2019-04-15T13:58:50Z","Recipients": {"recipientsDetails"}}
```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)

##### Send Message
```shell
$ omg run send -a from=<PHONE_NUMBER> -a to=<PHONE_NUMBER> -a message=<MESSAGE_TEXT_BODY> -e API_KEY=<API_KEY>
```
NOTE : Enter PHONE_NUMBER with country code example(INDIA) : "+919123456789"

**Note**: the OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/omg-services/messagebird/blob/master/LICENSE).
