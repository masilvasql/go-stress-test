package cli

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/masilvasql/go-stress-test/usecase"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var stressTestCmd = &cobra.Command{
	Use:   "stressTest",
	Short: "You can use this command to stress test a URL",
	Long:  `This command will send a number of requests to a given URL and return the response time and status code of each request.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		color.Yellow("Stress Test is running")
	},
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetInt("requests")
		concurrency, _ := cmd.Flags().GetInt("concurrency")

		usecase := usecase.NewStressTest(url, requests, concurrency)

		output, err := usecase.Execute()
		if err != nil {
			color.Red("Error: %v", err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Descrição", "Valor"})

		table.Append([]string{"Total de requisições", fmt.Sprintf("%d", output.TotalRequest)})
		table.Append([]string{"Tempo médio de resposta (ms)", fmt.Sprintf("%.2f", output.AvgResponseTime)})
		table.Append([]string{"Tempo total de resposta (ms)", fmt.Sprintf("%.2f", output.TotalResponseTime)})
		table.Append([]string{"Tempo total de resposta (s)", fmt.Sprintf("%.2f", output.TotalResponseTime/1000)})
		table.Append([]string{"Total de requisições bem-sucedidas", fmt.Sprintf("%d", output.StatusCounts[200])})

		table.Append([]string{"Distribuição de status codes", ""})
		for code, count := range output.StatusCounts {
			table.Append([]string{fmt.Sprintf("%d", code), fmt.Sprintf("%d", count)})
		}

		table.Render()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		color.Green("Stress Test is done")
	},
}

func init() {
	rootCmd.AddCommand(stressTestCmd)

	stressTestCmd.Flags().StringP("url", "u", "", "The URL to send the requests to")
	stressTestCmd.MarkFlagRequired("url")
	stressTestCmd.Flags().IntP("requests", "r", 10, "The number of requests to send")
	stressTestCmd.Flags().IntP("concurrency", "c", 2, "The number of requests to send concurrently")
}
