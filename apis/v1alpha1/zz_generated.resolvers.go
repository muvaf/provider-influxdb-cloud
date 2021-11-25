/*
Copyright 2021 The Crossplane Authors.

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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Bucket.
func (mg *Bucket) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OrgID),
		Extract:      OrganizationID(),
		Reference:    mg.Spec.ForProvider.OrgIDRef,
		Selector:     mg.Spec.ForProvider.OrgIDSelector,
		To: reference.To{
			List:    &OrganizationList{},
			Managed: &Organization{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OrgID")
	}
	mg.Spec.ForProvider.OrgID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OrgIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DatabaseRetentionPolicyMapping.
func (mg *DatabaseRetentionPolicyMapping) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.ForProvider.BucketID,
		Extract:      OrganizationID(),
		Reference:    mg.Spec.ForProvider.BucketIDRef,
		Selector:     mg.Spec.ForProvider.BucketIDSelector,
		To: reference.To{
			List:    &OrganizationList{},
			Managed: &Organization{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BucketID")
	}
	mg.Spec.ForProvider.BucketID = rsp.ResolvedValue
	mg.Spec.ForProvider.BucketIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.ForProvider.Org,
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.OrgRef,
		Selector:     mg.Spec.ForProvider.OrgSelector,
		To: reference.To{
			List:    &OrganizationList{},
			Managed: &Organization{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Org")
	}
	mg.Spec.ForProvider.Org = rsp.ResolvedValue
	mg.Spec.ForProvider.OrgRef = rsp.ResolvedReference

	return nil
}
