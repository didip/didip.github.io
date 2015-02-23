package templates

import (
	"fmt"
	"io/ioutil"
)

func Generate(title string, html []byte, outpath string) error {
	generated := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset=utf-8 />
    <meta name="viewport" content="width=device-width, initial-scale=1.0">

    <title>%v</title>

    <link rel="stylesheet" href="//yui.yahooapis.com/pure/0.5.0/pure-min.css">
    <link rel="stylesheet" href="//yui.yahooapis.com/pure/0.5.0/grids-responsive-min.css">
    <link rel="stylesheet" href="/css/layouts/blog.css">

    <!--[if IE]>
    <script src="//html5shiv.googlecode.com/svn/trunk/html5.js"></script>
    <![endif]-->
</head>

<body>

<div id="layout" class="pure-g">
    <div class="sidebar pure-u-1 pure-u-md-1-4">
        <div class="header">
            <h1 class="brand-title">%v</h1>

            <nav class="nav">
                <ul class="nav-list">
                    <li class="nav-item">
                        <a href="//twitter.com/didip">Twitter</a>
                    </li>
                    <li class="nav-item">
                        <a href="//github.com/didip">GitHub</a>
                    </li>
                </ul>
            </nav>
        </div>
    </div>

    <div class="content pure-u-1 pure-u-md-3-4">
        <div>
            <!-- A wrapper for all the blog posts -->
            <div class="posts">
                <section class="post">
                    %v
                </section>
            </div>
        </div>
    </div>
</div>
</body>
</html>`, title, title, string(html))

	return ioutil.WriteFile(outpath, []byte(generated), 0644)
}
