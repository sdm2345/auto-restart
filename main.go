package main

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func main() {

	var delay int
	var maxCount int
	var forceRestart bool
	var debug bool
	rootCmd := &cobra.Command{
		Use:     "auto-restart",
		Example: "auto-restart -- curl https://google.com",
		Args:    cobra.MinimumNArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			if debug {
				log.Printf("auto-start will to start process with delay:%d,try max count:%d, args:%v", delay, maxCount, args)
			}

			var t1, t2 time.Time
			for count := 0; ; time.Sleep(time.Second * time.Duration(delay)) {
				count++
				if maxCount > 0 && count > maxCount {
					log.Println("reach max count, exit")
					break
				}
				if debug {

					log.Printf("start process with count:%d,try max count:%d", count, maxCount)
				}
				t1 = time.Now()
				err := StartSubProcess(args)
				t2 = time.Now()
				if err != nil {
					if debug {

						log.Printf("start process get error:%v,last run time:%.2fs", err, t2.Sub(t1).Seconds())
					}
					continue
				}
				if forceRestart {
					if debug {
						log.Printf("force restart process get error:%v,last run time:%.2fs", err, t2.Sub(t1).Seconds())

					}
					continue
				}
				break
			}
		},
	}
	rootCmd.Flags().IntVarP(&delay, "delay-time", "t", 3, "auto restart delay time")
	rootCmd.Flags().BoolVarP(&forceRestart, "force", "f", false, "Force restart regardless of exit code 0")
	rootCmd.Flags().BoolVarP(&debug, "debug", "d", false, "debug flag")
	rootCmd.Flags().IntVarP(&maxCount, "max-count", "m", 0, "max retry count, 0 means no limit")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
func StartSubProcess(args []string) error {
	var err error
	var cmd = exec.Command(args[0], args[1:]...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err = cmd.Start(); err != nil {
		log.Println("start process get error:", args[0])
		return err
	}
	if err = cmd.Wait(); err != nil {
		log.Println("wait process get error:", args[0])
		return err
	}
	return nil
}
