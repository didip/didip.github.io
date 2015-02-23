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
    <title>%v</title>
    <link rel="stylesheet" type="text/css" media="screen" href="css/simple.css" />
    <!--[if IE]>
    <script src="http://html5shiv.googlecode.com/svn/trunk/html5.js"></script>
    <![endif]-->
</head>

<body>
    <h1>%v</h1>

    %v
</body>
</html>`, title, title, string(html))

	return ioutil.WriteFile(outpath, []byte(generated), 0644)
}
