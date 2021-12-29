// Package swagger embeds (hopefully) latest version of Swagger UI static files
// into a http.Handler
//
// Usage example:
//
// import (
//    "github.com/go-mixins/swagger-ui"
// )
//
// http.Handle(swagger.Handler(func()([]byte, error) {
//    return []byte(```
//    ... swagger.json contents, can be embedded as well ...
//    ```)
// })
//
// ...
//
// http.ListenAndServe(":8080", nil)
//
package swagger

import (
	"embed"
	"io/fs"
	"mime"
	"net/http"
	"os"
	"strconv"
)

//go:embed swagger-ui
var SwaggerFS embed.FS

// Handler provides http.Handler that serves swagger-ui directory contents or
// swagger.json as a result of specified function
func Handler(json func() ([]byte, error)) http.HandlerFunc {
	mime.AddExtensionType(".svg", "image/svg+xml")
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/swagger.json" {
			data, err := json()
			if os.IsNotExist(err) {
				http.Error(w, err.Error(), 404)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Content-Length", strconv.Itoa(len(data)))
			w.Write(data)
			return
		}
		fsys, err := fs.Sub(SwaggerFS, "swagger-ui")
		if err != nil {
			return
		}
		http.FileServer(http.FS(fsys)).ServeHTTP(w, r)
	}
}
