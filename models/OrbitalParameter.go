package models

type OrbitalParameter struct {
    tableName           struct{}     `pg:"orbital_parameters"`
    ID                  int64        `pg:"id,pk"`
    BodyID              int64
    SideralOrbit        float32      `pg:",use_zero"`
    SideralRotation     float32      `pg:",use_zero"`
}
