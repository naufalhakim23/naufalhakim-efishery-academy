job "hakim-echo" {
  datacenters = ["dc1"]
  type = "service"

  group "web" {
    count = 1

    network {
      port "http" {
        to = 1323
      }
    }

    task "hakim-echo" {
      driver = "docker"

      config {
        image = "hkimm/go-echo:v1"
        ports = ["http"]
      }

      resources {
        cpu    = 100
        memory = 128
      }
    }

    service {
      name = "hakim-echo"
      port = "http"
      tags = [
        "traefik.enable=true",
        "traefik.http.routers.hakim-echo-demo.rule=Host(\"hakim.cupang.efishery.ai\")",
      ]
      check {
        port        = "http"
        type        = "tcp"
        interval    = "15s"
        timeout     = "14s"
      }
    }

  }
}