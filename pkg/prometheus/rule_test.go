package prometheus

import (
	"testing"

	prometheusv1 "github.com/linclaus/prometheus-operator/api/v1"
	"github.com/spf13/viper"
)

func TestRule(t *testing.T) {
	viper.SetDefault("ruleFilePath", "./")
	viper.SetDefault("ruleFileSuffix", "-rule.yaml")
	r := prometheusv1.PrometheusRule{
		Spec: prometheusv1.PrometheusRuleSpec{
			Alert: "alert-test",
			Expr:  "vector(1)",
			For:   "5m",
			Labels: map[string]string{
				"key1": "value1",
			},
			Annotations: map[string]string{
				"key1": "value1",
			},
		},
	}
	GenerateRuleAndWriteFile(r)
}
