package mpipe

import (
	"bytes"
	"fmt"
	"testing"
)

var telegraflp_data = []string{
	`cpu,host=server\ 01,region=uswest value=1,msg="all systems, abc nominal" 1434055562000010000`,
	`cpu,host=server\ 01,region=us\,west value_int=1i 1434055562000010000`,
	`cpu,host=server01,region=uswest value=1 1434055562000000000`,
	`cpu,host=server02,region=uswest value=3 1434055562000010000`,
	`temperature,machine=unit42,type=assembly internal=32,external=100 1434055562000000035`,
	`temperature,machine=unit143,type=assembly internal=22,external=130 1434055562005000035`,
}

var telegraflp_expectedResult = []string{
	`{"@timestamp":"2017-04-20T10:07:10+09:00","t":{"database":"db0","host":"Ameba01gm","port":"6381","replication_role":"master","server":"localhost"},"m":{"redis_keyspace":{"avg_ttl":0,"expires":0,"keys":1}}}`,
	`{"@timestamp":"2017-04-20T10:07:10+09:00","t":{"host":"Ameba01gm","port":"6381","replication_role":"master","server":"localhost"},"m":{"redis":{"aof_current_rewrite_time_sec":-1,"aof_enabled":0,"aof_last_bgrewrite_status":"ok","aof_last_rewrite_time_sec":-1,"aof_last_write_status":"ok","aof_rewrite_in_progress":0,"aof_rewrite_scheduled":0,"blocked_clients":0,"client_biggest_input_buf":5,"client_longest_output_list":0,"clients":5,"cluster_enabled":0,"connected_slaves":1,"evicted_keys":0,"expired_keys":0,"instantaneous_input_kbps":0.24,"instantaneous_ops_per_sec":4,"instantaneous_output_kbps":0.55,"keyspace_hitrate":0,"keyspace_hits":0,"keyspace_misses":0,"latest_fork_usec":184,"loading":0,"lru_clock":16255421,"master_repl_offset":1972551115,"maxmemory":3000000000,"maxmemory_policy":"allkeys-lru","mem_fragmentation_ratio":1.82,"migrate_cached_sockets":0,"pubsub_channels":1,"pubsub_patterns":0,"rdb_bgsave_in_progress":0,"rdb_changes_since_last_save":0,"rdb_current_bgsave_time_sec":-1,"rdb_last_bgsave_status":"ok","rdb_last_bgsave_time_sec":0,"rdb_last_save_time":1484098806,"rdb_last_save_time_elapsed":8551624,"rejected_connections":0,"repl_backlog_active":1,"repl_backlog_first_byte_offset":1971502540,"repl_backlog_histlen":1048576,"repl_backlog_size":1048576,"slave0":"ip=172.30.218.140,port=6381,state=online,offset=1972550827,lag=1","sync_full":1,"sync_partial_err":0,"sync_partial_ok":0,"total_commands_processed":58056337,"total_connections_received":839126,"total_net_input_bytes":2973005006,"total_net_output_bytes":14435284764,"total_system_memory":33626107904,"uptime":13976225,"used_cpu_sys":5793.73,"used_cpu_sys_children":0,"used_cpu_user":3561.78,"used_cpu_user_children":0.01,"used_memory":2304312,"used_memory_lua":37888,"used_memory_peak":2617400,"used_memory_rss":4186112}}}`,
	`{"@timestamp":"2017-04-20T10:07:10+09:00","t":{"database":"db0","host":"Ameba01gm","port":"8001","replication_role":"master","server":"localhost"},"m":{"redis_keyspace":{"avg_ttl":0,"expires":0,"keys":20398223}}}`,
	`{"@timestamp":"2017-04-20T10:07:10+09:00","t":{"database":"db0","host":"Ameba01gm","port":"7000","replication_role":"slave","server":"localhost"},"m":{"redis_keyspace":{"avg_ttl":1127630700,"expires":11,"keys":3170}}}`,
	`{"@timestamp":"2017-04-20T10:07:10+09:00","t":{"database":"db0","host":"Ameba01gm","port":"7001","replication_role":"master","server":"localhost"},"m":{"redis_keyspace":{"avg_ttl":712715018,"expires":10,"keys":3102}}}`,
	`{"@timestamp":"2017-04-20T10:07:10+09:00","t":{"database":"db0","host":"Ameba01gm","port":"6382","replication_role":"master","server":"localhost"},"m":{"redis_keyspace":{"avg_ttl":372942585,"expires":26,"keys":29}}}`,
}

func readNextInfluxLPToken(lp string, begin int) (string, int) {
	var buffer bytes.Buffer
	i := begin
	for ; i < len(lp); i++ {
		switch lp[i] {
		case ',':
			return buffer.String(), i
		case ' ':
			return buffer.String(), i
		case '=':
			return buffer.String(), i
		case '\\':
			i++
			buffer.WriteByte(lp[i])
		case '"':
			i++
			for lp[i] != '"' {
				buffer.WriteByte(lp[i])
				i++
			}
		default:
			buffer.WriteByte(lp[i])
		}
	}

	return buffer.String(), i
}

func parseLP(lp string) {
	// len := len(lp)

	measure, cur := readNextInfluxLPToken(lp, 0)
	fmt.Println(measure)

	for lp[cur] != ' ' {
		var tagName, tagValue string
		tagName, cur = readNextInfluxLPToken(lp, cur+1)
		tagValue, cur = readNextInfluxLPToken(lp, cur+1)

		fmt.Printf("%s = %s, ", tagName, tagValue)
	}

	fmt.Printf("\n")
	cur++

	for lp[cur] != ' ' {
		var fieldName, fieldValue string
		fieldName, cur = readNextInfluxLPToken(lp, cur+1)
		fieldValue, cur = readNextInfluxLPToken(lp, cur+1)

		fmt.Printf("%s = %s, ", fieldName, fieldValue)
	}

	fmt.Printf("\n")
	cur++

	ts, cur := readNextInfluxLPToken(lp, cur)
	fmt.Println(ts)
}

func TestTelegrafLPParse(t *testing.T) {
	parseLP(telegraflp_data[0])
	parseLP(telegraflp_data[1])
	parseLP(telegraflp_data[2])
	parseLP(telegraflp_data[3])
}
