package main

import (
    "database/sql"
    "encoding/json"
    "github.com/joho/godotenv"
    "net/http"
    "fmt"
    "log"
    "os"

    "github.com/gin-gonic/gin"
    _ "github.com/heroku/x/hmetrics/onload"
    _ "github.com/lib/pq"
    cors "github.com/rs/cors/wrapper/gin"
)

type Question struct {
    ID      int64  `json:"id"`
    Question   string `json:"question"`
    Dimension string `json:"dimension"`
    // created_at time.Time `json:"created_at"`
    // updated_at time.Time `json:"updated_at"`
}

type Answer struct {
    ID      int64  `json:"id"`
    Answer   int64 `json:"answer"`
    Question   int64 `json:"question"`
    User   int64 `json:"user"`
    // Resonance string `json:"resonance"`
    // created_at time.Time `json:"created_at"`
    // updated_at time.Time `json:"updated_at"`
}

type User struct {
    ID      int64  `json:"id"`
    Email   string `json:"email"`
    // created_at time.Time `json:"created_at"`
    // updated_at time.Time `json:"updated_at"`
}

func getAnswers(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        query := c.Request.URL.Query()
        fmt.Println(query)
        
        // rows, err := db.Query("SELECT * FROM answers WHERE userid=$1")
        // if err != nil {
        //     c.String(http.StatusInternalServerError,
        //         fmt.Sprintf("Error reading questions: %q", err))
        //     return
        // }
        // defer rows.Close()

        // payload := make([]*Answer, 0)
        // for rows.Next() {
        //     data := new(Answer)

        //     err := rows.Scan(
        //         &data.ID,
        //         &data.Question,
        //         &data.Answer,
        //     )
        //     if err != nil {
        //         fmt.Println(err)
        //     }
        //     payload = append(payload, data)
        // }
        // if err := rows.Err(); err != nil {
        //     log.Fatalf("[x] Error when getting the list of questions. Reason: %s", err.Error())
        // }

        // c.JSON(http.StatusOK, payload)
    }
}

func getQuestions(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        rows, err := db.Query("SELECT * FROM questions")
        if err != nil {
            c.String(http.StatusInternalServerError,
                fmt.Sprintf("Error reading questions: %q", err))
            return
        }
        defer rows.Close()

        payload := make([]*Question, 0)
        for rows.Next() {
            data := new(Question)

            err := rows.Scan(
                &data.ID,
                &data.Question,
                &data.Dimension,
            )
            if err != nil {
                fmt.Println(err)
            }
            payload = append(payload, data)
        }
        if err := rows.Err(); err != nil {
            log.Fatalf("[x] Error when getting the list of questions. Reason: %s", err.Error())
        }

        c.JSON(http.StatusOK, payload)
    }
}

func createUser(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        buf := make([]byte, 1024)
        num, _ := c.Request.Body.Read(buf)
        reqBody := string(buf[0:num])

        user := User{}

        json.Unmarshal([]byte(reqBody), &user)

        stmt, err := db.Prepare("INSERT INTO users(email) VALUES ($1) RETURNING id")
        fmt.Println(stmt)

        if err != nil {
            log.Fatalf("[x] Error. Reason: %s", err.Error())
        }

        var id int
        errr := stmt.QueryRow(user.Email).Scan(&id)

        if errr != nil {
            fmt.Println(err)
        }

        defer stmt.Close()

        c.JSON(http.StatusOK, id)
    }
}

func createAnswers(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        buf := make([]byte, 1024)
        num, _ := c.Request.Body.Read(buf)
        reqBody := string(buf[0:num])

        answers := []Answer{}

        json.Unmarshal([]byte(reqBody), &answers)

        stmt, _ := db.Prepare("INSERT INTO answers(question, answer, userid) VALUES ($1, $2, $3)")

        for _, answer := range answers {
          _, err := stmt.Exec(answer.Question, answer.Answer, answer.User)
          if err != nil {
            fmt.Println(err)
          }
        }

        defer stmt.Close()
        // c.JSON(http.StatusOK)
    }
}

// init is invoked before main()
func init() {
    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

func main() {
    port, portExists := os.LookupEnv("PORT")

    if portExists {
        fmt.Println(port)
    }

    if port == "" {
        log.Fatal("$PORT must be set")
    }

    dbUrl, dbExists := os.LookupEnv("DATABASE_URL")

    if dbExists {
        fmt.Println(dbUrl)
    }

    db, err := sql.Open("postgres", dbUrl)
    if err != nil {
        log.Fatalf("Error opening database: %q", err)
    }

    router := gin.Default()

    router.Use(cors.AllowAll())
    router.Use(gin.Logger())

    router.GET("/questions", getQuestions(db))
    router.GET("/answers", getAnswers(db))
    router.POST("/users", createUser(db))
    router.POST("/answers", createAnswers(db))

    router.Run(":" + port)
}