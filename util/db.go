package util

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/satori/go.uuid"
	"os"
	"time"
)

type GormModel struct {
	ID uuid.UUID `gorm:"primary_key;type:char(36);column:id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type User struct {
	GormModel
	Username string `gorm:"type:varchar(35);column:username;"`
	Password string `gorm:"type:varchar(100);column:password;"`
	ModelID uuid.UUID `gorm:"index;type:char(36);column:model_id;"`
	ModelType string `gorm:"type:varchar(36);column:model_type;'"`
}

type Student struct {
	GormModel
	User User `gorm:"polymorphic:User;"`
	FullName string `gorm:"type:varchar(50);column:full_name;"`
	Address string `gorm:"type:varchar(255);column:address;"`
	PhoneNumber string `gorm:"type:varchar(14);column:phone_number;"`
	Age int `gorm:"column:age;"`
}

type Admin struct {
	GormModel
	User User `gorm:"polymorphic:User;"`
	FullName string `gorm:"type:varchar(50);column:full_name;"`
	Address string `gorm:"type:varchar(255);column:address;"`
	Age int `gorm:"column:age;"`
}

type Exam struct {
	GormModel
	Title string `gorm:"type:varchar(50);column:title;"`
	WrongAnswerWeighted int `gorm:"column:wrong_answer_weighted;"`
	RightAnswerWeighted int `gorm:"column:right_answer_weighted;"`
	DurationLimit int `gorm:"column:duration_limit;"`
	Questions []Question `gorm:"foreignkey:ExamID;"`
}

type Question struct {
	GormModel
	Exam Exam `gorm:"foreignkey:ExamID;"`
	ExamID uuid.UUID `gorm:"index;type:char(36);column:exam_id;"`
	Text string `gorm:"type:varchar(255);column:text;"`
	Choices []Choice `gorm:"foreignkey:QuestionID;"`
}

type Choice struct {
	GormModel
	Exam Exam `gorm:"foreignkey:ExamID;"`
	ExamID uuid.UUID `gorm:"index;type:char(36);column:exam_id;"`
	Question Question `gorm:"foreignkey:QuestionID;"`
	QuestionID uuid.UUID `gorm:"index;type:char(36);column:question_id;"`
	Text string `gorm:"type:varchar(255);column:text;"`
	IsCorrect bool `gorm:"column:is_correct;"`
}

type Attempt struct {
	GormModel
	UserID uuid.UUID `gorm:"index;type:char(36);column:user_id;"`
	ExamID uuid.UUID `gorm:"index;type:char(36);column:exam_id;"`
	StartAt time.Time `gorm:"column:start_at"`
	EndAt time.Time `gorm:"column:end_at"`
}

type UserAnswer struct {
	GormModel
	UserID uuid.UUID `gorm:"index;type:char(36);column:user_id;"`
	ExamID uuid.UUID `gorm:"index;type:char(36);column:exam_id;"`
	QuestionID uuid.UUID `gorm:"index;type:char(36);column:question_id;"`
	ChoiceID uuid.UUID `gorm:"index;type:char(36);column:choice_id;"`
	AttemptID uuid.UUID `gorm:"index;type:char(36);column:attempt_id;"`
	IsCorrect bool `gorm:"column:is_correct;"`
	AnswerWeighted int `gorm:"column:answer_weighted;"`
}

type Submission struct {
	GormModel
	User User `gorm:"foreignkey:UserID;"`
	UserID uuid.UUID `gorm:"index;type:char(36);column:user_id;"`
	Exam Exam `gorm:"foreignkey:ExamID;"`
	ExamID uuid.UUID `gorm:"index;type:char(36);column:exam_id;"`
	Attempt Attempt `gorm:"foreignkey:AttemptID;"`
	AttemptID uuid.UUID `gorm:"index;type:char(36);column:attempt_id;"`
	WrongAnswerTotal int `gorm:"column:wrong_answer_total;"`
	RightAnswerTotal int `gorm:"column:right_answer_total;"`
	NotAnswerTotal int `gorm:"column:not_answer_total;"`
	Percentage float32 `gorm:"type:decimal(10,2);column:percentage;"`
	TotalScore int `gorm:"column:total_score;"`
	DurationComplete string `gorm:"column:duration_complete;"`
}

const (
	dialect = "postgres"
	myhost = "postgres"
	myport = "5432"
	myuser = "postgres"
	mydbname = "otoy"
	mypassword = "runwols123"
	mysslmode = "disable"
)

var (
	DB = initDB()
)

func initDB() *gorm.DB {
	db, err := gorm.Open(dialect, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", myhost, myport, myuser, mydbname, mypassword, mysslmode))
	if err != nil {
		fmt.Printf("DB connection err: %s\n", err)
		os.Exit(-1)
	}
	fmt.Print("YEAY! connect database :D\n")
	return db
}

func Migrate() {
	DB.
		DropTableIfExists(new(User),
			new(Student),
			new(Admin),
			new(Exam),
			new(Question),
			new(Choice),
			new(Attempt),
			new(UserAnswer),
			new(Submission)).
		AutoMigrate(new(User),
			new(Student),
			new(Admin),
			new(Exam),
			new(Question),
			new(Choice),
			new(Attempt),
			new(UserAnswer),
			new(Submission))
}
