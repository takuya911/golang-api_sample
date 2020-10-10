# Graphql Service
resource "google_cloud_run_service" "graphql" {
  name     = "graphql"
  location = var.region_tokyo
  template {
    spec {
      container_concurrency = var.gql_svc_port
      timeout_seconds       = 30
      containers {
        image = var.gql_img
        env {
          name  = "GRAPHQL_SERVICE_PORT"
          value = var.gql_svc_port
        }
        env {
          name  = "USER_SERVICE_NAME"
          value = var.grpc_svc_name
        }
        env {
          name  = "USER_SERVICE_PORT"
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
