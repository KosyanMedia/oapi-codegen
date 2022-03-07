// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package server

import (
	"github.com/labstack/echo/v4"
	"sync"
)

var (
	lockEchoRouterMockCONNECT sync.RWMutex
	lockEchoRouterMockDELETE  sync.RWMutex
	lockEchoRouterMockGET     sync.RWMutex
	lockEchoRouterMockHEAD    sync.RWMutex
	lockEchoRouterMockOPTIONS sync.RWMutex
	lockEchoRouterMockPATCH   sync.RWMutex
	lockEchoRouterMockPOST    sync.RWMutex
	lockEchoRouterMockPUT     sync.RWMutex
	lockEchoRouterMockTRACE   sync.RWMutex
)

// Ensure, that EchoRouterMock does implement EchoRouter.
// If this is not the case, regenerate this file with moq.
var _ EchoRouter = &EchoRouterMock{}

// EchoRouterMock is a mock implementation of EchoRouter.
//
//     func TestSomethingThatUsesEchoRouter(t *testing.T) {
//
//         // make and configure a mocked EchoRouter
//         mockedEchoRouter := &EchoRouterMock{
//             CONNECTFunc: func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
// 	               panic("mock out the CONNECT method")
//             },
//             DELETEFunc: func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
// 	               panic("mock out the DELETE method")
//             },
//             GETFunc: func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
// 	               panic("mock out the GET method")
//             },
//             HEADFunc: func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
// 	               panic("mock out the HEAD method")
//             },
//             OPTIONSFunc: func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
// 	               panic("mock out the OPTIONS method")
//             },
//             PATCHFunc: func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
// 	               panic("mock out the PATCH method")
//             },
//             POSTFunc: func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
// 	               panic("mock out the POST method")
//             },
//             PUTFunc: func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
// 	               panic("mock out the PUT method")
//             },
//             TRACEFunc: func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
// 	               panic("mock out the TRACE method")
//             },
//         }
//
//         // use mockedEchoRouter in code that requires EchoRouter
//         // and then make assertions.
//
//     }
type EchoRouterMock struct {
	// CONNECTFunc mocks the CONNECT method.
	CONNECTFunc func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route

	// DELETEFunc mocks the DELETE method.
	DELETEFunc func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route

	// GETFunc mocks the GET method.
	GETFunc func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route

	// HEADFunc mocks the HEAD method.
	HEADFunc func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route

	// OPTIONSFunc mocks the OPTIONS method.
	OPTIONSFunc func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route

	// PATCHFunc mocks the PATCH method.
	PATCHFunc func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route

	// POSTFunc mocks the POST method.
	POSTFunc func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route

	// PUTFunc mocks the PUT method.
	PUTFunc func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route

	// TRACEFunc mocks the TRACE method.
	TRACEFunc func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route

	// calls tracks calls to the methods.
	calls struct {
		// CONNECT holds details about calls to the CONNECT method.
		CONNECT []struct {
			// Path is the path argument value.
			Path string
			// H is the h argument value.
			H echo.HandlerFunc
			// M is the m argument value.
			M []echo.MiddlewareFunc
		}
		// DELETE holds details about calls to the DELETE method.
		DELETE []struct {
			// Path is the path argument value.
			Path string
			// H is the h argument value.
			H echo.HandlerFunc
			// M is the m argument value.
			M []echo.MiddlewareFunc
		}
		// GET holds details about calls to the GET method.
		GET []struct {
			// Path is the path argument value.
			Path string
			// H is the h argument value.
			H echo.HandlerFunc
			// M is the m argument value.
			M []echo.MiddlewareFunc
		}
		// HEAD holds details about calls to the HEAD method.
		HEAD []struct {
			// Path is the path argument value.
			Path string
			// H is the h argument value.
			H echo.HandlerFunc
			// M is the m argument value.
			M []echo.MiddlewareFunc
		}
		// OPTIONS holds details about calls to the OPTIONS method.
		OPTIONS []struct {
			// Path is the path argument value.
			Path string
			// H is the h argument value.
			H echo.HandlerFunc
			// M is the m argument value.
			M []echo.MiddlewareFunc
		}
		// PATCH holds details about calls to the PATCH method.
		PATCH []struct {
			// Path is the path argument value.
			Path string
			// H is the h argument value.
			H echo.HandlerFunc
			// M is the m argument value.
			M []echo.MiddlewareFunc
		}
		// POST holds details about calls to the POST method.
		POST []struct {
			// Path is the path argument value.
			Path string
			// H is the h argument value.
			H echo.HandlerFunc
			// M is the m argument value.
			M []echo.MiddlewareFunc
		}
		// PUT holds details about calls to the PUT method.
		PUT []struct {
			// Path is the path argument value.
			Path string
			// H is the h argument value.
			H echo.HandlerFunc
			// M is the m argument value.
			M []echo.MiddlewareFunc
		}
		// TRACE holds details about calls to the TRACE method.
		TRACE []struct {
			// Path is the path argument value.
			Path string
			// H is the h argument value.
			H echo.HandlerFunc
			// M is the m argument value.
			M []echo.MiddlewareFunc
		}
	}
}

