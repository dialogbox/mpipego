package influxlp

import (
	"testing"
)

var data = []string{
	`cpu,host=server\ 01,region=uswest value=1,msg="all systems, abc nominal" 1508140809000000000`,
	`cpu,host=server\ 01,region=us\,west value_int=1i 1508140809000000000`,
	`cpu,host=server01,region=uswest value=1 1508140809000000000`,
	`cpu,host=server02,region=uswest value=3 1508140809000000000`,
	`temperature,machine=unit42,type=assembly internal=32,external=100 1508140809000000000`,
	`temperature,machine=unit143,type=assembly internal=22,external=130 1508140809000000000`,
	`mem,host=LM-SEL-00650622 inactive=5493116928i,used_percent=58.50780010223389,total=17179869184i,available=7128305664i,free=1635188736i,buffered=0i,used=10051563520i,cached=0i,active=6800138240i,available_percent=41.49219989776611 1508140809000000000`,
	`cpu,cpu=cpu0,host=LM-SEL-00650622 usage_idle=83,usage_nice=0,usage_irq=0,usage_softirq=0,usage_user=5,usage_system=12,usage_iowait=0,usage_steal=0,usage_guest=0,usage_guest_nice=0 1508140809000000000`,
	`cpu,cpu=cpu1,host=LM-SEL-00650622 usage_guest=0,usage_guest_nice=0,usage_user=1.0101010101010102,usage_idle=94.94949494949495,usage_iowait=0,usage_irq=0,usage_softirq=0,usage_system=4.040404040404041,usage_nice=0,usage_steal=0 1508140809000000000`,
	`cpu,cpu=cpu2,host=LM-SEL-00650622 usage_user=4.9504950495049505,usage_nice=0,usage_iowait=0,usage_guest=0,usage_guest_nice=0,usage_system=6.930693069306931,usage_idle=88.11881188118812,usage_irq=0,usage_softirq=0,usage_steal=0 1508140809000000000`,
	`cpu,cpu=cpu3,host=LM-SEL-00650622 usage_idle=94.05940594059406,usage_iowait=0,usage_irq=0,usage_softirq=0,usage_steal=0,usage_guest=0,usage_guest_nice=0,usage_user=2.9702970297029703,usage_nice=0,usage_system=2.9702970297029703 1508140809000000000`,
	`cpu,cpu=cpu-total,host=LM-SEL-00650622 usage_idle=90.02493765586036,usage_iowait=0,usage_irq=0,usage_softirq=0,usage_user=3.491271820448878,usage_system=6.483790523690773,usage_nice=0,usage_steal=0,usage_guest=0,usage_guest_nice=0 1508140809000000000`,
	`diskio,name=disk0,host=LM-SEL-00650622 read_bytes=64059693056i,read_time=1545954i,io_time=3407518i,iops_in_progress=0i,writes=1673519i,write_bytes=70562168832i,write_time=1861564i,weighted_io_time=0i,reads=2574625i 1508140809000000000`,
	`disk,path=/,device=disk1,fstype=hfs,host=LM-SEL-00650622 used=130693246976i,used_percent=52.39962371826942,inodes_total=4294967279i,inodes_free=4292144070i,inodes_used=2823209i,total=249678528512i,free=118723137536i 1508140809000000000`,
	`internal_memstats,host=LM-SEL-00650622 alloc_bytes=6063792i,pointer_lookups=245i,heap_alloc_bytes=6063792i,heap_in_use_bytes=7643136i,heap_objects=22589i,total_alloc_bytes=7547472i,frees=15414i,num_gc=2i,sys_bytes=12757240i,heap_idle_bytes=1040384i,heap_released_bytes=0i,mallocs=38003i,heap_sys_bytes=8683520i 1508140809000000000`,
	`internal_gather,host=LM-SEL-00650622,input=net metrics_gathered=20i,gather_time_ns=16403323i 1508140809000000000`,
	`internal_gather,input=cpu,host=LM-SEL-00650622 gather_time_ns=2685002i,metrics_gathered=10i 1508140809000000000`,
	`internal_gather,input=system,host=LM-SEL-00650622 metrics_gathered=4i,gather_time_ns=59415103i 1508140809000000000`,
	`internal_gather,input=netstat,host=LM-SEL-00650622 gather_time_ns=42913420i,metrics_gathered=2i 1508140809000000000`,
	`internal_write,output=file,host=LM-SEL-00650622 metrics_filtered=0i,buffer_size=71i,buffer_limit=10000i,write_time_ns=738236i,metrics_written=71i 1508140809000000000`,
	`internal_gather,input=disk,host=LM-SEL-00650622 gather_time_ns=2982343i,metrics_gathered=3i 1508140809000000000`,
	`internal_gather,input=swap,host=LM-SEL-00650622 metrics_gathered=4i,gather_time_ns=19893457i 1508140809000000000`,
	`internal_gather,input=diskio,host=LM-SEL-00650622 metrics_gathered=3i,gather_time_ns=5997383i 1508140809000000000`,
	`internal_gather,input=internal,host=LM-SEL-00650622 metrics_gathered=31i,gather_time_ns=4060442i 1508140809000000000`,
	`internal_gather,input=elasticsearch,host=LM-SEL-00650622 metrics_gathered=0i,gather_time_ns=3608560i 1508140809000000000`,
	`internal_gather,input=kernel,host=LM-SEL-00650622 metrics_gathered=0i,gather_time_ns=10950i 1508140809000000000`,
	`internal_gather,input=mem,host=LM-SEL-00650622 metrics_gathered=3i,gather_time_ns=325953i 1508140809000000000`,
	`internal_gather,host=LM-SEL-00650622,input=processes gather_time_ns=26821893i,metrics_gathered=2i 1508140809000000000`,
	`internal_agent,host=LM-SEL-00650622 metrics_gathered=82i,gather_errors=2i,metrics_written=78i,metrics_dropped=0i 1508140809000000000`,
	`net,interface=en0,host=LM-SEL-00650622 bytes_sent=355017658i,bytes_recv=818320774i,packets_recv=770229i,err_in=0i,packets_sent=571649i,err_out=0i,drop_in=0i,drop_out=51i 1508140809000000000`,
	`swap,host=LM-SEL-00650622 free=965476352i,used_percent=10.0830078125,total=1073741824i,used=108265472i 1508140809000000000`,
	`swap,host=LM-SEL-00650622 in=0i,out=0i 1508140809000000000`,
	`net,interface=en1,host=LM-SEL-00650622 drop_in=0i,bytes_sent=0i,err_in=0i,packets_recv=0i,err_out=0i,drop_out=0i,bytes_recv=0i,packets_sent=0i 1508140809000000000`,
	`net,interface=en2,host=LM-SEL-00650622 bytes_sent=0i,bytes_recv=0i,packets_sent=0i,packets_recv=0i,err_in=0i,err_out=0i,drop_in=0i,drop_out=0i 1508140809000000000`,
	`net,interface=p2p0,host=LM-SEL-00650622 packets_sent=0i,bytes_recv=0i,packets_recv=0i,err_in=0i,err_out=0i,drop_in=0i,drop_out=0i,bytes_sent=0i 1508140809000000000`,
	`net,host=LM-SEL-00650622,interface=awdl0 packets_recv=682i,err_in=0i,drop_in=0i,bytes_recv=149375i,packets_sent=718i,drop_out=0i,bytes_sent=186840i,err_out=0i 1508140809000000000`,
	`net,interface=utun0,host=LM-SEL-00650622 packets_recv=0i,bytes_recv=0i,packets_sent=3i,err_in=0i,err_out=0i,drop_in=0i,drop_out=0i,bytes_sent=268i 1508140809000000000`,
	`net,interface=utun1,host=LM-SEL-00650622 drop_in=0i,drop_out=0i,bytes_sent=268i,bytes_recv=0i,packets_sent=3i,err_in=0i,packets_recv=0i,err_out=0i 1508140809000000000`,
	`net,interface=utun2,host=LM-SEL-00650622 bytes_sent=11085i,packets_sent=40i,err_in=0i,drop_out=0i,bytes_recv=2634i,packets_recv=18i,err_out=0i,drop_in=0i 1508140809000000000`,
	`net,interface=en4,host=LM-SEL-00650622 drop_out=0i,bytes_sent=18820283i,bytes_recv=88642815i,err_in=0i,err_out=198i,packets_sent=108502i,packets_recv=311566i,drop_in=0i 1508140809000000000`,
	`processes,host=LM-SEL-00650622 blocked=1i,zombies=0i,stopped=0i,running=3i,sleeping=337i,total=341i,unknown=0i,idle=0i 1508140809000000000`,
	`netstat,host=LM-SEL-00650622 tcp_syn_sent=0i,tcp_none=0i,tcp_time_wait=0i,tcp_fin_wait1=0i,tcp_fin_wait2=0i,tcp_close_wait=0i,tcp_last_ack=0i,tcp_listen=2i,udp_socket=3i,tcp_established=17i,tcp_syn_recv=0i,tcp_close=0i,tcp_closing=0i 1508140809000000000`,
	`system,host=LM-SEL-00650622 n_users=1i,n_cpus=4i,load1=2.54,load5=1.91,load15=1.7 1508140809000000000`,
	`system,host=LM-SEL-00650622 uptime=280557i,uptime_format="3 days,  5:55" 1508140809000000000`,
}

