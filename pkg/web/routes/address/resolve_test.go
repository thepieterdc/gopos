package address

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/thepieterdc/gopos/pkg/web"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Address
const apple = "Apple%2010955%20N%20Tantau%20Ave,%20Cupertino,%20CA%2095014,United%20States"

func TestResolveViaLibPostal(t *testing.T) {
	// The resolver name.
	const resolverName = "libpostal"

	// Build the request.
	e := echo.New()
	e.Validator = &web.Validator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/address/resolve?resolver=%s&query=%s", resolverName, apple), nil)
	res := httptest.NewRecorder()
	ctx := e.NewContext(req, res)

	// Assertions.
	if assert.NoError(t, ResolveHandler(ctx)) {
		assert.Equal(t, http.StatusOK, res.Code)

		// Parse the response.
		var response resolveResponse
		assert.NoError(t, json.Unmarshal(res.Body.Bytes(), &response))

		// Validate the response.
		assert.Equal(t, resolverName, response.Resolver)
		assert.NotEmpty(t, response.DisplayAddress)
		assert.Equal(t, "cupertino", response.AddressInfo.City)
		assert.Equal(t, "united states", response.AddressInfo.Country)
		assert.Equal(t, "california", response.AddressInfo.StateOrProvince)
		assert.Equal(t, "north tantau avenue", response.AddressInfo.Street)
		assert.Equal(t, "10955", response.AddressInfo.StreetNumber)
	}
}
