// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package server

import (
	"github.com/labstack/echo/v4"
	"sync"
)

// Ensure, that ServerInterfaceMock does implement ServerInterface.
// If this is not the case, regenerate this file with moq.
var _ ServerInterface = &ServerInterfaceMock{}

// ServerInterfaceMock is a mock implementation of ServerInterface.
//
// 	func TestSomethingThatUsesServerInterface(t *testing.T) {
//
// 		// make and configure a mocked ServerInterface
// 		mockedServerInterface := &ServerInterfaceMock{
// 			CreateEveryTypeOptionalFunc: func(ctx echo.Context, params CreateEveryTypeOptionalParams, requestBody EveryTypeOptional) (int, error) {
// 				panic("mock out the CreateEveryTypeOptional method")
// 			},
// 			CreateResourceFunc: func(ctx echo.Context, argument string, requestBody EveryTypeRequired) (*CreateResourceResponse, error) {
// 				panic("mock out the CreateResource method")
// 			},
// 			CreateResource2Func: func(ctx echo.Context, inlineArgument int, params CreateResource2Params, requestBody Resource) (*CreateResource2Response, error) {
// 				panic("mock out the CreateResource2 method")
// 			},
// 			ErrorFunc: func(err error) (int, Error) {
// 				panic("mock out the Error method")
// 			},
// 			GetEveryTypeOptionalFunc: func(ctx echo.Context) (*GetEveryTypeOptionalResponse, error) {
// 				panic("mock out the GetEveryTypeOptional method")
// 			},
// 			GetReservedKeywordFunc: func(ctx echo.Context) (*GetReservedKeywordResponse, error) {
// 				panic("mock out the GetReservedKeyword method")
// 			},
// 			GetResponseWithReferenceFunc: func(ctx echo.Context) (*GetResponseWithReferenceResponse, error) {
// 				panic("mock out the GetResponseWithReference method")
// 			},
// 			GetSimpleFunc: func(ctx echo.Context) (*GetSimpleResponse, error) {
// 				panic("mock out the GetSimple method")
// 			},
// 			GetWithArgsFunc: func(ctx echo.Context, params GetWithArgsParams) (*GetWithArgsResponse, error) {
// 				panic("mock out the GetWithArgs method")
// 			},
// 			GetWithContentTypeFunc: func(ctx echo.Context, contentType GetWithContentTypeParamsContentType) (*GetWithContentTypeResponse, error) {
// 				panic("mock out the GetWithContentType method")
// 			},
// 			GetWithReferencesFunc: func(ctx echo.Context, globalArgument int64, argument string) (*GetWithReferencesResponse, error) {
// 				panic("mock out the GetWithReferences method")
// 			},
// 			UpdateResource3Func: func(ctx echo.Context, pFallthrough int, requestBody UpdateResource3JSONBody) (int, error) {
// 				panic("mock out the UpdateResource3 method")
// 			},
// 		}
//
// 		// use mockedServerInterface in code that requires ServerInterface
// 		// and then make assertions.
//
// 	}
type ServerInterfaceMock struct {
	// CreateEveryTypeOptionalFunc mocks the CreateEveryTypeOptional method.
	CreateEveryTypeOptionalFunc func(ctx echo.Context, params CreateEveryTypeOptionalParams, requestBody EveryTypeOptional) (int, error)

	// CreateResourceFunc mocks the CreateResource method.
	CreateResourceFunc func(ctx echo.Context, argument string, requestBody EveryTypeRequired) (*CreateResourceResponse, error)

	// CreateResource2Func mocks the CreateResource2 method.
	CreateResource2Func func(ctx echo.Context, inlineArgument int, params CreateResource2Params, requestBody Resource) (*CreateResource2Response, error)

	// ErrorFunc mocks the Error method.
	ErrorFunc func(err error) (int, Error)

	// GetEveryTypeOptionalFunc mocks the GetEveryTypeOptional method.
	GetEveryTypeOptionalFunc func(ctx echo.Context) (*GetEveryTypeOptionalResponse, error)

	// GetReservedKeywordFunc mocks the GetReservedKeyword method.
	GetReservedKeywordFunc func(ctx echo.Context) (*GetReservedKeywordResponse, error)

	// GetResponseWithReferenceFunc mocks the GetResponseWithReference method.
	GetResponseWithReferenceFunc func(ctx echo.Context) (*GetResponseWithReferenceResponse, error)

	// GetSimpleFunc mocks the GetSimple method.
	GetSimpleFunc func(ctx echo.Context) (*GetSimpleResponse, error)

	// GetWithArgsFunc mocks the GetWithArgs method.
	GetWithArgsFunc func(ctx echo.Context, params GetWithArgsParams) (*GetWithArgsResponse, error)

	// GetWithContentTypeFunc mocks the GetWithContentType method.
	GetWithContentTypeFunc func(ctx echo.Context, contentType GetWithContentTypeParamsContentType) (*GetWithContentTypeResponse, error)

	// GetWithReferencesFunc mocks the GetWithReferences method.
	GetWithReferencesFunc func(ctx echo.Context, globalArgument int64, argument string) (*GetWithReferencesResponse, error)

	// UpdateResource3Func mocks the UpdateResource3 method.
	UpdateResource3Func func(ctx echo.Context, pFallthrough int, requestBody UpdateResource3JSONBody) (int, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateEveryTypeOptional holds details about calls to the CreateEveryTypeOptional method.
		CreateEveryTypeOptional []struct {
			// Ctx is the ctx argument value.
			Ctx echo.Context
			// Params is the params argument value.
			Params CreateEveryTypeOptionalParams
			// RequestBody is the requestBody argument value.
			RequestBody EveryTypeOptional
		}
		// CreateResource holds details about calls to the CreateResource method.
		CreateResource []struct {
			// Ctx is the ctx argument value.
			Ctx echo.Context
			// Argument is the argument argument value.
			Argument string
			// RequestBody is the requestBody argument value.
			RequestBody EveryTypeRequired
		}
		// CreateResource2 holds details about calls to the CreateResource2 method.
		CreateResource2 []struct {
			// Ctx is the ctx argument value.
			Ctx echo.Context
			// InlineArgument is the inlineArgument argument value.
			InlineArgument int
			// Params is the params argument value.
			Params CreateResource2Params
			// RequestBody is the requestBody argument value.
			RequestBody Resource
		}
		// Error holds details about calls to the Error method.
		Error []struct {
			// Err is the err argument value.
			Err error
		}
		// GetEveryTypeOptional holds details about calls to the GetEveryTypeOptional method.
		GetEveryTypeOptional []struct {
			// Ctx is the ctx argument value.
			Ctx echo.Context
		}
		// GetReservedKeyword holds details about calls to the GetReservedKeyword method.
		GetReservedKeyword []struct {
			// Ctx is the ctx argument value.
			Ctx echo.Context
		}
		// GetResponseWithReference holds details about calls to the GetResponseWithReference method.
		GetResponseWithReference []struct {
			// Ctx is the ctx argument value.
			Ctx echo.Context
		}
		// GetSimple holds details about calls to the GetSimple method.
		GetSimple []struct {
			// Ctx is the ctx argument value.
			Ctx echo.Context
		}
		// GetWithArgs holds details about calls to the GetWithArgs method.
		GetWithArgs []struct {
			// Ctx is the ctx argument value.
			Ctx echo.Context
			// Params is the params argument value.
			Params GetWithArgsParams
		}
		// GetWithContentType holds details about calls to the GetWithContentType method.
		GetWithContentType []struct {
			// Ctx is the ctx argument value.
			Ctx echo.Context
			// ContentType is the contentType argument value.
			ContentType GetWithContentTypeParamsContentType
		}
		// GetWithReferences holds details about calls to the GetWithReferences method.
		GetWithReferences []struct {
			// Ctx is the ctx argument value.
			Ctx echo.Context
			// GlobalArgument is the globalArgument argument value.
			GlobalArgument int64
			// Argument is the argument argument value.
			Argument string
		}
		// UpdateResource3 holds details about calls to the UpdateResource3 method.
		UpdateResource3 []struct {
			// Ctx is the ctx argument value.
			Ctx echo.Context
			// PFallthrough is the pFallthrough argument value.
			PFallthrough int
			// RequestBody is the requestBody argument value.
			RequestBody UpdateResource3JSONBody
		}
	}
	lockCreateEveryTypeOptional  sync.RWMutex
	lockCreateResource           sync.RWMutex
	lockCreateResource2          sync.RWMutex
	lockError                    sync.RWMutex
	lockGetEveryTypeOptional     sync.RWMutex
	lockGetReservedKeyword       sync.RWMutex
	lockGetResponseWithReference sync.RWMutex
	lockGetSimple                sync.RWMutex
	lockGetWithArgs              sync.RWMutex
	lockGetWithContentType       sync.RWMutex
	lockGetWithReferences        sync.RWMutex
	lockUpdateResource3          sync.RWMutex
}

