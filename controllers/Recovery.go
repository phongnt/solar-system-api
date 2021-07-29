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

type PhysicalCharacteristicData struct {
    ID              int64     `csv:"id"`
    BodyID          int64     `csv:"body_id"`
    Density         float32   `csv:"density"`
    Gravity         float32   `csv:"gravity"`
    MassValue       float32   `csv:"mass_value"`
    MassExponent    int16     `csv:"mass_exponent"`
    VolumeValue     float32   `csv:"volume_value"`
    VolumeExponent  int16     `csv:"volume_exponent"`
}

type OrbitalParameterData struct {
    ID              int64     `csv:"id"`
    BodyID          int64     `csv:"body_id"`
    SideralOrbit    float32   `csv:"sideral_orbit"`
    SideralRotation float32   `csv:"sideral_rotation"`
}

func (r *RecoveryController) RecoverData(c *gin.Context) {
    var bodyFile, physicalFile, orbitalFile *multipart.FileHeader
    var err error

    if bodyFile, err = c.FormFile("body"); err != nil {
        c.JSON(400, gin.H{
            "message": "CSV file 'body' is required",
        })
        return
    }

    if physicalFile, err = c.FormFile("physicalData"); err != nil {
       c.JSON(400, gin.H{
           "message": "CSV file 'physicalData' is required",
       })
       return
    }

    if orbitalFile, err = c.FormFile("orbitalData"); err != nil {
       c.JSON(400, gin.H{
           "message": "CSV file 'orbitalData' is required",
       })
       return
    }

    bodyData, _ := bodyFile.Open()

    physicalCharacteristicData, _ := physicalFile.Open()
    orbitalData, _ := orbitalFile.Open()

    defer bodyData.Close()

    defer physicalCharacteristicData.Close()
    defer orbitalData.Close()

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

    physicalCharacteristics := make([]*PhysicalCharacteristicData, 0)
    if err := gocsv.Unmarshal(physicalCharacteristicData, &physicalCharacteristics); err != nil {
        panic(err)
    }
    for _, physicalCharacteristic := range physicalCharacteristics {
        model := &models.PhysicalCharacteristic{
            ID: physicalCharacteristic.ID,
            BodyID: physicalCharacteristic.BodyID,
            Density: physicalCharacteristic.Density,
            Gravity: physicalCharacteristic.Gravity,
            MassValue: physicalCharacteristic.MassValue,
            MassExponent: physicalCharacteristic.MassExponent,
            VolumeValue: physicalCharacteristic.VolumeValue,
            VolumeExponent: physicalCharacteristic.VolumeExponent,
        }

        fmt.Println(model)

        //TODO: Handle error - what if the ID already exists?
        _, err = db.GetDBObject().Model(model).Insert()
        if err != nil {
            panic(err)
        }
    }

    orbitalParams := make([] *OrbitalParameterData, 0)
    if err := gocsv.Unmarshal(orbitalData, &orbitalParams); err != nil {
        panic(err)
    }
    for _, orbitalParam := range orbitalParams {
        model := &models.OrbitalParameter{
            ID: orbitalParam.ID,
            BodyID: orbitalParam.BodyID,
            SideralOrbit: orbitalParam.SideralOrbit,
            SideralRotation: orbitalParam.SideralRotation,
        }

        fmt.Println(model)

        //TODO: Handle error - what if the ID already exists?
        _, err = db.GetDBObject().Model(model).Insert()
        if err != nil {
            panic(err)
        }
    }

    c.JSON(200, gin.H{
        "message": "Recovery successful",
    })
}
