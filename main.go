package main

import (
	"flag"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"quizes/commandApi"
	"quizes/handler"
	"quizes/queryApi"
	"quizes/util"
)

func main() {
	util.DB.LogMode(true)
	act := flag.String("act", "rest", "Either: rest or consumer")
	flag.Parse()

	fmt.Printf("Welcome to quizes service: %s\n\n", *act)

	switch *act {
	case "rest":
		fmt.Print("Welcome to REST\n")

		router := echo.New()
		u := router.Group("users")
		u.POST("/admins", commandApi.CreateAdminHandler)
		u.GET("/admins", queryApi.ListAdminHandler)
		u.POST("/admins/exams", commandApi.CreateExamHandler)
		u.POST("/admins/exams/questions", commandApi.CreateQuestionHandler)
		u.POST("/admins/exams/questions/choices", commandApi.CreateChoiceHandler)

		u.POST("/students", commandApi.CreateStudentHandler)
		u.GET("/students", queryApi.ListStudentHandler)
		u.POST("/students/attempts", commandApi.CreateAttemptHandler)
		u.POST("/students/attempts/user-answers", commandApi.CreateUserAnswerHandler)
		u.POST("/students/exams/attempts/submissions", commandApi.CreateSubmissionHandler)
		u.GET("/students/:user_id/exams/:exam_id/attempts/:attempt_id/submissions", queryApi.DetailSubmissionHandler)

		e := router.Group("exams")
		e.GET("", queryApi.ListExamHandler)
		e.GET("/:exam_id", queryApi.DetailExamHandler)
		e.GET("/:exam_id/submissions", queryApi.ListRankExam)


		router.Logger.Fatal(router.Start(":8080"))
	case "consumer":
		fmt.Print("Welcome to CONSUMER\n")
		handler.Listen("g-message", "message")
	case "testing":
		fmt.Print("Welcome to TESTING\n")
		fmt.Printf("write this you want!")
	case "migrate":
		fmt.Print("Welcome to Migrate\n")
		util.Migrate();
	}
}
