package controllers

import (
	"context"
	mlopsv1 "github.com/cnvrg-operator/api/v1"
	"github.com/cnvrg-operator/pkg/desired"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"time"
)

// +kubebuilder:docs-gen:collapse=Imports

var enabledInfraTests = map[string]bool{
	"monitoring": true,
	"storage":    true,
	"networking": true,
}

var _ = Describe("CnvrgInfra controller", func() {

	const (
		timeout  = time.Second * 10
		duration = time.Second * 10
		interval = time.Millisecond * 250
	)

	if enabledInfraTests["monitoring"] {
		Context("Test Monitoring", func() {
			It("Prometheus Labels", func() {
				ns := createNs()
				ctx := context.Background()
				labels := map[string]string{"foo": "bar"}
				infra := getDefaultTestInfraSpec(ns)
				infra.Spec.Monitoring.Prometheus.Enabled = &defaultTrue
				infra.Spec.InfraNamespace = ns
				infra.Spec.Labels = labels
				Expect(k8sClient.Create(ctx, infra)).Should(Succeed())
				prom := &unstructured.Unstructured{}
				prom.SetGroupVersionKind(desired.Kinds[desired.PrometheusGVR])
				Eventually(func() bool {
					err := k8sClient.Get(ctx, types.NamespacedName{Name: "cnvrg-infra-prometheus", Namespace: ns}, prom)
					if err != nil {
						return false
					}
					return true
				}, timeout, interval).Should(BeTrue())
				a := prom.Object["spec"].(map[string]interface{})["podMetadata"].(map[string]interface{})["annotations"]
				l := prom.Object["spec"].(map[string]interface{})["podMetadata"].(map[string]interface{})["labels"]
				Expect(a).Should(BeNil())
				Expect(l).Should(HaveKeyWithValue("foo", "bar"))
			})
			It("Prometheus Operator Labels/Annotations", func() {
				ns := createNs()
				ctx := context.Background()
				labels := map[string]string{"foo": "bar"}
				annotations := map[string]string{"foo1": "bar1"}
				infra := getDefaultTestInfraSpec(ns)
				infra.Spec.Monitoring.PrometheusOperator.Enabled = &defaultTrue
				infra.Spec.InfraNamespace = ns
				infra.Spec.Labels = labels
				infra.Spec.Annotations = annotations
				Expect(k8sClient.Create(ctx, infra)).Should(Succeed())
				deployment := v1.Deployment{}
				Eventually(func() bool {
					err := k8sClient.Get(ctx, types.NamespacedName{Name: "cnvrg-prometheus-operator", Namespace: ns}, &deployment)
					if err != nil {
						return false
					}
					return true
				}, timeout, interval).Should(BeTrue())
				Expect(deployment.Labels).Should(HaveKeyWithValue("foo", "bar"))
				Expect(deployment.Annotations).Should(HaveKeyWithValue("foo1", "bar1"))
			})
			It("Prometheus NodeExporter Labels/Annotations", func() {
				ns := createNs()
				ctx := context.Background()
				labels := map[string]string{"foo": "bar"}
				annotations := map[string]string{"foo1": "bar1"}
				infra := getDefaultTestInfraSpec(ns)
				infra.Spec.Monitoring.NodeExporter.Enabled = &defaultTrue
				infra.Spec.InfraNamespace = ns
				infra.Spec.Labels = labels
				infra.Spec.Annotations = annotations
				Expect(k8sClient.Create(ctx, infra)).Should(Succeed())
				ds := v1.DaemonSet{}
				Eventually(func() bool {
					err := k8sClient.Get(ctx, types.NamespacedName{Name: "node-exporter", Namespace: ns}, &ds)
					if err != nil {
						return false
					}
					return true
				}, timeout, interval).Should(BeTrue())
				Expect(ds.Labels).Should(HaveKeyWithValue("foo", "bar"))
				Expect(ds.Annotations).Should(HaveKeyWithValue("foo1", "bar1"))
			})
			It("Kube State Metrics Labels/Annotations", func() {
				ns := createNs()
				ctx := context.Background()
				labels := map[string]string{"foo": "bar"}
				annotations := map[string]string{"foo1": "bar1"}
				infra := getDefaultTestInfraSpec(ns)
				infra.Spec.Monitoring.KubeStateMetrics.Enabled = &defaultTrue
				infra.Spec.InfraNamespace = ns
				infra.Spec.Labels = labels
				infra.Spec.Annotations = annotations
				Expect(k8sClient.Create(ctx, infra)).Should(Succeed())
				deployment := v1.Deployment{}
				Eventually(func() bool {
					err := k8sClient.Get(ctx, types.NamespacedName{Name: "kube-state-metrics", Namespace: ns}, &deployment)
					if err != nil {
						return false
					}
					return true
				}, timeout, interval).Should(BeTrue())
				Expect(deployment.Labels).Should(HaveKeyWithValue("foo", "bar"))
				Expect(deployment.Annotations).Should(HaveKeyWithValue("foo1", "bar1"))
			})
			It("Grafana Labels/Annotations", func() {
				ns := createNs()
				ctx := context.Background()
				labels := map[string]string{"foo": "bar"}
				annotations := map[string]string{"foo1": "bar1"}
				infra := getDefaultTestInfraSpec(ns)
				infra.Spec.Monitoring.Grafana.Enabled = &defaultTrue
				infra.Spec.Monitoring.Prometheus.Enabled = &defaultTrue
				infra.Spec.InfraNamespace = ns
				infra.Spec.Labels = labels
				infra.Spec.Annotations = annotations
				Expect(k8sClient.Create(ctx, infra)).Should(Succeed())
				deployment := v1.Deployment{}
				Eventually(func() bool {
					err := k8sClient.Get(ctx, types.NamespacedName{Name: "grafana", Namespace: ns}, &deployment)
					if err != nil {
						return false
					}
					return true
				}, timeout, interval).Should(BeTrue())
				Expect(deployment.Labels).Should(HaveKeyWithValue("foo", "bar"))
				Expect(deployment.Annotations).Should(HaveKeyWithValue("foo1", "bar1"))
			})
			It("Dcgm exporter Labels/Annotations", func() {
				ns := createNs()
				ctx := context.Background()
				labels := map[string]string{"foo": "bar"}
				annotations := map[string]string{"foo1": "bar1"}
				infra := getDefaultTestInfraSpec(ns)
				infra.Spec.Monitoring.DcgmExporter.Enabled = &defaultTrue
				infra.Spec.InfraNamespace = ns
				infra.Spec.Labels = labels
				infra.Spec.Annotations = annotations
				Expect(k8sClient.Create(ctx, infra)).Should(Succeed())
				ds := v1.DaemonSet{}
				Eventually(func() bool {
					err := k8sClient.Get(ctx, types.NamespacedName{Name: "dcgm-exporter", Namespace: ns}, &ds)
					if err != nil {
						return false
					}
					return true
				}, timeout, interval).Should(BeTrue())
				Expect(ds.Labels).Should(HaveKeyWithValue("foo", "bar"))
				Expect(ds.Annotations).Should(HaveKeyWithValue("foo1", "bar1"))
			})
		})
	}

	if enabledInfraTests["storage"] {
		Context("Test Storage", func() {
			It("Hostpath provisioner Labels/Annotations", func() {
				ns := createNs()
				ctx := context.Background()
				labels := map[string]string{"foo": "bar"}
				annotations := map[string]string{"foo1": "bar1"}
				infra := getDefaultTestInfraSpec(ns)
				infra.Spec.Storage.Hostpath.Enabled = &defaultTrue
				infra.Spec.InfraNamespace = ns
				infra.Spec.Labels = labels
				infra.Spec.Annotations = annotations
				Expect(k8sClient.Create(ctx, infra)).Should(Succeed())
				ds := v1.DaemonSet{}
				Eventually(func() bool {
					err := k8sClient.Get(ctx, types.NamespacedName{Name: "hostpath-provisioner", Namespace: ns}, &ds)
					if err != nil {
						return false
					}
					return true
				}, timeout, interval).Should(BeTrue())
				Expect(ds.Labels).Should(HaveKeyWithValue("foo", "bar"))
				Expect(ds.Annotations).Should(HaveKeyWithValue("foo1", "bar1"))
			})
			It("Nfs provisioner Labels/Annotations", func() {
				ns := createNs()
				ctx := context.Background()
				labels := map[string]string{"foo": "bar"}
				annotations := map[string]string{"foo1": "bar1"}
				infra := getDefaultTestInfraSpec(ns)
				infra.Spec.Storage.Nfs.Enabled = &defaultTrue
				infra.Spec.Storage.Nfs.Path = "/nfs-path"
				infra.Spec.Storage.Nfs.Server = "nfs-server"
				infra.Spec.InfraNamespace = ns
				infra.Spec.Labels = labels
				infra.Spec.Annotations = annotations
				Expect(k8sClient.Create(ctx, infra)).Should(Succeed())
				d := v1.Deployment{}
				Eventually(func() bool {
					err := k8sClient.Get(ctx, types.NamespacedName{Name: "nfs-client-provisioner", Namespace: ns}, &d)
					if err != nil {
						return false
					}
					return true
				}, timeout, interval).Should(BeTrue())
				Expect(d.Labels).Should(HaveKeyWithValue("foo", "bar"))
				Expect(d.Annotations).Should(HaveKeyWithValue("foo1", "bar1"))
			})

		})
	}

	if enabledInfraTests["networking"] {
		Context("Test networking", func() {
			It("Istio Labels/Annotations", func() {
				ns := createNs()
				ctx := context.Background()
				labels := map[string]string{"foo": "bar"}
				infra := getDefaultTestInfraSpec(ns)
				infra.Spec.Networking.Istio.Enabled = &defaultTrue
				infra.Spec.InfraNamespace = ns
				infra.Spec.Labels = labels
				Expect(k8sClient.Create(ctx, infra)).Should(Succeed())
				istio := &unstructured.Unstructured{}
				istio.SetGroupVersionKind(desired.Kinds[desired.IstioGVR])
				Eventually(func() bool {
					err := k8sClient.Get(ctx, types.NamespacedName{Name: "cnvrg-istio", Namespace: ns}, istio)
					if err != nil {
						return false
					}
					return true
				}, timeout, interval).Should(BeTrue())
				Expect(istio.GetAnnotations()).Should(BeNil())
				Expect(istio.GetLabels()).Should(HaveKeyWithValue("foo", "bar"))
			})
			It("Istio Service Annotations", func() {
				ns := createNs()
				ctx := context.Background()
				labels := map[string]string{"foo": "bar"}
				infra := getDefaultTestInfraSpec(ns)
				infra.Spec.Networking.Istio.Enabled = &defaultTrue
				infra.Spec.Networking.Istio.IngressSvcAnnotations = map[string]string{"foo": "bar"}
				infra.Spec.InfraNamespace = ns
				infra.Spec.Labels = labels
				Expect(k8sClient.Create(ctx, infra)).Should(Succeed())
				istio := &unstructured.Unstructured{}
				istio.SetGroupVersionKind(desired.Kinds[desired.IstioGVR])
				Eventually(func() bool {
					err := k8sClient.Get(ctx, types.NamespacedName{Name: "cnvrg-istio", Namespace: ns}, istio)
					if err != nil {
						return false
					}
					return true
				}, timeout, interval).Should(BeTrue())
				serviceAnnotations := istio.Object["spec"].(map[string]interface{})["components"].(map[string]interface{})["ingressGateways"].([]interface{})[0].(map[string]interface{})["k8s"].(map[string]interface{})["serviceAnnotations"]
				Expect(serviceAnnotations).Should(HaveKeyWithValue("foo", "bar"))
			})
			It("Istio LoadBalancer Source Ranges", func() {
				ns := createNs()
				ctx := context.Background()
				labels := map[string]string{"foo": "bar"}
				infra := getDefaultTestInfraSpec(ns)
				infra.Spec.Networking.Istio.Enabled = &defaultTrue
				infra.Spec.Networking.Istio.LBSourceRanges = []string{"1.1.1.1/32", "2.2.2.2/24"}
				infra.Spec.InfraNamespace = ns
				infra.Spec.Labels = labels
				Expect(k8sClient.Create(ctx, infra)).Should(Succeed())
				istio := &unstructured.Unstructured{}
				istio.SetGroupVersionKind(desired.Kinds[desired.IstioGVR])
				Eventually(func() bool {
					err := k8sClient.Get(ctx, types.NamespacedName{Name: "cnvrg-istio", Namespace: ns}, istio)
					if err != nil {
						return false
					}
					return true
				}, timeout, interval).Should(BeTrue())
				lbSources := []string{
					istio.Object["spec"].(map[string]interface{})["components"].(map[string]interface{})["ingressGateways"].([]interface{})[0].(map[string]interface{})["k8s"].(map[string]interface{})["service"].(map[string]interface{})["loadBalancerSourceRanges"].([]interface{})[0].(string),
					istio.Object["spec"].(map[string]interface{})["components"].(map[string]interface{})["ingressGateways"].([]interface{})[0].(map[string]interface{})["k8s"].(map[string]interface{})["service"].(map[string]interface{})["loadBalancerSourceRanges"].([]interface{})[1].(string),
				}
				Expect(lbSources).Should(Equal([]string{"1.1.1.1/32", "2.2.2.2/24"}))
			})
			It("Istio ExternalIps", func() {
				ns := createNs()
				ctx := context.Background()
				labels := map[string]string{"foo": "bar"}
				exIps := []string{"1.1.1.1", "2.2.2.2"}
				infra := getDefaultTestInfraSpec(ns)
				infra.Spec.Networking.Istio.Enabled = &defaultTrue
				infra.Spec.Networking.Istio.ExternalIP = exIps
				infra.Spec.InfraNamespace = ns
				infra.Spec.Labels = labels
				Expect(k8sClient.Create(ctx, infra)).Should(Succeed())
				istio := &unstructured.Unstructured{}
				istio.SetGroupVersionKind(desired.Kinds[desired.IstioGVR])
				Eventually(func() bool {
					err := k8sClient.Get(ctx, types.NamespacedName{Name: "cnvrg-istio", Namespace: ns}, istio)
					if err != nil {
						return false
					}
					return true
				}, timeout, interval).Should(BeTrue())
				externalIps := []string{
					istio.Object["spec"].(map[string]interface{})["components"].(map[string]interface{})["ingressGateways"].([]interface{})[0].(map[string]interface{})["k8s"].(map[string]interface{})["service"].(map[string]interface{})["externalIPs"].([]interface{})[0].(string),
					istio.Object["spec"].(map[string]interface{})["components"].(map[string]interface{})["ingressGateways"].([]interface{})[0].(map[string]interface{})["k8s"].(map[string]interface{})["service"].(map[string]interface{})["externalIPs"].([]interface{})[1].(string),
				}
				Expect(externalIps).Should(Equal(exIps))
			})
			It("Istio ExtraPorts", func() {
				ns := createNs()
				ctx := context.Background()
				labels := map[string]string{"foo": "bar"}
				exPorts := []int{1111, 2222}
				infra := getDefaultTestInfraSpec(ns)
				infra.Spec.Networking.Istio.Enabled = &defaultTrue
				infra.Spec.Networking.Istio.IngressSvcExtraPorts = exPorts
				infra.Spec.InfraNamespace = ns
				infra.Spec.Labels = labels
				Expect(k8sClient.Create(ctx, infra)).Should(Succeed())
				istio := &unstructured.Unstructured{}
				istio.SetGroupVersionKind(desired.Kinds[desired.IstioGVR])
				Eventually(func() bool {
					err := k8sClient.Get(ctx, types.NamespacedName{Name: "cnvrg-istio", Namespace: ns}, istio)
					if err != nil {
						return false
					}
					return true
				}, timeout, interval).Should(BeTrue())
				port1111 := istio.Object["spec"].(map[string]interface{})["components"].(map[string]interface{})["ingressGateways"].([]interface{})[0].(map[string]interface{})["k8s"].(map[string]interface{})["service"].(map[string]interface{})["ports"].([]interface{})[2]
				port2222 := istio.Object["spec"].(map[string]interface{})["components"].(map[string]interface{})["ingressGateways"].([]interface{})[0].(map[string]interface{})["k8s"].(map[string]interface{})["service"].(map[string]interface{})["ports"].([]interface{})[3]
				var port1111int64 int64 = 1111
				var port2222int64 int64 = 2222
				Expect(port1111).Should(HaveKeyWithValue("name", "port1111"))
				Expect(port1111).Should(HaveKeyWithValue("port", port1111int64))
				Expect(port2222).Should(HaveKeyWithValue("name", "port2222"))
				Expect(port2222).Should(HaveKeyWithValue("port", port2222int64))
			})
		})
	}

})

func getDefaultTestInfraSpec(ns string) *mlopsv1.CnvrgInfra {
	testSpec := mlopsv1.DefaultCnvrgInfraSpec()
	return &mlopsv1.CnvrgInfra{
		TypeMeta: metav1.TypeMeta{
			Kind:       "CnvrgInfra",
			APIVersion: "mlops.cnvrg.io/v1"},

		ObjectMeta: metav1.ObjectMeta{
			Name:      ns,
			Namespace: ns,
		},
		Spec: testSpec,
	}
}