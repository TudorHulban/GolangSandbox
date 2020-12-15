### Minimal application
Default metrics exposed at http://localhost:2112/metrics.
```
package main

import (
        "net/http"

        "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
        http.Handle("/metrics", promhttp.Handler())
        http.ListenAndServe(":2112", nil)
}
```

## Resources
```
https://prometheus.io/docs/guides/go-application/
https://github.com/zsais/go-gin-prometheus
```
