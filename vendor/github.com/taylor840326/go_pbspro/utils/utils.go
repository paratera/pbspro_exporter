package utils

/*
#cgo CFLAGS: -g
#cgo LDFLAGS: -L/opt/pbspro/lib -lpbs
#include <stdlib.h>
#include "/opt/pbspro/include/pbs_error.h"
#include "/opt/pbspro/include/pbs_ifl.h"
*/
import "C"
import (
	"unsafe"

	"github.com/juju/errors"
)

// Manner defines how the server should be terminated
type Manner int

// Hold defines the type of job hold to place on a job
type Hold string

// MessageStream which output stream should be written to
type MessageStream int

// Operator defines types of logical comparator
type Operator int

type Command int

type ObjectType int

const (
	SHUT_IMMEDIATE Manner        = C.SHUT_IMMEDIATE
	SHUT_DELAY     Manner        = C.SHUT_DELAY
	USER_HOLD      Hold          = C.USER_HOLD
	OTHER_HOLD     Hold          = C.OTHER_HOLD
	SYSTEM_HOLD    Hold          = C.SYSTEM_HOLD
	MSG_ERR        MessageStream = C.MSG_ERR
	MSG_OUT        MessageStream = C.MSG_OUT
	MGR_CMD_CREATE Command       = C.MGR_CMD_CREATE
	MGR_CMD_DELETE Command       = C.MGR_CMD_DELETE
	MGR_CMD_SET    Command       = C.MGR_CMD_SET
	MGR_CMD_UNSET  Command       = C.MGR_CMD_UNSET
	MGR_CMD_LIST   Command       = C.MGR_CMD_LIST
	MGR_CMD_PRINT  Command       = C.MGR_CMD_PRINT
	MGR_CMD_ACTIVE Command       = C.MGR_CMD_ACTIVE
	MGR_OBJ_NONE   ObjectType    = C.MGR_OBJ_NONE
	MGR_OBJ_SERVER ObjectType    = C.MGR_OBJ_SERVER
	MGR_OBJ_QUEUE  ObjectType    = C.MGR_OBJ_QUEUE
	MGR_OBJ_JOB    ObjectType    = C.MGR_OBJ_JOB
	MGR_OBJ_NODE   ObjectType    = C.MGR_OBJ_NODE
	ATTR_a         string        = C.ATTR_a
	ATTR_c         string        = C.ATTR_c
	ATTR_e         string        = C.ATTR_e
	//ATTR_f                         string        = C.ATTR_f
	ATTR_g string = C.ATTR_g
	ATTR_h string = C.ATTR_h
	ATTR_j string = C.ATTR_j
	ATTR_k string = C.ATTR_k
	ATTR_l string = C.ATTR_l
	ATTR_m string = C.ATTR_m
	ATTR_o string = C.ATTR_o
	ATTR_p string = C.ATTR_p
	ATTR_q string = C.ATTR_q
	ATTR_r string = C.ATTR_r
	//ATTR_t                         string        = C.ATTR_t
	ATTR_array_id string = C.ATTR_array_id
	ATTR_u        string = C.ATTR_u
	ATTR_v        string = C.ATTR_v
	ATTR_A        string = C.ATTR_A
	//ATTR_args                      string        = C.ATTR_args
	ATTR_M        string = C.ATTR_M
	ATTR_N        string = C.ATTR_N
	ATTR_S        string = C.ATTR_S
	ATTR_depend   string = C.ATTR_depend
	ATTR_inter    string = C.ATTR_inter
	ATTR_stagein  string = C.ATTR_stagein
	ATTR_stageout string = C.ATTR_stageout
	//ATTR_jobtype                   string        = C.ATTR_jobtype
	//ATTR_submit_host               string        = C.ATTR_submit_host
	//ATTR_init_work_dir             string        = C.ATTR_init_work_dir
	ATTR_ctime    string = C.ATTR_ctime
	ATTR_exechost string = C.ATTR_exechost
	//ATTR_execport                  string        = C.ATTR_execport
	ATTR_mtime      string = C.ATTR_mtime
	ATTR_qtime      string = C.ATTR_qtime
	ATTR_session    string = C.ATTR_session
	ATTR_euser      string = C.ATTR_euser
	ATTR_egroup     string = C.ATTR_egroup
	ATTR_hashname   string = C.ATTR_hashname
	ATTR_hopcount   string = C.ATTR_hopcount
	ATTR_security   string = C.ATTR_security
	ATTR_sched_hint string = C.ATTR_sched_hint
	ATTR_substate   string = C.ATTR_substate
	ATTR_name       string = C.ATTR_name
	ATTR_owner      string = C.ATTR_owner
	ATTR_used       string = C.ATTR_used
	ATTR_state      string = C.ATTR_state
	ATTR_queue      string = C.ATTR_queue
	ATTR_server     string = C.ATTR_server
	ATTR_maxrun     string = C.ATTR_maxrun
	//ATTR_maxreport                 string        = C.ATTR_maxreport
	ATTR_total   string = C.ATTR_total
	ATTR_comment string = C.ATTR_comment
	ATTR_cookie  string = C.ATTR_cookie
	ATTR_qrank   string = C.ATTR_qrank
	ATTR_altid   string = C.ATTR_altid
	ATTR_etime   string = C.ATTR_etime
	//ATTR_exitstat                  string        = C.ATTR_exitstat
	//ATTR_forwardx11                string        = C.ATTR_forwardx11
	//ATTR_submit_args               string        = C.ATTR_submit_args
	//ATTR_tokens                    string        = C.ATTR_tokens
	//ATTR_netcounter                string        = C.ATTR_netcounter
	ATTR_umask string = C.ATTR_umask
	//ATTR_start_time                string        = C.ATTR_start_time
	//ATTR_start_count               string        = C.ATTR_start_count
	//ATTR_checkpoint_dir            string        = C.ATTR_checkpoint_dir
	//ATTR_checkpoint_name           string        = C.ATTR_checkpoint_name
	//ATTR_checkpoint_time           string        = C.ATTR_checkpoint_time
	//ATTR_checkpoint_restart_status string        = C.ATTR_checkpoint_restart_status
	//ATTR_restart_name              string        = C.ATTR_restart_name
	//ATTR_comp_time                 string        = C.ATTR_comp_time
	//ATTR_reported                  string        = C.ATTR_reported
	//ATTR_intcmd                    string        = C.ATTR_intcmd
	//ATTR_P                         string        = C.ATTR_P
	//ATTR_node_exclusive            string        = C.ATTR_node_exclusive
	//ATTR_exec_gpus                 string        = C.ATTR_exec_gpus
	ATTR_J string   = C.ATTR_J
	SET    Operator = C.SET
	UNSET  Operator = C.UNSET
	INCR   Operator = C.INCR
	DECR   Operator = C.DECR
	EQ     Operator = C.EQ
	NE     Operator = C.NE
	GE     Operator = C.GE
	GT     Operator = C.GT
	LE     Operator = C.LE
	LT     Operator = C.LT
	DFLT   Operator = C.DFLT
	//MERGE                          Operator      = C.MERGE
	//INCR_OLD                       Operator      = C.INCR_OLD
)

