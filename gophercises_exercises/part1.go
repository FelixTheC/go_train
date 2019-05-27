package main
import (
  "bufio"
  "encoding/csv"
  "fmt"
  "os"
  "strconv"
  "strings"
)

type Problem struct {
  Question string
  Answer string
}

type UserAnswers struct {
  Question int
  Anwser bool
}

func ReadCSVFile(filename string, delimiter string) (records [][]string, err error) {

  file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
  if err != nil {
  	fmt.Println("File error: ", err)
  }

  reader := csv.NewReader(file)
  switch delimiter {
    case ",":
      reader.Comma = ','
    case ";":
      reader.Comma = ';'
    case "\t":
      reader.Comma = '\t'
    }
  csvData, err := reader.ReadAll()
  return csvData, err
}

func GetQuestions(filename string) [] Problem {
  var questions [] Problem

  csvFile, err := ReadCSVFile(filename, ",")
  if err != nil {
    fmt.Println("Error in csv file: ", err)
  } else {
    for index, line := range csvFile {
      if index > 0 {
        problem := Problem {
          Question: line[0],
          Answer: line[1],
        }
          questions = append(questions, problem)
      }
    }
  }
  return questions
}

func checkAnswer(x string, y string) bool {
  result := false

  a, _ := strconv.Atoi(x)
  b, _ := strconv.Atoi(y)
  if a == b {
    result = true
  }
  return result
}

func main() {
  var answers [] UserAnswers
  for index, problem := range GetQuestions("problems.csv") {

    reader:= bufio.NewReader(os.Stdin)
    fmt.Print(problem.Question + ": ")
    answer, _ := reader.ReadString('\n')
    answer = strings.Replace(answer, "\n", "", -1)
    userAnswer := UserAnswers {
      Question: index,
      Anwser: checkAnswer(answer, problem.Answer),
    }
    answers = append(answers, userAnswer)
  }

  fmt.Println("Your results: ")
  // i := 0
  for _, result := range answers {
    fmt.Println("Answer: ", result.Question, result.Anwser)
  }
}
