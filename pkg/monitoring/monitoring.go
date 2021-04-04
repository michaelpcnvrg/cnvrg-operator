package monitoring

import (
	mlopsv1 "github.com/cnvrg-operator/api/v1"
	"github.com/cnvrg-operator/pkg/desired"
	"github.com/markbates/pkger"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"os"
)

const path = "/pkg/monitoring/tmpl"

func prometheusOperatorState() []*desired.State {
	return []*desired.State{
		{
			TemplatePath:   path + "/prometheus/operator/clusterrole.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.ClusterRoleGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/prometheus/operator/clusterrolebinding.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.ClusterRoleBindingGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/prometheus/operator/sa.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SaGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/prometheus/operator/dep.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.DeploymentGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/prometheus/operator/svc.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SvcGVR],
			Own:            true,
		},
	}
}

func infraPrometheusInstanceState() []*desired.State {
	return []*desired.State{
		{
			TemplatePath:   path + "/prometheus/instance/infra/clusterrole.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.ClusterRoleGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/prometheus/instance/infra/clusterrolebinding.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.ClusterRoleBindingGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/prometheus/instance/infra/sa.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SaGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/prometheus/instance/infra/prometheus.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.PrometheusGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/prometheus/instance/infra/svc.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SvcGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/prometheus/instance/infra/vs.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.IstioVsGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/prometheus/instance/infra/rules.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.PrometheusRuleGVR],
			Own:            true,
		},
	}
}

func kubeStateMetricsState() []*desired.State {
	return []*desired.State{
		{
			TemplatePath:   path + "/kube-state-metrics/clusterrole.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.ClusterRoleGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/kube-state-metrics/clusterrolebinding.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.ClusterRoleBindingGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/kube-state-metrics/dep.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.DeploymentGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/kube-state-metrics/sa.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SaGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/kube-state-metrics/servicemonitor.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.ServiceMonitorGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/kube-state-metrics/svc.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SvcGVR],
			Own:            true,
		},
	}
}

func grafanaState() []*desired.State {
	return []*desired.State{
		{
			TemplatePath:   path + "/grafana/sa.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SaGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/grafana/svc.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SvcGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/grafana/dep.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.DeploymentGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/grafana/datasource.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SecretGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/grafana/dashboards.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.ConfigMapGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/grafana/vs.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.IstioVsGVR],
			Own:            true,
		},
	}

}

func nodeExporterState() []*desired.State {
	return []*desired.State{
		{
			TemplatePath:   path + "/node-exporter/clusterrole.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.ClusterRoleGVR],
			Own:            true,
		},

		{
			TemplatePath:   path + "/node-exporter/clusterrolebinding.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.ClusterRoleBindingGVR],
			Own:            true,
		},

		{
			TemplatePath:   path + "/node-exporter/sa.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SaGVR],
			Own:            true,
		},

		{
			TemplatePath:   path + "/node-exporter/svc.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SvcGVR],
			Own:            true,
		},

		{
			TemplatePath:   path + "/node-exporter/ds.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.DaemonSetGVR],
			Own:            true,
		},

		{
			TemplatePath:   path + "/node-exporter/servicemonitor.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.ServiceMonitorGVR],
			Own:            true,
		},
	}
}

func ccpPrometheusInstance() []*desired.State {
	return []*desired.State{
		{
			TemplatePath:   path + "/prometheus/instance/ccp/role.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.RoleGVR],
			Own:            false,
		},
		{
			TemplatePath:   path + "/prometheus/instance/ccp/rolebinding.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.RoleBindingGVR],
			Own:            false,
		},
		{
			TemplatePath:   path + "/prometheus/instance/ccp/sa.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SaGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/prometheus/instance/ccp/prometheus.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.PrometheusGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/prometheus/instance/ccp/svc.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SvcGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/prometheus/instance/ccp/vs.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.IstioVsGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/prometheus/instance/ccp/staticconfig.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SecretGVR],
			Own:            true,
		},
	}
}

