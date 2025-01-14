package app

import (
	"os"

	"github.com/weilinux/go-gin-skeleton-auth/helper"
	"github.com/weilinux/go-gin-skeleton-auth/model"
)

// allowed app env name
const (
	EnvProd = "prod"
	EnvPre  = "pre"
	EnvTest = "test"
	EnvDev  = "dev"
)

// for application
const (
	Timezone   = "PRC"
	DateFormat = "2006-01-02 15:04:05"

	Timeout     = 10
	PageSize    = 20
	PageSizeStr = "20"
	MaxPageSize = 100

	configSuffix = ".toml"
)

var (
	Name    = "github.com/weilinux/go-gin-skeleton-auth"
	EnvName = EnvDev

	Debug    bool
	Hostname string
	RootPath string
	// AbsPath always return abs path.
	AbsPath = helper.GetRootPath()

	GitInfo model.AppInfo

	HttpPort = 9550
)

// the app work dir path
var WorkDir, _ = os.Getwd()

// IsEnv current env name check
func IsEnv(env string) bool {
	return env == EnvName
}
