// Copyright 2020-2021 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-twirp v8.1.0, DO NOT EDIT.
// source: buf/alpha/registry/v1alpha1/repository_tag.proto

package registryv1alpha1

import context "context"
import fmt "fmt"
import http "net/http"
import ioutil "io/ioutil"
import json "encoding/json"
import strconv "strconv"
import strings "strings"

import protojson "google.golang.org/protobuf/encoding/protojson"
import proto "google.golang.org/protobuf/proto"
import twirp "github.com/twitchtv/twirp"
import ctxsetters "github.com/twitchtv/twirp/ctxsetters"

// Version compatibility assertion.
// If the constant is not defined in the package, that likely means
// the package needs to be updated to work with this generated code.
// See https://twitchtv.github.io/twirp/docs/version_matrix.html
const _ = twirp.TwirpPackageMinVersion_8_1_0

// ==============================
// RepositoryTagService Interface
// ==============================

// RepositoryTagService is the Repository tag service.
type RepositoryTagService interface {
	// CreateRepositoryTag creates a new repository tag.
	CreateRepositoryTag(context.Context, *CreateRepositoryTagRequest) (*CreateRepositoryTagResponse, error)

	// ListRepositoryTags lists the repository tags associated with a Repository.
	ListRepositoryTags(context.Context, *ListRepositoryTagsRequest) (*ListRepositoryTagsResponse, error)
}

// ====================================
// RepositoryTagService Protobuf Client
// ====================================

