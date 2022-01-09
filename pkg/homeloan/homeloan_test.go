package homeloan_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/AxiomSamarth/home-loan-optimizer/pkg/homeloan"
)

var _ = Describe("Homeloan", func() {
	var (
		req     *http.Request
		handler http.Handler
		writer  *httptest.ResponseRecorder
	)
	Describe("HandlePOST", func() {
		BeforeEach(func() {
			handler = http.HandlerFunc(homeloan.Strategies)
			data := homeloan.Data{
				AnnualInterestRate: 7.15,
				LoanAmount:         5000000,
				LoanTenure:         15,
			}
			dataBytes, _ := json.Marshal(data)
			req = httptest.NewRequest(http.MethodPost, "/strategies", bytes.NewBuffer(dataBytes))
			writer = httptest.NewRecorder()
		})
		It("Processes a POST request successfully", func() {
			handler.ServeHTTP(writer, req)
			Expect(writer.Code).To(Equal(http.StatusOK))
		})
	})
})
