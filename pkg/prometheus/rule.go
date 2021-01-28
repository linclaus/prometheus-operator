package prometheus

import (
	"bufio"
	"fmt"
	"os"

	prometheusv1 "github.com/linclaus/prometheus-operator/api/v1"
	"github.com/prometheus/prometheus/pkg/rulefmt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

func GenerateRuleAndWriteFile(r prometheusv1.PrometheusRule) error {
	rule := rulefmt.Rule{
		Alert:       r.Spec.Alert,
		Expr:        r.Spec.Expr,
		Annotations: r.Spec.Annotations,
		Labels:      r.Spec.Labels,
	}
	fileName := r.Spec.Alert + viper.GetString("ruleFileSuffix")
	return addRulesToGroups(fileName, rule)
}

func DeleteRule(r prometheusv1.PrometheusRule) error {
	return deleteRuleFile(r.Spec.Alert + viper.GetString("ruleFileSuffix"))
}
func deleteRuleFile(fileName string) error {
	return os.Remove(viper.GetString("ruleFilePath") + fileName)
}

func addRulesToGroups(fileName string, rule rulefmt.Rule) error {
	rgs, err := rulefmt.ParseFile(viper.GetString("ruleFilePath") + fileName)
	if err != nil {
		fmt.Printf("file doesn't exist, %s will be created\n", fileName)
		rgs = &rulefmt.RuleGroups{
			Groups: []rulefmt.RuleGroup{
				{
					Name:  fileName,
					Rules: []rulefmt.Rule{},
				},
			},
		}
	}

	rgs.Groups[0].Rules = mergeOrAddRule(rgs.Groups[0].Rules, rule)
	return writeRuleToFile(*rgs, fileName)
}

func mergeOrAddRule(rules []rulefmt.Rule, rule rulefmt.Rule) []rulefmt.Rule {
	for _, r := range rules {
		if r.Alert == rule.Alert {
			return rules
		}
	}
	return append(rules, rule)
}

func writeRuleToFile(rgs rulefmt.RuleGroups, fileName string) error {
	res, _ := yaml.Marshal(rgs)
	file, err := os.OpenFile(viper.GetString("ruleFilePath")+fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("文件打开失败", err)
		return err
	}
	write := bufio.NewWriter(file)
	write.WriteString(string(res))
	write.Flush()
	// reloadPrometheus()
	return nil
}
