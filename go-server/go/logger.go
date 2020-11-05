/*
 * OpenTerps
 *
 * An Open API that contains information about terpenes, the effects, and the cannabis varieties that contain them.
 *
 * API version: 0.0.2
 * Contact: benjamminredden@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
    "log"
    "net/http"
    "time"
)

func Logger(inner http.Handler, name string) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        inner.ServeHTTP(w, r)

        log.Printf(
            "%s %s %s %s",
            r.Method,
            r.RequestURI,
            name,
            time.Since(start),
        )
    })
}