type repositoryTagServiceProtobufClient struct {
	client      HTTPClient
	urls        [2]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewRepositoryTagServiceProtobufClient creates a Protobuf client that implements the RepositoryTagService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewRepositoryTagServiceProtobufClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) RepositoryTagService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Using ReadOpt allows backwards and forwads compatibility with new options in the future
	literalURLs := false
	_ = clientOpts.ReadOpt("literalURLs", &literalURLs)
	var pathPrefix string
	if ok := clientOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(pathPrefix, "buf.alpha.registry.v1alpha1", "RepositoryTagService")
	urls := [2]string{
		serviceURL + "CreateRepositoryTag",
		serviceURL + "ListRepositoryTags",
	}

	return &repositoryTagServiceProtobufClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *repositoryTagServiceProtobufClient) CreateRepositoryTag(ctx context.Context, in *CreateRepositoryTagRequest) (*CreateRepositoryTagResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "buf.alpha.registry.v1alpha1")
	ctx = ctxsetters.WithServiceName(ctx, "RepositoryTagService")
	ctx = ctxsetters.WithMethodName(ctx, "CreateRepositoryTag")
	caller := c.callCreateRepositoryTag
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *CreateRepositoryTagRequest) (*CreateRepositoryTagResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateRepositoryTagRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateRepositoryTagRequest) when calling interceptor")
					}
					return c.callCreateRepositoryTag(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateRepositoryTagResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateRepositoryTagResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *repositoryTagServiceProtobufClient) callCreateRepositoryTag(ctx context.Context, in *CreateRepositoryTagRequest) (*CreateRepositoryTagResponse, error) {
	out := new(CreateRepositoryTagResponse)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *repositoryTagServiceProtobufClient) ListRepositoryTags(ctx context.Context, in *ListRepositoryTagsRequest) (*ListRepositoryTagsResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "buf.alpha.registry.v1alpha1")
	ctx = ctxsetters.WithServiceName(ctx, "RepositoryTagService")
	ctx = ctxsetters.WithMethodName(ctx, "ListRepositoryTags")
	caller := c.callListRepositoryTags
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *ListRepositoryTagsRequest) (*ListRepositoryTagsResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ListRepositoryTagsRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ListRepositoryTagsRequest) when calling interceptor")
					}
					return c.callListRepositoryTags(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*ListRepositoryTagsResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*ListRepositoryTagsResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *repositoryTagServiceProtobufClient) callListRepositoryTags(ctx context.Context, in *ListRepositoryTagsRequest) (*ListRepositoryTagsResponse, error) {
	out := new(ListRepositoryTagsResponse)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ================================
// RepositoryTagService JSON Client
// ================================

type repositoryTagServiceJSONClient struct {
	client      HTTPClient
	urls        [2]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewRepositoryTagServiceJSONClient creates a JSON client that implements the RepositoryTagService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewRepositoryTagServiceJSONClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) RepositoryTagService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Using ReadOpt allows backwards and forwads compatibility with new options in the future
	literalURLs := false
	_ = clientOpts.ReadOpt("literalURLs", &literalURLs)
	var pathPrefix string
	if ok := clientOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(pathPrefix, "buf.alpha.registry.v1alpha1", "RepositoryTagService")
	urls := [2]string{
		serviceURL + "CreateRepositoryTag",
		serviceURL + "ListRepositoryTags",
	}

	return &repositoryTagServiceJSONClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *repositoryTagServiceJSONClient) CreateRepositoryTag(ctx context.Context, in *CreateRepositoryTagRequest) (*CreateRepositoryTagResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "buf.alpha.registry.v1alpha1")
	ctx = ctxsetters.WithServiceName(ctx, "RepositoryTagService")
	ctx = ctxsetters.WithMethodName(ctx, "CreateRepositoryTag")
	caller := c.callCreateRepositoryTag
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *CreateRepositoryTagRequest) (*CreateRepositoryTagResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateRepositoryTagRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateRepositoryTagRequest) when calling interceptor")
					}
					return c.callCreateRepositoryTag(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateRepositoryTagResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateRepositoryTagResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *repositoryTagServiceJSONClient) callCreateRepositoryTag(ctx context.Context, in *CreateRepositoryTagRequest) (*CreateRepositoryTagResponse, error) {
	out := new(CreateRepositoryTagResponse)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *repositoryTagServiceJSONClient) ListRepositoryTags(ctx context.Context, in *ListRepositoryTagsRequest) (*ListRepositoryTagsResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "buf.alpha.registry.v1alpha1")
	ctx = ctxsetters.WithServiceName(ctx, "RepositoryTagService")
	ctx = ctxsetters.WithMethodName(ctx, "ListRepositoryTags")
	caller := c.callListRepositoryTags
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *ListRepositoryTagsRequest) (*ListRepositoryTagsResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ListRepositoryTagsRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ListRepositoryTagsRequest) when calling interceptor")
					}
					return c.callListRepositoryTags(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*ListRepositoryTagsResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*ListRepositoryTagsResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *repositoryTagServiceJSONClient) callListRepositoryTags(ctx context.Context, in *ListRepositoryTagsRequest) (*ListRepositoryTagsResponse, error) {
	out := new(ListRepositoryTagsResponse)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ===================================
// RepositoryTagService Server Handler
// ===================================

type repositoryTagServiceServer struct {
	RepositoryTagService
	interceptor      twirp.Interceptor
	hooks            *twirp.ServerHooks
	pathPrefix       string // prefix for routing
	jsonSkipDefaults bool   // do not include unpopulated fields (default values) in the response
	jsonCamelCase    bool   // JSON fields are serialized as lowerCamelCase rather than keeping the original proto names
}

// NewRepositoryTagServiceServer builds a TwirpServer that can be used as an http.Handler to handle
// HTTP requests that are routed to the right method in the provided svc implementation.
// The opts are twirp.ServerOption modifiers, for example twirp.WithServerHooks(hooks).
func NewRepositoryTagServiceServer(svc RepositoryTagService, opts ...interface{}) TwirpServer {
	serverOpts := newServerOpts(opts)

	// Using ReadOpt allows backwards and forwads compatibility with new options in the future
	jsonSkipDefaults := false
	_ = serverOpts.ReadOpt("jsonSkipDefaults", &jsonSkipDefaults)
	jsonCamelCase := false
	_ = serverOpts.ReadOpt("jsonCamelCase", &jsonCamelCase)
	var pathPrefix string
	if ok := serverOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	return &repositoryTagServiceServer{
		RepositoryTagService: svc,
		hooks:                serverOpts.Hooks,
		interceptor:          twirp.ChainInterceptors(serverOpts.Interceptors...),
		pathPrefix:           pathPrefix,
		jsonSkipDefaults:     jsonSkipDefaults,
		jsonCamelCase:        jsonCamelCase,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *repositoryTagServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// handleRequestBodyError is used to handle error when the twirp server cannot read request
func (s *repositoryTagServiceServer) handleRequestBodyError(ctx context.Context, resp http.ResponseWriter, msg string, err error) {
	if context.Canceled == ctx.Err() {
		s.writeError(ctx, resp, twirp.NewError(twirp.Canceled, "failed to read request: context canceled"))
		return
	}
	if context.DeadlineExceeded == ctx.Err() {
		s.writeError(ctx, resp, twirp.NewError(twirp.DeadlineExceeded, "failed to read request: deadline exceeded"))
		return
	}
	s.writeError(ctx, resp, twirp.WrapError(malformedRequestError(msg), err))
}

// RepositoryTagServicePathPrefix is a convenience constant that may identify URL paths.
// Should be used with caution, it only matches routes generated by Twirp Go clients,
// with the default "/twirp" prefix and default CamelCase service and method names.
// More info: https://twitchtv.github.io/twirp/docs/routing.html
const RepositoryTagServicePathPrefix = "/twirp/buf.alpha.registry.v1alpha1.RepositoryTagService/"

func (s *repositoryTagServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "buf.alpha.registry.v1alpha1")
	ctx = ctxsetters.WithServiceName(ctx, "RepositoryTagService")
	ctx = ctxsetters.WithResponseWriter(ctx, resp)

	var err error
	ctx, err = callRequestReceived(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	if req.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", req.Method)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	// Verify path format: [<prefix>]/<package>.<Service>/<Method>
	prefix, pkgService, method := parseTwirpPath(req.URL.Path)
	if pkgService != "buf.alpha.registry.v1alpha1.RepositoryTagService" {
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
	if prefix != s.pathPrefix {
		msg := fmt.Sprintf("invalid path prefix %q, expected %q, on path %q", prefix, s.pathPrefix, req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	switch method {
	case "CreateRepositoryTag":
		s.serveCreateRepositoryTag(ctx, resp, req)
		return
	case "ListRepositoryTags":
		s.serveListRepositoryTags(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
}

func (s *repositoryTagServiceServer) serveCreateRepositoryTag(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveCreateRepositoryTagJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveCreateRepositoryTagProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *repositoryTagServiceServer) serveCreateRepositoryTagJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CreateRepositoryTag")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	d := json.NewDecoder(req.Body)
	rawReqBody := json.RawMessage{}
	if err := d.Decode(&rawReqBody); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}
	reqContent := new(CreateRepositoryTagRequest)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.RepositoryTagService.CreateRepositoryTag
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *CreateRepositoryTagRequest) (*CreateRepositoryTagResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateRepositoryTagRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateRepositoryTagRequest) when calling interceptor")
					}
					return s.RepositoryTagService.CreateRepositoryTag(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateRepositoryTagResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateRepositoryTagResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *CreateRepositoryTagResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *CreateRepositoryTagResponse and nil error while calling CreateRepositoryTag. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	marshaler := &protojson.MarshalOptions{UseProtoNames: !s.jsonCamelCase, EmitUnpopulated: !s.jsonSkipDefaults}
	respBytes, err := marshaler.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *repositoryTagServiceServer) serveCreateRepositoryTagProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CreateRepositoryTag")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.handleRequestBodyError(ctx, resp, "failed to read request body", err)
		return
	}
	reqContent := new(CreateRepositoryTagRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.RepositoryTagService.CreateRepositoryTag
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *CreateRepositoryTagRequest) (*CreateRepositoryTagResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateRepositoryTagRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateRepositoryTagRequest) when calling interceptor")
					}
					return s.RepositoryTagService.CreateRepositoryTag(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateRepositoryTagResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateRepositoryTagResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *CreateRepositoryTagResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *CreateRepositoryTagResponse and nil error while calling CreateRepositoryTag. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *repositoryTagServiceServer) serveListRepositoryTags(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveListRepositoryTagsJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveListRepositoryTagsProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *repositoryTagServiceServer) serveListRepositoryTagsJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "ListRepositoryTags")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	d := json.NewDecoder(req.Body)
	rawReqBody := json.RawMessage{}
	if err := d.Decode(&rawReqBody); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}
	reqContent := new(ListRepositoryTagsRequest)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.RepositoryTagService.ListRepositoryTags
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *ListRepositoryTagsRequest) (*ListRepositoryTagsResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ListRepositoryTagsRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ListRepositoryTagsRequest) when calling interceptor")
					}
					return s.RepositoryTagService.ListRepositoryTags(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*ListRepositoryTagsResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*ListRepositoryTagsResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *ListRepositoryTagsResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *ListRepositoryTagsResponse and nil error while calling ListRepositoryTags. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	marshaler := &protojson.MarshalOptions{UseProtoNames: !s.jsonCamelCase, EmitUnpopulated: !s.jsonSkipDefaults}
	respBytes, err := marshaler.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *repositoryTagServiceServer) serveListRepositoryTagsProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "ListRepositoryTags")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.handleRequestBodyError(ctx, resp, "failed to read request body", err)
		return
	}
	reqContent := new(ListRepositoryTagsRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.RepositoryTagService.ListRepositoryTags
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *ListRepositoryTagsRequest) (*ListRepositoryTagsResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ListRepositoryTagsRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ListRepositoryTagsRequest) when calling interceptor")
					}
					return s.RepositoryTagService.ListRepositoryTags(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*ListRepositoryTagsResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*ListRepositoryTagsResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *ListRepositoryTagsResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *ListRepositoryTagsResponse and nil error while calling ListRepositoryTags. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *repositoryTagServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor13, 0
}

