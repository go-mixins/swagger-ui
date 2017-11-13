# Package swagger-ui

Embeds (hopefully) latest version of Swagger UI static files
into a http.Handler

Usage example:

```go
import (
   "github.com/go-mixins/swagger-ui"
)

http.Handle(swagger.Handler(func()([]byte, error) {
   return []byte(```
   ... swagger.json contents, can be embedded as well ...
   ```)
})

...

http.ListenAndServe(":8080", nil)
```

## Tools used

To re-generate embedded UI files [esc](https://github.com/mjibson/esc) tool must be installed.

## LICENSE

See [Swagger UI license](https://github.com/swagger-api/swagger-ui/blob/master/LICENSE).
