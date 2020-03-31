package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

// Chapter type has three field ie. Title, Paragraphs of []string type and Options of []Option Type.
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option type is used for the option used for the navigation purpose. It is used in the bottom of our application.
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

// It is map of path-name and Chapter. We are parsing our json file into this Story map.
type Story map[string]Chapter

// JsonDecode is used to decode the json file of type io.Reader into Story map. It is called from the main()
func JsonDecode(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	err := d.Decode(&story)
	if err != nil {
		return nil, err
	}
	return story, nil
}

var tpl *template.Template

// default template is initialized at the load of the program.
func init() {
	tpl = template.Must(template.New("").Parse(defaultTemplate))
}

type HandlerOption func(h *handler)

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		if t != nil {
			h.t = t
		} else {
			h.t = tpl
		}

	}
}

// NewHandler types story and optional HandlerOptions, basis that it serves the template and story.
func NewHandler(s Story, opts ...HandlerOption) http.Handler {
	h := handler{s, tpl}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

type handler struct {
	s Story
	t *template.Template
}

// ServeHTTP method on h handler is used to find the url path, and then check the same path in the story map,
// if finds it then it serves the chapter. Otherwise returns a NotFoundErr.
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]

	if chapter, ok := h.s[path]; ok {
		err := h.t.Execute(w, chapter)
		if err != nil {
			log.Printf("%v\n", err)
			http.Error(w, "Error while Executing the template", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Path not found", http.StatusNotFound)
}

// defaultTemplate is the html-css template that is rendered for our web application.
var defaultTemplate = `<!DOCTYPE html>
			<html>
				<head>
					<meta charset="utf-8">
					<title>Choose Your own adventure</title>
				</head>

				<body >
				<section class="page">
					<h1>{{.Title}} </h1>
					{{range .Paragraphs}}
					<p>{{.}} </p>
					{{end}}

					<ul>
						{{range .Options}}
						<li><a href="/{{.Chapter}}">{{.Text}}</a> </li>
						{{end}}
					</ul>
					</section>
					<style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FFFCF6;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #777;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: none;
        color: #6295b5;
      }
      a:active,
      a:hover {
        color: #7792a2;
      }
      p {
        text-indent: 1em;
      }
    </style>
				</body>
			</html>`
