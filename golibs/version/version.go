package version

import (
	"fmt"
	"runtime"
	"strings"
)

// go build -ldflags "-X cqserver/golibs/version.BuildDate=`date '+%F-%T'` -X cqserver/golibs/version.GitCommit=`git rev-parse HEAD`";

var (
	GitCommit string
	BuildDate string
)

func VersionDetail(appName string, appVersion string) string {
	if len(strings.TrimSpace(GitCommit)) == 0 {
		GitCommit = "unknown"
	}
	if len(strings.TrimSpace(BuildDate)) == 0 {
		BuildDate = "0000-00-00"
	}
	return fmt.Sprintf("%s %s (%s %s %s %s %s)", appName, appVersion, runtime.Version(), runtime.GOOS, runtime.GOARCH, GitCommit, BuildDate)
}
