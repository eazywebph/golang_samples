// Sequence: 4
// Title: MySQL Database
// Source: https://gowebexamples.com/mysql-database/

/* 
Prerequisites on a new machine

Method 1 using traditional MySQL database.
1. Install the MySQL database driver.
2. $ go get -u github.com/go-sql-driver/mysql
3. Set your host, username and password.

Method 2 using Docker
1. Install Docker and create a new container instance for MySQL
2. $ sudo docker pull mysql/mysql-server:latest
3. Verify the versions pulled from the docker host
4. $ sudo docker images
5. Deploy the MySQL container
6. $ sudo docker run --name=[container_name] -d [image_tag_name]
7. example $ sudo docker run --name=mysql_golang -d mysql/mysql-server, where name is any and -d means as a background service
8. $ docker ps, to check if the container is running https://prnt.sc/1xi6fz5
9. $ docker logs, then check the password generated https://prnt.sc/1xi6yto
10. $ docker exec -it [container name] bash
11. bash#: mysql -uroot -p, then enter the generated password from #9
12. mysql>: ALTER USER 'root'@'localhost' IDENTIFIED BY [new password];
13. That's it, your mysql docker instance for golang is setup and running. Now create the database using queries.
https://phoenixnap.com/kb/mysql-docker-container

*/

package main

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "root:Passw0rd1!@(127.0.0.1:3306)/golang?parseTime=true") // (mysql, "username:password@(127.0.0.1:3306)/dbname?parseTime=true"
    if err != nil {
        log.Fatal(err)
    }
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }

    { // Create a new table
        query := `
            CREATE TABLE users (
                id INT AUTO_INCREMENT,
                username TEXT NOT NULL,
                password TEXT NOT NULL,
                created_at DATETIME,
                PRIMARY KEY (id)
            );`

        if _, err := db.Exec(query); err != nil {
            log.Fatal(err)
        }
    }

    { // Insert a new user
        username := "johndoe"
        password := "secret"
        createdAt := time.Now()

        result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
        if err != nil {
            log.Fatal(err)
        }

        id, err := result.LastInsertId()
        fmt.Println(id)
    }

    { // Query a single user
        var (
            id        int
            username  string
            password  string
            createdAt time.Time
        )

        query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
        if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
            log.Fatal(err)
        }

        fmt.Println(id, username, password, createdAt)
    }

    { // Query all users
        type user struct {
            id        int
            username  string
            password  string
            createdAt time.Time
        }

        rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
        if err != nil {
            log.Fatal(err)
        }
        defer rows.Close()

        var users []user
        for rows.Next() {
            var u user

            err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
            if err != nil {
                log.Fatal(err)
            }
            users = append(users, u)
        }
        if err := rows.Err(); err != nil {
            log.Fatal(err)
        }

        fmt.Printf("%#v", users)
    }

    {
        _, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
        if err != nil {
            log.Fatal(err)
        }
    }
}