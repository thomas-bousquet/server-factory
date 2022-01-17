package server_test

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/thomas-bousquet/server-factory/pkg/server"

	"github.com/jaswdr/faker"
)

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Server Suite")
}

var _ = Describe("Config tests", func() {
	f := faker.New()

	var appName string
	var appVersion string

	BeforeEach(func() {
		os.Clearenv()
		appName = f.Lorem().Word()
		appVersion = fmt.Sprintf("%d.%d.%d", f.RandomNumber(1), f.RandomNumber(1), f.RandomNumber(1))

		os.Setenv("X_APP_NAME", appName)
		os.Setenv("X_APP_VERSION", appVersion)
		os.Setenv("X_ENV", "dev")
	})

	When("Server creation succeeds", func() {
		It("should create a server", func() {
			server.NewServer()
		})
	})

	When("Server serve succeeds", func() {
		It("should respond 200 to built-in healthcheck", func() {
			server := server.NewServer()
			go func() {
				server.Serve()
			}()

			resp, err := http.Get("http://localhost:8080/health")
			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
		})
	})
})
