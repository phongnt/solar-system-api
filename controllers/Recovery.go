package controllers

import (
    "fmt"

    "github.com/gin-gonic/gin"
    "github.com/gocarina/gocsv"
    "mime/multipart"
    "topcoder.com/skill-builder/golang/db"
    "topcoder.com/skill-builder/golang/models"
)

type RecoveryController struct{}

type BodyData struct {
    ID              int64     `csv:"id"`
    Type            string    `csv:"type"`
    Name            string    `csv:"name"`
    Description     string    `csv:"description"`
    NumberOfMoons   int16     `csv:"number_of_moons"`
}

//TODO: Easy-250: Add fields to these structs
type PhysicalCharacteristicData struct {}
type OrbitalParameterData struct {}

func (r *RecoveryController) RecoverData(c *gin.Context) {
    var bodyFile /*, physicalFile, orbitalFile */ *multipart.FileHeader
    var err error

    if bodyFile, err = c.FormFile("body"); err != nil {
        c.JSON(400, gin.H{
            "message": "CSV file 'body' is required",
        })
        return
    }

    // if physicalFile, err := c.FormFile("physicalData"); err != nil {
    //    c.JSON(400, gin.H{
    //        "message": "CSV file 'physicalData' is required",
    //    })
    //    return
    // }

    // if orbitalFile, err := c.FormFile("orbitalData"); err != nil {
    //    c.JSON(400, gin.H{
    //        "message": "CSV file 'orbitalData' is required",
    //    })
    //    return
    // }

    bodyData, _ := bodyFile.Open()

    // physicalCharacteristicData, _ := physicalFile.Open()
    // orbitalData, _ := orbitalFile.Open()

    defer bodyData.Close()

    // defer physicalCharacteristicData.Close()
    // defer orbitalData.Close()

    bodies := make([]*BodyData, 0)

    if err := gocsv.Unmarshal(bodyData, &bodies); err != nil {
        panic(err)
    }

    for _, body := range bodies {
        model := &models.Body{
            ID: body.ID,
            Type: body.Type,
            Name: body.Name,
            Description: body.Description,
            Moons: body.NumberOfMoons,
        }

        fmt.Println(model)

        //TODO: Handle error - what if the ID already exists?
        db.GetDBObject().Model(model).Insert()
    }

    //TODO: Easy-250: Insert PhysicalCharacteristics data into database
    // physicalCharacteristics := make([]*PhysicalCharacteristicData, 0)
    // for physicalCharacteristic, _ := range physicalCharacteristics {
    // }

    //TODO: Easy-250: Insert OrbitalParameters data into database
    // orbitalParams := make([] *OrbitalParameterData, 0)
    // for orbitalParam, _ := range orbitalParams {
    // }

    c.JSON(200, gin.H{
        "message": "Recovery successful",
    })
}
