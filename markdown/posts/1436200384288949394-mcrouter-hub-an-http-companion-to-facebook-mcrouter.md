## mcrouter-hub: An HTTP companion to Facebook's mcrouter

Gophercon 2015 is right around the corner and that sets up the mood for another open source project.

So, what is mcrouter? <a target="_blank" href="//code.facebook.com/posts/296442737213493/introducing-mcrouter-a-memcached-protocol-router-for-scaling-memcached-deployments/">Mcrouter</a> is a memcache router that helps Facebook scale their Memcache infrastructure to 5 billion requests per second.

It has a dozen or so, read and write strategies for customized scaling experience: Replication, hot-cold warmup, shard by prefix to a separate backend pool, and live reload config file.

The mcrouter daemon itself is tiny, on cold boot consumes 11MB of RAM, high performance C++ project.

For more information, visit <a target="_blank" href="//github.com/facebook/mcrouter/wiki">mcrouter wiki page</a>.


## Live reload config file

This one particular feature is ripe to have a complementary HTTP agent helper.

Imagine the ability to live update hundreds of mcrouter, how cool is that?

So today I am please to release <a target="_blank" href="//github.com/didip/mcrouter-hub" style="font-size: 20px; font-weight: bold">mcrouter-hub</a>, an HTTP companion to mcrouter. It allows you:

* In agent mode:

    * to expose the mcrouter config JSON via HTTP or HTTPS.

    * to modify the mcrouter config JSON. The POST method can be protected behind token-based security.

    * to expose mcrouter stats gathered from both telnet interface and file.

    * to export stats data to New Relic Insights.

* In central mode:

    * to receive config JSON from all agents via HTTP or HTTPS.

    * to receive stats JSON from all agents.

    * Both features can be protected behind token-based security as well.

For more information, visit <a target="_blank" href="//github.com/didip/mcrouter-hub">mcrouter-hub README page</a>.


## Getting started

Download the binary from <a target="_blank" href="//github.com/didip/mcrouter-hub/releases">this list</a> and run it.

<pre class="code"><code class="language-bash"># Agent mode
# Prerequisites:
#   * It must be run on the same box where mcrouter is running.
#   * netcat command line tool must be installed.
#
# MCRHUB_CENTRAL_URLS is only useful if you want to have central daemons.
# That setting is optional.
MCROUTER_ADDR=localhost:5000 \
MCROUTER_CONFIG_FILE=/path/to/mcrouter.json \
MCRHUB_CENTRAL_URLS=http://localhost:5002 \
/path/to/mcrouter-hub

# Central mode
MCRHUB_MODE=central /path/to/mcrouter-hub
</code></pre>

For those of you who run considerable Memcache infrastructure, I heartily recommend mcrouter. It provides greater flexibility when provisioning extra capacity to your cluster.

And when you do try mcrouter, feel free to try <a target="_blank" href="//github.com/didip/mcrouter-hub" style="font-size: 20px; font-weight: bold">mcrouter-hub</a> as well. See if you like it.
