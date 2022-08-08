package capture

import (
	"github.com/kumataahh/eagle-eye/pkg/tools"
)



func GetIronfish() map[string]string {

	irondata := make(map[string]string)

	irondata["node_graffi"] = string(tools.ParseCMD("docker exec -i iron-node bash -c \"ironfish config:show\" | grep blockGraffiti | awk '{print $2}' | cut -c2-8 | xargs echo -n"))
	syncStatus := string(tools.ParseCMD("docker exec -i iron-node bash -c \"ironfish status\" | grep Blockchain | awk '{print $2}' | xargs echo -n"))
	if "SYNCED" == syncStatus {
		irondata["node_height"] = string(tools.ParseCMD("docker exec -i iron-node bash -c \"ironfish status\" | grep Blockchain | awk '{print $6}' | cut -d '(' -f2 | cut -d ')' -f1 | xargs echo -n"))
	}else{
		irondata["node_height"] = string(tools.ParseCMD("docker exec -i iron-node bash -c \"ironfish status\" | grep Blockchain | awk '{print $7}' | cut -d '(' -f2 | cut -d ')' -f1 | xargs echo -n"))
	}
	irondata["node_version"] = string(tools.ParseCMD("docker exec -i iron-node bash -c \"ironfish status\" | grep Version |  awk '{print $2}' | xargs echo -n"))
	irondata["node_count"] = string(tools.ParseCMD("docker ps | grep iron-node | wc -l | xargs echo -n"))
	irondata["miner_count"] = string(tools.ParseCMD("docker ps | grep iron-miner | wc -l | xargs echo -n"))
	// irondata["account_pubkey"] = string(tools.ParseCMD("docker exec -i iron-node bash -c \"ironfish accounts:publickey\" | awk '{print $5}' | xargs echo -n"))

	return irondata
}

// func GetIronNode() map[string]string {

// 	nodedata := make(map[string]string)

// 	nodedata["Graffi"] = string(tools.ParseCMD("docker exec -i iron-node bash -c \"ironfish config:show\" | grep blockGraffiti | awk '{print $2}' | cut -c2-8 | xargs echo -n"))
// 	nodedata["Height"] = string(tools.ParseCMD("docker exec -i iron-node bash -c \"ironfish status\" | grep Blockchain | awk '{print $6}' | cut -d '(' -f2 | cut -d ')' -f1 | xargs echo -n"))
// 	nodedata["Version"] = string(tools.ParseCMD("docker exec -i iron-node bash -c \"ironfish status\" | grep Version |  awk '{print $2}' | xargs echo -n"))
// 	nodedata["nodeCount"] = string(tools.ParseCMD("docker ps | grep iron-node | wc -l | xargs echo -n"))

// 	return nodedata
// }


// func GetIronMiner() map[string]string{	
// 	minerdata := make(map[string]string)
// 	minerdata["MinerCount"] = string(tools.ParseCMD("docker ps | grep iron-miner | wc -l | xargs echo -n"))
// 	// minerdata["HashRate"] = string(tools.ParseCMD("docker ps | grep iron-miner | wc -l | xargs echo -n"))

// 	return minerdata
// }

// func GetIronAccount() map[string]string{	
// 	accountdata := make(map[string]string)
// 	accountdata["Pubkey"] = string(tools.ParseCMD("docker exec -i iron-node bash -c \"ironfish accounts:publickey\" | awk '{print $5}' | xargs echo -n"))

// 	return accountdata
// }

