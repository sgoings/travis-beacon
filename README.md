# travis-beacon

## Goal

Provide the ability to place status/quality badges in READMEs within
subdirectories of a repository which get their state from the repository's
build job(s).

## Summary

It would be really neat to show badges and shields on a per-chart basis within
helm's [charts repo][helm-charts]. [TravisCI][travis] (or any other CI system) doesn't yet support the
concept of a badge-per-subproject, which is essentially what we're up to here.
So... a sort of "shim" service is needed to bridge the gap between:
- various helm tests (lint, e2e) that are run within a [TravisCI][travis] build
- a [badge service][shields-web]

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

When something like the [shields.io][shields-web] queries the travis-beacon
service with that endpoint and parses the output, we can end up with a badge
in the `charts.git/redis-standalone/README.md` that looks like:

![helm-lint](https://img.shields.io/badge/helm--lint-95%25-green.svg)

## Why the name?

A beacon is a [device designed to attract attention to a specific location.][beacon]
In a similar way, this project provides users of particular helm charts more
focused visibility to the level of quality for individual helm charts while
surveying the [charts repo][helm-charts].

There's also an amusing rhyme of travis-beacon with the product [Mavis Beacon Teaches Typing][mavis].
That connection is unintentional, but puns can always be stretched to find amusement...

[beacon]: https://en.wikipedia.org/wiki/Beacon
[helm-charts]: https://github.com/helm/charts
[mavis]: http://www.vice.com/read/whats-mavis-beacon-up-to-these-days-nothing-shes-fake-926
[shields-web]: http://shields.io/
[travis]: https://travis-ci.org
