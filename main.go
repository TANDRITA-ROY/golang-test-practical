package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    PhoneNumber string `json:"phone_number"`
    City        string `json:"city"`
    State       string `json:"state"`
    Street1     string `json:"street1"`
    Street2     string `json:"street2"`
    ZipCode     string `json:"zip_code"`
}

func main() {
    router := gin.Default()

    // Database connection
    db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/cetec")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Ensure the connection is available
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }

    // Define routes
    router.GET("/person/:person_id/info", func(c *gin.Context) {
        personID := c.Param("person_id")
        person, err := getPersonInfo(db, personID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, person)
    })

    router.POST("/person/create", func(c *gin.Context) {
        var person Person
        if err := c.ShouldBindJSON(&person); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if err := createPerson(db, &person); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"status": "Person created successfully"})
    })

    router.Run(":8080")
}

func getPersonInfo(db *sql.DB, personID string) (*Person, error) {
    query := `
        SELECT p.name, ph.number, a.city, a.state, a.street1, a.street2, a.zip_code
        FROM person p
        JOIN phone ph ON p.id = ph.person_id
        JOIN address_join aj ON p.id = aj.person_id
        JOIN address a ON aj.address_id = a.id
        WHERE p.id = ?`

    var person Person
    err := db.QueryRow(query, personID).Scan(&person.Name, &person.PhoneNumber, &person.City, &person.State, &person.Street1, &person.Street2, &person.ZipCode)
    if err != nil {
        return nil, err
    }
    return &person, nil
}

func createPerson(db *sql.DB, person *Person) error {
    // Insert person
    personQuery := "INSERT INTO person(name, age) VALUES (?, 0)"
    res, err := db.Exec(personQuery, person.Name)
    if err != nil {
        return err
    }
    personID, err := res.LastInsertId()
    if err != nil {
        return err
    }

    // Insert phone
    phoneQuery := "INSERT INTO phone(person_id, number) VALUES (?, ?)"
    _, err = db.Exec(phoneQuery, personID, person.PhoneNumber)
    if err != nil {
        return err
    }

    // Insert address
    addressQuery := "INSERT INTO address(city, state, street1, street2, zip_code) VALUES (?, ?, ?, ?, ?)"
    res, err = db.Exec(addressQuery, person.City, person.State, person.Street1, person.Street2, person.ZipCode)
    if err != nil {
        return err
    }
    addressID, err := res.LastInsertId()
    if err != nil {
        return err
    }

    // Insert address_join
    addressJoinQuery := "INSERT INTO address_join(person_id, address_id) VALUES (?, ?)"
    _, err = db.Exec(addressJoinQuery, personID, addressID)
    if err != nil {
        return err
    }

    return nil
}
