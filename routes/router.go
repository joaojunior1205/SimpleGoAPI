package routes

import (
	"context"
	"net/http"
	"regexp"
	"strings"
)

type Middleware func(http.Handler) http.Handler

type Router struct {
	routes      []*routeEntry
	middlewares []Middleware
	prefix      string
}

type routeEntry struct {
	method  string
	pattern string
	regex   *regexp.Regexp
	params  []string
	handler http.Handler
}

// NewRouter cria um novo router
func NewRouter() *Router {
	return &Router{}
}

// Use adiciona middleware
func (r *Router) Use(mw Middleware) {
	r.middlewares = append(r.middlewares, mw)
}

// ServeHTTP implementa http.Handler
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	for _, rt := range r.routes {
		if rt.method != req.Method && rt.method != "ANY" {
			continue
		}
		matches := rt.regex.FindStringSubmatch(path)
		if matches != nil {
			params := map[string]string{}
			for i, name := range rt.params {
				params[name] = matches[i+1]
			}
			type contextKey string
			const paramsKey contextKey = "params"
			ctx := context.WithValue(req.Context(), paramsKey, params)
			rt.handler.ServeHTTP(w, req.WithContext(ctx))
			return
		}
	}
	http.NotFound(w, req)
}

// GET, POST, PUT, DELETE helpers
func (r *Router) GET(path string, handler http.HandlerFunc) {
	r.Handle("GET", path, handler)
}

func (r *Router) POST(path string, handler http.HandlerFunc) {
	r.Handle("POST", path, handler)
}

func (r *Router) PUT(path string, handler http.HandlerFunc) {
	r.Handle("PUT", path, handler)
}

func (r *Router) DELETE(path string, handler http.HandlerFunc) {
	r.Handle("DELETE", path, handler)
}

func (r *Router) Any(path string, handler http.HandlerFunc) {
	r.Handle("ANY", path, handler)
}

// Handle registra a rota
func (r *Router) Handle(method, path string, handler http.HandlerFunc) {
	fullPath := r.prefix + path
	regex, params := pathToRegex(fullPath)

	h := http.Handler(handler)
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		h = r.middlewares[i](h)
	}

	r.routes = append(r.routes, &routeEntry{
		method:  method,
		pattern: fullPath,
		regex:   regex,
		params:  params,
		handler: h,
	})
}

// Group cria um sub-router com prefixo
func (r *Router) Group(prefix string, fn func(sub *Router)) {
	sub := &Router{
		middlewares: append([]Middleware{}, r.middlewares...),
		prefix:      r.prefix + prefix,
	}
	fn(sub)
	r.routes = append(r.routes, sub.routes...)
}

// converte /users/:id em regex e lista de parâmetros
func pathToRegex(path string) (*regexp.Regexp, []string) {
	parts := strings.Split(path, "/")
	params := []string{}
	for i, part := range parts {
		if strings.HasPrefix(part, ":") {
			params = append(params, part[1:])
			parts[i] = "([^/]+)"
		} else {
			parts[i] = regexp.QuoteMeta(part)
		}
	}
	regex := regexp.MustCompile("^" + strings.Join(parts, "/") + "$")
	return regex, params
}

// Pega parâmetro da requisição
func Param(r *http.Request, name string) string {
	params := r.Context().Value("params")
	if paramsMap, ok := params.(map[string]string); ok {
		return paramsMap[name]
	}
	return ""
}
