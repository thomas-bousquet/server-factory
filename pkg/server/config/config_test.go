package config_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/thomas-bousquet/server-factory/pkg/server/config"

	"github.com/jaswdr/faker"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}

var _ = Describe("Config tests", func() {
	f := faker.New()

	var appName string
	var appVersion string
	var env string

	BeforeEach(func() {
		os.Clearenv()
		appName = f.Lorem().Word()
		appVersion = fmt.Sprintf("%d.%d.%d", f.RandomNumber(1), f.RandomNumber(1), f.RandomNumber(1))
		env = f.Lorem().Word()
	})

	When("Config creation succeeds", func() {
		It("should create config with minimum required env variables", func() {
			os.Setenv("X_APP_NAME", appName)
			os.Setenv("X_APP_VERSION", appVersion)
			os.Setenv("X_ENV", env)

			c, err := config.NewConfig()
			Expect(err).Error().To(BeNil())
			Expect(c.AppName).To(Equal(appName))
			Expect(c.AppVersion).To(Equal(appVersion))
			Expect(c.Env).To(Equal(env))
		})

		It("should create config with all default values", func() {
			os.Setenv("X_APP_NAME", appName)
			os.Setenv("X_APP_VERSION", appVersion)
			os.Setenv("X_ENV", env)

			c, err := config.NewConfig()
			Expect(err).Error().To(BeNil())
			Expect(c.AppPort).To(Equal(8080))
			Expect(c.LogLevel).To(Equal("info"))
			Expect(c.ReadTimeout).To(Equal(5 * time.Second))
			Expect(c.WriteTimeout).To(Equal(5 * time.Second))
			Expect(c.IdleTimeout).To(Equal(5 * time.Second))
			Expect(c.GracefullShutdownTimeout).To(Equal(10 * time.Second))
		})

		It("should create config with all possible values set as env vars", func() {
			appPort := f.RandomNumber(4)

			os.Setenv("X_APP_NAME", appName)
			os.Setenv("X_APP_VERSION", appVersion)
			os.Setenv("X_ENV", env)
			os.Setenv("X_APP_PORT", strconv.Itoa(appPort))
			os.Setenv("X_LOG_LEVEL", "info")
			os.Setenv("X_READ_TIMEOUT", "1s")
			os.Setenv("X_WRITE_TIMEOUT", "2s")
			os.Setenv("X_IDLE_TIMEOUT", "3s")
			os.Setenv("X_GRACEFULL_SHUTDOWN_TIMEOUT", "4s")

			c, err := config.NewConfig()
			Expect(err).Error().To(BeNil())
			Expect(c.AppName).To(Equal(appName))
			Expect(c.AppVersion).To(Equal(appVersion))
			Expect(c.Env).To(Equal(env))
			Expect(c.AppPort).To(Equal(appPort))
			Expect(c.LogLevel).To(Equal("info"))
			Expect(c.ReadTimeout).To(Equal(1 * time.Second))
			Expect(c.WriteTimeout).To(Equal(2 * time.Second))
			Expect(c.IdleTimeout).To(Equal(3 * time.Second))
			Expect(c.GracefullShutdownTimeout).To(Equal(4 * time.Second))
		})
	})

	When("Config creation fails", func() {
		When("Required X_APP_NAME is missing", func() {
			It("should return error for required X_APP_NAME", func() {
				os.Setenv("X_APP_VERSION", appVersion)
				os.Setenv("X_ENV", env)

				_, err := config.NewConfig()
				Expect(err).Error().NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("X_APP_NAME"))
				Expect(err.Error()).To(ContainSubstring("required"))
			})
		})

		When("Required X_APP_VERSION is missing", func() {
			It("should return error for required X_APP_VERSION", func() {
				os.Setenv("X_APP_NAME", appName)
				os.Setenv("X_ENV", env)

				_, err := config.NewConfig()
				Expect(err).Error().NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("X_APP_VERSION"))
				Expect(err.Error()).To(ContainSubstring("required"))
			})
		})

		When("Required X_ENV is missing", func() {
			It("should return error for required X_ENV", func() {
				os.Setenv("X_APP_NAME", appName)
				os.Setenv("X_APP_VERSION", appVersion)

				_, err := config.NewConfig()
				Expect(err).Error().NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("X_ENV"))
				Expect(err.Error()).To(ContainSubstring("required"))
			})
		})
	})
})
