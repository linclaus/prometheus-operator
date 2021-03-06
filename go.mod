module github.com/linclaus/prometheus-operator

go 1.13

require (
	github.com/go-logr/logr v0.1.0
	github.com/onsi/ginkgo v1.11.0
	github.com/onsi/gomega v1.8.1
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/prometheus/prometheus v0.0.0-20191017095924-6f92ce560538
	github.com/spf13/viper v1.3.2
	gopkg.in/yaml.v2 v2.2.4
	k8s.io/apimachinery v0.17.2
	k8s.io/client-go v0.17.2
	sigs.k8s.io/controller-runtime v0.5.0
)