// CONNECT calls CONNECTFunc.
func (mock *EchoRouterMock) CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	if mock.CONNECTFunc == nil {
		panic("EchoRouterMock.CONNECTFunc: method is nil but EchoRouter.CONNECT was just called")
	}
	callInfo := struct {
		Path string
		H    echo.HandlerFunc
		M    []echo.MiddlewareFunc
	}{
		Path: path,
		H:    h,
		M:    m,
	}
	lockEchoRouterMockCONNECT.Lock()
	mock.calls.CONNECT = append(mock.calls.CONNECT, callInfo)
	lockEchoRouterMockCONNECT.Unlock()
	return mock.CONNECTFunc(path, h, m...)
}

// CONNECTCalls gets all the calls that were made to CONNECT.
// Check the length with:
//     len(mockedEchoRouter.CONNECTCalls())
func (mock *EchoRouterMock) CONNECTCalls() []struct {
	Path string
	H    echo.HandlerFunc
	M    []echo.MiddlewareFunc
} {
	var calls []struct {
		Path string
		H    echo.HandlerFunc
		M    []echo.MiddlewareFunc
	}
	lockEchoRouterMockCONNECT.RLock()
	calls = mock.calls.CONNECT
	lockEchoRouterMockCONNECT.RUnlock()
	return calls
}

// DELETE calls DELETEFunc.
func (mock *EchoRouterMock) DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	if mock.DELETEFunc == nil {
		panic("EchoRouterMock.DELETEFunc: method is nil but EchoRouter.DELETE was just called")
	}
	callInfo := struct {
		Path string
		H    echo.HandlerFunc
		M    []echo.MiddlewareFunc
	}{
		Path: path,
		H:    h,
		M:    m,
	}
	lockEchoRouterMockDELETE.Lock()
	mock.calls.DELETE = append(mock.calls.DELETE, callInfo)
	lockEchoRouterMockDELETE.Unlock()
	return mock.DELETEFunc(path, h, m...)
}

// DELETECalls gets all the calls that were made to DELETE.
// Check the length with:
//     len(mockedEchoRouter.DELETECalls())
func (mock *EchoRouterMock) DELETECalls() []struct {
	Path string
	H    echo.HandlerFunc
	M    []echo.MiddlewareFunc
} {
	var calls []struct {
		Path string
		H    echo.HandlerFunc
		M    []echo.MiddlewareFunc
	}
	lockEchoRouterMockDELETE.RLock()
	calls = mock.calls.DELETE
	lockEchoRouterMockDELETE.RUnlock()
	return calls
}

// GET calls GETFunc.
func (mock *EchoRouterMock) GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	if mock.GETFunc == nil {
		panic("EchoRouterMock.GETFunc: method is nil but EchoRouter.GET was just called")
	}
	callInfo := struct {
		Path string
		H    echo.HandlerFunc
		M    []echo.MiddlewareFunc
	}{
		Path: path,
		H:    h,
		M:    m,
	}
	lockEchoRouterMockGET.Lock()
	mock.calls.GET = append(mock.calls.GET, callInfo)
	lockEchoRouterMockGET.Unlock()
	return mock.GETFunc(path, h, m...)
}

