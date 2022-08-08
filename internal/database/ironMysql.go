package database

import (  
	_ "github.com/go-sql-driver/mysql"
	"github.com/kumataahh/eagle-eye/internal/models"
    "database/sql"
	"fmt"
)

var db *sql.DB

func PostIronMysql(dataGroup models.Collections){

	db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/ironfish?charset=utf8&parseTime=True")  
	if err != nil {  
	fmt.Println("connect mysql error!", err)  
	}
	defer db.Close()

    _, delete_err := db.Exec("DELETE FROM miner_ironfish WHERE machine_id = ?", dataGroup.MachineId)
    if delete_err!=nil{
        fmt.Println("delete err=",delete_err)
        return
    }
	fmt.Println("delete row success!")

	_, insert_err := db.Exec("INSERT INTO miner_ironfish(machine_id, ip_address, machine_tag, work_status, cpu_usage, mem_usage, miner_count, node_count, node_graffi, node_height, node_version, update_time) VALUES(?,?,?,?,?,?,?,?,?,?,?,?);",
	dataGroup.MachineId, 
	dataGroup.IPAddress, 
	dataGroup.MachineTag,  
	dataGroup.WorkStatus, 
	dataGroup.CPUUsage, 
	dataGroup.MemUsage, 
	dataGroup.MinerCount, 
	dataGroup.NodeCount,
	dataGroup.NodeGraffi, 
	dataGroup.NodeHeight, 
	dataGroup.NodeVersion, 
	dataGroup.UpdateTime)

    if insert_err!=nil{
        fmt.Println("insert err=",insert_err)
        return
    }
	fmt.Println("insert row success!")
}