// Copyright 2019-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package partitiongroupmembership

import (
	"context"
	"github.com/atomix/kubernetes-controller/pkg/apis/cloud/v1beta3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("partition_group_membership_controller")

// Add creates a new Database controller and adds it to the Manager. The Manager will set fields on the
// controller and Start it when the Manager is Started.
func Add(mgr manager.Manager, eventCh chan<- types.NamespacedName) error {
	r := &Reconciler{
		client:  mgr.GetClient(),
		scheme:  mgr.GetScheme(),
		config:  mgr.GetConfig(),
		eventCh: eventCh,
	}

	// Create a new controller
	c, err := controller.New("partition-group-membership-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource PartitionGroupMembership
	err = c.Watch(&source.Kind{Type: &v1beta3.PartitionGroupMembership{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource Member and requeue the PartitionGroupMembership
	err = c.Watch(&source.Kind{Type: &v1beta3.Member{}}, &handler.EnqueueRequestsFromMapFunc{
		ToRequests: &memberMapper{
			client: mgr.GetClient(),
		},
	})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource PartitionGroup and requeue the PartitionGroupMembership
	err = c.Watch(&source.Kind{Type: &v1beta3.PartitionGroup{}}, &handler.EnqueueRequestsFromMapFunc{
		ToRequests: &partitionGroupMapper{
			client: mgr.GetClient(),
		},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &Reconciler{}

// Reconciler reconciles a PartitionGroup object
type Reconciler struct {
	client  client.Client
	scheme  *runtime.Scheme
	config  *rest.Config
	eventCh chan<- types.NamespacedName
}

func (r *Reconciler) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	logger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	logger.Info("Reconciling PartitionGroupMembership")

	// Fetch the PartitionGroupMembership instance
	partitionGroupMembership := &v1beta3.PartitionGroupMembership{}
	err := r.client.Get(context.TODO(), request.NamespacedName, partitionGroupMembership)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	if partitionGroupMembership.DeletionTimestamp == nil {
		addFinalizer := true
		for _, finalizer := range partitionGroupMembership.Finalizers {
			if finalizer == "partition-group-membership-controller" {
				addFinalizer = false
				break
			}
		}
		if addFinalizer {
			partitionGroupMembership.Finalizers = append(partitionGroupMembership.Finalizers, "partition-group-membership-controller")
			err = r.client.Update(context.TODO(), partitionGroupMembership)
			return reconcile.Result{}, err
		}
	}

	defer func() {
		go func() {
			r.eventCh <- types.NamespacedName{
				Namespace: partitionGroupMembership.Namespace,
				Name:      partitionGroupMembership.Bind.Group,
			}
		}()
	}()

	if partitionGroupMembership.DeletionTimestamp == nil {
		member := &v1beta3.Member{}
		memberName := types.NamespacedName{
			Namespace: partitionGroupMembership.Namespace,
			Name:      partitionGroupMembership.Bind.Member,
		}
		err = r.client.Get(context.TODO(), memberName, member)
		if err != nil {
			if !errors.IsNotFound(err) {
				return reconcile.Result{}, err
			}
			err = r.client.Delete(context.TODO(), partitionGroupMembership)
			return reconcile.Result{}, err
		}

		partitionGroup := &v1beta3.PartitionGroup{}
		partitionGroupName := types.NamespacedName{
			Namespace: partitionGroupMembership.Namespace,
			Name:      partitionGroupMembership.Bind.Group,
		}
		err = r.client.Get(context.TODO(), partitionGroupName, partitionGroup)
		if err != nil {
			if !errors.IsNotFound(err) {
				return reconcile.Result{}, err
			}
			err = r.client.Delete(context.TODO(), partitionGroupMembership)
			return reconcile.Result{}, err
		}
	} else {
		finalize := false
		for _, finalizer := range partitionGroupMembership.Finalizers {
			if finalizer == "partition-group-membership-controller" {
				finalize = true
				break
			}
		}
		if finalize {
			finalizers := make([]string, 0, len(partitionGroupMembership.Finalizers)-1)
			for _, finalizer := range partitionGroupMembership.Finalizers {
				if finalizer != "partition-group-membership-controller" {
					finalizers = append(finalizers, finalizer)
				}
			}
			partitionGroupMembership.Finalizers = finalizers
			err = r.client.Update(context.TODO(), partitionGroupMembership)
			return reconcile.Result{}, err
		}
	}
	return reconcile.Result{}, nil
}

type memberMapper struct {
	client client.Client
}

func (m *memberMapper) Map(object handler.MapObject) []reconcile.Request {
	member := object.Object.(*v1beta3.Member)
	memberships := &v1beta3.PartitionGroupMembershipList{}
	err := m.client.List(context.TODO(), memberships, &client.ListOptions{Namespace: member.Namespace})
	if err != nil {
		return []reconcile.Request{}
	}

	requests := make([]reconcile.Request, 0)
	for _, membership := range memberships.Items {
		if membership.Bind.Member == member.Name {
			requests = append(requests, reconcile.Request{
				NamespacedName: types.NamespacedName{
					Namespace: membership.Namespace,
					Name:      membership.Name,
				},
			})
		}
	}
	return requests
}

type partitionGroupMapper struct {
	client client.Client
}

func (m *partitionGroupMapper) Map(object handler.MapObject) []reconcile.Request {
	partitionGroup := object.Object.(*v1beta3.PartitionGroup)
	memberships := &v1beta3.PartitionGroupMembershipList{}
	err := m.client.List(context.TODO(), memberships, &client.ListOptions{Namespace: partitionGroup.Namespace})
	if err != nil {
		return []reconcile.Request{}
	}

	requests := make([]reconcile.Request, 0)
	for _, membership := range memberships.Items {
		if membership.Bind.Group == partitionGroup.Name {
			requests = append(requests, reconcile.Request{
				NamespacedName: types.NamespacedName{
					Namespace: membership.Namespace,
					Name:      membership.Name,
				},
			})
		}
	}
	return requests
}
