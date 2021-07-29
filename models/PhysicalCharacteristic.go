package models

type PhysicalCharacteristic struct {
    tableName           struct{}     `pg:"physical_data"`
    ID                  int64        `pg:"id,pk"`
    BodyID              int64
    Density             float32      `pg:",use_zero"`
    Gravity             float32      `pg:",use_zero"`
    MassValue           float32      `pg:",use_zero"`
    MassExponent        int16        `pg:",use_zero"`
    VolumeValue         float32      `pg:",use_zero"`
    VolumeExponent      int16        `pg:",use_zero"`
}