func (s *repositoryTagServiceServer) ProtocGenTwirpVersion() string {
	return "v8.1.0"
}

// PathPrefix returns the base service path, in the form: "/<prefix>/<package>.<Service>/"
// that is everything in a Twirp route except for the <Method>. This can be used for routing,
// for example to identify the requests that are targeted to this service in a mux.
func (s *repositoryTagServiceServer) PathPrefix() string {
	return baseServicePath(s.pathPrefix, "buf.alpha.registry.v1alpha1", "RepositoryTagService")
}

var twirpFileDescriptor13 = []byte{
	// 580 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0x25, 0xd9, 0x75, 0xdd, 0xbd, 0xb5, 0x5d, 0x18, 0x45, 0x62, 0x8a, 0xb6, 0x54, 0xd0, 0xc5,
	0x87, 0xc4, 0xad, 0xe0, 0x4a, 0xf6, 0x69, 0xe3, 0x83, 0x08, 0x2a, 0x35, 0x2d, 0x82, 0x52, 0x28,
	0x93, 0x76, 0x9a, 0x0e, 0x36, 0x9d, 0x38, 0x33, 0x29, 0xee, 0xbe, 0x0b, 0x7e, 0x81, 0xe0, 0x93,
	0xe0, 0xdb, 0xfa, 0x29, 0x7e, 0x86, 0x8f, 0x7e, 0x85, 0x64, 0xd2, 0xa9, 0x8d, 0xb6, 0x81, 0xbe,
	0x25, 0xe7, 0x9c, 0x7b, 0x72, 0xee, 0xbd, 0x93, 0x81, 0x87, 0x61, 0x3a, 0x76, 0xf1, 0x34, 0x99,
	0x60, 0x97, 0x93, 0x88, 0x0a, 0xc9, 0xcf, 0xdd, 0xf9, 0xb1, 0x02, 0x8e, 0x5d, 0x4e, 0x12, 0x26,
	0xa8, 0x64, 0xfc, 0x7c, 0x20, 0x71, 0xe4, 0x24, 0x9c, 0x49, 0x86, 0xea, 0x61, 0x3a, 0x76, 0x94,
	0xc0, 0xd1, 0x15, 0x8e, 0xae, 0xb0, 0x1b, 0x11, 0x63, 0xd1, 0x94, 0xb8, 0x4a, 0x9a, 0x59, 0x4b,
	0x1a, 0x13, 0x21, 0x71, 0x9c, 0xe4, 0xd5, 0xad, 0x4b, 0x03, 0xaa, 0xc1, 0xd2, 0xb6, 0x87, 0x23,
	0x54, 0x03, 0x93, 0x8e, 0x2c, 0xa3, 0x69, 0x1c, 0x1d, 0x04, 0x26, 0x1d, 0xa1, 0x53, 0xa8, 0x0c,
	0x39, 0xc1, 0x92, 0x0c, 0xb2, 0x5a, 0xcb, 0x6c, 0x1a, 0x47, 0x95, 0xb6, 0xed, 0xe4, 0xc6, 0x8e,
	0x36, 0x76, 0x7a, 0xda, 0x38, 0x80, 0x5c, 0x9e, 0x01, 0x08, 0xc1, 0xee, 0x0c, 0xc7, 0xc4, 0xda,
	0x55, 0x76, 0xea, 0x19, 0x35, 0xa0, 0x32, 0x64, 0x71, 0x4c, 0xe5, 0x40, 0x51, 0x57, 0x14, 0x05,
	0x39, 0xf4, 0x2a, 0x13, 0xdc, 0x84, 0x3d, 0x9c, 0xca, 0x09, 0xe3, 0xd6, 0x9e, 0xe2, 0x16, 0x6f,
	0xad, 0x39, 0xd8, 0x4f, 0x95, 0x75, 0x21, 0x70, 0x40, 0x3e, 0xa4, 0x44, 0x48, 0x74, 0x17, 0xaa,
	0x2b, 0xf3, 0x59, 0xb6, 0x70, 0xed, 0x2f, 0xf8, 0x7c, 0xb4, 0xcc, 0x63, 0x6e, 0xce, 0xb3, 0xf3,
	0x6f, 0x9e, 0x56, 0x02, 0xf5, 0xb5, 0xdf, 0x15, 0x09, 0x9b, 0x09, 0x82, 0x5e, 0x43, 0xad, 0xb8,
	0x18, 0xf5, 0xe5, 0x4a, 0xfb, 0x81, 0x53, 0xb2, 0x19, 0xa7, 0xe8, 0xb5, 0x12, 0xbd, 0x87, 0xa3,
	0xd6, 0x17, 0x03, 0x6e, 0xbd, 0xa0, 0x42, 0x16, 0x44, 0x62, 0xab, 0x4e, 0xeb, 0x70, 0x90, 0xe0,
	0x88, 0x0c, 0x04, 0xbd, 0xc8, 0xdb, 0xad, 0x06, 0xfb, 0x19, 0xd0, 0xa5, 0x17, 0x04, 0xdd, 0x06,
	0x50, 0xa4, 0x64, 0xef, 0xc9, 0x6c, 0xd1, 0xb1, 0x92, 0xf7, 0x32, 0x00, 0x59, 0x70, 0x95, 0x93,
	0x39, 0xe1, 0x22, 0x5f, 0xdc, 0x7e, 0xa0, 0x5f, 0x5b, 0x5f, 0x0d, 0xb0, 0xd7, 0x05, 0x5b, 0x8c,
	0xa2, 0x0b, 0x87, 0xc5, 0x51, 0x08, 0xcb, 0x68, 0xee, 0x6c, 0x39, 0x8b, 0x5a, 0x61, 0x16, 0x02,
	0xdd, 0x83, 0xc3, 0x19, 0xf9, 0x28, 0x07, 0x2b, 0x89, 0xf3, 0xf5, 0x55, 0x33, 0xb8, 0xa3, 0x53,
	0xb7, 0x2f, 0x4d, 0xb8, 0x51, 0x70, 0xea, 0x12, 0x3e, 0xa7, 0x43, 0x82, 0x3e, 0x1b, 0x70, 0x7d,
	0xcd, 0x02, 0xd1, 0x49, 0x69, 0xa8, 0xcd, 0x47, 0xcd, 0x7e, 0xb2, 0x7d, 0xe1, 0x62, 0x40, 0x9f,
	0x0c, 0x40, 0xff, 0xcf, 0x0f, 0x3d, 0x2e, 0x35, 0xdc, 0x78, 0x12, 0xec, 0x93, 0xad, 0xeb, 0xf2,
	0x1c, 0xfe, 0x37, 0x13, 0x1a, 0x43, 0x16, 0x97, 0x95, 0xfb, 0xa8, 0x50, 0xdb, 0xc9, 0x7e, 0xf4,
	0x8e, 0xf1, 0xee, 0x6d, 0x44, 0xe5, 0x24, 0x0d, 0x9d, 0x21, 0x8b, 0xdd, 0x30, 0x1d, 0x87, 0x29,
	0x9d, 0x8e, 0xb2, 0x07, 0x37, 0xe1, 0x74, 0x8e, 0x25, 0x71, 0x23, 0x32, 0xcb, 0xaf, 0x1b, 0x37,
	0x62, 0x6e, 0xc9, 0x6d, 0x76, 0xaa, 0x11, 0x0d, 0x7c, 0x37, 0x77, 0xfc, 0xb3, 0xe0, 0x87, 0x59,
	0xf7, 0xd3, 0xb1, 0x73, 0xa6, 0x42, 0x05, 0x3a, 0xd4, 0x9b, 0x85, 0xe6, 0xa7, 0x62, 0xfb, 0x8a,
	0xed, 0x6b, 0xb6, 0xaf, 0xd9, 0x5f, 0xe6, 0xfd, 0x12, 0xb6, 0xff, 0xac, 0xe3, 0xbf, 0x24, 0x12,
	0x8f, 0xb0, 0xc4, 0xbf, 0xcd, 0x3b, 0x7e, 0x3a, 0xf6, 0x3c, 0x25, 0xf5, 0x3c, 0xad, 0xf5, 0x3c,
	0x2d, 0x0e, 0xf7, 0x54, 0x0f, 0x8f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x55, 0xd8, 0x68,
	0x91, 0x05, 0x00, 0x00,
}
