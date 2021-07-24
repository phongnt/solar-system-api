package models

type Body struct {
    tableName    struct{}     `pg:"body"`
    ID           int64        `pg:"id,pk"`
    Type         string
    Name         string
    Description  string
    Moons        int16        `pg:",use_zero"` // go-pg marshals Go zero values (empty string, 0, zero time, nil map, and nil slice) as SQL NULL; we want to store 0 if number of moons is zero 
}
