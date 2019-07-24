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

// var createTableStatements = []string{
//     `CREATE TABLE IF NOT EXISTS questions (
//         id SERIAL PRIMARY KEY,
//         question VARCHAR(255) NULL,
//         dimension VARCHAR(255) NULL
//     )`,
// }

// var seedDataStatements = []string{
//     `INSERT INTO questions (
//         question,
//         dimension
//     )

//     VALUES (
//         'I feel invigorated from my time being around other people.',
//         'extraversion'
//     ),
//     (
//         'I feel comfortable working in groups of people and enjoy it.',
//         'extraversion'
//     ),
//     (
//         'Others may describe me as ‘reserved’ or ‘reflective.',
//         'introversion'
//     ),
//     (
//         'Sometimes I spend too much time reflecting and do not take action quickly enough.',
//         'introversion'
//     );`,
// }

// // createTable creates the table, and if necessary, the database.
// func createTable(conn *sql.DB) error {
//     for _, stmt := range createTableStatements {
//         _, err := conn.Exec(stmt)
//         if err != nil {
//             return err
//         }
//     }
//     return nil
// }

// func seedData(conn *sql.DB) error {
//     for _, stmt := range seedDataStatements {
//         _, err := conn.Exec(stmt)
//         if err != nil {
//             return err
//         }
//     }
//     return nil
// }

type Question struct {
    ID      int64  `json:"id"`
    Question   string `json:"question"`
    Dimension string `json:"dimension"`
    // created_at time.Time `json:"created_at"`
    // updated_at time.Time `json:"updated_at"`
}

type Answer struct {
    ID      int64  `json:"id"`
    User   int64 `json:"user"`
    Question   int64 `json:"question"`
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
        fmt.Println(user.Email)

        stmt, err := db.Prepare("INSERT INTO users(email) VALUES ($1) RETURNING id")
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
        answers := c.Param("answers")

        fmt.Println(answers)
        // fmt.Println(c.Param)
        fmt.Println(c.Param("answers"))



        // answers = make([]*Answer, 0)
        // for _, answer := range mList {
        //   _, err := stmt.Exec(int64(answer.Question), int64(answer.User))
        //   if err != nil {
        //     log.Fatal(err)
        //   }
        // }

        // query := "insert INTO users(email) values(?)"

        // stmt, err := m.Conn.PrepareContext(ctx, query)
        // if err != nil {
        //     return -1, err
        // }

        // user := &User{
        //     Email: email,
        // }

        // res, err := stmt.ExecContext(ctx, user.Email)
        // fmt.Println(err)

        // defer stmt.Close()

        // if err != nil {
        //     return -1, err
        // }

        // c.JSON(http.StatusOK, res.LastInsertId())

        // stmt, _ := db.Prepare(pq.CopyIn("answers", "question", "user")) // answers is the table name
        // m := &Answer{
        //   Question:          123456,
        //   User:       123434,
        // }
        // mList := make([]*Answer, 0)
        // for i:=0 ; i<100 ; i++ {
        //   mList = append(mList, m)
        // }
        // fmt.Println(stmt)
        // fmt.Println(mList)
        // for _, answer := range mList {
        //   _, err := stmt.Exec(int64(answer.Question), int64(answer.User))
        //   if err != nil {
        //     log.Fatal(err)
        //   }
        // }
        // err = stmt.Close()
        // if err != nil {
        //   log.Fatal(err)
        // }
        // err = txn.Commit()
        // if err != nil {
        //   log.Fatal(err)
        // }
        // for _, answer := range mList {
        //   _, err := stmt.Exec(int64(answer.Question), int64(answer.User))
        //   if err != nil {
        //     log.Fatal(err)
        //   }
        // }
        // _, err = stmt.Exec()
        // if err != nil {
        //   log.Fatal(err)
        // }
        // err = stmt.Close()
        // if err != nil {
        //   log.Fatal(err)
        // }
        // err = txn.Commit()
        // if err != nil {
        //   log.Fatal(err)
        // }
        // query := "insert INTO answers(question, user, resonance) values(?, ?, ?, ?)"

        // stmt, err := m.Conn.PrepareContext(ctx, query)
        // if err != nil {
        //     return -1, err
        // }

        // payload, err := stmt.ExecContext(ctx, p.Question, p.User, p.Resonance)

        // defer stmt.Close()

        // if err != nil {
        //     return -1, err
        // }

        // c.JSON(http.StatusOK, payload)
    }
}

// func ensureTableExists() error {
//     conn, err := sql.Open("postgres", "")
//     if err != nil {
//         return fmt.Errorf("mysql: could not get a connection: %v", err)
//     }
//     defer conn.Close()

//     // // Check the connection.
//     // if conn.Ping() == driver.ErrBadConn {
//     //     return fmt.Errorf("mysql: could not connect to the database. " +
//     //         "could be bad address, or this address is not whitelisted for access.")
//     // }

//     if _, err := conn.Exec("DESCRIBE questions"); err != nil {
//         createTable(conn)
//         return seedData(conn)
//     }
//     return nil
// }

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
    router.POST("/users", createUser(db))
    router.POST("/answers", createAnswers(db))

    router.Run(":" + port)
}