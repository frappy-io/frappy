terraform {
  required_providers {
    docker = {
        source = "kreuzwerker/docker"
        version = ">= 2.13.0"
    }
  }
}

provider "docker" {
  host = ""
}

resource "docker_image" "go-win-app" {
  name          = "go-win-app:latest"
  keep_locally  = false 
}

resource "docker_container" "go-win-app" {
    image = docker_image.go-win-app.latest
    name = "frappy-app"
    ports {
        internal = 80
        externa = 8000
    }
}