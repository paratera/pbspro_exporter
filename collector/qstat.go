package collector

import "github.com/prometheus/client_golang/prometheus"

const (
	qstatCollectorSubSystem = "qstat"
)

var (
	// server_state
	pbsproServerState = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_state"),
		"pbspro_exporter: server state.",
		[]string{"collector"},
		nil,
	)
	// server_host
	pbsproServerHost = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_state"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// scheduling
	pbsproServerScheduling = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_scheduling"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// total_jobs
	pbsproServerTotalJobs = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_total_jobs"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// transit_state_count
	pbsproServerTransitStateCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_transit_state_count"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// queued_state_count
	pbsproServerQueuedStateCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_queued_state_count"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// held_state_count
	pbsproServerHeldStateCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_held_state_count"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// waiting_state_count
	pbsproServerWaitingStateCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_waiting_state_count"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// running_state_count
	pbsproServerRunningStateCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_running_state_count"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// exiting_state_count
	pbsproServerExitingStateCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_exiting_state_count"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// begun_state_count
	pbsproServerBegunStateCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_begun_state_count"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// default_queue
	pbsproServerDefaultQueue = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_default_queue"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// log_events
	pbsproServerLogEvents = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_log_events"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// mail_from
	pbsproServerMailFrom = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_mail_from"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// query_other_jobs
	pbsproServerQueryOtherJobs = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_query_other_jobs"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// resources_default_ncpus
	pbsproServerResourcesDefaultNcpus = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_resources_default_ncpus"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// default_chunk_ncpus
	pbsproServerDefaultChunkNcpus = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_default_chunk_ncpus"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// resources_assigned_ncpus
	pbsproServerResourcesAssignedNcpus = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_resources_assigned_ncpus"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// resources_assigned_nodect
	pbsproServerResourcesAssignedNodect = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_resources_assigned_nodect"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// scheduler_iteration
	pbsproServerSchedulerIteration = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_scheduler_iteration"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// flicenses
	pbsproServerFlicenses = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_flicenses"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	//resv_enable
	pbsproServerResvEnable = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_resv_enable"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// node_fail_requeue
	pbsproServerNodeFailRequeue = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_node_fail_requeue"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// max_array_size
	pbsproServerMaxArraySize = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_max_array_size"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// pbs_license_min
	pbsproServerPbsLicenseMin = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_pbs_license_min"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// pbs_license_max
	pbsproServerPbsLicenseMax = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_pbs_license_max"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// pbs_license_linger_time
	pbsproServerPbsLicenseLingerTime = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_pbs_license_linger_time"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// license_count_avail_global
	pbsproServerLicenseCountAvailGlobal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_license_count_avail_global"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// license_count_used
	pbsproServerLicenseCountUsed = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_license_used"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// license_count_high_use
	pbsproServerLicenseCountHighUse = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_license_count_high_use"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// pbs_version
	pbsproServerPbsVersion = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_pbs_version"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// eligible_time_enable
	pbsproServerEligibleTimeEnable = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_eligible_time_enable"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// job_history_enable
	pbsproServerJobHistoryEnable = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_job_history_enable"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// job_history_duration
	pbsproServerJobHistoryDuration = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_job_history_duration"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	// max_concurrent_provision
	pbsproServerMaxConcurrentProvision = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_max_concurrent_provision"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
)
