[![godoc.org](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/{{ .Repo.URL }})
[![pkg.go.dev](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/{{ .Repo.URL }}?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/{{ .Repo.URL }})](https://goreportcard.com/report/{{ .Repo.URL }})

# Vue.js SPA / Go API template

## Features

* Generate `README.md` using template `README.md.tpl`, injecting frontend and backend dependency information, and table of contents (ToC).
* Flat `frontend` directory.
* SASS instead of CSS.

## Usage

```sh
git clone git@github.com:virvum/vuejs-golang-template.git app
cd app
rm -rf .git
vi app.yaml
make setup
git init
git add -A
git commit -m 'initial commit'
git remote add origin ...
git push -u origin master
```

__TOC__

## Dependencies

### During runtime

* ...

### Backend
{{ range .Deps.Backend }}
* [{{ .Name }}]({{ .URL }}) {{ .Version }}
{{- end }}

### Frontend
{{ range .Deps.Frontend }}
* [{{ .Name }}]({{ .URL }}) {{ .Version }}
{{- end }}

## Development commands

```sh
npm install		# frontend: install dependencies
npm update		# frontend: update dependencies
npm run serve		# frontend: compiles and hot-reloads for development
npm run build		# frontend: compiles and minifies for production
npm run lint		# frontend: lints and fixes files

make build		# Build production binary
```
