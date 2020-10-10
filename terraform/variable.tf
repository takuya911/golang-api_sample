# よく変更する部分
variable "gql_img" {
  default = "gcr.io/golang-portfolio/golang-api_sample/graphql@sha256:a2cf3d76928d5409ec1a17fc7c8afb83e5fd72fbce368cab4cb6b09376acf480"
}

variable "grpc_img" {
  default = "gcr.io/golang-portfolio/golang-api_sample/grpc@sha256:c131955494d11a5df15855586a82dc679fe7f8a48312a6605a8c63b0e7a403a8"
}

# 基本変更しない部分
variable "project_name" {
  default = "golang-portfolio"
}

variable "zone_tokyo" {
  default = "asia-northeast1-a"
}

variable "region_tokyo" {
  default = "asia-northeast1"
}

variable "gql_svc_port" {
  default = "80"
}

variable "grpc_svc_name" {
  default="grpc-repftyfivq-an.a.run.app"
}

variable "grpc_svc_port" {
  default="443"
}