// GETCalls gets all the calls that were made to GET.
// Check the length with:
//     len(mockedEchoRouter.GETCalls())
func (mock *EchoRouterMock) GETCalls() []struct {
	Path string
	H    echo.HandlerFunc
	M    []echo.MiddlewareFunc
} {
	var calls []struct {
		Path string
		H    echo.HandlerFunc
		M    []echo.MiddlewareFunc
	}
	lockEchoRouterMockGET.RLock()
	calls = mock.calls.GET
	lockEchoRouterMockGET.RUnlock()
	return calls
}

// HEAD calls HEADFunc.
func (mock *EchoRouterMock) HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	if mock.HEADFunc == nil {
		panic("EchoRouterMock.HEADFunc: method is nil but EchoRouter.HEAD was just called")
	}
	callInfo := struct {
		Path string
		H    echo.HandlerFunc
		M    []echo.MiddlewareFunc
	}{
		Path: path,
		H:    h,
		M:    m,
	}
	lockEchoRouterMockHEAD.Lock()
	mock.calls.HEAD = append(mock.calls.HEAD, callInfo)
	lockEchoRouterMockHEAD.Unlock()
	return mock.HEADFunc(path, h, m...)
}

// HEADCalls gets all the calls that were made to HEAD.
// Check the length with:
//     len(mockedEchoRouter.HEADCalls())
func (mock *EchoRouterMock) HEADCalls() []struct {
	Path string
	H    echo.HandlerFunc
	M    []echo.MiddlewareFunc
} {
	var calls []struct {
		Path string
		H    echo.HandlerFunc
		M    []echo.MiddlewareFunc
	}
	lockEchoRouterMockHEAD.RLock()
	calls = mock.calls.HEAD
	lockEchoRouterMockHEAD.RUnlock()
	return calls
}

// OPTIONS calls OPTIONSFunc.
func (mock *EchoRouterMock) OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	if mock.OPTIONSFunc == nil {
		panic("EchoRouterMock.OPTIONSFunc: method is nil but EchoRouter.OPTIONS was just called")
	}
	callInfo := struct {
		Path string
		H    echo.HandlerFunc
		M    []echo.MiddlewareFunc
	}{
		Path: path,
		H:    h,
		M:    m,
	}
	lockEchoRouterMockOPTIONS.Lock()
	mock.calls.OPTIONS = append(mock.calls.OPTIONS, callInfo)
	lockEchoRouterMockOPTIONS.Unlock()
	return mock.OPTIONSFunc(path, h, m...)
}

// OPTIONSCalls gets all the calls that were made to OPTIONS.
// Check the length with:
//     len(mockedEchoRouter.OPTIONSCalls())
func (mock *EchoRouterMock) OPTIONSCalls() []struct {
	Path string
	H    echo.HandlerFunc
	M    []echo.MiddlewareFunc
} {
	var calls []struct {
		Path string
		H    echo.HandlerFunc
		M    []echo.MiddlewareFunc
	}
	lockEchoRouterMockOPTIONS.RLock()
	calls = mock.calls.OPTIONS
	lockEchoRouterMockOPTIONS.RUnlock()
	return calls
}

// PATCH calls PATCHFunc.
func (mock *EchoRouterMock) PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	if mock.PATCHFunc == nil {
		panic("EchoRouterMock.PATCHFunc: method is nil but EchoRouter.PATCH was just called")
	}
	callInfo := struct {
		Path string
		H    echo.HandlerFunc
		M    []echo.MiddlewareFunc
	}{
		Path: path,
		H:    h,
		M:    m,
	}
	lockEchoRouterMockPATCH.Lock()
	mock.calls.PATCH = append(mock.calls.PATCH, callInfo)
	lockEchoRouterMockPATCH.Unlock()
	return mock.PATCHFunc(path, h, m...)
}

