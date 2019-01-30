package collector

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/taylor840326/go_pbspro/qstat"
)

const (
	qstatCollectorSubSystem = "qstat"
)

func init() {
	registerCollector(qstatCollectorSubSystem, defaultEnabled, NewQstatCollector)
}

type qstatCollector struct {
	server_state string
}

func (c *qstatCollector) Update(ch chan<- prometheus.Metric) error {
	c.server_state = "haha"
	return nil
}

type qstatMetric struct {
	name            string
	desc            string
	value           float64
	metricType      prometheus.ValueType
	extraLabel      []string
	extraLabelValue string
}

func NewQstatCollector() (Collector, error) {
	qc := new(qstatCollector)
	return &qstatCollector{server_state: qc.server_state}, nil
}

func (c *qstatCollector) updateQstat(ch chan<- prometheus.Metric) {

	var allMetrics []qstatMetric
	var metrics []qstatMetric

	qstat, err := qstat.NewQstat("172.18.7.10")
	if err != nil {
		fmt.Println(err.Error())
	}

	qstat.SetAttribs(nil)
	qstat.SetExtend("")

	err = qstat.ConnectPBS()
	if err != nil {
		fmt.Println("ConnectPBS Error")
	}

	err = qstat.PbsServerState()
	if err != nil {
		fmt.Println(err.Error())
	}

	allMetrics = []qstatMetric{
		{
			name:       "server_name",
			desc:       "pbspro_exporter: server name.",
			value:      qstat.ServerState[0].ServerName,
			metricType: prometheus.UntypedValue,
		},
		{
			name:       "server_state",
			desc:       "pbspro_exporter: server state.",
			value:      float(0)
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_host",
			desc:       "pbspro_exporter: Server Host.",
			value:      float(0)
			metricType: prometheus.GaugeValue,
		},
		{

			name:       "server_scheduling",
			desc:       "pbspro_exporter: Server Scheduling.",
			value:      float(0)
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_total_jobs",
			desc:       "pbspro_exporter: Server Total Jobs.",
			value:      float64(0),
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_transit_state_count",
			desc:       "pbspro_exporter: Server Transit State Count.",
			value:      float64(0),
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_queued_state_count",
			desc:       "pbspro_exporter: Server Queued State Count.",
			value:      float64(0),
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_held_state_count",
			desc:       "pbspro_exporter: Server Held State Count.",
			value:      float64(0),
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_waiting_state_count",
			desc:       "pbspro_exporter: Server Waiting State Count.",
			value:      float64(0),
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_running_state_count",
			desc:       "pbspro_exporter: Server Running State Count.",
			value:      float64(0),
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_exiting_state_count",
			desc:       "pbspro_exporter: Server Exiting State Count.",
			value:      float64(0),
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_begun_state_count",
			desc:       "pbspro_exporter: Server Begun State Count.",
			value:      float64(0),
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_default_queue",
			desc:       "pbspro_exporter: Server Default Queue.",
			value:      float(0)
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_log_events",
			desc:       "pbspro_exporter: Server Log Events.",
			value:      float64(0),
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_mail_from",
			desc:       "pbspro_exporter: Server Mail From.",
			value:      float(0)
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_query_other_jobs",
			desc:       "pbspro_exporter: Server Query Other Jobs.",
			value:      float(0)
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_resources_default_ncpus",
			desc:       "pbspro_exporter: Server Resources Default Ncpus.",
			value:      float64(0),
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_default_chunk_ncpus",
			desc:       "pbspro_exporter: Server Default Chunk Ncpus.",
			value:      float64(0),
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_resources_assigned_ncpus",
			desc:       "pbspro_exporter: Server Resources Assigned Ncpus.",
			value:      float64(0),
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_resources_assigned_nodect",
			desc:       "pbspro_exporter: Server Resources Assigned Nodect.",
			value:      float64(0),
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_scheduler_iteration",
			desc:       "pbspro_exporter: Server Scheudler Iteration.",
			value:      float64(0),
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_flicenses",
			desc:       "pbspro_exporter: Server Flicense.",
			value:      float64(0),
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_resv_enable",
			desc:       "pbspro_exporter: Server Resv Enable.",
			value:      float(0)
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_node_fail_requeue",
			desc:       "pbspro_exporter: Server Node Fail Requeue.",
			value:      float64(0),
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_max_array_size",
			desc:       "pbspro_exporter: Server Max Array Size.",
			value:      float64(0),
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_pbs_license_min",
			desc:       "pbspro_exporter: Server PBS License Min.",
			value:      float64(0),
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_pbs_license_max",
			desc:       "pbspro_exporter: Server PBS License Max.",
			value:      float64(0),
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_pbs_license_linger_time",
			desc:       "pbspro_exporter: Server PBS License Linger Time.",
			value:      float64(0),
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_license_count_avail_global",
			desc:       "pbspro_exporter: Server License Count Avail Global.",
			value:      float64(0),
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_license_used",
			desc:       "pbspro_exporter: Server License Used.",
			value:      float64(0),
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_license_count_high_use",
			desc:       "pbspro_exporter: Server License Count High Use.",
			value:      float64(0),
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_pbs_version",
			desc:       "pbspro_exporter: Server PBS Version.",
			value:      float(0)
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_eligible_time_enable",
			desc:       "pbspro_exporter: Server Eligible Time Enable.",
			value:      float(0)
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_job_history_enable",
			desc:       "pbspro_exporter: Server Job History Enable.",
			value:      float(0)
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_job_history_duration",
			desc:       "pbspro_exporter: Server Job History Duration.",
			value:      float(0)
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_max_concurrent_provision",
			desc:       "pbspro_exporter: Server Max Concurrent Provision.",
			value:      float64(0),
			metricType: prometheus.GaugeValue,
		},
		{
			name:       "server_power_provisioning",
			desc:       "pbspro_exporter: Server Power Provisioning.",
			value:      float(0)
			metricType: prometheus.GaugeValue,
		},
	}

	err = qstat.DisconnectPBS()
	if err != nil {
		fmt.Println("DisconnectPBS Error")
	}

}
