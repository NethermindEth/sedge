/*
Copyright 2022 Nethermind

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cli

import (
	"fmt"
	"strconv"

	// "runtime/metrics"
	// "internal/pkg/hardware/hardware"
	log "github.com/sirupsen/logrus"

	"github.com/NethermindEth/sedge/internal/pkg/hardware"
	"github.com/spf13/cobra"
)

const MIN_RAM =4
const MIN_CPU =2; 
const MIN_SPACE =200;
func HardwareCommand() *cobra.Command{
	cmd:= &cobra.Command{
		Use:"hardware",
		Short:"Manage the hardware of the host device ",
		Long:"Manage , check and resolve hardware requirements",
	}
	// cmd.AddCommand(hardwareCheckCommand());
	cmd.AddCommand(hardwareMetricsCheckCommand());
	cmd.AddCommand(hardwareMetricsReqDisplayCommand())
	return cmd;
	
}

func hardwareMetricsCheckCommand() *cobra.Command {
	cmd:= &cobra.Command{
		Use: "check",
		Short: "Verifiy minimum requirements",
		Long:"Check whther host meets the minimum requirement needed to run Sedge",
		RunE: func(cmd *cobra.Command,args []string) error {
			//Memory free
			hardwaremetrics,err:= hardware.GetHardwareMetrics();
			var eligible = true
			if  err!=nil {
				return fmt.Errorf("error in hardware.go",err.Error())

			}
			fmt.Println(hardwaremetrics);
			if(hardwaremetrics.CPUs<MIN_CPU){
				eligible=false
				log.Errorf("Available Number of CPUs:- %s ; Required Number of CPUs:- %s ; Status : FAILED ",strconv.Itoa(hardwaremetrics.CPUs),strconv.Itoa(MIN_CPU))
				// fmt.Printf("Minimum CPU not available. Device has %s number of CPUs.Required %s number of CPUs to run Sedge !\n", strconv.Itoa(MIN_CPU))
			} else{
				log.Infof("Available Number of CPUs:- %s ; Required Number of CPUs:- %s ; Status : OK ",strconv.Itoa(hardwaremetrics.CPUs),strconv.Itoa(MIN_CPU))


			}
			if(hardwaremetrics.FreeRAM<MIN_RAM ){
				eligible=false
				log.Errorf("Available Free Physical RAM:- %s GB ; Required Free Physical RAM:- %s GB; Status : FAILED ",strconv.Itoa(int(hardwaremetrics.FreeRAM)),strconv.Itoa(MIN_RAM))
				// fmt.Printf("Minimum CPU not available. Device has %s number of CPUs.Required %s number of CPUs to run Sedge !\n", strconv.Itoa(MIN_CPU))
			} else{
				log.Infof("Available Free Physical RAM:- %s GB; Required Free Physical RAM:- %s GB; Status : OK ",strconv.Itoa(int(hardwaremetrics.FreeRAM)),strconv.Itoa(MIN_RAM))


			}
		var s=""
			for _, v := range hardwaremetrics.Disks {
				if(v.Free>=MIN_SPACE){
					s+=v.Name
					s+=","
				}
			}
			
			if(len(s)==0 ){
				eligible=false
				log.Errorf(" No Disk meet the minimum Free Disk Space . Required :- %s GB; Status : FAILED ",strconv.Itoa(MIN_SPACE))
				// fmt.Printf("Minimum CPU not available. Device has %s number of CPUs.Required %s number of CPUs to run Sedge !\n", strconv.Itoa(MIN_CPU))
			} else{
				log.Infof("Disks %s meet the Minimum Disk Requirement . Required :- %s GB; Status : OK",s,strconv.Itoa(MIN_SPACE))


			}
			if(eligible){
log.Infof("\n Device is ELIGIBLE to run Sedge ");
			}else{
				log.Infof("\n Device is NOT ELIGIBLE to run Sedge ");

			}
			return nil;
			

			
		},
	}
	return cmd;

}
func hardwareMetricsReqDisplayCommand() *cobra.Command{
	cmd:= &cobra.Command{
		Use: "display",
		Short: "Display minimumn requirements and available requirements",
		Long: "Display the minimum hardware requirements involving RAM ,CPU ,Disk to run Sedge.",
		RunE: func(cmd * cobra.Command,args []string) error {
			hardwaremetrics,err:= hardware.GetHardwareMetrics();
			var result ="";
			if  err!=nil {
				return fmt.Errorf("error in hardware.go",err.Error())

			}
			// fmt.Println(hardwaremetrics);
		
				result+=("\n\n Number of CPUs:- \n Available = "+strconv.Itoa(hardwaremetrics.CPUs)+" ; Required = "+strconv.Itoa(MIN_CPU)+" ; ")


			
		
				result+=" \n\n Free Physical Memory (RAM):- \n Available =  "+strconv.Itoa(int(hardwaremetrics.FreeRAM))+" GB; Required = "+strconv.Itoa(MIN_RAM)+" GB; "
				result+= " \n\n Total Available Physical Memory (RAM):-  "+strconv.Itoa(int(hardwaremetrics.TotalRam))+"GB; "


			
		var s=""
			for _, v := range hardwaremetrics.Disks {
					s+=(v.Name + " "+v.Mount+"="+strconv.Itoa(int(v.Free))+"GB")
					s+="\n"
			}
			
			
				result+=("\n\nDisks detail:- \n"+s+" \n. Required :- "+strconv.Itoa(MIN_SPACE)+" GB; ")

log.Infof(result)
			
		
			return nil;

		},
	}
	return cmd;

}