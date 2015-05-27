## Tollbooth: An HTTP rate limiter middleware in Go

Another great week leads to another OSS project. I am pleased to announce [Tollbooth: #golang HTTP rate limiter middleware](https://github.com/didip/tollbooth).

It allows you to limit access to each one of your request handlers.

For example, you may want to allow unlimited access to `/` but limit access to `POST /login` for as much as 10 requests per second per remote IP.


## Why do I need one?

I don't think it's controversial to say that every web application/service needs a rate limiter.

The fact that the internet is a hostile environment full of potential abusers ready to attack your site is pretty obvious.

What's not obvious is that the cloud is also hostile to your environment especially if you are using Service Oriented Architecture. Without throttling, a brief spike in network latency can cause one of your service to innocently attack your other services, causing a domino effect.

HTTP rate limiter can act as first line of defense in keeping your services in line.

Allow me to give you a few more real life use cases:

* Limiting `POST` on `/login` and `/signup` helps to reduce spammer from creating too many accounts.

* Limiting `POST` on `/upload` endpoint allows you to control how much memory your application is spending on buffering+streaming blobs.

* Limiting `GET` on some of your public API endpoints protects you from your customers accessing too much data in super short period of time.

* Public API endpoints need DDoS protection anyway. This middleware won't help you if your link is saturated through other means of attack.

* As mentioned above, limiting your private API endpoints protects you from your own services DoS attack. This accidental DoS can even come from your own Javascript.

* Limiting API endpoints is usually an easy way to receive cookies and pies from DB team.


## Sounds good, how can I use it on my Go application?

Great question! Here's the five minutes tutorial:

<pre class="code"><code class="language-go">package main

import (
    "github.com/didip/tollbooth"
    "net/http"
    "time"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("Hello, World!"))
}

func main() {
    // You can create a generic limiter for all your handlers
    // or one for each handler. Your choice.
    // This limiter basically says: allow at most 1 request per 1 second.
    limiter := tollbooth.NewLimiter(1, time.Second)

    // This is an example on how to limit only GET and POST requests.
    limiter.Methods = []string{"GET", "POST"}

    // You can also limit by specific request headers, containing certain values.
    // Typically, you prefetched these values from the database.
    limiter.Headers = make(map[string][]string)
    limiter.Headers["X-Access-Token"] = []string{"abc123", "xyz098"}

    // And finally, you can limit access based on basic auth usernames.
    // Typically, you prefetched these values from the database as well.
    limiter.BasicAuthUsers = []string{"bob", "joe", "didip"}

    // Example on how to wrap your request handler.
    http.Handle("/", tollbooth.LimitFuncHandler(limiter, HelloHandler))
    http.ListenAndServe(":12345", nil)
}
</code></pre>


## Looks too easy to be true, what's the catch?

Some of you sceptics might think: Surely I need to store some metadata on the database to make this work. If I need a database, I'm introducing another point of failure, right?

Good thinking, but no. You do not need a database to use this middleware thanks to this little algorithm called [Token Bucket. Click on this link to read more on Wikipedia.](http://en.wikipedia.org/wiki/Token_bucket).

So give it a try: <a target="_blank" href="//github.com/didip/tollbooth" style="font-size: 20px; font-weight: bold">github.com/didip/tollbooth</a>. See if you like it. Feel free to follow the discussions on <a target="_blank" href="//news.ycombinator.com/item?id=9588178" style="font-size: 17px; font-weight: bold">Hacker News</a> and <a target="_blank" href="//www.reddit.com/r/golang/comments/36uicf/tollbooth_an_http_rate_limiter_middleware_in_go/" style="font-size: 17px; font-weight: bold">/r/golang</a>.

<a aria-label="Star didip/tollbooth on GitHub" data-count-aria-label="# stargazers on GitHub" data-count-api="/repos/didip/tollbooth#stargazers_count" data-count-href="/didip/tollbooth/stargazers" data-style="mega" data-icon="octicon-star" href="//github.com/didip/tollbooth" class="github-button">Star</a>
<a aria-label="Fork didip/tollbooth on GitHub" data-count-aria-label="# forks on GitHub" data-count-api="/repos/didip/tollbooth#forks_count" data-count-href="/didip/tollbooth/network" data-style="mega" data-icon="octicon-git-branch" href="//github.com/didip/tollbooth/fork" class="github-button">Fork</a>
<a aria-label="Watch didip/tollbooth on GitHub" data-count-aria-label="# watchers on GitHub" data-count-api="/repos/didip/tollbooth#subscribers_count" data-count-href="/didip/tollbooth/watchers" data-style="mega" data-icon="octicon-eye" href="//github.com/didip/tollbooth" class="github-button">Watch</a>

<script async defer id="github-bjs" src="//buttons.github.io/buttons.js"></script>