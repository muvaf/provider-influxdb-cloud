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

package controller

import (
	"github.com/crossplane-contrib/provider-influxdb/internal/controller/bucket"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane-contrib/provider-influxdb/internal/controller/organization"
	"github.com/crossplane-contrib/provider-influxdb/internal/controller/providerconfig"
)

// Setup creates all Template controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, wl workqueue.RateLimiter) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter) error{
		providerconfig.Setup,
		organization.Setup,
		bucket.Setup,
	} {
		if err := setup(mgr, l, wl); err != nil {
			return err
		}
	}
	return nil
}