// CreateEveryTypeOptional calls CreateEveryTypeOptionalFunc.
func (mock *ServerInterfaceMock) CreateEveryTypeOptional(ctx echo.Context, params CreateEveryTypeOptionalParams, requestBody EveryTypeOptional) (int, error) {
	if mock.CreateEveryTypeOptionalFunc == nil {
		panic("ServerInterfaceMock.CreateEveryTypeOptionalFunc: method is nil but ServerInterface.CreateEveryTypeOptional was just called")
	}
	callInfo := struct {
		Ctx         echo.Context
		Params      CreateEveryTypeOptionalParams
		RequestBody EveryTypeOptional
	}{
		Ctx:         ctx,
		Params:      params,
		RequestBody: requestBody,
	}
	mock.lockCreateEveryTypeOptional.Lock()
	mock.calls.CreateEveryTypeOptional = append(mock.calls.CreateEveryTypeOptional, callInfo)
	mock.lockCreateEveryTypeOptional.Unlock()
	return mock.CreateEveryTypeOptionalFunc(ctx, params, requestBody)
}

// CreateEveryTypeOptionalCalls gets all the calls that were made to CreateEveryTypeOptional.
// Check the length with:
//     len(mockedServerInterface.CreateEveryTypeOptionalCalls())
func (mock *ServerInterfaceMock) CreateEveryTypeOptionalCalls() []struct {
	Ctx         echo.Context
	Params      CreateEveryTypeOptionalParams
	RequestBody EveryTypeOptional
} {
	var calls []struct {
		Ctx         echo.Context
		Params      CreateEveryTypeOptionalParams
		RequestBody EveryTypeOptional
	}
	mock.lockCreateEveryTypeOptional.RLock()
	calls = mock.calls.CreateEveryTypeOptional
	mock.lockCreateEveryTypeOptional.RUnlock()
	return calls
}

