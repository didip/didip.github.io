## Go is beating the averages

In April 2001, rev. April 2003, Paul Graham wrote an article called ["Beating the Averages"](http://www.paulgraham.com/avg.html).

This blog post is about how Go, following the rationale of that article, is the secret weapon that all startups should have in their toolbelt.


## The Secret Weapon

This paragraph highlights everything that's great about Go:

<blockquote>
Software is a very competitive business, prone to natural monopolies. A company that gets software written faster and better will, all other things being equal, put its competitors out of business. And when you're starting a startup, you feel this very keenly. Startups tend to be an all or nothing proposition. You either get rich, or you get nothing. In a startup, if you bet on the wrong technology, your competitors will crush you.
</blockquote>

Go is not a language full of features. It does not have Generics (That said, you can use code generation to build custom data structures), its inheritance is not like what people are used to, it only have partial tail-call optimizations, etc.

The small language specs and its tooling: `go get`, `go fmt`, `go vet` allows you to get things done at a very high velocity and the resulting code is:


## 1. Easy to read

Thanks to go fmt, most Go codes look the same. This is an important trait as your engineering department grows, newcomers can get up to speed quickly.

And if the newcomers don't know Go? No big deal, Go has a small specs, if they are good engineers, they can pick Go up in 1-3 weekends.


## 2. Easy to deploy

The single binary artifact allows you to ignore a whole class of deploy problems.

* You never have to think about differing dependencies between development and production environments.

* Your Vagrant setup suddenly becomes a lot simpler and this allows you to do more cross OS testings.


## 3. Production ready

In most cases, Go performance is on par with Java and it consumes an order of magnitude less memory.

A lot of startups use Rails or Django to get that rapid prototyping velocity. Sounds good, but once you need to deploy to production environment, now what? These are what people typically do:

* Spins up five to ten medium EC2 instances.

* Write complicated deploy pipeline in `cap` or `fab`.

* Setup Nginx/HAProxy in front of the app because Unicorn or Django cannot handle direct internet traffic.

* What about websocket? Oh, that's a separate Node app, with equally complicated deploy pipeline.

* What about background workers? Same problem, another separate app with its own deploy pipeline.

What if your startup can do away with all that and increase your velocity even higher:

* Spins up only 2 micro EC2 instances.

* Write a simple script to deploy. Three commands only: Rsync, symlink, and run.

* You don't need Nginx/HAProxy.

* Websocket? Just install `github.com/gorilla/websocket`.

* Background workers? For a lot of cases, just create goroutines and queue internally using channels.


## The Blub Paradox

This entire section, Lisp vs Blubs, correlates strongly with Go community.

There's a flood of these hypothetical Blub programmers looking down at Go because it does not have x features.

