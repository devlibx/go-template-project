package e2e

import (
	"fmt"
	"github.com/devlibx/gox-base"
	"github.com/devlibx/gox-base/serialization"
	"github.com/zeebo/assert"
	"testing"
)

func (s *e2eTestSuite) TestPostApi() {
	s.T().Run("Get Post - Success", func(t *testing.T) {
		resp, err := s.restyClient.R().
			SetHeader("Content-Type", "application/json").
			Get("/post/1")
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())
		respMap := gox.StringObjectMap{}
		err = serialization.JsonBytesToObject(resp.Body(), &respMap)
		assert.NoError(t, err)
		fmt.Println("Get Post - Success Result\n", respMap.JsonPrettyStringIgnoreError())
		assert.Equal(t, 1, respMap.IntOrZero("id"))
	})
}
