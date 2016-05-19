# project-doc

A simple, pretty tool, convert markdown to html. No config! Only one file in release!


### Install

If you have Go installed, do this:

```
go get github.com/simpleelegant/project-doc
```

Or, directly download the binary in [releases](https://github.com/simpleelegant/project-doc/releases).


### Usage

```bash
mkdir -p docs/src
cd docs/src

# edit other documents
vim README.md
vim News.md
vim News_list.md

# converting, output files will in docs/out
project-doc ../
```
