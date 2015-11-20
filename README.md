# travis-beacon

## Goal

Provide the ability to place status/quality badges in READMEs within
subdirectories of a repository which get their state from the repository's
build job(s).

## Summary

It would be really neat to show badges and shields on a per-chart basis within
helm's charts repo. TravisCI (or any other CI system) doesn't yet support the
concept of multiple badges per repository. So... a sort of "shim" service is
needed to bridge the gap between:
- various helm tests (lint, e2e) that are run within a TravisCI build
- a badge service

Then you can build an endpoint to the badge service which can query
travis-beacon in more complex ways, such as in our case:

GET /charts/redis-standalone

Which can return a json object such as:

    {
      "name": "redis-standalone",
      "lint": {
        "lint_score": 95
        "lint_failures": {
          "lint-picky-1": "lint-picky-1 description",
          "lint-serious-5", "lint-serious-5 description"
        }  
      }
      "e2e_pass": false
    }

## Why the name?

A beacon is a device [designed to attract attention to a specific location.][beacon]
Therefore, this project provides a way for us to identify and organize problems
within the tests of helm charts run by TravisCI.

There's also an amusing rhyme of travis-beacon with the product [Mavis Beacon Teaches Typing][mavis].
That connection is unintentional, but the pun could always be stretched.

[beacon]: https://en.wikipedia.org/wiki/Beacon
[mavis]: http://www.vice.com/read/whats-mavis-beacon-up-to-these-days-nothing-shes-fake-926