// CreateResource calls CreateResourceFunc.
func (mock *ServerInterfaceMock) CreateResource(ctx echo.Context, argument string, requestBody EveryTypeRequired) (*CreateResourceResponse, error) {
	if mock.CreateResourceFunc == nil {
		panic("ServerInterfaceMock.CreateResourceFunc: method is nil but ServerInterface.CreateResource was just called")
	}
	callInfo := struct {
		Ctx         echo.Context
		Argument    string
		RequestBody EveryTypeRequired
	}{
		Ctx:         ctx,
		Argument:    argument,
		RequestBody: requestBody,
	}
	mock.lockCreateResource.Lock()
	mock.calls.CreateResource = append(mock.calls.CreateResource, callInfo)
	mock.lockCreateResource.Unlock()
	return mock.CreateResourceFunc(ctx, argument, requestBody)
}

// CreateResourceCalls gets all the calls that were made to CreateResource.
// Check the length with:
//     len(mockedServerInterface.CreateResourceCalls())
func (mock *ServerInterfaceMock) CreateResourceCalls() []struct {
	Ctx         echo.Context
	Argument    string
	RequestBody EveryTypeRequired
} {
	var calls []struct {
		Ctx         echo.Context
		Argument    string
		RequestBody EveryTypeRequired
	}
	mock.lockCreateResource.RLock()
	calls = mock.calls.CreateResource
	mock.lockCreateResource.RUnlock()
	return calls
}

// CreateResource2 calls CreateResource2Func.
func (mock *ServerInterfaceMock) CreateResource2(ctx echo.Context, inlineArgument int, params CreateResource2Params, requestBody Resource) (*CreateResource2Response, error) {
	if mock.CreateResource2Func == nil {
		panic("ServerInterfaceMock.CreateResource2Func: method is nil but ServerInterface.CreateResource2 was just called")
	}
	callInfo := struct {
		Ctx            echo.Context
		InlineArgument int
		Params         CreateResource2Params
		RequestBody    Resource
	}{
		Ctx:            ctx,
		InlineArgument: inlineArgument,
		Params:         params,
		RequestBody:    requestBody,
	}
	mock.lockCreateResource2.Lock()
	mock.calls.CreateResource2 = append(mock.calls.CreateResource2, callInfo)
	mock.lockCreateResource2.Unlock()
	return mock.CreateResource2Func(ctx, inlineArgument, params, requestBody)
}

