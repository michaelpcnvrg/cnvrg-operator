apiVersion: monitoring.coreos.com/v1
kind: PrometheusRule
metadata:
  labels:
    app: cnvrg-infra-prometheus
    role: alert-rules
  name: prometheus-k8s-rules
  namespace: {{ ns . }}
spec:
  groups:
    - name: node-exporter.rules
      rules:
        - expr: |
            count without (cpu) (
              count without (mode) (
                node_cpu_seconds_total{job="node-exporter"}
              )
            )
          record: instance:node_num_cpu:sum
        - expr: |
            1 - avg without (cpu, mode) (
              rate(node_cpu_seconds_total{job="node-exporter", mode="idle"}[1m])
            )
          record: instance:node_cpu_utilisation:rate1m
        - expr: |
            (
              node_load1{job="node-exporter"}
            /
              instance:node_num_cpu:sum{job="node-exporter"}
            )
          record: instance:node_load1_per_cpu:ratio
        - expr: |
            1 - (
              node_memory_MemAvailable_bytes{job="node-exporter"}
            /
              node_memory_MemTotal_bytes{job="node-exporter"}
            )
          record: instance:node_memory_utilisation:ratio
        - expr: |
            rate(node_vmstat_pgmajfault{job="node-exporter"}[1m])
          record: instance:node_vmstat_pgmajfault:rate1m
        - expr: |
            rate(node_disk_io_time_seconds_total{job="node-exporter", device=~"mmcblk.p.+|nvme.+|rbd.+|sd.+|vd.+|xvd.+|dm-.+|dasd.+"}[1m])
          record: instance_device:node_disk_io_time_seconds:rate1m
        - expr: |
            rate(node_disk_io_time_weighted_seconds_total{job="node-exporter", device=~"mmcblk.p.+|nvme.+|rbd.+|sd.+|vd.+|xvd.+|dm-.+|dasd.+"}[1m])
          record: instance_device:node_disk_io_time_weighted_seconds:rate1m
        - expr: |
            sum without (device) (
              rate(node_network_receive_bytes_total{job="node-exporter", device!="lo"}[1m])
            )
          record: instance:node_network_receive_bytes_excluding_lo:rate1m
        - expr: |
            sum without (device) (
              rate(node_network_transmit_bytes_total{job="node-exporter", device!="lo"}[1m])
            )
          record: instance:node_network_transmit_bytes_excluding_lo:rate1m
        - expr: |
            sum without (device) (
              rate(node_network_receive_drop_total{job="node-exporter", device!="lo"}[1m])
            )
          record: instance:node_network_receive_drop_excluding_lo:rate1m
        - expr: |
            sum without (device) (
              rate(node_network_transmit_drop_total{job="node-exporter", device!="lo"}[1m])
            )
          record: instance:node_network_transmit_drop_excluding_lo:rate1m
    - name: kube-apiserver.rules
      rules:
        - expr: |
            (
              (
                # too slow
                sum(rate(apiserver_request_duration_seconds_count{job="apiserver",verb=~"LIST|GET"}[1d]))
                -
                (
                  (
                    sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"LIST|GET",scope=~"resource|",le="0.1"}[1d]))
                    or
                    vector(0)
                  )
                  +
                  sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"LIST|GET",scope="namespace",le="0.5"}[1d]))
                  +
                  sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"LIST|GET",scope="cluster",le="5"}[1d]))
                )
              )
              +
              # errors
              sum(rate(apiserver_request_total{job="apiserver",verb=~"LIST|GET",code=~"5.."}[1d]))
            )
            /
            sum(rate(apiserver_request_total{job="apiserver",verb=~"LIST|GET"}[1d]))
          labels:
            verb: read
          record: apiserver_request:burnrate1d
        - expr: |
            (
              (
                # too slow
                sum(rate(apiserver_request_duration_seconds_count{job="apiserver",verb=~"LIST|GET"}[1h]))
                -
                (
                  (
                    sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"LIST|GET",scope=~"resource|",le="0.1"}[1h]))
                    or
                    vector(0)
                  )
                  +
                  sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"LIST|GET",scope="namespace",le="0.5"}[1h]))
                  +
                  sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"LIST|GET",scope="cluster",le="5"}[1h]))
                )
              )
              +
              # errors
              sum(rate(apiserver_request_total{job="apiserver",verb=~"LIST|GET",code=~"5.."}[1h]))
            )
            /
            sum(rate(apiserver_request_total{job="apiserver",verb=~"LIST|GET"}[1h]))
          labels:
            verb: read
          record: apiserver_request:burnrate1h
        - expr: |
            (
              (
                # too slow
                sum(rate(apiserver_request_duration_seconds_count{job="apiserver",verb=~"LIST|GET"}[2h]))
                -
                (
                  (
                    sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"LIST|GET",scope=~"resource|",le="0.1"}[2h]))
                    or
                    vector(0)
                  )
                  +
                  sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"LIST|GET",scope="namespace",le="0.5"}[2h]))
                  +
                  sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"LIST|GET",scope="cluster",le="5"}[2h]))
                )
              )
              +
              # errors
              sum(rate(apiserver_request_total{job="apiserver",verb=~"LIST|GET",code=~"5.."}[2h]))
            )
            /
            sum(rate(apiserver_request_total{job="apiserver",verb=~"LIST|GET"}[2h]))
          labels:
            verb: read
          record: apiserver_request:burnrate2h
        - expr: |
            (
              (
                # too slow
                sum(rate(apiserver_request_duration_seconds_count{job="apiserver",verb=~"LIST|GET"}[30m]))
                -
                (
                  (
                    sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"LIST|GET",scope=~"resource|",le="0.1"}[30m]))
                    or
                    vector(0)
                  )
                  +
                  sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"LIST|GET",scope="namespace",le="0.5"}[30m]))
                  +
                  sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"LIST|GET",scope="cluster",le="5"}[30m]))
                )
              )
              +
              # errors
              sum(rate(apiserver_request_total{job="apiserver",verb=~"LIST|GET",code=~"5.."}[30m]))
            )
            /
            sum(rate(apiserver_request_total{job="apiserver",verb=~"LIST|GET"}[30m]))
          labels:
            verb: read
          record: apiserver_request:burnrate30m
        - expr: |
            (
              (
                # too slow
                sum(rate(apiserver_request_duration_seconds_count{job="apiserver",verb=~"LIST|GET"}[3d]))
                -
                (
                  (
                    sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"LIST|GET",scope=~"resource|",le="0.1"}[3d]))
                    or
                    vector(0)
                  )
                  +
                  sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"LIST|GET",scope="namespace",le="0.5"}[3d]))
                  +
                  sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"LIST|GET",scope="cluster",le="5"}[3d]))
                )
              )
              +
              # errors
              sum(rate(apiserver_request_total{job="apiserver",verb=~"LIST|GET",code=~"5.."}[3d]))
            )
            /
            sum(rate(apiserver_request_total{job="apiserver",verb=~"LIST|GET"}[3d]))
          labels:
            verb: read
          record: apiserver_request:burnrate3d
        - expr: |
            (
              (
                # too slow
                sum(rate(apiserver_request_duration_seconds_count{job="apiserver",verb=~"LIST|GET"}[5m]))
                -
                (
                  (
                    sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"LIST|GET",scope=~"resource|",le="0.1"}[5m]))
                    or
                    vector(0)
                  )
                  +
                  sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"LIST|GET",scope="namespace",le="0.5"}[5m]))
                  +
                  sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"LIST|GET",scope="cluster",le="5"}[5m]))
                )
              )
              +
              # errors
              sum(rate(apiserver_request_total{job="apiserver",verb=~"LIST|GET",code=~"5.."}[5m]))
            )
            /
            sum(rate(apiserver_request_total{job="apiserver",verb=~"LIST|GET"}[5m]))
          labels:
            verb: read
          record: apiserver_request:burnrate5m
        - expr: |
            (
              (
                # too slow
                sum(rate(apiserver_request_duration_seconds_count{job="apiserver",verb=~"LIST|GET"}[6h]))
                -
                (
                  (
                    sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"LIST|GET",scope=~"resource|",le="0.1"}[6h]))
                    or
                    vector(0)
                  )
                  +
                  sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"LIST|GET",scope="namespace",le="0.5"}[6h]))
                  +
                  sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"LIST|GET",scope="cluster",le="5"}[6h]))
                )
              )
              +
              # errors
              sum(rate(apiserver_request_total{job="apiserver",verb=~"LIST|GET",code=~"5.."}[6h]))
            )
            /
            sum(rate(apiserver_request_total{job="apiserver",verb=~"LIST|GET"}[6h]))
          labels:
            verb: read
          record: apiserver_request:burnrate6h
        - expr: |
            (
              (
                # too slow
                sum(rate(apiserver_request_duration_seconds_count{job="apiserver",verb=~"POST|PUT|PATCH|DELETE"}[1d]))
                -
                sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"POST|PUT|PATCH|DELETE",le="1"}[1d]))
              )
              +
              sum(rate(apiserver_request_total{job="apiserver",verb=~"POST|PUT|PATCH|DELETE",code=~"5.."}[1d]))
            )
            /
            sum(rate(apiserver_request_total{job="apiserver",verb=~"POST|PUT|PATCH|DELETE"}[1d]))
          labels:
            verb: write
          record: apiserver_request:burnrate1d
        - expr: |
            (
              (
                # too slow
                sum(rate(apiserver_request_duration_seconds_count{job="apiserver",verb=~"POST|PUT|PATCH|DELETE"}[1h]))
                -
                sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"POST|PUT|PATCH|DELETE",le="1"}[1h]))
              )
              +
              sum(rate(apiserver_request_total{job="apiserver",verb=~"POST|PUT|PATCH|DELETE",code=~"5.."}[1h]))
            )
            /
            sum(rate(apiserver_request_total{job="apiserver",verb=~"POST|PUT|PATCH|DELETE"}[1h]))
          labels:
            verb: write
          record: apiserver_request:burnrate1h
        - expr: |
            (
              (
                # too slow
                sum(rate(apiserver_request_duration_seconds_count{job="apiserver",verb=~"POST|PUT|PATCH|DELETE"}[2h]))
                -
                sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"POST|PUT|PATCH|DELETE",le="1"}[2h]))
              )
              +
              sum(rate(apiserver_request_total{job="apiserver",verb=~"POST|PUT|PATCH|DELETE",code=~"5.."}[2h]))
            )
            /
            sum(rate(apiserver_request_total{job="apiserver",verb=~"POST|PUT|PATCH|DELETE"}[2h]))
          labels:
            verb: write
          record: apiserver_request:burnrate2h
        - expr: |
            (
              (
                # too slow
                sum(rate(apiserver_request_duration_seconds_count{job="apiserver",verb=~"POST|PUT|PATCH|DELETE"}[30m]))
                -
                sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"POST|PUT|PATCH|DELETE",le="1"}[30m]))
              )
              +
              sum(rate(apiserver_request_total{job="apiserver",verb=~"POST|PUT|PATCH|DELETE",code=~"5.."}[30m]))
            )
            /
            sum(rate(apiserver_request_total{job="apiserver",verb=~"POST|PUT|PATCH|DELETE"}[30m]))
          labels:
            verb: write
          record: apiserver_request:burnrate30m
        - expr: |
            (
              (
                # too slow
                sum(rate(apiserver_request_duration_seconds_count{job="apiserver",verb=~"POST|PUT|PATCH|DELETE"}[3d]))
                -
                sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"POST|PUT|PATCH|DELETE",le="1"}[3d]))
              )
              +
              sum(rate(apiserver_request_total{job="apiserver",verb=~"POST|PUT|PATCH|DELETE",code=~"5.."}[3d]))
            )
            /
            sum(rate(apiserver_request_total{job="apiserver",verb=~"POST|PUT|PATCH|DELETE"}[3d]))
          labels:
            verb: write
          record: apiserver_request:burnrate3d
        - expr: |
            (
              (
                # too slow
                sum(rate(apiserver_request_duration_seconds_count{job="apiserver",verb=~"POST|PUT|PATCH|DELETE"}[5m]))
                -
                sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"POST|PUT|PATCH|DELETE",le="1"}[5m]))
              )
              +
              sum(rate(apiserver_request_total{job="apiserver",verb=~"POST|PUT|PATCH|DELETE",code=~"5.."}[5m]))
            )
            /
            sum(rate(apiserver_request_total{job="apiserver",verb=~"POST|PUT|PATCH|DELETE"}[5m]))
          labels:
            verb: write
          record: apiserver_request:burnrate5m
        - expr: |
            (
              (
                # too slow
                sum(rate(apiserver_request_duration_seconds_count{job="apiserver",verb=~"POST|PUT|PATCH|DELETE"}[6h]))
                -
                sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"POST|PUT|PATCH|DELETE",le="1"}[6h]))
              )
              +
              sum(rate(apiserver_request_total{job="apiserver",verb=~"POST|PUT|PATCH|DELETE",code=~"5.."}[6h]))
            )
            /
            sum(rate(apiserver_request_total{job="apiserver",verb=~"POST|PUT|PATCH|DELETE"}[6h]))
          labels:
            verb: write
          record: apiserver_request:burnrate6h
        - expr: |
            sum by (code,resource) (rate(apiserver_request_total{job="apiserver",verb=~"LIST|GET"}[5m]))
          labels:
            verb: read
          record: code_resource:apiserver_request_total:rate5m
        - expr: |
            sum by (code,resource) (rate(apiserver_request_total{job="apiserver",verb=~"POST|PUT|PATCH|DELETE"}[5m]))
          labels:
            verb: write
          record: code_resource:apiserver_request_total:rate5m
        - expr: |
            histogram_quantile(0.99, sum by (le, resource) (rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"LIST|GET"}[5m]))) > 0
          labels:
            quantile: "0.99"
            verb: read
          record: cluster_quantile:apiserver_request_duration_seconds:histogram_quantile
        - expr: |
            histogram_quantile(0.99, sum by (le, resource) (rate(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"POST|PUT|PATCH|DELETE"}[5m]))) > 0
          labels:
            quantile: "0.99"
            verb: write
          record: cluster_quantile:apiserver_request_duration_seconds:histogram_quantile
        - expr: |
            histogram_quantile(0.99, sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",subresource!="log",verb!~"LIST|WATCH|WATCHLIST|DELETECOLLECTION|PROXY|CONNECT"}[5m])) without(instance, pod))
          labels:
            quantile: "0.99"
          record: cluster_quantile:apiserver_request_duration_seconds:histogram_quantile
        - expr: |
            histogram_quantile(0.9, sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",subresource!="log",verb!~"LIST|WATCH|WATCHLIST|DELETECOLLECTION|PROXY|CONNECT"}[5m])) without(instance, pod))
          labels:
            quantile: "0.9"
          record: cluster_quantile:apiserver_request_duration_seconds:histogram_quantile
        - expr: |
            histogram_quantile(0.5, sum(rate(apiserver_request_duration_seconds_bucket{job="apiserver",subresource!="log",verb!~"LIST|WATCH|WATCHLIST|DELETECOLLECTION|PROXY|CONNECT"}[5m])) without(instance, pod))
          labels:
            quantile: "0.5"
          record: cluster_quantile:apiserver_request_duration_seconds:histogram_quantile
    - interval: 3m
      name: kube-apiserver-availability.rules
      rules:
        - expr: |
            1 - (
              (
                # write too slow
                sum(increase(apiserver_request_duration_seconds_count{verb=~"POST|PUT|PATCH|DELETE"}[30d]))
                -
                sum(increase(apiserver_request_duration_seconds_bucket{verb=~"POST|PUT|PATCH|DELETE",le="1"}[30d]))
              ) +
              (
                # read too slow
                sum(increase(apiserver_request_duration_seconds_count{verb=~"LIST|GET"}[30d]))
                -
                (
                  (
                    sum(increase(apiserver_request_duration_seconds_bucket{verb=~"LIST|GET",scope=~"resource|",le="0.1"}[30d]))
                    or
                    vector(0)
                  )
                  +
                  sum(increase(apiserver_request_duration_seconds_bucket{verb=~"LIST|GET",scope="namespace",le="0.5"}[30d]))
                  +
                  sum(increase(apiserver_request_duration_seconds_bucket{verb=~"LIST|GET",scope="cluster",le="5"}[30d]))
                )
              ) +
              # errors
              sum(code:apiserver_request_total:increase30d{code=~"5.."} or vector(0))
            )
            /
            sum(code:apiserver_request_total:increase30d)
          labels:
            verb: all
          record: apiserver_request:availability30d
        - expr: |
            1 - (
              sum(increase(apiserver_request_duration_seconds_count{job="apiserver",verb=~"LIST|GET"}[30d]))
              -
              (
                # too slow
                (
                  sum(increase(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"LIST|GET",scope=~"resource|",le="0.1"}[30d]))
                  or
                  vector(0)
                )
                +
                sum(increase(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"LIST|GET",scope="namespace",le="0.5"}[30d]))
                +
                sum(increase(apiserver_request_duration_seconds_bucket{job="apiserver",verb=~"LIST|GET",scope="cluster",le="5"}[30d]))
              )
              +
              # errors
              sum(code:apiserver_request_total:increase30d{verb="read",code=~"5.."} or vector(0))
            )
            /
            sum(code:apiserver_request_total:increase30d{verb="read"})
          labels:
            verb: read
          record: apiserver_request:availability30d
        - expr: |
            1 - (
              (
                # too slow
                sum(increase(apiserver_request_duration_seconds_count{verb=~"POST|PUT|PATCH|DELETE"}[30d]))
                -
                sum(increase(apiserver_request_duration_seconds_bucket{verb=~"POST|PUT|PATCH|DELETE",le="1"}[30d]))
              )
              +
              # errors
              sum(code:apiserver_request_total:increase30d{verb="write",code=~"5.."} or vector(0))
            )
            /
            sum(code:apiserver_request_total:increase30d{verb="write"})
          labels:
            verb: write
          record: apiserver_request:availability30d
        - expr: |
            sum by (code, verb) (increase(apiserver_request_total{job="apiserver",verb="LIST",code=~"2.."}[30d]))
          record: code_verb:apiserver_request_total:increase30d
        - expr: |
            sum by (code, verb) (increase(apiserver_request_total{job="apiserver",verb="GET",code=~"2.."}[30d]))
          record: code_verb:apiserver_request_total:increase30d
        - expr: |
            sum by (code, verb) (increase(apiserver_request_total{job="apiserver",verb="POST",code=~"2.."}[30d]))
          record: code_verb:apiserver_request_total:increase30d
        - expr: |
            sum by (code, verb) (increase(apiserver_request_total{job="apiserver",verb="PUT",code=~"2.."}[30d]))
          record: code_verb:apiserver_request_total:increase30d
        - expr: |
            sum by (code, verb) (increase(apiserver_request_total{job="apiserver",verb="PATCH",code=~"2.."}[30d]))
          record: code_verb:apiserver_request_total:increase30d
        - expr: |
            sum by (code, verb) (increase(apiserver_request_total{job="apiserver",verb="DELETE",code=~"2.."}[30d]))
          record: code_verb:apiserver_request_total:increase30d
        - expr: |
            sum by (code, verb) (increase(apiserver_request_total{job="apiserver",verb="LIST",code=~"3.."}[30d]))
          record: code_verb:apiserver_request_total:increase30d
        - expr: |
            sum by (code, verb) (increase(apiserver_request_total{job="apiserver",verb="GET",code=~"3.."}[30d]))
          record: code_verb:apiserver_request_total:increase30d
        - expr: |
            sum by (code, verb) (increase(apiserver_request_total{job="apiserver",verb="POST",code=~"3.."}[30d]))
          record: code_verb:apiserver_request_total:increase30d
        - expr: |
            sum by (code, verb) (increase(apiserver_request_total{job="apiserver",verb="PUT",code=~"3.."}[30d]))
          record: code_verb:apiserver_request_total:increase30d
        - expr: |
            sum by (code, verb) (increase(apiserver_request_total{job="apiserver",verb="PATCH",code=~"3.."}[30d]))
          record: code_verb:apiserver_request_total:increase30d
        - expr: |
            sum by (code, verb) (increase(apiserver_request_total{job="apiserver",verb="DELETE",code=~"3.."}[30d]))
          record: code_verb:apiserver_request_total:increase30d
        - expr: |
            sum by (code, verb) (increase(apiserver_request_total{job="apiserver",verb="LIST",code=~"4.."}[30d]))
          record: code_verb:apiserver_request_total:increase30d
        - expr: |
            sum by (code, verb) (increase(apiserver_request_total{job="apiserver",verb="GET",code=~"4.."}[30d]))
          record: code_verb:apiserver_request_total:increase30d
        - expr: |
            sum by (code, verb) (increase(apiserver_request_total{job="apiserver",verb="POST",code=~"4.."}[30d]))
          record: code_verb:apiserver_request_total:increase30d
        - expr: |
            sum by (code, verb) (increase(apiserver_request_total{job="apiserver",verb="PUT",code=~"4.."}[30d]))
          record: code_verb:apiserver_request_total:increase30d
        - expr: |
            sum by (code, verb) (increase(apiserver_request_total{job="apiserver",verb="PATCH",code=~"4.."}[30d]))
          record: code_verb:apiserver_request_total:increase30d
        - expr: |
            sum by (code, verb) (increase(apiserver_request_total{job="apiserver",verb="DELETE",code=~"4.."}[30d]))
          record: code_verb:apiserver_request_total:increase30d
        - expr: |
            sum by (code, verb) (increase(apiserver_request_total{job="apiserver",verb="LIST",code=~"5.."}[30d]))
          record: code_verb:apiserver_request_total:increase30d
        - expr: |
            sum by (code, verb) (increase(apiserver_request_total{job="apiserver",verb="GET",code=~"5.."}[30d]))
          record: code_verb:apiserver_request_total:increase30d
        - expr: |
            sum by (code, verb) (increase(apiserver_request_total{job="apiserver",verb="POST",code=~"5.."}[30d]))
          record: code_verb:apiserver_request_total:increase30d
        - expr: |
            sum by (code, verb) (increase(apiserver_request_total{job="apiserver",verb="PUT",code=~"5.."}[30d]))
          record: code_verb:apiserver_request_total:increase30d
        - expr: |
            sum by (code, verb) (increase(apiserver_request_total{job="apiserver",verb="PATCH",code=~"5.."}[30d]))
          record: code_verb:apiserver_request_total:increase30d
        - expr: |
            sum by (code, verb) (increase(apiserver_request_total{job="apiserver",verb="DELETE",code=~"5.."}[30d]))
          record: code_verb:apiserver_request_total:increase30d
        - expr: |
            sum by (code) (code_verb:apiserver_request_total:increase30d{verb=~"LIST|GET"})
          labels:
            verb: read
          record: code:apiserver_request_total:increase30d
        - expr: |
            sum by (code) (code_verb:apiserver_request_total:increase30d{verb=~"POST|PUT|PATCH|DELETE"})
          labels:
            verb: write
          record: code:apiserver_request_total:increase30d
    - name: k8s.rules
      rules:
        - expr: |
            sum by (cluster, namespace, pod, container) (
              rate(container_cpu_usage_seconds_total{job="kubelet", metrics_path="/metrics/cadvisor", image!="", container!="POD"}[5m])
            ) * on (cluster, namespace, pod) group_left(node) topk by (cluster, namespace, pod) (
              1, max by(cluster, namespace, pod, node) (kube_pod_info{node!=""})
            )
          record: node_namespace_pod_container:container_cpu_usage_seconds_total:sum_rate
        - expr: |
            container_memory_working_set_bytes{job="kubelet", metrics_path="/metrics/cadvisor", image!=""}
            * on (namespace, pod) group_left(node) topk by(namespace, pod) (1,
              max by(namespace, pod, node) (kube_pod_info{node!=""})
            )
          record: node_namespace_pod_container:container_memory_working_set_bytes
        - expr: |
            container_memory_rss{job="kubelet", metrics_path="/metrics/cadvisor", image!=""}
            * on (namespace, pod) group_left(node) topk by(namespace, pod) (1,
              max by(namespace, pod, node) (kube_pod_info{node!=""})
            )
          record: node_namespace_pod_container:container_memory_rss
        - expr: |
            container_memory_cache{job="kubelet", metrics_path="/metrics/cadvisor", image!=""}
            * on (namespace, pod) group_left(node) topk by(namespace, pod) (1,
              max by(namespace, pod, node) (kube_pod_info{node!=""})
            )
          record: node_namespace_pod_container:container_memory_cache
        - expr: |
            container_memory_swap{job="kubelet", metrics_path="/metrics/cadvisor", image!=""}
            * on (namespace, pod) group_left(node) topk by(namespace, pod) (1,
              max by(namespace, pod, node) (kube_pod_info{node!=""})
            )
          record: node_namespace_pod_container:container_memory_swap
        - expr: |
            sum by (namespace) (
                sum by (namespace, pod) (
                    max by (namespace, pod, container) (
                        kube_pod_container_resource_requests_memory_bytes{job="kube-state-metrics"}
                    ) * on(namespace, pod) group_left() max by (namespace, pod) (
                        kube_pod_status_phase{phase=~"Pending|Running"} == 1
                    )
                )
            )
          record: namespace:kube_pod_container_resource_requests_memory_bytes:sum
        - expr: |
            sum by (namespace) (
                sum by (namespace, pod) (
                    max by (namespace, pod, container) (
                        kube_pod_container_resource_requests_cpu_cores{job="kube-state-metrics"}
                    ) * on(namespace, pod) group_left() max by (namespace, pod) (
                      kube_pod_status_phase{phase=~"Pending|Running"} == 1
                    )
                )
            )
          record: namespace:kube_pod_container_resource_requests_cpu_cores:sum
        - expr: |
            max by (cluster, namespace, workload, pod) (
              label_replace(
                label_replace(
                  kube_pod_owner{job="kube-state-metrics", owner_kind="ReplicaSet"},
                  "replicaset", "$1", "owner_name", "(.*)"
                ) * on(replicaset, namespace) group_left(owner_name) topk by(replicaset, namespace) (
                  1, max by (replicaset, namespace, owner_name) (
                    kube_replicaset_owner{job="kube-state-metrics"}
                  )
                ),
                "workload", "$1", "owner_name", "(.*)"
              )
            )
          labels:
            workload_type: deployment
          record: namespace_workload_pod:kube_pod_owner:relabel
        - expr: |
            max by (cluster, namespace, workload, pod) (
              label_replace(
                kube_pod_owner{job="kube-state-metrics", owner_kind="DaemonSet"},
                "workload", "$1", "owner_name", "(.*)"
              )
            )
          labels:
            workload_type: daemonset
          record: namespace_workload_pod:kube_pod_owner:relabel
        - expr: |
            max by (cluster, namespace, workload, pod) (
              label_replace(
                kube_pod_owner{job="kube-state-metrics", owner_kind="StatefulSet"},
                "workload", "$1", "owner_name", "(.*)"
              )
            )
          labels:
            workload_type: statefulset
          record: namespace_workload_pod:kube_pod_owner:relabel
    - name: kube-scheduler.rules
      rules:
        - expr: |
            histogram_quantile(0.99, sum(rate(scheduler_e2e_scheduling_duration_seconds_bucket{job="kube-scheduler"}[5m])) without(instance, pod))
          labels:
            quantile: "0.99"
          record: cluster_quantile:scheduler_e2e_scheduling_duration_seconds:histogram_quantile
        - expr: |
            histogram_quantile(0.99, sum(rate(scheduler_scheduling_algorithm_duration_seconds_bucket{job="kube-scheduler"}[5m])) without(instance, pod))
          labels:
            quantile: "0.99"
          record: cluster_quantile:scheduler_scheduling_algorithm_duration_seconds:histogram_quantile
        - expr: |
            histogram_quantile(0.99, sum(rate(scheduler_binding_duration_seconds_bucket{job="kube-scheduler"}[5m])) without(instance, pod))
          labels:
            quantile: "0.99"
          record: cluster_quantile:scheduler_binding_duration_seconds:histogram_quantile
        - expr: |
            histogram_quantile(0.9, sum(rate(scheduler_e2e_scheduling_duration_seconds_bucket{job="kube-scheduler"}[5m])) without(instance, pod))
          labels:
            quantile: "0.9"
          record: cluster_quantile:scheduler_e2e_scheduling_duration_seconds:histogram_quantile
        - expr: |
            histogram_quantile(0.9, sum(rate(scheduler_scheduling_algorithm_duration_seconds_bucket{job="kube-scheduler"}[5m])) without(instance, pod))
          labels:
            quantile: "0.9"
          record: cluster_quantile:scheduler_scheduling_algorithm_duration_seconds:histogram_quantile
        - expr: |
            histogram_quantile(0.9, sum(rate(scheduler_binding_duration_seconds_bucket{job="kube-scheduler"}[5m])) without(instance, pod))
          labels:
            quantile: "0.9"
          record: cluster_quantile:scheduler_binding_duration_seconds:histogram_quantile
        - expr: |
            histogram_quantile(0.5, sum(rate(scheduler_e2e_scheduling_duration_seconds_bucket{job="kube-scheduler"}[5m])) without(instance, pod))
          labels:
            quantile: "0.5"
          record: cluster_quantile:scheduler_e2e_scheduling_duration_seconds:histogram_quantile
        - expr: |
            histogram_quantile(0.5, sum(rate(scheduler_scheduling_algorithm_duration_seconds_bucket{job="kube-scheduler"}[5m])) without(instance, pod))
          labels:
            quantile: "0.5"
          record: cluster_quantile:scheduler_scheduling_algorithm_duration_seconds:histogram_quantile
        - expr: |
            histogram_quantile(0.5, sum(rate(scheduler_binding_duration_seconds_bucket{job="kube-scheduler"}[5m])) without(instance, pod))
          labels:
            quantile: "0.5"
          record: cluster_quantile:scheduler_binding_duration_seconds:histogram_quantile
    - name: node.rules
      rules:
        - expr: |
            topk by(namespace, pod) (1,
              max by (node, namespace, pod) (
                label_replace(kube_pod_info{job="kube-state-metrics",node!=""}, "pod", "$1", "pod", "(.*)")
            ))
          record: 'node_namespace_pod:kube_pod_info:'
        - expr: |
            count by (cluster, node) (sum by (node, cpu) (
              node_cpu_seconds_total{job="node-exporter"}
            * on (namespace, pod) group_left(node)
              node_namespace_pod:kube_pod_info:
            ))
          record: node:node_num_cpu:sum
        - expr: |
            sum(
              node_memory_MemAvailable_bytes{job="node-exporter"} or
              (
                node_memory_Buffers_bytes{job="node-exporter"} +
                node_memory_Cached_bytes{job="node-exporter"} +
                node_memory_MemFree_bytes{job="node-exporter"} +
                node_memory_Slab_bytes{job="node-exporter"}
              )
            ) by (cluster)
          record: :node_memory_MemAvailable_bytes:sum
    - name: kubelet.rules
      rules:
        - expr: |
            histogram_quantile(0.99, sum(rate(kubelet_pleg_relist_duration_seconds_bucket[5m])) by (instance, le) * on(instance) group_left(node) kubelet_node_name{job="kubelet", metrics_path="/metrics"})
          labels:
            quantile: "0.99"
          record: node_quantile:kubelet_pleg_relist_duration_seconds:histogram_quantile
        - expr: |
            histogram_quantile(0.9, sum(rate(kubelet_pleg_relist_duration_seconds_bucket[5m])) by (instance, le) * on(instance) group_left(node) kubelet_node_name{job="kubelet", metrics_path="/metrics"})
          labels:
            quantile: "0.9"
          record: node_quantile:kubelet_pleg_relist_duration_seconds:histogram_quantile
        - expr: |
            histogram_quantile(0.5, sum(rate(kubelet_pleg_relist_duration_seconds_bucket[5m])) by (instance, le) * on(instance) group_left(node) kubelet_node_name{job="kubelet", metrics_path="/metrics"})
          labels:
            quantile: "0.5"
          record: node_quantile:kubelet_pleg_relist_duration_seconds:histogram_quantile
    - name: kube-prometheus-node-recording.rules
      rules:
        - expr: sum(rate(node_cpu_seconds_total{mode!="idle",mode!="iowait",mode!="steal"}[3m]))
            BY (instance)
          record: instance:node_cpu:rate:sum
        - expr: sum(rate(node_network_receive_bytes_total[3m])) BY (instance)
          record: instance:node_network_receive_bytes:rate:sum
        - expr: sum(rate(node_network_transmit_bytes_total[3m])) BY (instance)
          record: instance:node_network_transmit_bytes:rate:sum
        - expr: sum(rate(node_cpu_seconds_total{mode!="idle",mode!="iowait",mode!="steal"}[5m]))
            WITHOUT (cpu, mode) / ON(instance) GROUP_LEFT() count(sum(node_cpu_seconds_total)
            BY (instance, cpu)) BY (instance)
          record: instance:node_cpu:ratio
        - expr: sum(rate(node_cpu_seconds_total{mode!="idle",mode!="iowait",mode!="steal"}[5m]))
          record: cluster:node_cpu:sum_rate5m
        - expr: cluster:node_cpu_seconds_total:rate5m / count(sum(node_cpu_seconds_total)
            BY (instance, cpu))
          record: cluster:node_cpu:ratio
    - name: kube-prometheus-general.rules
      rules:
        - expr: count without(instance, pod, node) (up == 1)
          record: count:up1
        - expr: count without(instance, pod, node) (up == 0)
          record: count:up0