// Attrib represents the attrl and attropl structures
type (
	Attrib struct {
		Name     string
		Resource string
		Value    string
		Op       Operator
	}

	// BatchStatus represents the batch_status structure
	BatchStatus struct {
		Name       string
		Text       string
		Attributes []Attrib
	}
)

// Pbs_default reports the default torque server
func Pbs_default() string {
	// char* from pbs_default is statically allocated, so can't be freed
	return C.GoString(C.pbs_default())
}

// Pbs_connect makes a connection to server, or if server is an empty string, the default server. The returned handle is used by subsequent calls to the functions in this package to identify the server.
func Pbs_connect(server string) (int, error) {
	str := C.CString(server)
	defer C.free(unsafe.Pointer(str))

	handle := C.pbs_connect(str)
	if handle < 0 {
		return 0, errors.New(Pbs_strerror(int(C.pbs_errno)))
	}

	return int(handle), nil
}

func Pbs_disconnect(handle int) error {
	ret := C.pbs_disconnect(C.int(handle))
	if ret != 0 {
		return errors.New(Pbs_strerror(int(C.pbs_errno)))
	}
	return nil
}

func getLastError() error {
	return errors.New(Pbs_strerror(int(C.pbs_errno)))
}

func Pbs_geterrmsg(handle int) string {
	s := C.pbs_geterrmsg(C.int(handle))
	// char* from pbs_geterrmsg is statically allocated, so can't be freed
	return C.GoString(s)
}

func Pbs_strerror(errno int) string {
	// char* from pbs_strerror is statically allocated, so can't be freed
	return C.GoString(C.pbse_to_txt(C.int(errno)))
}
