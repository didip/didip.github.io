package templates

import (
	"fmt"
	"io/ioutil"
)

func Generate(blogTitle string, postTitle string, html []byte, outpath string) error {
	generated := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset=utf-8 />
    <meta name="viewport" content="width=device-width, initial-scale=1.0">

    <title>%v - %v</title>

    <link rel="stylesheet" href="//yui.yahooapis.com/pure/0.5.0/pure-min.css">
    <link rel="stylesheet" href="//yui.yahooapis.com/pure/0.5.0/grids-responsive-min.css">
    <link rel="stylesheet" href="//cdnjs.cloudflare.com/ajax/libs/font-awesome/4.3.0/css/font-awesome.min.css">
    <link rel="stylesheet" href="//fonts.googleapis.com/css?family=Roboto:400,500" type="text/css">
    <link rel="stylesheet" href="/css/prism.css">
    <link rel="stylesheet" href="/css/layouts/blog.css">

    <script src="/js/prism.js"></script>
    <!--[if IE]>
    <script src="//html5shiv.googlecode.com/svn/trunk/html5.js"></script>
    <![endif]-->
</head>

<body>

<div id="layout" class="pure-g">
    <div class="sidebar pure-u-1 pure-u-md-1-4">
        <div class="self">
            <img src="https://pbs.twimg.com/profile_images/519365815116627968/4Ob_nn_o_400x400.jpeg" alt="Didip Kerabat">

            <div class="self-info">
                <h3 class="fullname">Didip Kerabat</h3>

                <nav class="nav">
                    <ul class="nav-list">
                        <li class="nav-item">
                            <a href="//twitter.com/didip"><i class="fa fa-twitter"></i> Twitter</a>
                        </li>
                        <li class="nav-item">
                            <a href="//github.com/didip"><i class="fa fa-github"></i> GitHub</a>
                        </li>
                    </ul>
                </nav>
            </div>
        </div>

        <div class="header">
            <h1 class="brand-title"><a href="/">%v</a></h1>
            <h2 class="brand-tagline">They go together</h2>
            <h3 class="self-mobile">Didip Kerabat</h3>
        </div>
    </div>

    <div class="content pure-u-1 pure-u-md-3-4">
        <!-- A wrapper for all the blog posts -->
        <div class="posts">
            <section class="post">
                %v
            </section>
        </div>
    </div>
</div>
</body>
</html>`, blogTitle, postTitle, blogTitle, string(html))

	return ioutil.WriteFile(outpath, []byte(generated), 0644)
}
