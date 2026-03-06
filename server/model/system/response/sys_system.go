package response

import "hz-admin-base/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
