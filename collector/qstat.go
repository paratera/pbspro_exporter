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
	c.updateQstat(ch)
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
	//var metrics []qstatMetric
	var labels []string{}

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
	defer qstat.DisconnectPBS()

	err = qstat.PbsServerState()
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, ss := range qstat.ServerState {
		allMetrics = []qstatMetric{
			{
				name:       "server_state",
				desc:       "pbspro_exporter: server state.",
				value:      float64(ss.ServerState),
				metricType: prometheus.GaugeValue,
			},
			{

				name:       "server_scheduling",
				desc:       "pbspro_exporter: Server Scheduling.",
				value:      float64(ss.ServerScheduling),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_total_jobs",
				desc:       "pbspro_exporter: Server Total Jobs.",
				value:      float64(ss.TotalJobs),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_transit_state_count",
				desc:       "pbspro_exporter: Server Transit State Count.",
				value:      float64(ss.StateCountTransit),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_queued_state_count",
				desc:       "pbspro_exporter: Server Queued State Count.",
				value:      float64(ss.StateCountQueued),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_held_state_count",
				desc:       "pbspro_exporter: Server Held State Count.",
				value:      float64(ss.StateCountHeld),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_waiting_state_count",
				desc:       "pbspro_exporter: Server Waiting State Count.",
				value:      float64(ss.StateCountWaiting),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_running_state_count",
				desc:       "pbspro_exporter: Server Running State Count.",
				value:      float64(ss.StateCountRunning),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_exiting_state_count",
				desc:       "pbspro_exporter: Server Exiting State Count.",
				value:      float64(ss.StateCountExiting),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_begun_state_count",
				desc:       "pbspro_exporter: Server Begun State Count.",
				value:      float64(ss.StateCountBegun),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_log_events",
				desc:       "pbspro_exporter: Server Log Events.",
				value:      float64(ss.LogEvents),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_query_other_jobs",
				desc:       "pbspro_exporter: Server Query Other Jobs.",
				value:      float64(ss.QueryOtherJobs),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_resources_default_ncpus",
				desc:       "pbspro_exporter: Server Resources Default Ncpus.",
				value:      float64(ss.ResourcesDefaultNcpus),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_default_chunk_ncpus",
				desc:       "pbspro_exporter: Server Default Chunk Ncpus.",
				value:      float64(ss.DefaultChunkNcpus),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_resources_assigned_ncpus",
				desc:       "pbspro_exporter: Server Resources Assigned Ncpus.",
				value:      float64(ss.ResourcesAssignedNcpus),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_resources_assigned_nodect",
				desc:       "pbspro_exporter: Server Resources Assigned Nodect.",
				value:      float64(ss.ResourcesAssignedNodect),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_scheduler_iteration",
				desc:       "pbspro_exporter: Server Scheudler Iteration.",
				value:      float64(ss.SchedulerIteration),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_flicenses",
				desc:       "pbspro_exporter: Server Flicense.",
				value:      float64(ss.Flicense),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_resv_enable",
				desc:       "pbspro_exporter: Server Resv Enable.",
				value:      float64(ss.ResvEnable),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_node_fail_requeue",
				desc:       "pbspro_exporter: Server Node Fail Requeue.",
				value:      float64(ss.NodeFailRequeue),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_max_array_size",
				desc:       "pbspro_exporter: Server Max Array Size.",
				value:      float64(ss.MaxArraySize),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_pbs_license_min",
				desc:       "pbspro_exporter: Server PBS License Min.",
				value:      float64(ss.PBSLicenseMin),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_pbs_license_max",
				desc:       "pbspro_exporter: Server PBS License Max.",
				value:      float64(ss.PBSLicenseMax),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_pbs_license_linger_time",
				desc:       "pbspro_exporter: Server PBS License Linger Time.",
				value:      float64(ss.PBSLicenseLingerTime),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_license_count_avail_global",
				desc:       "pbspro_exporter: Server License Count Avail Global.",
				value:      float64(ss.LicenseCountAvailGlobal),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_license_count_avail_local",
				desc:       "pbspro_exporter: Server License Count Avail Global.",
				value:      float64(ss.LicenseCountAvailLocal),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_license_count_used",
				desc:       "pbspro_exporter: Server License Used.",
				value:      float64(ss.LicenseCountUsed),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_license_count_high_use",
				desc:       "pbspro_exporter: Server License Count High Use.",
				value:      float64(ss.LicenseCountHighUse),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_eligible_time_enable",
				desc:       "pbspro_exporter: Server Eligible Time Enable.",
				value:      float64(ss.EligibleTimeEnable),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_job_history_enable",
				desc:       "pbspro_exporter: Server Job History Enable.",
				value:      float64(ss.JobHistoryEnable),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_job_history_duration",
				desc:       "pbspro_exporter: Server Job History Duration.",
				value:      float64(ss.JobHistoryDuration),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_max_concurrent_provision",
				desc:       "pbspro_exporter: Server Max Concurrent Provision.",
				value:      float64(ss.MaxConcurrentProvision),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_power_provisioning",
				desc:       "pbspro_exporter: Server Power Provisioning.",
				value:      float64(ss.PowerProvisioning),
				metricType: prometheus.GaugeValue,
			},
		}
		labels = []string{qstat.ServerName,qstat.ServerHost,qstat.DefaultQueue,qstat.MailFrom,qstat.PBSVersion}
	}

	for _, m := range allMetrics {

		desc := prometheus.NewDesc(
			prometheus.BuildFQName(namespace, qstatCollectorSubSystem, m.name),
			m.desc,
			nil,
			nil,
		)

		ch <- prometheus.MustNewConstMetric(
			desc,
			m.metricType,
			m.value,
			labels,
		)
	}

}
