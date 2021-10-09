package filters

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/context"

	"github.com/prometheus/client_golang/prometheus"
)

// 定义监控指标
var (
	// 总的请求次数是一个 counter 类型不带 label
	totalRequest = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "cmdb_Request_total",
		Help: "cmdb 总请求次数检测",
	})

	// 每个 url 的请求次数也是一个 counter 类型而且是一个带可变 label
	urlRequest = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name:        "cmdb_urlRquest_total",
		Help:        "cmdb 请求 url 总次数检测",
		ConstLabels: prometheus.Labels{},
	}, []string{"url"})

	// 状态码统计 counter ，带可变 label
	statusCode = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cmdb_statusCode_total",
		Help: "cmdb 请求状态码检测",
	}, []string{"code"})

	// 每个 url 请求时间 histogram 带可变 label
	elapsedTime = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "cmdb_elapsedTime",
		Help: "cmdb url 请求消耗时间",
	}, []string{"url"})
)

// 定义 BeferExec 过滤器，因为在 controller 执行的时候还会执行该过滤器
func BeferExec(ctx *context.Context) {
	// 如果执行 controller ，totalRequest 就 inc ，inc 是累加
	totalRequest.Inc()

	// 更新指标采样值,必须要设置 WithLabelValues 值，通过 ctx.Input.URL() 拿到用户请求的 url 并作累加
	urlRequest.WithLabelValues(ctx.Input.URL()).Inc()

	// 设置开始时间
	ctx.Input.SetData("stime", time.Now())
}

// 定义 afterexec 过滤器
func AfterExec(ctx *context.Context) {
	// statusCode 肯定在 afterexec 过滤器执行之后拿到
	statusCode.WithLabelValues(strconv.Itoa(ctx.ResponseWriter.Status)).Inc()

	// 设置结束时间
	stime := ctx.Input.GetData("stime")

	// 这里通过 stime 来判断是否等于 nil，如果不等于 nil 类型断言成时间类型
	if stime != nil {
		if t, ok := stime.(time.Time); ok {

			// 然后用程序执行的总时间减掉程序开始执行的时间，得到程序消耗时间 elapsed
			elapsed := time.Now().Sub(t)

			// 将拿到的时间通过 Observe 放到 WithLabelValues 中
			elapsedTime.WithLabelValues(ctx.Input.URL()).Observe((float64(elapsed)))
		}
	}
}

// 注册 Prometheus 监控
func init() {
	prometheus.MustRegister(totalRequest, urlRequest, statusCode, elapsedTime)
}
