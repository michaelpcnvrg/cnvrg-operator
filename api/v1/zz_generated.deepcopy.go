// +build !ignore_autogenerated

/*
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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDbs) DeepCopyInto(out *AppDbs) {
	*out = *in
	in.Pg.DeepCopyInto(&out.Pg)
	in.Redis.DeepCopyInto(&out.Redis)
	in.Minio.DeepCopyInto(&out.Minio)
	in.Es.DeepCopyInto(&out.Es)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDbs.
func (in *AppDbs) DeepCopy() *AppDbs {
	if in == nil {
		return nil
	}
	out := new(AppDbs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BaseConfig) DeepCopyInto(out *BaseConfig) {
	*out = *in
	if in.FeatureFlags != nil {
		in, out := &in.FeatureFlags, &out.FeatureFlags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaseConfig.
func (in *BaseConfig) DeepCopy() *BaseConfig {
	if in == nil {
		return nil
	}
	out := new(BaseConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cnvrg) DeepCopyInto(out *Cnvrg) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cnvrg.
func (in *Cnvrg) DeepCopy() *Cnvrg {
	if in == nil {
		return nil
	}
	out := new(Cnvrg)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnvrgApp) DeepCopyInto(out *CnvrgApp) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnvrgApp.
func (in *CnvrgApp) DeepCopy() *CnvrgApp {
	if in == nil {
		return nil
	}
	out := new(CnvrgApp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CnvrgApp) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnvrgAppList) DeepCopyInto(out *CnvrgAppList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CnvrgApp, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnvrgAppList.
func (in *CnvrgAppList) DeepCopy() *CnvrgAppList {
	if in == nil {
		return nil
	}
	out := new(CnvrgAppList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CnvrgAppList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnvrgAppLogging) DeepCopyInto(out *CnvrgAppLogging) {
	*out = *in
	out.Elastalert = in.Elastalert
	in.Kibana.DeepCopyInto(&out.Kibana)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnvrgAppLogging.
func (in *CnvrgAppLogging) DeepCopy() *CnvrgAppLogging {
	if in == nil {
		return nil
	}
	out := new(CnvrgAppLogging)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnvrgAppMonitoring) DeepCopyInto(out *CnvrgAppMonitoring) {
	*out = *in
	out.Prometheus = in.Prometheus
	in.Grafana.DeepCopyInto(&out.Grafana)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnvrgAppMonitoring.
func (in *CnvrgAppMonitoring) DeepCopy() *CnvrgAppMonitoring {
	if in == nil {
		return nil
	}
	out := new(CnvrgAppMonitoring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnvrgAppNetworking) DeepCopyInto(out *CnvrgAppNetworking) {
	*out = *in
	out.Ingress = in.Ingress
	out.HTTPS = in.HTTPS
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnvrgAppNetworking.
func (in *CnvrgAppNetworking) DeepCopy() *CnvrgAppNetworking {
	if in == nil {
		return nil
	}
	out := new(CnvrgAppNetworking)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnvrgAppSpec) DeepCopyInto(out *CnvrgAppSpec) {
	*out = *in
	in.ControlPlane.DeepCopyInto(&out.ControlPlane)
	out.Registry = in.Registry
	in.Dbs.DeepCopyInto(&out.Dbs)
	out.Networking = in.Networking
	in.Logging.DeepCopyInto(&out.Logging)
	in.Monitoring.DeepCopyInto(&out.Monitoring)
	out.SSO = in.SSO
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnvrgAppSpec.
func (in *CnvrgAppSpec) DeepCopy() *CnvrgAppSpec {
	if in == nil {
		return nil
	}
	out := new(CnvrgAppSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnvrgInfra) DeepCopyInto(out *CnvrgInfra) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnvrgInfra.
func (in *CnvrgInfra) DeepCopy() *CnvrgInfra {
	if in == nil {
		return nil
	}
	out := new(CnvrgInfra)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CnvrgInfra) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnvrgInfraList) DeepCopyInto(out *CnvrgInfraList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CnvrgInfra, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnvrgInfraList.
func (in *CnvrgInfraList) DeepCopy() *CnvrgInfraList {
	if in == nil {
		return nil
	}
	out := new(CnvrgInfraList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CnvrgInfraList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnvrgInfraLogging) DeepCopyInto(out *CnvrgInfraLogging) {
	*out = *in
	out.Fluentbit = in.Fluentbit
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnvrgInfraLogging.
func (in *CnvrgInfraLogging) DeepCopy() *CnvrgInfraLogging {
	if in == nil {
		return nil
	}
	out := new(CnvrgInfraLogging)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnvrgInfraMonitoring) DeepCopyInto(out *CnvrgInfraMonitoring) {
	*out = *in
	out.PrometheusOperator = in.PrometheusOperator
	out.Prometheus = in.Prometheus
	out.NodeExporter = in.NodeExporter
	out.KubeStateMetrics = in.KubeStateMetrics
	in.Grafana.DeepCopyInto(&out.Grafana)
	out.DcgmExporter = in.DcgmExporter
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnvrgInfraMonitoring.
func (in *CnvrgInfraMonitoring) DeepCopy() *CnvrgInfraMonitoring {
	if in == nil {
		return nil
	}
	out := new(CnvrgInfraMonitoring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnvrgInfraNetworking) DeepCopyInto(out *CnvrgInfraNetworking) {
	*out = *in
	out.Ingress = in.Ingress
	out.HTTPS = in.HTTPS
	out.Istio = in.Istio
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnvrgInfraNetworking.
func (in *CnvrgInfraNetworking) DeepCopy() *CnvrgInfraNetworking {
	if in == nil {
		return nil
	}
	out := new(CnvrgInfraNetworking)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnvrgInfraSpec) DeepCopyInto(out *CnvrgInfraSpec) {
	*out = *in
	in.Monitoring.DeepCopyInto(&out.Monitoring)
	out.Networking = in.Networking
	out.Logging = in.Logging
	out.Registry = in.Registry
	out.Storage = in.Storage
	in.Dbs.DeepCopyInto(&out.Dbs)
	out.SSO = in.SSO
	out.Gpu = in.Gpu
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnvrgInfraSpec.
func (in *CnvrgInfraSpec) DeepCopy() *CnvrgInfraSpec {
	if in == nil {
		return nil
	}
	out := new(CnvrgInfraSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsistentHash) DeepCopyInto(out *ConsistentHash) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsistentHash.
func (in *ConsistentHash) DeepCopy() *ConsistentHash {
	if in == nil {
		return nil
	}
	out := new(ConsistentHash)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlane) DeepCopyInto(out *ControlPlane) {
	*out = *in
	in.WebApp.DeepCopyInto(&out.WebApp)
	out.Sidekiq = in.Sidekiq
	out.Searchkiq = in.Searchkiq
	out.Systemkiq = in.Systemkiq
	out.Hyper = in.Hyper
	out.Seeder = in.Seeder
	in.BaseConfig.DeepCopyInto(&out.BaseConfig)
	out.Ldap = in.Ldap
	out.Rbac = in.Rbac
	out.SMTP = in.SMTP
	out.Tenancy = in.Tenancy
	out.ObjectStorage = in.ObjectStorage
	in.Mpi.DeepCopyInto(&out.Mpi)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlane.
func (in *ControlPlane) DeepCopy() *ControlPlane {
	if in == nil {
		return nil
	}
	out := new(ControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DcgmExporter) DeepCopyInto(out *DcgmExporter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DcgmExporter.
func (in *DcgmExporter) DeepCopy() *DcgmExporter {
	if in == nil {
		return nil
	}
	out := new(DcgmExporter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultServiceMonitors) DeepCopyInto(out *DefaultServiceMonitors) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultServiceMonitors.
func (in *DefaultServiceMonitors) DeepCopy() *DefaultServiceMonitors {
	if in == nil {
		return nil
	}
	out := new(DefaultServiceMonitors)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Elastalert) DeepCopyInto(out *Elastalert) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Elastalert.
func (in *Elastalert) DeepCopy() *Elastalert {
	if in == nil {
		return nil
	}
	out := new(Elastalert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Es) DeepCopyInto(out *Es) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Es.
func (in *Es) DeepCopy() *Es {
	if in == nil {
		return nil
	}
	out := new(Es)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Fluentbit) DeepCopyInto(out *Fluentbit) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Fluentbit.
func (in *Fluentbit) DeepCopy() *Fluentbit {
	if in == nil {
		return nil
	}
	out := new(Fluentbit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Gpu) DeepCopyInto(out *Gpu) {
	*out = *in
	out.NvidiaDp = in.NvidiaDp
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Gpu.
func (in *Gpu) DeepCopy() *Gpu {
	if in == nil {
		return nil
	}
	out := new(Gpu)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Grafana) DeepCopyInto(out *Grafana) {
	*out = *in
	in.OauthProxy.DeepCopyInto(&out.OauthProxy)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Grafana.
func (in *Grafana) DeepCopy() *Grafana {
	if in == nil {
		return nil
	}
	out := new(Grafana)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPS) DeepCopyInto(out *HTTPS) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPS.
func (in *HTTPS) DeepCopy() *HTTPS {
	if in == nil {
		return nil
	}
	out := new(HTTPS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hostpath) DeepCopyInto(out *Hostpath) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hostpath.
func (in *Hostpath) DeepCopy() *Hostpath {
	if in == nil {
		return nil
	}
	out := new(Hostpath)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HugePages) DeepCopyInto(out *HugePages) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HugePages.
func (in *HugePages) DeepCopy() *HugePages {
	if in == nil {
		return nil
	}
	out := new(HugePages)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hyper) DeepCopyInto(out *Hyper) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hyper.
func (in *Hyper) DeepCopy() *Hyper {
	if in == nil {
		return nil
	}
	out := new(Hyper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdleMetricsExporter) DeepCopyInto(out *IdleMetricsExporter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdleMetricsExporter.
func (in *IdleMetricsExporter) DeepCopy() *IdleMetricsExporter {
	if in == nil {
		return nil
	}
	out := new(IdleMetricsExporter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Images) DeepCopyInto(out *Images) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Images.
func (in *Images) DeepCopy() *Images {
	if in == nil {
		return nil
	}
	out := new(Images)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraDbs) DeepCopyInto(out *InfraDbs) {
	*out = *in
	in.Redis.DeepCopyInto(&out.Redis)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraDbs.
func (in *InfraDbs) DeepCopy() *InfraDbs {
	if in == nil {
		return nil
	}
	out := new(InfraDbs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ingress) DeepCopyInto(out *Ingress) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ingress.
func (in *Ingress) DeepCopy() *Ingress {
	if in == nil {
		return nil
	}
	out := new(Ingress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Istio) DeepCopyInto(out *Istio) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Istio.
func (in *Istio) DeepCopy() *Istio {
	if in == nil {
		return nil
	}
	out := new(Istio)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kibana) DeepCopyInto(out *Kibana) {
	*out = *in
	in.OauthProxy.DeepCopyInto(&out.OauthProxy)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kibana.
func (in *Kibana) DeepCopy() *Kibana {
	if in == nil {
		return nil
	}
	out := new(Kibana)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeStateMetrics) DeepCopyInto(out *KubeStateMetrics) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeStateMetrics.
func (in *KubeStateMetrics) DeepCopy() *KubeStateMetrics {
	if in == nil {
		return nil
	}
	out := new(KubeStateMetrics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ldap) DeepCopyInto(out *Ldap) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ldap.
func (in *Ldap) DeepCopy() *Ldap {
	if in == nil {
		return nil
	}
	out := new(Ldap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Limits) DeepCopyInto(out *Limits) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Limits.
func (in *Limits) DeepCopy() *Limits {
	if in == nil {
		return nil
	}
	out := new(Limits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsServer) DeepCopyInto(out *MetricsServer) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsServer.
func (in *MetricsServer) DeepCopy() *MetricsServer {
	if in == nil {
		return nil
	}
	out := new(MetricsServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Minio) DeepCopyInto(out *Minio) {
	*out = *in
	out.SharedStorage = in.SharedStorage
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Minio.
func (in *Minio) DeepCopy() *Minio {
	if in == nil {
		return nil
	}
	out := new(Minio)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinioExporter) DeepCopyInto(out *MinioExporter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinioExporter.
func (in *MinioExporter) DeepCopy() *MinioExporter {
	if in == nil {
		return nil
	}
	out := new(MinioExporter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Mpi) DeepCopyInto(out *Mpi) {
	*out = *in
	if in.ExtraArgs != nil {
		in, out := &in.ExtraArgs, &out.ExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Registry = in.Registry
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Mpi.
func (in *Mpi) DeepCopy() *Mpi {
	if in == nil {
		return nil
	}
	out := new(Mpi)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Nfs) DeepCopyInto(out *Nfs) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Nfs.
func (in *Nfs) DeepCopy() *Nfs {
	if in == nil {
		return nil
	}
	out := new(Nfs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeExporter) DeepCopyInto(out *NodeExporter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeExporter.
func (in *NodeExporter) DeepCopy() *NodeExporter {
	if in == nil {
		return nil
	}
	out := new(NodeExporter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NvidiaDp) DeepCopyInto(out *NvidiaDp) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NvidiaDp.
func (in *NvidiaDp) DeepCopy() *NvidiaDp {
	if in == nil {
		return nil
	}
	out := new(NvidiaDp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OauthProxyServiceConf) DeepCopyInto(out *OauthProxyServiceConf) {
	*out = *in
	if in.SkipAuthRegex != nil {
		in, out := &in.SkipAuthRegex, &out.SkipAuthRegex
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OauthProxyServiceConf.
func (in *OauthProxyServiceConf) DeepCopy() *OauthProxyServiceConf {
	if in == nil {
		return nil
	}
	out := new(OauthProxyServiceConf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStorage) DeepCopyInto(out *ObjectStorage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStorage.
func (in *ObjectStorage) DeepCopy() *ObjectStorage {
	if in == nil {
		return nil
	}
	out := new(ObjectStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pg) DeepCopyInto(out *Pg) {
	*out = *in
	out.HugePages = in.HugePages
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pg.
func (in *Pg) DeepCopy() *Pg {
	if in == nil {
		return nil
	}
	out := new(Pg)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Prometheus) DeepCopyInto(out *Prometheus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Prometheus.
func (in *Prometheus) DeepCopy() *Prometheus {
	if in == nil {
		return nil
	}
	out := new(Prometheus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusOperator) DeepCopyInto(out *PrometheusOperator) {
	*out = *in
	out.Images = in.Images
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusOperator.
func (in *PrometheusOperator) DeepCopy() *PrometheusOperator {
	if in == nil {
		return nil
	}
	out := new(PrometheusOperator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rbac) DeepCopyInto(out *Rbac) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rbac.
func (in *Rbac) DeepCopy() *Rbac {
	if in == nil {
		return nil
	}
	out := new(Rbac)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Redis) DeepCopyInto(out *Redis) {
	*out = *in
	out.Limits = in.Limits
	out.Requests = in.Requests
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Redis.
func (in *Redis) DeepCopy() *Redis {
	if in == nil {
		return nil
	}
	out := new(Redis)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Registry) DeepCopyInto(out *Registry) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Registry.
func (in *Registry) DeepCopy() *Registry {
	if in == nil {
		return nil
	}
	out := new(Registry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Requests) DeepCopyInto(out *Requests) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Requests.
func (in *Requests) DeepCopy() *Requests {
	if in == nil {
		return nil
	}
	out := new(Requests)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SMTP) DeepCopyInto(out *SMTP) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SMTP.
func (in *SMTP) DeepCopy() *SMTP {
	if in == nil {
		return nil
	}
	out := new(SMTP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSO) DeepCopyInto(out *SSO) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSO.
func (in *SSO) DeepCopy() *SSO {
	if in == nil {
		return nil
	}
	out := new(SSO)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Searchkiq) DeepCopyInto(out *Searchkiq) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Searchkiq.
func (in *Searchkiq) DeepCopy() *Searchkiq {
	if in == nil {
		return nil
	}
	out := new(Searchkiq)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Seeder) DeepCopyInto(out *Seeder) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Seeder.
func (in *Seeder) DeepCopy() *Seeder {
	if in == nil {
		return nil
	}
	out := new(Seeder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedStorage) DeepCopyInto(out *SharedStorage) {
	*out = *in
	out.ConsistentHash = in.ConsistentHash
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedStorage.
func (in *SharedStorage) DeepCopy() *SharedStorage {
	if in == nil {
		return nil
	}
	out := new(SharedStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sidekiq) DeepCopyInto(out *Sidekiq) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sidekiq.
func (in *Sidekiq) DeepCopy() *Sidekiq {
	if in == nil {
		return nil
	}
	out := new(Sidekiq)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SidekiqExporter) DeepCopyInto(out *SidekiqExporter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SidekiqExporter.
func (in *SidekiqExporter) DeepCopy() *SidekiqExporter {
	if in == nil {
		return nil
	}
	out := new(SidekiqExporter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Status) DeepCopyInto(out *Status) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Status.
func (in *Status) DeepCopy() *Status {
	if in == nil {
		return nil
	}
	out := new(Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Storage) DeepCopyInto(out *Storage) {
	*out = *in
	out.Hostpath = in.Hostpath
	out.Nfs = in.Nfs
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Storage.
func (in *Storage) DeepCopy() *Storage {
	if in == nil {
		return nil
	}
	out := new(Storage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Systemkiq) DeepCopyInto(out *Systemkiq) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Systemkiq.
func (in *Systemkiq) DeepCopy() *Systemkiq {
	if in == nil {
		return nil
	}
	out := new(Systemkiq)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tenancy) DeepCopyInto(out *Tenancy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tenancy.
func (in *Tenancy) DeepCopy() *Tenancy {
	if in == nil {
		return nil
	}
	out := new(Tenancy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebApp) DeepCopyInto(out *WebApp) {
	*out = *in
	in.OauthProxy.DeepCopyInto(&out.OauthProxy)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebApp.
func (in *WebApp) DeepCopy() *WebApp {
	if in == nil {
		return nil
	}
	out := new(WebApp)
	in.DeepCopyInto(out)
	return out
}
