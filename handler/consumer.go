package handler

import (
	"fmt"
	"quizes/command"
	"quizes/util"
)

var (
	key string
)

func Listen(groupID string, topics string) {
	consumer := util.NewKafkaConsumer(groupID, topics)
	defer func() {
		if err := consumer.Close(); err != nil {
			fmt.Printf("Kafka err: %s\n", err)
		}
	}()

	go func() {
		for {
			select {
			case err, more := <-consumer.Errors():
				if more {
					fmt.Printf("Kafka err: %s\n", err)
				}
			case ntf, more := <-consumer.Notifications():
				if more {
					fmt.Printf("Rebalanced: %+v\n", ntf)
				}
			}
		}
	}()

	for {
		select {
		case msg := <-consumer.Messages():
			key = string(msg.Key)
			consumer.MarkOffset(msg, "")

			switch key {
			case "CreateStudentCommand":
				cmd := new(command.CreateStudentCommand)
				util.Decode(msg.Value, cmd)
				cmd.Process()
			case "CreateAdminCommand":
				cmd := new(command.CreateAdminCommand)
				util.Decode(msg.Value, cmd)
				cmd.Process()
			case "CreateExamCommand":
				cmd := new(command.CreateExamCommand)
				util.Decode(msg.Value, cmd)
				cmd.Process()
			case "CreateQuestionCommand":
				cmd := new(command.CreateQuestionCommand)
				util.Decode(msg.Value, cmd)
				cmd.Process()
			case "CreateChoiceCommand":
				cmd := new(command.CreateChoiceCommand)
				util.Decode(msg.Value, cmd)
				cmd.Process()
			case "AttemptCommand":
				cmd := new(command.AttemptCommand)
				util.Decode(msg.Value, cmd)
				cmd.Process()
			case "UserAnswerCommand":
				cmd := new(command.UserAnswerCommand)
				util.Decode(msg.Value, cmd)
				cmd.Process()
			case "SubmissionCommand":
				cmd := new(command.SubmissionCommand)
				util.Decode(msg.Value, cmd)
				cmd.Process()
			}
		}
	}
}