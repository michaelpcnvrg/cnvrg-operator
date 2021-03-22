kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: leader-locking-nfs-client-provisioner
  namespace: {{ ns . }}
subjects:
  - kind: ServiceAccount
    name: nfs-client-provisioner
    namespace: {{ ns . }}
roleRef:
  kind: Role
  name: leader-locking-nfs-client-provisioner
  apiGroup: rbac.authorization.k8s.io