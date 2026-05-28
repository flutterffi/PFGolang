package tenantx

import "net/http"

func Extract(r *http.Request) string {
	tenant := r.Header.Get("X-Tenant-ID")
	if tenant == "" {
		return "default"
	}
	return tenant
}
