package webservice

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestDiceServiceError tests the situation where the external service we depend on return an error.
func TestDiceServiceError(t *testing.T) {
	// Testing objects
	mockSixSidedDie := &mockDie{}
	webservice := New(mockSixSidedDie)
	ts := httptest.NewServer(http.HandlerFunc(webservice.Roll))
	defer ts.Close()

	// Mock out external service values
	const testFakeErrorText = "fake error"
	mockSixSidedDie.setMockValues(0, errors.New(testFakeErrorText))

	// Test action
	responseText := getTestServerResponse(t, ts)

	// Assert
	assert.Equal(t, testFakeErrorText+"\n", responseText)
}

// TestDiceServicePassThrough is a table tests which tests a series of possible successful values we may get from our external service.
func TestDiceServicePassThrough(t *testing.T) {
	t.Parallel()

	for name, testCase := range map[string]struct {
		returnValue    int
		expectedOutput string
	}{
		"Negative value": {
			returnValue:    -5,
			expectedOutput: "-5\n",
		},
		"Zero value": {
			returnValue:    0,
			expectedOutput: "0\n",
		},
		"Positive - 1": {
			returnValue:    1,
			expectedOutput: "1\n",
		},
		"Positive integer in range": {
			returnValue:    4,
			expectedOutput: "4\n",
		},
		"Positive above range": {
			returnValue:    7,
			expectedOutput: "7\n",
		},
	} {
		testCase := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Testing objects
			mockSixSidedDie := &mockDie{}
			webservice := New(mockSixSidedDie)
			ts := httptest.NewServer(http.HandlerFunc(webservice.Roll))
			defer ts.Close()

			// Assertions
			mockSixSidedDie.setMockValues(testCase.returnValue, nil)
			assert.Equal(t, testCase.expectedOutput, getTestServerResponse(t, ts))
		})
	}
}

// A helper function to get the response from our test server.
func getTestServerResponse(t *testing.T, ts *httptest.Server) string {
	resp, err := http.Get(ts.URL)
	if err != nil {
		require.NoError(t, err)
	}

	responseText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		require.NoError(t, err)
	}

	require.NoError(t, resp.Body.Close())

	return string(responseText)
}

// A mocked out version of our external service.
type mockDie struct {
	returnVal int
	returnErr error
}

func (m *mockDie) setMockValues(returnVal int, returnErr error) {
	m.returnVal = returnVal
	m.returnErr = returnErr
}

func (m *mockDie) Roll6() (int, error) {
	return m.returnVal, m.returnErr
}