// CreateResource2Calls gets all the calls that were made to CreateResource2.
// Check the length with:
//     len(mockedServerInterface.CreateResource2Calls())
func (mock *ServerInterfaceMock) CreateResource2Calls() []struct {
	Ctx            echo.Context
	InlineArgument int
	Params         CreateResource2Params
	RequestBody    Resource
} {
	var calls []struct {
		Ctx            echo.Context
		InlineArgument int
		Params         CreateResource2Params
		RequestBody    Resource
	}
	mock.lockCreateResource2.RLock()
	calls = mock.calls.CreateResource2
	mock.lockCreateResource2.RUnlock()
	return calls
}

// Error calls ErrorFunc.
func (mock *ServerInterfaceMock) Error(err error) (int, Error) {
	if mock.ErrorFunc == nil {
		panic("ServerInterfaceMock.ErrorFunc: method is nil but ServerInterface.Error was just called")
	}
	callInfo := struct {
		Err error
	}{
		Err: err,
	}
	mock.lockError.Lock()
	mock.calls.Error = append(mock.calls.Error, callInfo)
	mock.lockError.Unlock()
	return mock.ErrorFunc(err)
}

// ErrorCalls gets all the calls that were made to Error.
// Check the length with:
//     len(mockedServerInterface.ErrorCalls())
func (mock *ServerInterfaceMock) ErrorCalls() []struct {
	Err error
} {
	var calls []struct {
		Err error
	}
	mock.lockError.RLock()
	calls = mock.calls.Error
	mock.lockError.RUnlock()
	return calls
}

