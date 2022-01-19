package polo

import (
	"net/http"
	"testing"

	"github.com/aemloviji/golang-guide-microservices/src/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "polo", polo)
}

func TestPolo(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/marco", nil)
	response, c := test_utils.GetMockedContext(request)

	Marco(c)

	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "polo", response.Body.String())
}
