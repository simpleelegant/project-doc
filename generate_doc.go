package main

import (
	"fmt"
	"io/ioutil"

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
            border: 1px solid #CECECE;
            border-radius: 4px;
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
            border: 0;
        }

        .body {
            height: auto;
            max-width: 980px;
            margin: 0px auto;
            position: relative;
        }

        .left {
            position: absolute;
            top: 0;
            left: 0;
            right: 250px;
        }

        .right {
            position: absolute;
            top: 20px;
            right: 0;
            border: 1px solid #DDD;
            border-radius: 4px;
            overflow: hidden;
            width: 220px;
            background-color: #F5F5F5;
			font-size: 0.9em;
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
            line-height: 32px;
            padding: 0 10px 0 2em;
            border-bottom: 1px solid #E8E8E8;
        }

        .right a.h {
            color: #555;
            padding-left: 10px;
            background-color: #FFF;
			font-size: 1.2em;
        }
    </style>
</head>

<body>
    <div class="body">
        <div class="left">%s</div>
        <div class="right">%s</div>
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

	// write file
	return ioutil.WriteFile(outName, []byte(fmt.Sprintf(tmpl, html, contents)), 0644)
}
