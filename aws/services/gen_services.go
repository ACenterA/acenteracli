// Auto generated implementation for the AWS cloud service

/*
Copyright 2017 WALLIX

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package awsservices

// DO NOT EDIT - This file was automatically generated with go generate

import (
	"context"
	"errors"
	"sync"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/service/sts/stsiface"
	awsfetch "github.com/wallix/awless/aws/fetch"
	"github.com/wallix/awless/cloud"
	"github.com/wallix/awless/fetch"
	"github.com/wallix/awless/graph"
	"github.com/wallix/awless/logger"
	tstore "github.com/wallix/triplestore"
)

const accessDenied = "Access Denied"

var ServiceNames = []string{
	"access",
}

var ResourceTypes = []string{
	"user",
	"group",
	"role",
	"policy",
	"accesskey",
	"instanceprofile",
	"mfadevice",
}

var ServicePerAPI = map[string]string{
	"iam": "access",
	"sts": "access",
}

var ServicePerResourceType = map[string]string{
	"user":            "access",
	"group":           "access",
	"role":            "access",
	"policy":          "access",
	"accesskey":       "access",
	"instanceprofile": "access",
	"mfadevice":       "access",
}

var APIPerResourceType = map[string]string{
	"user":            "iam",
	"group":           "iam",
	"role":            "iam",
	"policy":          "iam",
	"accesskey":       "iam",
	"instanceprofile": "iam",
	"mfadevice":       "iam",
}

type Access struct {
	fetcher         fetch.Fetcher
	region, profile string
	config          map[string]interface{}
	log             *logger.Logger
	iamiface.IAMAPI
	stsiface.STSAPI
}

func NewAccess(sess *session.Session, profile string, extraConf map[string]interface{}, log *logger.Logger) cloud.Service {
	region := "global"
	iamAPI := iam.New(sess)
	stsAPI := sts.New(sess)

	fetchConfig := awsfetch.NewConfig(
		iamAPI,
		stsAPI,
	)
	fetchConfig.Extra = extraConf
	fetchConfig.Log = log

	return &Access{
		IAMAPI:  iamAPI,
		STSAPI:  stsAPI,
		fetcher: fetch.NewFetcher(awsfetch.BuildAccessFetchFuncs(fetchConfig)),
		config:  extraConf,
		region:  region,
		profile: profile,
		log:     log,
	}
}

func (s *Access) Name() string {
	return "access"
}

func (s *Access) Region() string {
	return s.region
}

func (s *Access) Profile() string {
	return s.profile
}

func (s *Access) ResourceTypes() []string {
	return []string{
		"user",
		"group",
		"role",
		"policy",
		"accesskey",
		"instanceprofile",
		"mfadevice",
	}
}

func (s *Access) Fetch(ctx context.Context) (cloud.GraphAPI, error) {
	if s.IsSyncDisabled() {
		return graph.NewGraph(), nil
	}

	allErrors := new(fetch.Error)

	gph, err := s.fetcher.Fetch(context.WithValue(ctx, "region", s.region))
	defer s.fetcher.Reset()

	for _, e := range *fetch.WrapError(err) {
		switch ee := e.(type) {
		case awserr.RequestFailure:
			switch ee.Message() {
			case accessDenied:
				allErrors.Add(cloud.ErrFetchAccessDenied)
			default:
				allErrors.Add(ee)
			}
		case nil:
			continue
		default:
			allErrors.Add(ee)
		}
	}

	if err := gph.AddResource(graph.InitResource(cloud.Region, s.region)); err != nil {
		return gph, err
	}

	snap := gph.AsRDFGraphSnaphot()

	errc := make(chan error)
	var wg sync.WaitGroup
	if getBool(s.config, "aws.access.user.sync", true) {
		list, err := s.fetcher.Get("user_objects")
		if err != nil {
			return gph, err
		}
		if _, ok := list.([]*iam.UserDetail); !ok {
			return gph, errors.New("cannot cast to '[]*iam.UserDetail' type from fetch context")
		}
		for _, r := range list.([]*iam.UserDetail) {
			for _, fn := range addParentsFns["user"] {
				wg.Add(1)
				go func(f addParentFn, snap tstore.RDFGraph, region string, res *iam.UserDetail) {
					defer wg.Done()
					err := f(gph, snap, region, res)
					if err != nil {
						errc <- err
						return
					}
				}(fn, snap, s.region, r)
			}
		}
	}
	if getBool(s.config, "aws.access.group.sync", true) {
		list, err := s.fetcher.Get("group_objects")
		if err != nil {
			return gph, err
		}
		if _, ok := list.([]*iam.GroupDetail); !ok {
			return gph, errors.New("cannot cast to '[]*iam.GroupDetail' type from fetch context")
		}
		for _, r := range list.([]*iam.GroupDetail) {
			for _, fn := range addParentsFns["group"] {
				wg.Add(1)
				go func(f addParentFn, snap tstore.RDFGraph, region string, res *iam.GroupDetail) {
					defer wg.Done()
					err := f(gph, snap, region, res)
					if err != nil {
						errc <- err
						return
					}
				}(fn, snap, s.region, r)
			}
		}
	}
	if getBool(s.config, "aws.access.role.sync", true) {
		list, err := s.fetcher.Get("role_objects")
		if err != nil {
			return gph, err
		}
		if _, ok := list.([]*iam.RoleDetail); !ok {
			return gph, errors.New("cannot cast to '[]*iam.RoleDetail' type from fetch context")
		}
		for _, r := range list.([]*iam.RoleDetail) {
			for _, fn := range addParentsFns["role"] {
				wg.Add(1)
				go func(f addParentFn, snap tstore.RDFGraph, region string, res *iam.RoleDetail) {
					defer wg.Done()
					err := f(gph, snap, region, res)
					if err != nil {
						errc <- err
						return
					}
				}(fn, snap, s.region, r)
			}
		}
	}
	if getBool(s.config, "aws.access.policy.sync", true) {
		list, err := s.fetcher.Get("policy_objects")
		if err != nil {
			return gph, err
		}
		if _, ok := list.([]*iam.Policy); !ok {
			return gph, errors.New("cannot cast to '[]*iam.Policy' type from fetch context")
		}
		for _, r := range list.([]*iam.Policy) {
			for _, fn := range addParentsFns["policy"] {
				wg.Add(1)
				go func(f addParentFn, snap tstore.RDFGraph, region string, res *iam.Policy) {
					defer wg.Done()
					err := f(gph, snap, region, res)
					if err != nil {
						errc <- err
						return
					}
				}(fn, snap, s.region, r)
			}
		}
	}
	if getBool(s.config, "aws.access.accesskey.sync", true) {
		list, err := s.fetcher.Get("accesskey_objects")
		if err != nil {
			return gph, err
		}
		if _, ok := list.([]*iam.AccessKeyMetadata); !ok {
			return gph, errors.New("cannot cast to '[]*iam.AccessKeyMetadata' type from fetch context")
		}
		for _, r := range list.([]*iam.AccessKeyMetadata) {
			for _, fn := range addParentsFns["accesskey"] {
				wg.Add(1)
				go func(f addParentFn, snap tstore.RDFGraph, region string, res *iam.AccessKeyMetadata) {
					defer wg.Done()
					err := f(gph, snap, region, res)
					if err != nil {
						errc <- err
						return
					}
				}(fn, snap, s.region, r)
			}
		}
	}
	if getBool(s.config, "aws.access.instanceprofile.sync", true) {
		list, err := s.fetcher.Get("instanceprofile_objects")
		if err != nil {
			return gph, err
		}
		if _, ok := list.([]*iam.InstanceProfile); !ok {
			return gph, errors.New("cannot cast to '[]*iam.InstanceProfile' type from fetch context")
		}
		for _, r := range list.([]*iam.InstanceProfile) {
			for _, fn := range addParentsFns["instanceprofile"] {
				wg.Add(1)
				go func(f addParentFn, snap tstore.RDFGraph, region string, res *iam.InstanceProfile) {
					defer wg.Done()
					err := f(gph, snap, region, res)
					if err != nil {
						errc <- err
						return
					}
				}(fn, snap, s.region, r)
			}
		}
	}
	if getBool(s.config, "aws.access.mfadevice.sync", true) {
		list, err := s.fetcher.Get("mfadevice_objects")
		if err != nil {
			return gph, err
		}
		if _, ok := list.([]*iam.VirtualMFADevice); !ok {
			return gph, errors.New("cannot cast to '[]*iam.VirtualMFADevice' type from fetch context")
		}
		for _, r := range list.([]*iam.VirtualMFADevice) {
			for _, fn := range addParentsFns["mfadevice"] {
				wg.Add(1)
				go func(f addParentFn, snap tstore.RDFGraph, region string, res *iam.VirtualMFADevice) {
					defer wg.Done()
					err := f(gph, snap, region, res)
					if err != nil {
						errc <- err
						return
					}
				}(fn, snap, s.region, r)
			}
		}
	}

	go func() {
		wg.Wait()
		close(errc)
	}()

	for err := range errc {
		if err != nil {
			allErrors.Add(err)
		}
	}

	if allErrors.Any() {
		return gph, allErrors
	}

	return gph, nil
}

func (s *Access) FetchByType(ctx context.Context, t string) (cloud.GraphAPI, error) {
	defer s.fetcher.Reset()
	return s.fetcher.FetchByType(context.WithValue(ctx, "region", s.region), t)
}

func (s *Access) IsSyncDisabled() bool {
	return !getBool(s.config, "aws.access.sync", true)
}
