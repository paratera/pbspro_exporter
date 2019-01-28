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
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_host"),
		"pbspro_exporter: Server Host.",
		[]string{"collector"},
		nil,
	)
	// scheduling
	pbsproServerScheduling = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_scheduling"),
		"pbspro_exporter: Server Scheduling.",
		[]string{"collector"},
		nil,
	)
	// total_jobs
	pbsproServerTotalJobs = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_total_jobs"),
		"pbspro_exporter: Server Total Jobs.",
		[]string{"collector"},
		nil,
	)
	// transit_state_count
	pbsproServerTransitStateCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_transit_state_count"),
		"pbspro_exporter: Server Transit State Count.",
		[]string{"collector"},
		nil,
	)
	// queued_state_count
	pbsproServerQueuedStateCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_queued_state_count"),
		"pbspro_exporter: Server Queued State Count.",
		[]string{"collector"},
		nil,
	)
	// held_state_count
	pbsproServerHeldStateCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_held_state_count"),
		"pbspro_exporter: Server Held State Count.",
		[]string{"collector"},
		nil,
	)
	// waiting_state_count
	pbsproServerWaitingStateCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_waiting_state_count"),
		"pbspro_exporter: Server Waiting State Count.",
		[]string{"collector"},
		nil,
	)
	// running_state_count
	pbsproServerRunningStateCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_running_state_count"),
		"pbspro_exporter: Server Running State Count.",
		[]string{"collector"},
		nil,
	)
	// exiting_state_count
	pbsproServerExitingStateCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_exiting_state_count"),
		"pbspro_exporter: Server Exiting State Count.",
		[]string{"collector"},
		nil,
	)
	// begun_state_count
	pbsproServerBegunStateCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_begun_state_count"),
		"pbspro_exporter: Server Begun State Count.",
		[]string{"collector"},
		nil,
	)
	// default_queue
	pbsproServerDefaultQueue = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_default_queue"),
		"pbspro_exporter: Server Default Queue.",
		[]string{"collector"},
		nil,
	)
	// log_events
	pbsproServerLogEvents = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_log_events"),
		"pbspro_exporter: Server Log Events.",
		[]string{"collector"},
		nil,
	)
	// mail_from
	pbsproServerMailFrom = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_mail_from"),
		"pbspro_exporter: Server Mail From.",
		[]string{"collector"},
		nil,
	)
	// query_other_jobs
	pbsproServerQueryOtherJobs = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_query_other_jobs"),
		"pbspro_exporter: Server Query Other Jobs.",
		[]string{"collector"},
		nil,
	)
	// resources_default_ncpus
	pbsproServerResourcesDefaultNcpus = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_resources_default_ncpus"),
		"pbspro_exporter: Server Resources Default Ncpus.",
		[]string{"collector"},
		nil,
	)
	// default_chunk_ncpus
	pbsproServerDefaultChunkNcpus = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_default_chunk_ncpus"),
		"pbspro_exporter: Server Default Chunk Ncpus.",
		[]string{"collector"},
		nil,
	)
	// resources_assigned_ncpus
	pbsproServerResourcesAssignedNcpus = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_resources_assigned_ncpus"),
		"pbspro_exporter: Server Resources Assigned Ncpus.",
		[]string{"collector"},
		nil,
	)
	// resources_assigned_nodect
	pbsproServerResourcesAssignedNodect = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_resources_assigned_nodect"),
		"pbspro_exporter: Server Resources Assigned Nodect.",
		[]string{"collector"},
		nil,
	)
	// scheduler_iteration
	pbsproServerSchedulerIteration = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_scheduler_iteration"),
		"pbspro_exporter: Server Scheudler Iteration.",
		[]string{"collector"},
		nil,
	)
	// flicenses
	pbsproServerFlicenses = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_flicenses"),
		"pbspro_exporter: Server Flicense.",
		[]string{"collector"},
		nil,
	)
	//resv_enable
	pbsproServerResvEnable = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_resv_enable"),
		"pbspro_exporter: Server Resv Enable.",
		[]string{"collector"},
		nil,
	)
	// node_fail_requeue
	pbsproServerNodeFailRequeue = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_node_fail_requeue"),
		"pbspro_exporter: Server Node Fail Requeue.",
		[]string{"collector"},
		nil,
	)
	// max_array_size
	pbsproServerMaxArraySize = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_max_array_size"),
		"pbspro_exporter: Server Max Array Size.",
		[]string{"collector"},
		nil,
	)
	// pbs_license_min
	pbsproServerPbsLicenseMin = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_pbs_license_min"),
		"pbspro_exporter: Server PBS License Min.",
		[]string{"collector"},
		nil,
	)
	// pbs_license_max
	pbsproServerPbsLicenseMax = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_pbs_license_max"),
		"pbspro_exporter: Server PBS License Max.",
		[]string{"collector"},
		nil,
	)
	// pbs_license_linger_time
	pbsproServerPbsLicenseLingerTime = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_pbs_license_linger_time"),
		"pbspro_exporter: Server PBS License Linger Time.",
		[]string{"collector"},
		nil,
	)
	// license_count_avail_global
	pbsproServerLicenseCountAvailGlobal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_license_count_avail_global"),
		"pbspro_exporter: Server License Count Avail Global.",
		[]string{"collector"},
		nil,
	)
	// license_count_used
	pbsproServerLicenseCountUsed = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_license_used"),
		"pbspro_exporter: Server License Used.",
		[]string{"collector"},
		nil,
	)
	// license_count_high_use
	pbsproServerLicenseCountHighUse = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_license_count_high_use"),
		"pbspro_exporter: Server License Count High Use.",
		[]string{"collector"},
		nil,
	)
	// pbs_version
	pbsproServerPbsVersion = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_pbs_version"),
		"pbspro_exporter: Server PBS Version.",
		[]string{"collector"},
		nil,
	)
	// eligible_time_enable
	pbsproServerEligibleTimeEnable = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_eligible_time_enable"),
		"pbspro_exporter: Server Eligible Time Enable.",
		[]string{"collector"},
		nil,
	)
	// job_history_enable
	pbsproServerJobHistoryEnable = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_job_history_enable"),
		"pbspro_exporter: Server Job History Enable.",
		[]string{"collector"},
		nil,
	)
	// job_history_duration
	pbsproServerJobHistoryDuration = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_job_history_duration"),
		"pbspro_exporter: Server Job History Duration.",
		[]string{"collector"},
		nil,
	)
	// max_concurrent_provision
	pbsproServerMaxConcurrentProvision = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, qstatCollectorSubSystem, "server_max_concurrent_provision"),
		"pbspro_exporter: Server Max Concurrent Provision.",
		[]string{"collector"},
		nil,
	)
)
