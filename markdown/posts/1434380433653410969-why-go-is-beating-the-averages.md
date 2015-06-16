## Why Go is beating the averages

In April 2001, rev. April 2003, Paul Graham wrote an article called <a href="http://www.paulgraham.com/avg.html" target="_blank">"Beating the Averages"</a>

This blog post is about how Go, following the rationale of that article, is the secret weapon that all startups should have.


## The secret weapon

<blockquote>
Software is a very competitive business, prone to natural monopolies. A company that gets software written faster and better will, all other things being equal, put its competitors out of business... In a startup, if you bet on the wrong technology, your competitors will crush you.
</blockquote>

That quoted paragraph describes how having a high velocity is very important for a startup to succeed. Makes sense, since every startup always have limited amount of time and money. Achieving high velocity, is precisely where Go shines.

Go is not a language full of features. It does not have Generics (That said, you can use code generation to build custom data structures), its inheritance is not like what people are used to, it only have partial tail-call optimizations, etc. But that's what makes Go great! <b>Small language specs and great tooling</b>, allow you to get things done at a very high velocity and the resulting code is usually:


### 1. Easy to write

* Go standard library is rich, easy to read, high quality, and performant. You'll be surprised how far you can get without using 3rd party libraries. As a side note, did you know that `database/sql` gives you connection pooling for free?

* `go get` is not perfect, but it allows you to pull down third party libraries without any setup at all and this is hugely important during development phase.

* The type system hits the sweet spot for me, it eliminates a whole class of problems with minimal ceremony.


### 2. Easy to read

Thanks to `go fmt`, most Go codes look the same. This is an important trait as your engineering department grows, newcomers can get up to speed quickly.

And if the newcomers don't know Go? No big deal. Thanks to the minimal specs, good engineers can pick up Go in 1-3 weekends and have a good grasps on all the features.


### 3. Easy to deploy

Nowadays, a lot of startups use Rails or Django to get that rapid prototyping velocity. Seems sensible, but once you need to deploy to production environment, now what? These are what people typically do:

* Spins up five to ten medium EC2 instances right of the bat.

* Write complicated deploy pipeline in `cap` or `fab`.

* Setup Nginx/HAProxy in front of the app because Unicorn or Gunicorn cannot handle direct internet traffic.

* What about websocket? Oh, that's a separate Node app, with equally complicated deploy pipeline.

* What about background workers? Same problem, another separate app with its own deploy pipeline.

What if your startup can do away with all that and increase the velocity even higher:

* High performance and small memory footprint app means less money spent on computing resources, which in turns mean longer runway.

* Thanks to single binary artifact, you rarely have to think about differing dependencies between development and production environments.

* Single binary artifact makes your deployment pipeline a lot shorter. At bare minimum you only need three commands: rsync, symlink, and run. If you versioned your binary, rollback is as simple as flipping symlinks.

* Single binary artifact also makes your Vagrant setup a lot simpler and this allows you to do more cross OS testings.

* You don't need Nginx/HAProxy as connection buffer. Built-in `net/http` server handles high traffic just fine.

* Websocket? Thanks to Go evented design, you can run regular HTTP handlers and Websocket handlers in the same app. Just install `github.com/gorilla/websocket`.

* Background workers? For a lot of cases, just create goroutines and queue internally using channels.


### 4. Production ready

In most cases, Go performance is on par with Java and it consumes an order of magnitude less memory.


## The blub paradox

The entire section, Lisp vs Blubs, got me thinking about the Go community.

There is no shortage of these hypothetical Blub programmers looking down at Go because it does not have X features.

Should Go community care about them? I don't think so.

Go is quickly shaping the startup community. Here's the non exhaustive, and ever growing, list of companies that use Go: Bit.ly, Baidu, CloudFlare, DigitalOcean, Disqus, Dropbox, GitHub, Heroku, New York Times, Parse, Square, Twitch, Tumblr, Twitter, etc.

Soon enough, some of the nay-sayers will convert to Go anyway because they are following the job market.


## Back to the original point

Given all the arguments above, startup founders, y'all should really consider Go seriously.

There's so much to win and zero to lose. Why wait? Call <a style="font-size: 17px; font-weight: bold" href="//play.golang.org/" target="_blank">1-800-SWITCH-GO</a> and reap all the benefits now.
