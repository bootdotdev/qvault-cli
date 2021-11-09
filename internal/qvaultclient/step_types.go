package qvaultclient

import "net/http"

// CheckTypes
const (
	CheckTypeGoVersion       = "GO_VERSION"
	CheckTypeGoBuild         = "GO_BUILD"
	CheckTypeGoTest          = "GO_TEST"
	CheckTypeHTTPLocalGet    = "HTTP_LOCAL_GET"
	CheckTypeHTTPLocalPost   = "HTTP_LOCAL_POST"
	CheckTypeHTTPLocalPut    = "HTTP_LOCAL_PUT"
	CheckTypeHTTPLocalDelete = "HTTP_LOCAL_DELETE"
	CheckTypeProjectComplete = "PROJECT_COMPLETE"
)

// GoTestExpectedOutputs
const (
	GoTestExpectedOutputPass = "PASS"
	GoTestExpectedOutputFail = "FAIL"
	GoTestExpectedOutputNone = "NONE"
)

// CheckCurrentStepResponseBody -
type CheckCurrentStepResponseBody struct {
	Passed       bool
	ErrorMessage string
}

// CheckCurrentStepCheckDataGoTest -
type CheckCurrentStepCheckDataGoTest struct {
	Modules map[string]struct {
		ExpectedOutput string
	}
}

// CheckCurrentStepCheckData -
type CheckCurrentStepCheckData struct {
	Checks []struct {
		CheckData map[string]interface{}
	}
}

// CheckCurrentStepCheckDataHTTPLocalGet -
type CheckCurrentStepCheckDataHTTPLocalGet struct {
	BodyString   string
	ResponseCode int
	Headers      http.Header
}

// CheckCurrentStepCheckDataHTTPLocalPost -
type CheckCurrentStepCheckDataHTTPLocalPost struct {
	BodyString   string
	ResponseCode int
	Headers      http.Header
}

// CheckCurrentStepCheckDataHTTPLocalPut -
type CheckCurrentStepCheckDataHTTPLocalPut struct {
	BodyString   string
	ResponseCode int
	Headers      http.Header
}

// CheckCurrentStepCheckDataHTTPLocalDelete -
type CheckCurrentStepCheckDataHTTPLocalDelete struct {
	BodyString   string
	ResponseCode int
	Headers      http.Header
}

// GetCurrentStepResponseBody -
type GetCurrentStepResponseBody struct {
	StepNumber    int
	NumberOfSteps int
	Checks        []struct {
		CheckType string
		CheckData map[string]interface{}
	}
}

// GetCurrentStepCheckDataHTTPLocalGet -
type GetCurrentStepCheckDataHTTPLocalGet struct {
	FullURL string
	Headers http.Header
}

// GetCurrentStepCheckDataHTTPLocalPost -
type GetCurrentStepCheckDataHTTPLocalPost struct {
	FullURL    string
	Headers    http.Header
	BodyString string
}

// GetCurrentStepCheckDataHTTPLocalPut -
type GetCurrentStepCheckDataHTTPLocalPut struct {
	FullURL    string
	Headers    http.Header
	BodyString string
}

// GetCurrentStepCheckDataHTTPLocalDelete -
type GetCurrentStepCheckDataHTTPLocalDelete struct {
	FullURL string
	Headers http.Header
}