func defaultServiceMonitors() []*desired.State {
	return []*desired.State{
		{
			TemplatePath:   path + "/default-servicemonitors/apiserver.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.ServiceMonitorGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/default-servicemonitors/controller-manager.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.ServiceMonitorGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/default-servicemonitors/coredns.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.ServiceMonitorGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/default-servicemonitors/kubelet.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.ServiceMonitorGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/default-servicemonitors/scheduler.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.ServiceMonitorGVR],
			Own:            true,
		},
	}
}

func dcgmExporter() []*desired.State {
	return []*desired.State{
		{
			TemplatePath:   path + "/dcgm-exporter/ds.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.DaemonSetGVR],
			Own:            true,
		},

		{
			TemplatePath:   path + "/dcgm-exporter/sa.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SaGVR],
			Own:            true,
		},

		{
			TemplatePath:   path + "/dcgm-exporter/servicemonitor.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.ServiceMonitorGVR],
			Own:            true,
		},

		{
			TemplatePath:   path + "/dcgm-exporter/svc.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SvcGVR],
			Own:            true,
		},
	}
}

func grafanaOauthProxy() []*desired.State {
	return []*desired.State{
		{
			TemplatePath:   path + "/grafana/oauth.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SecretGVR],
			Own:            true,
		},
	}
}

func appServiceMonitors() []*desired.State {
	return []*desired.State{
		{
			TemplatePath:   path + "/default-servicemonitors/cnvrg-idle-metrics.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.ServiceMonitorGVR],
			Own:            true,
		},
	}
}

func AppMonitoringState(cnvrgApp *mlopsv1.CnvrgApp) []*desired.State {

	// always add cnvrg idle metrics exporter
	var state = appServiceMonitors()

	if cnvrgApp.Spec.Monitoring.Enabled == "true" && cnvrgApp.Spec.Monitoring.Prometheus.Enabled == "true" {
		state = append(state, ccpPrometheusInstance()...)
	}
	if cnvrgApp.Spec.Monitoring.Enabled == "true" && cnvrgApp.Spec.Monitoring.Grafana.Enabled == "true" {
		state = append(state, grafanaState()...)
	}
	if cnvrgApp.Spec.Monitoring.Enabled == "true" && cnvrgApp.Spec.SSO.Enabled == "true" {
		state = append(state, grafanaOauthProxy()...)
	}

	return state
}

func InfraMonitoringState(infra *mlopsv1.CnvrgInfra) []*desired.State {
	var state []*desired.State
	if infra.Spec.Monitoring.Enabled == "true" && infra.Spec.Monitoring.PrometheusOperator.Enabled == "true" {
		state = append(state, prometheusOperatorState()...)
	}
	if infra.Spec.Monitoring.Enabled == "true" && infra.Spec.Monitoring.Prometheus.Enabled == "true" {
		state = append(state, infraPrometheusInstanceState()...)
	}
	if infra.Spec.Monitoring.Enabled == "true" && infra.Spec.Monitoring.KubeStateMetrics.Enabled == "true" {
		state = append(state, kubeStateMetricsState()...)
	}
	if infra.Spec.Monitoring.Enabled == "true" && infra.Spec.Monitoring.Grafana.Enabled == "true" {
		state = append(state, grafanaState()...)
	}
	if infra.Spec.Monitoring.Enabled == "true" && infra.Spec.Monitoring.NodeExporter.Enabled == "true" {
		state = append(state, nodeExporterState()...)
	}
	if infra.Spec.Monitoring.Enabled == "true" {
		state = append(state, defaultServiceMonitors()...)
	}
	if infra.Spec.Monitoring.Enabled == "true" && infra.Spec.Monitoring.DcgmExporter.Enabled == "true" {
		state = append(state, dcgmExporter()...)
	}
	if infra.Spec.Monitoring.Enabled == "true" && infra.Spec.SSO.Enabled == "true" {
		state = append(state, grafanaOauthProxy()...)
	}
	return state
}

func Crds() (crds []*desired.State) {
	err := pkger.Walk(path+"/crds", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		crd := &desired.State{

			TemplatePath:   path,
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.CrdGVR],
			Own:            false,
		}
		crds = append(crds, crd)
		return nil
	})
	if err != nil {
		zap.S().Error(err, "error loading prometheus crds")
	}
	return
}
