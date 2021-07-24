package main

import (
    "topcoder.com/skill-builder/golang/db"
    "topcoder.com/skill-builder/golang/server"
)

func main() {
    db.Init()
    server.Init()
}
