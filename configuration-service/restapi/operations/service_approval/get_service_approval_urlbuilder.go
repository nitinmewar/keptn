// Code generated by go-swagger; DO NOT EDIT.

package service_approval

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// GetServiceApprovalURL generates an URL for the get service approval operation
type GetServiceApprovalURL struct {
	ApprovalID  string
	ProjectName string
	ServiceName string
	StageName   string

	NextPageKey *string
	PageSize    *int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetServiceApprovalURL) WithBasePath(bp string) *GetServiceApprovalURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetServiceApprovalURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetServiceApprovalURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/project/{projectName}/stage/{stageName}/service/{serviceName}/approval/{approvalID}"

	approvalID := o.ApprovalID
	if approvalID != "" {
		_path = strings.Replace(_path, "{approvalID}", approvalID, -1)
	} else {
		return nil, errors.New("approvalId is required on GetServiceApprovalURL")
	}

	projectName := o.ProjectName
	if projectName != "" {
		_path = strings.Replace(_path, "{projectName}", projectName, -1)
	} else {
		return nil, errors.New("projectName is required on GetServiceApprovalURL")
	}

	serviceName := o.ServiceName
	if serviceName != "" {
		_path = strings.Replace(_path, "{serviceName}", serviceName, -1)
	} else {
		return nil, errors.New("serviceName is required on GetServiceApprovalURL")
	}

	stageName := o.StageName
	if stageName != "" {
		_path = strings.Replace(_path, "{stageName}", stageName, -1)
	} else {
		return nil, errors.New("stageName is required on GetServiceApprovalURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var nextPageKeyQ string
	if o.NextPageKey != nil {
		nextPageKeyQ = *o.NextPageKey
	}
	if nextPageKeyQ != "" {
		qs.Set("nextPageKey", nextPageKeyQ)
	}

	var pageSizeQ string
	if o.PageSize != nil {
		pageSizeQ = swag.FormatInt64(*o.PageSize)
	}
	if pageSizeQ != "" {
		qs.Set("pageSize", pageSizeQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetServiceApprovalURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetServiceApprovalURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetServiceApprovalURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetServiceApprovalURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetServiceApprovalURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetServiceApprovalURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
