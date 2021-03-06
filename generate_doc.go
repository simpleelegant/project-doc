package main

import (
	"io/ioutil"
	"strings"

	"github.com/russross/blackfriday"
	"github.com/simpleelegant/project-doc/list"
)

const tmpl = `
<html>

<head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <title>Project Documents</title>
    <style media="screen">
        body {
            font-family: "Fira Sans", "Trebuchet MS", Arial, "Microsoft Yahei", sans-serif;
        }

        code {
            font-family: Consolas, "Fira Mono", "Liberation Mono", Courier, "Microsoft Yahei", monospace;
            color: red;
            padding: 2px 4px;
            font-size: 0.9em;
        }

        pre {
            border: 1px solid #CACACA;
            padding: 10px;
            overflow: auto;
            border-radius: 3px;
            background-color: #FAFAFB;
            color: #393939;
            margin: 2em 0px;
        }
        pre code {
            padding: 0;
        }

		table {
			border: 1px solid #CCC;
			border-collapse: collapse;
			width: 100%;
			margin-bottom: 20px;
		}

		th {
			border: 1px solid #CCC;
			background-color: aliceblue;
			white-space: nowrap;
			text-align:left;
			padding: 4px;
		}

		td {
			border: 1px solid #CCC;
			padding: 4px;
			word-break: break-word;
		}

        .body {
            max-width: 980px;
            margin: 0px auto;
        }

        .left {
			margin: 30px 250px 20px 0;
		}
        .right {
			margin: 20px 0;
            border: 1px solid #DDD;
            border-radius: 4px;
            overflow: hidden;
            width: 220px;
            background-color: #F5F5F5;
			font-size: 0.9em;
			float: right;
        }

        .right ul {
            list-style-type: none;
            padding: 0;
            margin: 0;
        }

        .right a {
            display: block;
            text-decoration: none;
            color: #4183C4;
            padding: 8px 8px 8px 2em;
            border-bottom: 1px solid #E8E8E8;
        }

        .right a.h {
            color: #555;
            padding-left: 10px;
            background-color: #FFF;
			font-size: 1.2em;
        }
		.right a:hover {
			  background-color: antiquewhite;
		}
    </style>
</head>

<body>
    <div class="body">
        <div class="right">{{right}}</div>
        <div class="left">{{left}}</div>
    </div>
</body>

</html>
`

func generateDoc(base, src, out string) error {
	fileInfos, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	docs := list.New(fileInfos)
	docs.SetHomePage()
	docs.SetTopDoc()
	contents := docs.ToHTML()

	for _, v := range docs {
		if v.SrcName != "" {
			if err := md2HTML(src+v.SrcName, out+v.OutName, contents); err != nil {
				return err
			}
		}

		for _, w := range v.Sub {
			if err := md2HTML(src+w.SrcName, out+w.OutName, contents); err != nil {
				return err
			}
		}
	}

	return nil
}

func md2HTML(srcName, outName, contents string) error {
	data, err := ioutil.ReadFile(srcName)
	if err != nil {
		return err
	}

	// markdown to html
	html := blackfriday.MarkdownCommon(data)

	out := strings.Replace(
		strings.Replace(tmpl, "{{right}}", contents, 1),
		"{{left}}",
		string(html),
		1)

	// write file
	return ioutil.WriteFile(outName, []byte(out), 0644)
}
