package logger_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/thomas-bousquet/server-factory/pkg/server/config"
	"github.com/thomas-bousquet/server-factory/pkg/server/logger"
)

func TestLogger(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Logger Suite")
}

var _ = Describe("Logger tests", func() {
	When("Logger creation succeeds", func() {
		It("should create a logger", func() {
			logger, err := logger.NewLogger(config.NewTestConfig())
			Expect(err).Error().To(BeNil())
			Expect(logger).ToNot(BeNil())
		})
	})
})
