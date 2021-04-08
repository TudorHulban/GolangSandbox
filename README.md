# Golang Resources <a name="top"></a>
* [Why Go?](#why)
* [Getting Started](#start)
* [Directory Structure](#directory)
* [Error Handling](#error)
* [Modules](#mod)
* [OOP](#oop)
* [Constructors](#constructors)
* [Design Patterns](#pattern)
* [Architecture & Best Practices ](#arch)
* [SOLID](#solid)
* [Concurrency](#conc)
* [Interfaces](#interf)
* [Data Structures and Algorithms](#structs)
* [BlockChain](#block)
* [JSON](#json)
* [PR](#pr)
* [Security](#security)
* [Linters, Code Analysis, Style](#lint)
* [Tutorial Series](#tutorials)
* [Functional](#functional)
* [RDBMS](#rdbms)
* [CI / CD](#cicd)
* [Microservices](#micros)
* [Pointers](#point)
* [Context](#ctx)
* [Memory Management](#mems)
* [RPC](#rpc)
* [Authentication](#auth)
* [Testing](#tests)
* [Benchmarking](#bench)
* [Templates](#templ)
* [Logging](#log)
* [CI/CD](#ci)
* [Miscellaneous](#misc)
* [Variables](#vars)
* [String Formatting](#strings)
* [High Performance](#high)
* [AWS Lambdas](#lambda)


## Why Go? <a name="why"></a> 
```html
https://qarea.com/blog/golang-web-development-better-than-python
https://hackernoon.com/5-reasons-why-we-switched-from-python-to-go-4414d5f42690
https://qarea.com/blog/ruby-vs-golang-comparison-best-solution
https://stxnext.com/blog/2017/09/27/go-go-python-rangers-comparing-python-and-golang/?utm_campaign=Pat&utm_medium=Social&utm_source=Quora
https://qarea.com/blog/why-do-enterprises-go-golang
https://www.quora.com/Googles-Go-language-vs-Python-which-would-you-prefer-to-use-and-why
https://qarea.com/blog/golang-development-vs-node-js-who-wins
https://www.freecodecamp.org/news/here-are-some-amazing-advantages-of-go-that-you-dont-hear-much-about-1af99de3b23a/
https://blog.golang.org/survey2018-results
https://www.quora.com/Why-are-while-loops-and-classes-absent-from-the-Go-programming-language/answer/Roman-Scharkov?ch=10&share=91dd1192&srid=hdP1T
https://www.quora.com/Why-does-Go-avoid-inheritance-and-provides-first-class-support-for-composition-instead/answer/Roman-Scharkov?ch=10&share=8429a5f0&srid=hdP1T
```
## Getting Started <a name="start"></a> ([Up](#top))
```html
https://draft.dev/learn/technical-blogs/golang
https://www.quora.com/What-is-the-best-advice-for-a-beginner-who-wants-to-learn-Golang-as-a-first-language
http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/
https://dave.cheney.net/practical-go/presentations/qcon-china.html
https://peter.bourgon.org/go-best-practices-2016/
https://dave.cheney.net/category/golang
https://github.com/emluque/golang-internals-resources
https://docs.google.com/document/d/1yIAYmbvL3JxOKOjuCyon7JhW4cSv1wy5hC0ApeGMV9s/pub
https://dev.to/seattledataguy/the-interview-study-guide-for-software-engineers-764
https://12factor.net/
https://itnext.io/tips-for-writing-self-documenting-code-e54a15e9de2
https://ueokande.github.io/go-slice-tricks/
```
## Directory Structure <a name="directory"></a> ([Up](#top))
```
https://medium.com/dm03514-tech-blog/go-aws-lambda-project-structure-using-golang-98b6c0a5339d
```

## Error Handling <a name="error"></a> ([Up](#top))
```
import "github.com/pkg/errors"
```
```
https://blog.logrocket.com/error-handling-in-golang/
https://www.youtube.com/watch?v=4WIhhzTTd0Y
```
## Modules <a name="mod"></a> ([Up](#top))
Pick a different branch:
```
# by commit:
go get github.com/someone/some_module@af041234567
# by branch:
go get github.com/someone/some_module@mybranch
```
Clean cache:
```bash
go clean --modcache
```
```html
https://github.com/golang/go/wiki/Modules
https://bogotobogo.com/GoLang/GoLang_Modules_1_Creating_a_new_module.php
https://stackoverflow.com/questions/53682247/how-to-point-go-module-dependency-in-go-mod-to-a-latest-commit-in-a-repo
```
## OOP <a name="oop"></a> ([Up](#top))
```html
https://www.oodesign.com
https://code.tutsplus.com/tutorials/lets-go-object-oriented-programming-in-golang--cms-26540
```
## Constructors <a name="constructors"></a> ([Up](#top))
```html
https://medium.com/@j7mbo/constructors-in-go-b1d1513e1c1d
```
## Design Patterns <a name="pattern"></a> ([Up](#top))
### Singleton
```html
http://blog.ralch.com/tutorial/design-patterns/golang-singleton/ 
```
### Factory
```html
http://blog.ralch.com/articles/design-patterns/golang-factory-method/ 
https://www.sohamkamani.com/blog/golang/2018-06-20-golang-factory-patterns/
```
### Others
```html
https://medium.com/swlh/provider-model-in-go-and-why-you-should-use-it-clean-architecture-1d84cfe1b097
http://blog.ralch.com/tutorial/design-patterns/golang-design-patterns/
http://blog.ralch.com/articles/design-patterns/golang-decorator/ 
http://blog.ralch.com/articles/design-patterns/golang-adapter/ 
http://blog.ralch.com/tutorial/design-patterns/golang-bridge/ 
http://blog.ralch.com/tutorial/design-patterns/golang-composite/ 
http://blog.ralch.com/tutorial/design-patterns/golang-prototype/ 
http://blog.ralch.com/tutorial/design-patterns/golang-builder/ 
https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html 
https://github.com/bxcodec/go-clean-arch 
https://hackernoon.com/golang-clean-archithecture-efd6d7c43047 
https://dave.cheney.net/practical-go/presentations/qcon-china.html
http://www.golangpatterns.info/home
http://tmrts.com/go-patterns/
```
### Repository Pattern
```html
https://adodd.net/post/go-ddd-repository-pattern/
https://deviq.com/repository-pattern/
```
### Publish - Subscribe
```html
https://eli.thegreenplace.net/2020/pubsub-using-channels-in-go/
https://engineering.grab.com/plumbing-at-scale
```
## Architecture & Best Practices <a name="arch"></a> ([Up](#top))
```html
https://towardsdatascience.com/10-common-software-architectural-patterns-in-a-nutshell-a0b47a1e9013
https://github.com/OWASP/Go-SCP
https://doordash.engineering/2019/07/22/writing-delightful-http-middlewares-in-go/
https://dave.cheney.net/practical-go/presentations/qcon-china.html
https://www.joeldholmes.com/post/go-hex-arch/
```
## SOLID <a name="solid"></a> ([Up](#top))
```html
https://medium.com/swlh/improve-your-code-with-the-solid-principles-fcdcfdf21810
```
### S - Single responsibility principle
```html
https://medium.com/@felipedutratine/solid-single-responsibility-principle-in-golang-dc4a6be9bb3a
```
### O - Open–closed principle
Modifying a package behaviour comes by adding new code, never by modifying the old one.
```html
https://medium.com/@felipedutratine/solid-open-closed-principle-in-golang-3d23e8258a45
```
### L - Liskov substitution principle
```html
https://medium.com/@felipedutratine/solid-liskov-substitution-principle-fb01a3b51a68
```
### I - Interface segregation principle
```html
https://medium.com/@felipedutratine/solid-interface-segregation-principle-in-golang-49d4bbb4d3f7
```
### D - Dependency inversion principle
```html
https://en.wikipedia.org/wiki/Dependency_injection
https://www.captaincodeman.com/2015/03/05/dependency-injection-in-go-golang
```
## Concurrency <a name="conc"></a> ([Up](#top))
```html
https://www.youtube.com/watch?utm_campaign=Go+Full-Stack&utm_medium=email&utm_source=Revue+newsletter&v=S-MaTH8WpOM
https://blog.golang.org/race-detector
http://blog.ralch.com/articles/golang-concurrency-patterns-context/ 
https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html
https://player.vimeo.com/video/49718712   # Rob Pike - "Concurrency Is Not Parallelism"
https://12factor.net/
https://medium.com/a-journey-with-go/go-how-does-go-recycle-goroutines-f047a79ab352
https://medium.com/@hau12a1/go-worker-pool-vs-pool-of-workers-b7c0598b4a67
https://www.alexedwards.net/blog/how-to-rate-limit-http-requests?utm_campaign=The%20Go%20Gazette&utm_medium=email&utm_source=Revue%20newsletter
https://www.youtube.com/watch?v=QDDwwePbDtw
https://medium.com/swlh/why-go-so-fast-1c4a20d15bde
```
### Channels
```
https://medium.com/rungo/anatomy-of-channels-in-go-concurrency-in-go-1ec336086adb
```
## Interfaces <a name="interf"></a> ([Up](#top))
```html
http://www.golangprograms.com/golang/interface-type/
http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go 
https://gobyexample.com/interfaces 
https://stackoverflow.com/questions/39092925/why-are-interfaces-needed-in-golang
https://medium.com/better-programming/how-i-build-robust-scalable-go-applications-192a0df32d1
https://www.efekarakus.com/golang/2019/12/29/working-with-interfaces-in-go.html
https://medium.com/@matryer/golang-advent-calendar-day-seventeen-io-reader-in-depth-6f744bb4320b
```
## Data Structures and Algorithms <a name="structs"></a> ([Up](#top))
### Arrays
```
https://www.golangprograms.com/go-language/arrays.html
```
### Slices
```html
https://www.youtube.com/watch?v=5o97pT9A1d4
```
### Linked Lists
```html
https://ieftimov.com/post/golang-datastructures-linked-lists/
```
### Trees
```
https://ieftimov.com/post/golang-datastructures-trees/
```
## BlockChain <a name="block"></a> ([Up](#top))
```html
https://medium.com/@mycoralhealth/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc
```
## JSON <a name="json"></a> ([Up](#top))
```html
https://www.sohamkamani.com/blog/2017/10/18/parsing-json-in-golang/
https://thehoard.blog/using-golang-and-json-for-kafka-consumption-with-high-throughput-4cae28e08f90
```
## PR <a name="pr"></a> ([Up](#top))
```html
https://medium.freecodecamp.org/what-i-learned-from-an-old-github-project-that-won-3-000-stars-in-a-week-628349a5ee14
```
## Security <a name="security"></a> ([Up](#top))
### SSL
```html
https://www.ssllabs.com/
```
### Injection
```html
https://www.calhoun.io/what-is-sql-injection-and-how-do-i-avoid-it-in-go/
```

## Linters, Code Analysis, Style <a name="lint"></a> ([Up](#top))
Establishing good paradigms and consistent, accessible standards for writing clean code can help prevent developers from wasting many meaningless hours on trying to understand their own (or others') work.
### Clean Code
```
https://github.com/Pungyeon/clean-go-article
```
### Linters
```
https://golangci.com
https://go-critic.github.io/overview.html
```
### Static Code Analysis (SCA)
```
https://superhighway.dev/staticcheck-in-action
```
### Style
```
https://medium.com/@dgryski/idiomatic-go-resources-966535376dba
https://github.com/dgryski/awesome-go-style
https://wiki.crdb.io/wiki/spaces/CRDB/pages/73072807/Git+Commit+Messages
https://github.com/uber-go/guide/blob/master/style.md
https://github.com/golang/go/wiki/CodeReviewComments#go-code-review-comments
```
## Tutorial Series <a name="tutorials"></a> ([Up](#top))
```html
https://golangbot.com/learn-golang-series/
https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql
```
## Functional <a name="functional"></a> ([Up](#top))
```html
https://medium.com/swlh/functional-pipe-in-go-1c755467fe14
https://stackoverflow.com/questions/60102866/golang-trouble-understanding-a-function-as-a-receiver
https://kinbiko.com/posts/2021-01-10-function-types-in-go/
```
## RDBMS <a name="rdbms"></a> ([Up](#top))
```html
https://medium.com/@rocketlaunchr.cloud/canceling-mysql-in-go-827ed8f83b30
```
## CI / CD <a name="cicd"></a> ([Up](#top))
```html
https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
https://blog.alexellis.io/inject-build-time-vars-golang/
```
## Microservices <a name="micros"></a> ([Up](#top))
```html
https://rubygarage.org/blog/top-languages-for-microservices
https://ewanvalentine.io/microservices-in-golang-part-2/
https://github.com/mfornos/awesome-microservices#go
https://martinfowler.com/articles/microservices.html
https://medium.com/better-programming/deploy-a-full-stack-go-and-react-app-on-kubernetes-4f31cdd9a48b
```
## Pointers <a name="point"></a> ([Up](#top))
```html
http://piotrzurek.net/2013/09/20/pointers-in-go.html
https://www.youtube.com/watch?v=sTFJtxJXkaY
```
## Context <a name="ctx"></a> ([Up](#top))
```html
https://dev.to/georgeoffley/working-with-context-in-go-75e
http://p.agnihotry.com/post/understanding_the_context_package_in_golang/
https://github.com/golang/go/wiki/ExperienceReports#context
https://medium.com/@cep21/how-to-correctly-use-context-context-in-go-1-7-8f2c0fafdf39
https://www.alexedwards.net/blog/how-to-manage-database-timeouts-and-cancellations-in-go
```
## Memory Management <a name="mems"></a> ([Up](#top))
```html
https://povilasv.me/go-memory-management-part-2/#
```
## RPC <a name="rpc"></a> ([Up](#top))
```html
https://blog.twitch.tv/twirp-a-sweet-new-rpc-framework-for-go-5f2febbf35f
```
## Authentication <a name="auth"></a> ([Up](#top))
```html
https://blog.ellation.com/managing-cpu-load-in-golang-515b9356bc5
https://dev.to/wagslane/how-to-build-jwt-s-in-go-golang-i7c
```
## Testing <a name="tests"></a> ([Up](#top))
```html
https://github.com/golang/go/wiki/TestComments#equality-comparison-and-diffs
https://blog.codecentric.de/en/2017/08/gomock-tutorial/
https://dave.cheney.net/2019/05/07/prefer-table-driven-tests
https://apitest.dev/
https://golang.testcontainers.org/
```
### Integration Testing
```
https://hackandsla.sh/posts/2020-07-12-golang-integration-testing/
```
## Benchmarking <a name="bench"></a> ([Up](#top))
```html
https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
https://medium.com/justforfunc/analyzing-the-performance-of-go-functions-with-benchmarks-60b8162e61c6
https://blog.simon-frey.eu/known-length-slice-initialization-speed-golang-benchmark-wednesday?utm_campaign=The%20Go%20Gazette&utm_medium=email&utm_source=Revue%20newsletter
```
## Templates <a name="templ"></a> ([Up](#top))
```html
https://goinbigdata.com/example-of-using-templates-in-golang/
https://golangtutorials.blogspot.com/2011/06/go-templates.html
https://curtisvermeeren.github.io/2017/09/14/Golang-Templates-Cheatsheet
https://www.calhoun.io/intro-to-templates-p3-functions/
https://philipptanlak.com/web-frontends-in-go/
```
## Logging <a name="log"></a> ([Up](#top))
```html
https://blog.maddevs.io/how-to-start-with-logging-in-golang-projects-part-1-3e3a708b75be
https://dev-journal.in/2019/05/27/adding-uber-go-zap-logger-to-golang-project/
https://engineering.grab.com/structured-logging
```
## Miscellaneous <a name="misc"></a> ([Up](#top))
```html
https://medium.com/weareservian/automagical-https-with-docker-and-go-4953fdaf83d2
```
## Variables <a name="vars"></a> ([Up](#top))
GOSUMDB
```html
https://flak.tedunangst.com/post/please-note-gosumdb-caches-410
https://www.ardanlabs.com/blog/2020/02/modules-04-mirros-checksums-athens.html
```
## String Formatting <a name="strings"></a> ([Up](#top))
```html
https://medium.com/swlh/ultimate-golang-string-formatting-cheat-sheet-234ec92c97da
```
## High Performance <a name="high"></a> ([Up](#top))
```html
https://dave.cheney.net/high-performance-go-workshop/gophercon-2019.html
https://github.com/dgryski/go-perfbook
https://github.com/alecthomas/go_serialization_benchmarks
https://developer20.com/using-sync-pool/index.html
https://bionic.fullstory.com/why-you-should-be-using-errgroup-withcontext-in-golang-server-handlers/
https://dev.to/codenation/profiling-golang-applications-using-pprof-1c7o
https://www.freecodecamp.org/news/how-i-investigated-memory-leaks-in-go-using-pprof-on-a-large-codebase-4bec4325e192/
https://www.aaronraff.dev/blog/an-introduction-to-producing-and-consuming-kafka-messages-in-go
```
### Protocol Buffers
```html
https://kpbird.medium.com/golang-serialize-struct-using-gob-part-1-e927a6547c00
https://blog.golang.org/gob
https://appliedgo.net/networking/
```
## AWS Lambdas <a name="lambda"></a> ([Up](#top))
```
https://geekflare.com/aws-lambda-for-beginners/
https://docs.aws.amazon.com/lambda/latest/dg/lambda-golang.html
https://cloudnative.ly/lambdas-with-golang-a-technical-guide-6f381284897b
```