// GetEveryTypeOptional calls GetEveryTypeOptionalFunc.
func (mock *ServerInterfaceMock) GetEveryTypeOptional(ctx echo.Context) (*GetEveryTypeOptionalResponse, error) {
	if mock.GetEveryTypeOptionalFunc == nil {
		panic("ServerInterfaceMock.GetEveryTypeOptionalFunc: method is nil but ServerInterface.GetEveryTypeOptional was just called")
	}
	callInfo := struct {
		Ctx echo.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetEveryTypeOptional.Lock()
	mock.calls.GetEveryTypeOptional = append(mock.calls.GetEveryTypeOptional, callInfo)
	mock.lockGetEveryTypeOptional.Unlock()
	return mock.GetEveryTypeOptionalFunc(ctx)
}

// GetEveryTypeOptionalCalls gets all the calls that were made to GetEveryTypeOptional.
// Check the length with:
//     len(mockedServerInterface.GetEveryTypeOptionalCalls())
func (mock *ServerInterfaceMock) GetEveryTypeOptionalCalls() []struct {
	Ctx echo.Context
} {
	var calls []struct {
		Ctx echo.Context
	}
	mock.lockGetEveryTypeOptional.RLock()
	calls = mock.calls.GetEveryTypeOptional
	mock.lockGetEveryTypeOptional.RUnlock()
	return calls
}

// GetReservedKeyword calls GetReservedKeywordFunc.
func (mock *ServerInterfaceMock) GetReservedKeyword(ctx echo.Context) (*GetReservedKeywordResponse, error) {
	if mock.GetReservedKeywordFunc == nil {
		panic("ServerInterfaceMock.GetReservedKeywordFunc: method is nil but ServerInterface.GetReservedKeyword was just called")
	}
	callInfo := struct {
		Ctx echo.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetReservedKeyword.Lock()
	mock.calls.GetReservedKeyword = append(mock.calls.GetReservedKeyword, callInfo)
	mock.lockGetReservedKeyword.Unlock()
	return mock.GetReservedKeywordFunc(ctx)
}

// GetReservedKeywordCalls gets all the calls that were made to GetReservedKeyword.
// Check the length with:
//     len(mockedServerInterface.GetReservedKeywordCalls())
func (mock *ServerInterfaceMock) GetReservedKeywordCalls() []struct {
	Ctx echo.Context
} {
	var calls []struct {
		Ctx echo.Context
	}
	mock.lockGetReservedKeyword.RLock()
	calls = mock.calls.GetReservedKeyword
	mock.lockGetReservedKeyword.RUnlock()
	return calls
}

// GetResponseWithReference calls GetResponseWithReferenceFunc.
func (mock *ServerInterfaceMock) GetResponseWithReference(ctx echo.Context) (*GetResponseWithReferenceResponse, error) {
	if mock.GetResponseWithReferenceFunc == nil {
		panic("ServerInterfaceMock.GetResponseWithReferenceFunc: method is nil but ServerInterface.GetResponseWithReference was just called")
	}
	callInfo := struct {
		Ctx echo.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetResponseWithReference.Lock()
	mock.calls.GetResponseWithReference = append(mock.calls.GetResponseWithReference, callInfo)
	mock.lockGetResponseWithReference.Unlock()
	return mock.GetResponseWithReferenceFunc(ctx)
}

// GetResponseWithReferenceCalls gets all the calls that were made to GetResponseWithReference.
// Check the length with:
//     len(mockedServerInterface.GetResponseWithReferenceCalls())
func (mock *ServerInterfaceMock) GetResponseWithReferenceCalls() []struct {
	Ctx echo.Context
} {
	var calls []struct {
		Ctx echo.Context
	}
	mock.lockGetResponseWithReference.RLock()
	calls = mock.calls.GetResponseWithReference
	mock.lockGetResponseWithReference.RUnlock()
	return calls
}

// GetSimple calls GetSimpleFunc.
func (mock *ServerInterfaceMock) GetSimple(ctx echo.Context) (*GetSimpleResponse, error) {
	if mock.GetSimpleFunc == nil {
		panic("ServerInterfaceMock.GetSimpleFunc: method is nil but ServerInterface.GetSimple was just called")
	}
	callInfo := struct {
		Ctx echo.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetSimple.Lock()
	mock.calls.GetSimple = append(mock.calls.GetSimple, callInfo)
	mock.lockGetSimple.Unlock()
	return mock.GetSimpleFunc(ctx)
}

// GetSimpleCalls gets all the calls that were made to GetSimple.
// Check the length with:
//     len(mockedServerInterface.GetSimpleCalls())
func (mock *ServerInterfaceMock) GetSimpleCalls() []struct {
	Ctx echo.Context
} {
	var calls []struct {
		Ctx echo.Context
	}
	mock.lockGetSimple.RLock()
	calls = mock.calls.GetSimple
	mock.lockGetSimple.RUnlock()
	return calls
}

// GetWithArgs calls GetWithArgsFunc.
func (mock *ServerInterfaceMock) GetWithArgs(ctx echo.Context, params GetWithArgsParams) (*GetWithArgsResponse, error) {
	if mock.GetWithArgsFunc == nil {
		panic("ServerInterfaceMock.GetWithArgsFunc: method is nil but ServerInterface.GetWithArgs was just called")
	}
	callInfo := struct {
		Ctx    echo.Context
		Params GetWithArgsParams
	}{
		Ctx:    ctx,
		Params: params,
	}
	mock.lockGetWithArgs.Lock()
	mock.calls.GetWithArgs = append(mock.calls.GetWithArgs, callInfo)
	mock.lockGetWithArgs.Unlock()
	return mock.GetWithArgsFunc(ctx, params)
}

// GetWithArgsCalls gets all the calls that were made to GetWithArgs.
// Check the length with:
//     len(mockedServerInterface.GetWithArgsCalls())
func (mock *ServerInterfaceMock) GetWithArgsCalls() []struct {
	Ctx    echo.Context
	Params GetWithArgsParams
} {
	var calls []struct {
		Ctx    echo.Context
		Params GetWithArgsParams
	}
	mock.lockGetWithArgs.RLock()
	calls = mock.calls.GetWithArgs
	mock.lockGetWithArgs.RUnlock()
	return calls
}

// GetWithContentType calls GetWithContentTypeFunc.
func (mock *ServerInterfaceMock) GetWithContentType(ctx echo.Context, contentType GetWithContentTypeParamsContentType) (*GetWithContentTypeResponse, error) {
	if mock.GetWithContentTypeFunc == nil {
		panic("ServerInterfaceMock.GetWithContentTypeFunc: method is nil but ServerInterface.GetWithContentType was just called")
	}
	callInfo := struct {
		Ctx         echo.Context
		ContentType GetWithContentTypeParamsContentType
	}{
		Ctx:         ctx,
		ContentType: contentType,
	}
	mock.lockGetWithContentType.Lock()
	mock.calls.GetWithContentType = append(mock.calls.GetWithContentType, callInfo)
	mock.lockGetWithContentType.Unlock()
	return mock.GetWithContentTypeFunc(ctx, contentType)
}

// GetWithContentTypeCalls gets all the calls that were made to GetWithContentType.
// Check the length with:
//     len(mockedServerInterface.GetWithContentTypeCalls())
func (mock *ServerInterfaceMock) GetWithContentTypeCalls() []struct {
	Ctx         echo.Context
	ContentType GetWithContentTypeParamsContentType
} {
	var calls []struct {
		Ctx         echo.Context
		ContentType GetWithContentTypeParamsContentType
	}
	mock.lockGetWithContentType.RLock()
	calls = mock.calls.GetWithContentType
	mock.lockGetWithContentType.RUnlock()
	return calls
}

// GetWithReferences calls GetWithReferencesFunc.
func (mock *ServerInterfaceMock) GetWithReferences(ctx echo.Context, globalArgument int64, argument string) (*GetWithReferencesResponse, error) {
	if mock.GetWithReferencesFunc == nil {
		panic("ServerInterfaceMock.GetWithReferencesFunc: method is nil but ServerInterface.GetWithReferences was just called")
	}
	callInfo := struct {
		Ctx            echo.Context
		GlobalArgument int64
		Argument       string
	}{
		Ctx:            ctx,
		GlobalArgument: globalArgument,
		Argument:       argument,
	}
	mock.lockGetWithReferences.Lock()
	mock.calls.GetWithReferences = append(mock.calls.GetWithReferences, callInfo)
	mock.lockGetWithReferences.Unlock()
	return mock.GetWithReferencesFunc(ctx, globalArgument, argument)
}

// GetWithReferencesCalls gets all the calls that were made to GetWithReferences.
// Check the length with:
//     len(mockedServerInterface.GetWithReferencesCalls())
func (mock *ServerInterfaceMock) GetWithReferencesCalls() []struct {
	Ctx            echo.Context
	GlobalArgument int64
	Argument       string
} {
	var calls []struct {
		Ctx            echo.Context
		GlobalArgument int64
		Argument       string
	}
	mock.lockGetWithReferences.RLock()
	calls = mock.calls.GetWithReferences
	mock.lockGetWithReferences.RUnlock()
	return calls
}

// UpdateResource3 calls UpdateResource3Func.
func (mock *ServerInterfaceMock) UpdateResource3(ctx echo.Context, pFallthrough int, requestBody UpdateResource3JSONBody) (int, error) {
	if mock.UpdateResource3Func == nil {
		panic("ServerInterfaceMock.UpdateResource3Func: method is nil but ServerInterface.UpdateResource3 was just called")
	}
	callInfo := struct {
		Ctx          echo.Context
		PFallthrough int
		RequestBody  UpdateResource3JSONBody
	}{
		Ctx:          ctx,
		PFallthrough: pFallthrough,
		RequestBody:  requestBody,
	}
	mock.lockUpdateResource3.Lock()
	mock.calls.UpdateResource3 = append(mock.calls.UpdateResource3, callInfo)
	mock.lockUpdateResource3.Unlock()
	return mock.UpdateResource3Func(ctx, pFallthrough, requestBody)
}

// UpdateResource3Calls gets all the calls that were made to UpdateResource3.
// Check the length with:
//     len(mockedServerInterface.UpdateResource3Calls())
func (mock *ServerInterfaceMock) UpdateResource3Calls() []struct {
	Ctx          echo.Context
	PFallthrough int
	RequestBody  UpdateResource3JSONBody
} {
	var calls []struct {
		Ctx          echo.Context
		PFallthrough int
		RequestBody  UpdateResource3JSONBody
	}
	mock.lockUpdateResource3.RLock()
	calls = mock.calls.UpdateResource3
	mock.lockUpdateResource3.RUnlock()
	return calls
}