var expectedResult = []string{
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"cpu","t":{"host":"server 01","region":"uswest"},"m":{"cpu":{"value":1,"msg":"all systems, abc nominal"}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"cpu","t":{"host":"server 01","region":"us,west"},"m":{"cpu":{"value_int":1}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"cpu","t":{"host":"server01","region":"uswest"},"m":{"cpu":{"value":1}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"cpu","t":{"host":"server02","region":"uswest"},"m":{"cpu":{"value":3}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"temperature","t":{"machine":"unit42","type":"assembly"},"m":{"temperature":{"internal":32,"external":100}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"temperature","t":{"machine":"unit143","type":"assembly"},"m":{"temperature":{"internal":22,"external":130}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"mem","t":{"host":"LM-SEL-00650622"},"m":{"mem":{"inactive":5493116928,"used_percent":58.50780010223389,"total":17179869184,"available":7128305664,"free":1635188736,"buffered":0,"used":10051563520,"cached":0,"active":6800138240,"available_percent":41.49219989776611}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"cpu","t":{"cpu":"cpu0","host":"LM-SEL-00650622"},"m":{"cpu":{"usage_idle":83,"usage_nice":0,"usage_irq":0,"usage_softirq":0,"usage_user":5,"usage_system":12,"usage_iowait":0,"usage_steal":0,"usage_guest":0,"usage_guest_nice":0}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"cpu","t":{"cpu":"cpu1","host":"LM-SEL-00650622"},"m":{"cpu":{"usage_guest":0,"usage_guest_nice":0,"usage_user":1.0101010101010102,"usage_idle":94.94949494949495,"usage_iowait":0,"usage_irq":0,"usage_softirq":0,"usage_system":4.040404040404041,"usage_nice":0,"usage_steal":0}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"cpu","t":{"cpu":"cpu2","host":"LM-SEL-00650622"},"m":{"cpu":{"usage_user":4.9504950495049505,"usage_nice":0,"usage_iowait":0,"usage_guest":0,"usage_guest_nice":0,"usage_system":6.930693069306931,"usage_idle":88.11881188118812,"usage_irq":0,"usage_softirq":0,"usage_steal":0}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"cpu","t":{"cpu":"cpu3","host":"LM-SEL-00650622"},"m":{"cpu":{"usage_idle":94.05940594059406,"usage_iowait":0,"usage_irq":0,"usage_softirq":0,"usage_steal":0,"usage_guest":0,"usage_guest_nice":0,"usage_user":2.9702970297029703,"usage_nice":0,"usage_system":2.9702970297029703}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"cpu","t":{"cpu":"cpu-total","host":"LM-SEL-00650622"},"m":{"cpu":{"usage_idle":90.02493765586036,"usage_iowait":0,"usage_irq":0,"usage_softirq":0,"usage_user":3.491271820448878,"usage_system":6.483790523690773,"usage_nice":0,"usage_steal":0,"usage_guest":0,"usage_guest_nice":0}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"diskio","t":{"name":"disk0","host":"LM-SEL-00650622"},"m":{"diskio":{"read_bytes":64059693056,"read_time":1545954,"io_time":3407518,"iops_in_progress":0,"writes":1673519,"write_bytes":70562168832,"write_time":1861564,"weighted_io_time":0,"reads":2574625}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"disk","t":{"path":"/","device":"disk1","fstype":"hfs","host":"LM-SEL-00650622"},"m":{"disk":{"used":130693246976,"used_percent":52.39962371826942,"inodes_total":4294967279,"inodes_free":4292144070,"inodes_used":2823209,"total":249678528512,"free":118723137536}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"internal_memstats","t":{"host":"LM-SEL-00650622"},"m":{"internal_memstats":{"alloc_bytes":6063792,"pointer_lookups":245,"heap_alloc_bytes":6063792,"heap_in_use_bytes":7643136,"heap_objects":22589,"total_alloc_bytes":7547472,"frees":15414,"num_gc":2,"sys_bytes":12757240,"heap_idle_bytes":1040384,"heap_released_bytes":0,"mallocs":38003,"heap_sys_bytes":8683520}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"internal_gather","t":{"host":"LM-SEL-00650622","input":"net"},"m":{"internal_gather":{"metrics_gathered":20,"gather_time_ns":16403323}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"internal_gather","t":{"input":"cpu","host":"LM-SEL-00650622"},"m":{"internal_gather":{"gather_time_ns":2685002,"metrics_gathered":10}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"internal_gather","t":{"input":"system","host":"LM-SEL-00650622"},"m":{"internal_gather":{"metrics_gathered":4,"gather_time_ns":59415103}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"internal_gather","t":{"input":"netstat","host":"LM-SEL-00650622"},"m":{"internal_gather":{"gather_time_ns":42913420,"metrics_gathered":2}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"internal_write","t":{"output":"file","host":"LM-SEL-00650622"},"m":{"internal_write":{"metrics_filtered":0,"buffer_size":71,"buffer_limit":10000,"write_time_ns":738236,"metrics_written":71}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"internal_gather","t":{"input":"disk","host":"LM-SEL-00650622"},"m":{"internal_gather":{"gather_time_ns":2982343,"metrics_gathered":3}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"internal_gather","t":{"input":"swap","host":"LM-SEL-00650622"},"m":{"internal_gather":{"metrics_gathered":4,"gather_time_ns":19893457}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"internal_gather","t":{"input":"diskio","host":"LM-SEL-00650622"},"m":{"internal_gather":{"metrics_gathered":3,"gather_time_ns":5997383}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"internal_gather","t":{"input":"internal","host":"LM-SEL-00650622"},"m":{"internal_gather":{"metrics_gathered":31,"gather_time_ns":4060442}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"internal_gather","t":{"input":"elasticsearch","host":"LM-SEL-00650622"},"m":{"internal_gather":{"metrics_gathered":0,"gather_time_ns":3608560}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"internal_gather","t":{"input":"kernel","host":"LM-SEL-00650622"},"m":{"internal_gather":{"metrics_gathered":0,"gather_time_ns":10950}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"internal_gather","t":{"input":"mem","host":"LM-SEL-00650622"},"m":{"internal_gather":{"metrics_gathered":3,"gather_time_ns":325953}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"internal_gather","t":{"host":"LM-SEL-00650622","input":"processes"},"m":{"internal_gather":{"gather_time_ns":26821893,"metrics_gathered":2}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"internal_agent","t":{"host":"LM-SEL-00650622"},"m":{"internal_agent":{"metrics_gathered":82,"gather_errors":2,"metrics_written":78,"metrics_dropped":0}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"net","t":{"interface":"en0","host":"LM-SEL-00650622"},"m":{"net":{"bytes_sent":355017658,"bytes_recv":818320774,"packets_recv":770229,"err_in":0,"packets_sent":571649,"err_out":0,"drop_in":0,"drop_out":51}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"swap","t":{"host":"LM-SEL-00650622"},"m":{"swap":{"free":965476352,"used_percent":10.0830078125,"total":1073741824,"used":108265472}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"swap","t":{"host":"LM-SEL-00650622"},"m":{"swap":{"in":0,"out":0}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"net","t":{"interface":"en1","host":"LM-SEL-00650622"},"m":{"net":{"drop_in":0,"bytes_sent":0,"err_in":0,"packets_recv":0,"err_out":0,"drop_out":0,"bytes_recv":0,"packets_sent":0}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"net","t":{"interface":"en2","host":"LM-SEL-00650622"},"m":{"net":{"bytes_sent":0,"bytes_recv":0,"packets_sent":0,"packets_recv":0,"err_in":0,"err_out":0,"drop_in":0,"drop_out":0}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"net","t":{"interface":"p2p0","host":"LM-SEL-00650622"},"m":{"net":{"packets_sent":0,"bytes_recv":0,"packets_recv":0,"err_in":0,"err_out":0,"drop_in":0,"drop_out":0,"bytes_sent":0}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"net","t":{"host":"LM-SEL-00650622","interface":"awdl0"},"m":{"net":{"packets_recv":682,"err_in":0,"drop_in":0,"bytes_recv":149375,"packets_sent":718,"drop_out":0,"bytes_sent":186840,"err_out":0}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"net","t":{"interface":"utun0","host":"LM-SEL-00650622"},"m":{"net":{"packets_recv":0,"bytes_recv":0,"packets_sent":3,"err_in":0,"err_out":0,"drop_in":0,"drop_out":0,"bytes_sent":268}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"net","t":{"interface":"utun1","host":"LM-SEL-00650622"},"m":{"net":{"drop_in":0,"drop_out":0,"bytes_sent":268,"bytes_recv":0,"packets_sent":3,"err_in":0,"packets_recv":0,"err_out":0}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"net","t":{"interface":"utun2","host":"LM-SEL-00650622"},"m":{"net":{"bytes_sent":11085,"packets_sent":40,"err_in":0,"drop_out":0,"bytes_recv":2634,"packets_recv":18,"err_out":0,"drop_in":0}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"net","t":{"interface":"en4","host":"LM-SEL-00650622"},"m":{"net":{"drop_out":0,"bytes_sent":18820283,"bytes_recv":88642815,"err_in":0,"err_out":198,"packets_sent":108502,"packets_recv":311566,"drop_in":0}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"processes","t":{"host":"LM-SEL-00650622"},"m":{"processes":{"blocked":1,"zombies":0,"stopped":0,"running":3,"sleeping":337,"total":341,"unknown":0,"idle":0}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"netstat","t":{"host":"LM-SEL-00650622"},"m":{"netstat":{"tcp_syn_sent":0,"tcp_none":0,"tcp_time_wait":0,"tcp_fin_wait1":0,"tcp_fin_wait2":0,"tcp_close_wait":0,"tcp_last_ack":0,"tcp_listen":2,"udp_socket":3,"tcp_established":17,"tcp_syn_recv":0,"tcp_close":0,"tcp_closing":0}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"system","t":{"host":"LM-SEL-00650622"},"m":{"system":{"n_users":1,"n_cpus":4,"load1":2.54,"load5":1.91,"load15":1.7}}}`,
	`{"@timestamp":"2017-10-16T17:00:09+09:00","name":"system","t":{"host":"LM-SEL-00650622"},"m":{"system":{"uptime":280557,"uptime_format":"3 days,  5:55"}}}`,
}

func TestInfluxLPConverter(t *testing.T) {
	converter := NewConverter()

	for i := range data {
		r, err := converter.Convert([]byte(data[i]))
		if err != nil {
			t.Error(err)
		}

		jsonData := r.JSON()

		if jsonData != expectedResult[i] {
			t.Errorf("Expected:\n%s\nGot:\n%s", expectedResult[i], jsonData)
		}
	}
}

func BenchmarkInfluxLPConverter(b *testing.B) {
	b.ReportAllocs()
	converter := NewConverter()

	for i := 0; i < b.N; i++ {
		for i := range data {
			m, err := converter.Convert([]byte(data[i]))
			if err != nil {
				b.Error(err)
				continue
			}

			jsondata := m.JSON()
			if jsondata != expectedResult[i] {
				b.Logf("Expected:\n%s\nGot:\n%s", expectedResult[i], jsondata)
			}
		}
	}
}
