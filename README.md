[![Build Status](https://travis-ci.org/blacksails/mailspree.svg?branch=master)](https://travis-ci.org/blacksails/mailspree)
[![Coverage Status](https://coveralls.io/repos/github/blacksails/mailspree/badge.svg?branch=master)](https://coveralls.io/github/blacksails/mailspree?branch=master)
[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](http://godoc.org/github.com/blacksails/mailspree) 
[![Dockerhub](https://img.shields.io/badge/dockerhub-repo-blue.svg)](https://hub.docker.com/r/blacksails/mailspree)
# Mailspree

Mailspree is a mailing service which delegates mailsending to different mail
service providers. It aims to be fault tolerant, failing over to another mail
service provider if there was a problem with the one that tried to send the
email.

## Architeture

Mailspree is written in go which is great for writing services. Mailspree is
exposed as a http api, with token authentication. It currently supports
[mailgun](http://www.mailgun.com/) and [sendgrid](https://sendgrid.com/). It
has a simple circuit breaker, which prevents a mail service from being used for
some time, if it has been having problems.

## Installation

The easiest way to run the mailspree service is using docker. You can either
grab the latest image from dockerhub
[blacksails/mailspree](https://hub.docker.com/r/blacksails/mailspree), or
build the image from source using the following instructions:

```bash
git clone git@github.com:blacksails/mailspree.git
cd mailspree-ui
npm install
npm run build
docker build -t blacksails/mailspree .
```

Instead of using `docker run` you can save a create an environment varibale
file in the directory, called `mailspree-vars.env`, which contains the
environment variable. See the
[mailspree-vars.env.example](mailspree-vars.env.example) file for an
example. Then run the mailspree container with docker compose.

```base
docker-compose up
```

## Caveats

Currently the api allows that the consumer supplies the from email, which can
be problematic due to
[SPF](https://en.wikipedia.org/wiki/Sender_Policy_Framework) checks.

Say a user wants to send an eamil from test@test.com to someone, when we send
the email, the mail service that we used for sending the email will be checked
against the SPF record of test.com, and it might very well get bounced.

Therefore you should be sure that the mailing services you have configured with
mailspree, are allowed to send mails from the domain that you want as the
sender.

## Future ideas

I was making this project with a set time, and there are certainly areas in the
application that would be nice to improve. To read more about these, checkout
the issues.

## Demo

To see a demo of the project go to the
[mailspree-ui](https://github.com/blacksails/mailspree-ui) project and follow
the instructions.
