package n_rpc

import "github.com/abc463774475/bbTool/n_log"

var (
	G_RemotePool *RemoteFamilyMgr
)

func init() {
	G_RemotePool = CreateMgr()
	n_log.Debug("rpc  init ok")
}
