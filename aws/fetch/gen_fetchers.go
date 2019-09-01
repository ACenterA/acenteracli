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

package awsfetch

// DO NOT EDIT - This file was automatically generated with go generate

import (
	"context"

	"github.com/aws/aws-sdk-go/service/iam"
	awsconv "github.com/wallix/awless/aws/conv"
	"github.com/wallix/awless/fetch"
	"github.com/wallix/awless/graph"
)

func BuildAccessFetchFuncs(conf *Config) fetch.Funcs {
	funcs := make(map[string]fetch.Func)

	addManualAccessFetchFuncs(conf, funcs)

	funcs["instanceprofile"] = func(ctx context.Context, cache fetch.Cache) ([]*graph.Resource, interface{}, error) {
		var resources []*graph.Resource
		var objects []*iam.InstanceProfile

		if !conf.getBoolDefaultTrue("aws.access.instanceprofile.sync") && !getBoolFromContext(ctx, "force") {
			conf.Log.Verbose("sync: *disabled* for resource access[instanceprofile]")
			return resources, objects, nil
		}
		var badResErr error
		err := conf.APIs.Iam.ListInstanceProfilesPages(&iam.ListInstanceProfilesInput{},
			func(out *iam.ListInstanceProfilesOutput, lastPage bool) (shouldContinue bool) {
				for _, output := range out.InstanceProfiles {
					if badResErr != nil {
						return false
					}
					objects = append(objects, output)
					var res *graph.Resource
					if res, badResErr = awsconv.NewResource(output); badResErr != nil {
						return false
					}
					resources = append(resources, res)
				}
				return out.Marker != nil
			})
		if err != nil {
			return resources, objects, err
		}

		return resources, objects, badResErr
	}

	funcs["mfadevice"] = func(ctx context.Context, cache fetch.Cache) ([]*graph.Resource, interface{}, error) {
		var resources []*graph.Resource
		var objects []*iam.VirtualMFADevice

		if !conf.getBoolDefaultTrue("aws.access.mfadevice.sync") && !getBoolFromContext(ctx, "force") {
			conf.Log.Verbose("sync: *disabled* for resource access[mfadevice]")
			return resources, objects, nil
		}
		var badResErr error
		err := conf.APIs.Iam.ListVirtualMFADevicesPages(&iam.ListVirtualMFADevicesInput{},
			func(out *iam.ListVirtualMFADevicesOutput, lastPage bool) (shouldContinue bool) {
				for _, output := range out.VirtualMFADevices {
					if badResErr != nil {
						return false
					}
					objects = append(objects, output)
					var res *graph.Resource
					if res, badResErr = awsconv.NewResource(output); badResErr != nil {
						return false
					}
					resources = append(resources, res)
				}
				return out.Marker != nil
			})
		if err != nil {
			return resources, objects, err
		}

		return resources, objects, badResErr
	}
	return funcs
}
