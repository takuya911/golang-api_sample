# Graphql Service
resource "google_cloud_run_service" "graphql" {
  name     = "graphql"
  location = var.region_tokyo
  template {
    spec {
      timeout_seconds       = 30
      containers {
        image = var.gql_img
        env {
          name  = "GRAPHQL_SERVICE_PORT"
          value = var.gql_svc_port
        }
        env {
          name  = "GRPC_SERVICE_NAME"
          value = var.grpc_svc_name
        }
        env {
          name  = "GRPC_SERVICE_PORT"
          value = var.grpc_svc_port
        }
        resources {
          limits = {
            "cpu" : "1000m",
            "memory" : "128Mi"
          }
        }
        ports {
          container_port = var.gql_svc_port
        }
      }
    }
  }
  traffic {
    percent         = 100
    latest_revision = true
  }
}

# gRPC Service
resource "google_cloud_run_service" "grpc" {
  name     = "grpc"
  location = var.region_tokyo
  template {
    spec {
      timeout_seconds       = 30
      containers {
        image = var.grpc_img
        env {
          name  = "GRPC_SERVICE_PORT"
          value = var.grpc_svc_port
        }
        resources {
          limits = {
            "cpu" : "1000m",
            "memory" : "128Mi"
          }
        }
        ports {
          container_port = var.grpc_svc_port
        }
      }
    }
  }
  traffic {
    percent         = 100
    latest_revision = true
  }
}