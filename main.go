package main

import (
	"context"
	"fmt"

	// "io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/vmihailenco/taskq/v3"
	"github.com/vmihailenco/taskq/v3/memqueue"
)

func main() {
	fmt.Println("Hii")

	var QueueFactory = memqueue.NewFactory()
	var MainQueue = QueueFactory.RegisterQueue(&taskq.QueueOptions{
		Name: "api-worker",
		Handler: taskq.NewHandler(func(ctx context.Context) error {
			res, err := http.Post("http://localhost:6565/v1/updateRuleStatus", "", nil)
			if err != nil {
				return err
			}

			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
			defer res.Body.Close()

			fmt.Println("SERVER INFORMED ABOUT STATUS")

			return nil
		}),
	})

	// MainQueue.Add(&taskq.Message{
	// 	Name: "RuleState1",
	// 	Args: []interface{}{"RuleState1"},
	// })
	c := 0
	go GetRulesFromServer(MainQueue, &c)

	go SendStatusToServer(MainQueue)

	time.Sleep(time.Second * 1)
}

func GetRulesFromServer(mainQueue taskq.Queue, counter *int) {
	fmt.Println("IN GetRulesFromServer")
	var mTask = taskq.RegisterTask(&taskq.TaskOptions{
		Name: "RuleStateName" + strconv.Itoa(*counter),
		Handler: func() error {
			fmt.Println("IN HANDLER2")
			return nil
		},
	})

	for {
		if isInternetConnected() {
			res, err := http.Post("http://localhost:6565/v1/getRules", "", nil)
			if err != nil {
				return
			}

			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			defer res.Body.Close()

			//

			err = mainQueue.Add(mTask.WithArgs(context.Background()))
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("REQUEST ADDED IN QUEUE")

		}
	}
}

func SendStatusToServer(mainQueue taskq.Queue) {

	if isInternetConnected() {

		c := context.Background()
		mainQueue.Consumer().Start(c)
		mainQueue.Options().Handler.HandleMessage(taskq.NewMessage(c))

	}

}

func isInternetConnected() (ok bool) {
	_, err := http.Get("http://clients3.google.com/generate_204")
	return err == nil
}
