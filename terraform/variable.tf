# よく変更する部分
variable "gql_img" {
  default = "gcr.io/golang-portfolio/golang-api_sample/graphql@sha256:089e9fbfb2332b3787f231b78ad8161ea246e48f2e0a1a45ec20ecf66c1b420b"
}

variable "grpc_img" {
  default = "gcr.io/golang-portfolio/golang-api_sample/grpc@sha256:7351a01c0c2cc830e38b23b5c03f0fa897de9b258933d3866ab7b9f0d9660739"
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