# Solar System API

## How to test

* Build docker image

  ```bash
  docker build -t image-name .
  ```

* Run built image

  ```bash
  docker run -p 8081:8080 --name container-name image-name
  ```

* Verify by health checking

  ```bash
  curl localhost:8081/health
  ```
