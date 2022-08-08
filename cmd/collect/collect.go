package main

import (
	"fmt"
	"time"
	"strconv"
	"encoding/json"
	"github.com/kumataahh/eagle-eye/internal/collector"
	"github.com/kumataahh/eagle-eye/pkg/tools"

)

func main(){
	println("start monitor..")

	var logdate string
	// cpu usage ?%

	// check time limit
	nowdatetime := time.Now()
	dataGroup := make(map[string]interface{})
	nowtime := nowdatetime.Format("2006-01-02 15:04:05") // current time
	nowdate := nowdatetime.Format("2006-01-02")
	if logdate == "" || nowdate != logdate {
		logdate = nowdate
	}


	// dataGroup	
	dataGroup["update_time"] = nowtime
	systemdata := capture.GetSystemInfo()
	dataGroup["work_status"] = "stop"
	dataGroup["machine_tag"] = "null"
	dataGroup["ip_address"] = systemdata["ip_address"]
	dataGroup["cpu_usage"] = systemdata["cpu_usage"]
	dataGroup["mem_usage"] = systemdata["mem_usage"]
	irondata := capture.GetIronfish()
	dataGroup["node_graffi"] = irondata["node_graffi"]
	dataGroup["node_height"] = irondata["node_height"]
	dataGroup["node_version"] = irondata["node_version"]
	dataGroup["node_count"] = irondata["node_count"]
	dataGroup["miner_count"] = irondata["miner_count"]

	machine_id, err:= strconv.Atoi(systemdata["machine_id"])
	if err != nil {
		fmt.Println("can't convert machine_id to int")
		dataGroup["machine_id"] = -1
	} else {
		dataGroup["machine_id"] = machine_id
	}

	height := irondata["node_height"]
	node_count := irondata["node_count"]
	miner_count := irondata["miner_count"]

	if 2 == machine_id {
		dataGroup["machine_tag"] = "core"
		if height != "" && node_count == "1"{
			dataGroup["work_status"] = "running"
		}
	} else if machine_id > 9 && machine_id < 25 {
			dataGroup["machine_tag"] = "node"
			if height != "" && node_count == "1" && miner_count == "4"{
				dataGroup["work_status"] = "running"
			}
	} else {
		dataGroup["machine_tag"] = "miner"
		if miner_count == "4"{
			dataGroup["work_status"] = "running"
		}
	}
	
	jsonStr, err := json.Marshal(dataGroup)
	if err != nil {
		fmt.Println("MapToJsonDemo err: ", err)
	}
	// write file
	// tools.SaveToLocalFile(logdate + ".syslog",jsonStr)

	// curl -X POST http://127.0.0.1:9100/postjson -d '{"account_pubkey":"","cpu_cores":"8","cpu_usage":"1.00%","ip_address":"172.16.200.12","mem_total":"67.33G","mem_usage":"1.90%","miner_count":"0","node_count":"0","node_graffi":"","node_height":"","node_version":"","update_time":"2022-06-22 16:00:58"}'
	postCmd := "curl -X POST http://127.0.0.1:9100/postjson -d '" + string(jsonStr) + "'"
	postRs := tools.ParseCMD(postCmd)
	fmt.Printf("post response: %s\n", postRs)
	// postCmd := "curl -X POST http://192.168.1.1:9100/postjson -d '" + string(jsonStr) + "'"
	// fmt.Println(string(jsonStr))


	// time.Sleep(time.Duration(60)*time.Second)
}