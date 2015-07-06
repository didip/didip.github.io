[<h2>mcrouter-hub: An HTTP companion to Facebook's mcrouter</h2>](/posts/1436200384288949394-mcrouter-hub-an-http-companion-to-facebook-mcrouter.html)


Gophercon 2015 is right around the corner and that sets up the mood for another open source project.

So, what is mcrouter? <a target="_blank" href="//code.facebook.com/posts/296442737213493/introducing-mcrouter-a-memcached-protocol-router-for-scaling-memcached-deployments/">Mcrouter</a> is a memcache router that helps Facebook scale their Memcache infrastructure to 5 billion requests per second.

It has a dozen or so, read and write strategies for customized scaling experience: Replication, hot-cold warmup, shard by prefix to a separate backend pool, and live reload config file.


[<h2>Why Go is beating the averages</h2>](/posts/1434380433653410969-why-go-is-beating-the-averages.html)


In April 2001, rev. April 2003, Paul Graham wrote an article called <a href="http://www.paulgraham.com/avg.html" target="_blank">"Beating the Averages"</a>

This blog post is about how Go, following the rationale of that article, is the secret weapon that all startups should have.



[<h2>Tollbooth: An HTTP rate limiter middleware in Go</h2>](/posts/1432264032306462173-tollbooth-http-rate-limiter-middleware-in-go.html)


Another great week leads to another OSS project. I am pleased to announce [Tollbooth: #golang HTTP rate limiter middleware](https://github.com/didip/tollbooth).

It allows you to limit access to each one of your request handlers.

For example, you may want to allow unlimited access to `/` but limit access to `POST /login` for as much as 10 requests per second per remote IP.


[<h2>Can you build a web application in Go?</h2>](/posts/1430960969799974572-can-you-build-a-web-application-in-go.html)


Surely the answer is yes. Any yet, a surprising number of people asked me this question. Something must have prevented them from writing web project in Go.

So, I began the journey...



[<h2>Back to blogging again. This time in Go!</h2>](/posts/1424762251437171629-back-to-blogging.html)


I've always wanted to have more control over my own blog, reducing as much bells and whistles as possible, and focus on just the essence: The posts themselves.

All my friends know that I've been enamored with [Go](https://golang.org/) nowadays. This blog engine (static-site generator) is another perfect excuse to write even more Go.

For your reading pleasure, I'll share some of the blog's interesting parts you may find useful.

