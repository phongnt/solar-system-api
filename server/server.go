package server

func Init() {
    r := NewRouter()
    _ = r.Run("0.0.0.0:8080")
}
