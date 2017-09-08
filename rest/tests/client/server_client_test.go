package clienttest

import (
	check "gopkg.in/check.v1"

	"github.com/radanalyticsio/oshinko-cli/rest/version"
	"github.com/radanalyticsio/oshinko-cli/rest/helpers/info"
	"os"
	"fmt"
)

func (s *OshinkoRestTestSuite) TestServerInfo(c *check.C) {
	fmt.Println("starting server test")
	resp, _ := s.cli.Server.GetServerInfo(nil)
	fmt.Println("after request")

	expectedName := version.GetAppName()
	expectedVersion := version.GetVersion()
	expectedImage := info.GetSparkImage()

	observedName := resp.Payload.Application.Name
	observedVersion := resp.Payload.Application.Version
	observedImage := resp.Payload.Application.DefaultClusterImage

	c.Assert(*observedName, check.Equals, expectedName)
	c.Assert(*observedVersion, check.Equals, expectedVersion)
	c.Assert(*observedImage, check.Equals, expectedImage)

	os.Setenv("OSHINKO_CLUSTER_IMAGE", "bobby")
	expectedImage = "bobby"
	fmt.Println("second request")
	resp, _ = s.cli.Server.GetServerInfo(nil)
	fmt.Println("after second request")
	observedImage = resp.Payload.Application.DefaultClusterImage
	c.Assert(*observedImage, check.Equals, expectedImage)
	fmt.Println("ending server test " + *observedImage)
}
