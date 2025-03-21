//go:build all || !chat

package server

import "golang-clean-architecture/internal/chat"

func (r *Resource) InitModule() {
	chat.InitModule(r.Cfg, r.App, r.HTTPClient, r.DBConn)
}
