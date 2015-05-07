## Can you build a web application in Go?

Surely the answer is yes. Any yet, a surprising number of people asked me this question. Something must have prevented them from writing web project in Go.

So, I began the journey...


## The Beginning

Superficially, Go culture is similar to Node and Python in the web world.

Everyone is suggesting newcomers to start small, piece components one by one as you need them. Seems reasonable.

So that begs the question, what are the minimum features you need to build a modern web application?

Here's the non-exhaustive list:

* A database, usually RDBMS but not always.

* Database migration tool, some web frameworks provide it, some don't.

* Logging facility, all modern languages provide this.

* Testing tools, most modern languages provide this.

* HTTP router, usually comes by default if you use a web framework.

* HTTP middleware, some web frameworks provide this by default.

* Session, context, and cookie management, fairly standard stuff.

* Password hasher, you cannot have a web project without it.

* Healty amount of 3rd party libraries to build awesome things.

* A UI framework, most people settle with <a target="_blank" href="//getbootstrap.com/">Bootstrap</a> or <a target="_blank" href="//foundation.zurb.com/">Foundation</a>.

That's... quite a lot of stuff and professional work usually comes with a deadline attached. I can see how this may presents a barrier to new users.


## The Buildup

At this stage most newcomers would start googling "Golang Web Framework". And boy, do they have a lot of choices:

* <a target="_blank" href="//www.gorillatoolkit.org/">Gorilla Toolkit</a>

* <a target="_blank" href="//martini.codegangsta.io/">Martini</a> and it's sibling: <a target="_blank" href="//github.com/codegangsta/negroni">Negroni</a>.

* <a target="_blank" href="//gin-gonic.github.io/gin/">Gin Gonic</a>

* <a target="_blank" href="//goji.io/">Goji</a>

* <a target="_blank" href="//github.com/gocraft/web/">gocraft/web</a>

* <a target="_blank" href="//beego.me/">Beego</a>

* <a target="_blank" href="//revel.github.io/">Revel</a>

* Various small players like: <a target="_blank" href="//github.com/plimble/ace">Ace</a>, <a target="_blank" href="//github.com/nbio/hitch">Hitch</a>, <a target="_blank" href="//github.com/guregu/kami">Kami</a>, <a target="_blank" href="//github.com/rocwong/neko">Neko</a>

This reminded me of the bad old days of Python web frameworks before Django came along: Zope, Twisted, CherryPy, TurboGear, Spyce, Quixote, web.py, <a target="_blank" href="//wiki.python.org/moin/WebFrameworks">etc.</a>

By now, probably half of the newcomers who are interested bounced already and return to their comfort zone: Rails or Django.


## The Problem

That said, the high number of web frameworks is not the biggest problem. Instead, it's the difficulty to bootstrap a new project within a deadline.

<b>From 0-60, how fast can I get a new project up and running while still following the community's general advice; by starting small.</b>

If a newcomer decides to use one of the larger frameworks, the community will tell them to use smaller ones or even just net/http.

But realistically, a big web project needs a structure. A common set of tools or conventions for engineers to maximize their agility at work.


## The Solution

So what if, you can have a cake and eat it too? What if, you can have all the things you need for a modern web application without being forced to adhere to a big framework.

A project generator can do that for you. It removes the <a target="_blank" href="//en.wikipedia.org/wiki/Analysis_paralysis">analysis-paralysis</a> situation during the research phase and provides users with the ultimate freedom to use or remove some of the components.

To test this idea, I created <a target="_blank" href="//go-bootstrap.io/">go-bootstrap</a>.

It puts together common tools (following the requirements described above) and bootstrap a new project to the point where user can signup and login.

Once a project is generated, users have 100% control over the web stack. There's no need to worry about upgrade headache from version 2.0 to 3.0.

So give it a try: <a target="_blank" href="//go-bootstrap.io/" style="font-size: 20px; font-weight: bold">go-bootstrap.io</a>. See if you like it and hopefully you'll convert to <a target="_blank" href="//play.golang.org/">Go</a>, just like me.
