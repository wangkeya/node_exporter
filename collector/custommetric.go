package collector

import (
	"fmt"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	// 定义自定义数据指标的子系统名称
	customMetricsSubsystem = "metrics"
)

// 定义 customMetricsCollector 结构体
type customMetricsCollector struct {
	logger log.Logger
	//...
}

func init() {
	// 在该函数中调用 registerCollector() 函数,注册自定义 customMetricsCollector
	registerCollector("custom_metrics", defaultEnabled, NewCustomMetricsCollector)
}

// 定义 customMetricsCollector 的工厂函数,后续传入 registerCollector() 函数中,以便创建 customMetricsCollector 对象
func NewCustomMetricsCollector(logger log.Logger) (Collector, error) {
	return &customMetricsCollector{
		logger: logger,
	}, nil
}

// 实现 Update() 函数,以便在处理请求时被 Collector.Collect() 调用
func (c *customMetricsCollector) Update(ch chan<- prometheus.Metric) error {
	var metricType prometheus.ValueType
	var value = 1.1
	metricType = prometheus.CounterValue
	level.Debug(c.logger).Log("msg", "Set custom_metrics", "metrics", value)

	// 通过 `prometheus.MustNewConstMetric` 创建自定义 `prometheus.Metric` 接口对象 `prometheus.constMetric`,并将其传入`prometheus.Metric` 管道
	ch <- prometheus.MustNewConstMetric(
		// 需要传入 Metric 实现对象的描述信息,对象数据类型,值
		prometheus.NewDesc(
			// 描述信息包括 数据指标名称(由 `BuildFQName()`函数组合而成),帮助信息,变量标签,常量标签
			prometheus.BuildFQName(namespace, customMetricsSubsystem, "custom_metrics"),
			fmt.Sprintf("Custom metrics field %s.", "custom_metrics"),
			[]string{"host"}, map[string]string{"ip": "127.0.0.1"},
		),
		metricType, value, "localhost",
	)

	return nil
}
