/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/orangekame3/ibmq/model"
	"github.com/orangekame3/ibmq/pkg"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")
		runJob()
	},
}

func init() {
	jobCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runJob() {
	token := pkg.LoadCredentials()
	openQASM := `OPENQASM 2.0;
  include "qelib1.inc";
  qreg q[2];
  creg c[2];
  h q[0];
  cx q[0], q[1];
  measure q[0] -> c[0];
  measure q[1] -> c[1];`
	// token, err := pkg.Authenticate(token)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	jobID, err := submitJob(token, openQASM)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Job submitted. Job ID: %s\n", jobID)
	for {
		jobStatus, err := pkg.GetJobStatus(token, jobID)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Job status: %s\n", jobStatus.Status)

		if jobStatus.Status == "COMPLETED" {
			jobResult, err := pkg.GetJobResult(token, jobID)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Println("Job result:")
			for _, result := range jobResult.Results {
				fmt.Println(result.Counts.Values)
			}
			break
		} else if jobStatus.Status == "FAILED" || jobStatus.Status == "CANCELLED" {
			fmt.Printf("Job %s: %s\n", jobStatus.Status, jobID)
			break
		}

		// Wait for a while before checking the status again
		time.Sleep(5 * time.Second)
	}

}

func submitJob(token string, openQASM string) (string, error) {
	jobRequest := &model.JobRequest{QASM: openQASM}
	data, err := json.Marshal(jobRequest)
	if err != nil {
		return "", err
	}
	resp, err := pkg.PostRequest(token, "jobs", data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("failed to submit job: %s", body)
	}

	jobStatus := &model.JobStatus{}
	err = json.NewDecoder(resp.Body).Decode(jobStatus)
	if err != nil {
		return "", err
	}

	return jobStatus.QobjID, nil
}
