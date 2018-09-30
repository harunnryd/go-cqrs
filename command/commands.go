package command

import (
	"github.com/satori/go.uuid"
)

type Command struct {
	ID uuid.UUID
	Type string
}

type CreateStudentCommand struct {
	Command
	Username string
	Password string
	FullName string
	Address string
	PhoneNumber string
	Age int
}

func NewCreateStudentCommand(username, password, fullName, address, phoneNumber string, age int) CreateStudentCommand {
	cmd := new(CreateStudentCommand)
	cmd.ID = uuid.Must(uuid.NewV4())
	cmd.Type = "CreateStudentCommand"
	cmd.Username = username
	cmd.Password = password
	cmd.FullName = fullName
	cmd.Address = address
	cmd.PhoneNumber = phoneNumber
	cmd.Age = age
	return *cmd
}

type CreateAdminCommand struct {
	Command
	Username string
	Password string
	FullName string
	Address string
	Age int
}

func NewCreateAdminCommand(username, password, fullName, address string, age int) CreateAdminCommand {
	cmd := new(CreateAdminCommand)
	cmd.ID = uuid.Must(uuid.NewV4())
	cmd.Type = "CreateAdminCommand"
	cmd.Username = username
	cmd.Password = password
	cmd.FullName = fullName
	cmd.Address = address
	cmd.Age = age
	return *cmd
}

type CreateExamCommand struct {
	Command
	Title string
	WrongAnswerWeighted int
	RightAnswerWeighted int
	DurationLimit int
}

func NewCreateExamCommand(title string, wrongAnswerWeighted, rightAnswerWeighted, durationLimit int) CreateExamCommand {
	cmd := new(CreateExamCommand)
	cmd.ID = uuid.Must(uuid.NewV4())
	cmd.Type = "CreateExamCommand"
	cmd.Title = title
	cmd.WrongAnswerWeighted = wrongAnswerWeighted
	cmd.RightAnswerWeighted = rightAnswerWeighted
	cmd.DurationLimit = durationLimit
	return *cmd
}

type CreateQuestionCommand struct {
	Command
	ExamID uuid.UUID
	Text string
}

func NewCreateQuestionCommand(examID uuid.UUID, text string) CreateQuestionCommand {
	cmd := new(CreateQuestionCommand)
	cmd.ID = uuid.Must(uuid.NewV4())
	cmd.Type = "CreateQuestionCommand"
	cmd.ExamID = examID
	cmd.Text = text
	return *cmd
}

type CreateChoiceCommand struct {
	Command
	ExamID uuid.UUID
	QuestionID uuid.UUID
	Text string
	IsCorrect bool
}

func NewCreateChoiceCommand(examID, questionID uuid.UUID, text string, isCorrect bool) CreateChoiceCommand {
	cmd := new(CreateChoiceCommand)
	cmd.ID = uuid.Must(uuid.NewV4())
	cmd.Type = "CreateChoiceCommand"
	cmd.ExamID = examID
	cmd.QuestionID = questionID
	cmd.Text = text
	cmd.IsCorrect = isCorrect
	return *cmd
}

type AttemptCommand struct {
	Command
	UserID uuid.UUID
	ExamID uuid.UUID
}

func NewAttemptCommand(userID, examID uuid.UUID) AttemptCommand {
	cmd := new(AttemptCommand)
	cmd.ID = uuid.Must(uuid.NewV4())
	cmd.Type = "AttemptCommand"
	cmd.UserID = userID
	cmd.ExamID = examID
	return *cmd
}

type UserAnswerCommand struct {
	Command
	UserID uuid.UUID
	ExamID uuid.UUID
	QuestionID uuid.UUID
	ChoiceID uuid.UUID
	AttemptID uuid.UUID
}

func NewUserAnswerCommand(userID, examID, questionID, choiceID, attempID uuid.UUID) UserAnswerCommand {
	cmd := new(UserAnswerCommand)
	cmd.ID = uuid.Must(uuid.NewV4())
	cmd.Type = "UserAnswerCommand"
	cmd.UserID = userID
	cmd.ExamID = examID
	cmd.QuestionID = questionID
	cmd.ChoiceID = choiceID
	cmd.AttemptID = attempID
	return *cmd
}

type SubmissionCommand struct {
	Command
	UserID uuid.UUID
	ExamID uuid.UUID
	AttemptID uuid.UUID
}

func NewSubmissionCommand(userID, examID, attemptID uuid.UUID) SubmissionCommand {
	cmd := new(SubmissionCommand)
	cmd.ID = uuid.Must(uuid.NewV4())
	cmd.Type = "SubmissionCommand"
	cmd.UserID = userID
	cmd.ExamID = examID
	cmd.AttemptID = attemptID
	return *cmd
}