// PATCHCalls gets all the calls that were made to PATCH.
// Check the length with:
//     len(mockedEchoRouter.PATCHCalls())
func (mock *EchoRouterMock) PATCHCalls() []struct {
	Path string
	H    echo.HandlerFunc
	M    []echo.MiddlewareFunc
} {
	var calls []struct {
		Path string
		H    echo.HandlerFunc
		M    []echo.MiddlewareFunc
	}
	lockEchoRouterMockPATCH.RLock()
	calls = mock.calls.PATCH
	lockEchoRouterMockPATCH.RUnlock()
	return calls
}

// POST calls POSTFunc.
func (mock *EchoRouterMock) POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	if mock.POSTFunc == nil {
		panic("EchoRouterMock.POSTFunc: method is nil but EchoRouter.POST was just called")
	}
	callInfo := struct {
		Path string
		H    echo.HandlerFunc
		M    []echo.MiddlewareFunc
	}{
		Path: path,
		H:    h,
		M:    m,
	}
	lockEchoRouterMockPOST.Lock()
	mock.calls.POST = append(mock.calls.POST, callInfo)
	lockEchoRouterMockPOST.Unlock()
	return mock.POSTFunc(path, h, m...)
}

// POSTCalls gets all the calls that were made to POST.
// Check the length with:
//     len(mockedEchoRouter.POSTCalls())
func (mock *EchoRouterMock) POSTCalls() []struct {
	Path string
	H    echo.HandlerFunc
	M    []echo.MiddlewareFunc
} {
	var calls []struct {
		Path string
		H    echo.HandlerFunc
		M    []echo.MiddlewareFunc
	}
	lockEchoRouterMockPOST.RLock()
	calls = mock.calls.POST
	lockEchoRouterMockPOST.RUnlock()
	return calls
}

// PUT calls PUTFunc.
func (mock *EchoRouterMock) PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	if mock.PUTFunc == nil {
		panic("EchoRouterMock.PUTFunc: method is nil but EchoRouter.PUT was just called")
	}
	callInfo := struct {
		Path string
		H    echo.HandlerFunc
		M    []echo.MiddlewareFunc
	}{
		Path: path,
		H:    h,
		M:    m,
	}
	lockEchoRouterMockPUT.Lock()
	mock.calls.PUT = append(mock.calls.PUT, callInfo)
	lockEchoRouterMockPUT.Unlock()
	return mock.PUTFunc(path, h, m...)
}

// PUTCalls gets all the calls that were made to PUT.
// Check the length with:
//     len(mockedEchoRouter.PUTCalls())
func (mock *EchoRouterMock) PUTCalls() []struct {
	Path string
	H    echo.HandlerFunc
	M    []echo.MiddlewareFunc
} {
	var calls []struct {
		Path string
		H    echo.HandlerFunc
		M    []echo.MiddlewareFunc
	}
	lockEchoRouterMockPUT.RLock()
	calls = mock.calls.PUT
	lockEchoRouterMockPUT.RUnlock()
	return calls
}

// TRACE calls TRACEFunc.
func (mock *EchoRouterMock) TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	if mock.TRACEFunc == nil {
		panic("EchoRouterMock.TRACEFunc: method is nil but EchoRouter.TRACE was just called")
	}
	callInfo := struct {
		Path string
		H    echo.HandlerFunc
		M    []echo.MiddlewareFunc
	}{
		Path: path,
		H:    h,
		M:    m,
	}
	lockEchoRouterMockTRACE.Lock()
	mock.calls.TRACE = append(mock.calls.TRACE, callInfo)
	lockEchoRouterMockTRACE.Unlock()
	return mock.TRACEFunc(path, h, m...)
}

// TRACECalls gets all the calls that were made to TRACE.
// Check the length with:
//     len(mockedEchoRouter.TRACECalls())
func (mock *EchoRouterMock) TRACECalls() []struct {
	Path string
	H    echo.HandlerFunc
	M    []echo.MiddlewareFunc
} {
	var calls []struct {
		Path string
		H    echo.HandlerFunc
		M    []echo.MiddlewareFunc
	}
	lockEchoRouterMockTRACE.RLock()
	calls = mock.calls.TRACE
	lockEchoRouterMockTRACE.RUnlock()
	return calls
}
