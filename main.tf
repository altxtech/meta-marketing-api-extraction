# Variables
variable "gcp_project" {
  description = "Project to deploy to"
  type        = string
}

variable "gcp_region" {
  description = "Default Region to deploy to"
  type        = string
}

variable "state_bucket" {
  description = "The name of the Google Cloud Storage bucket"
  type        = string
}

variable "project_name" {
  description = "Name of the project. Used for namespacing resources"
  type        = string
}

variable "environment" {
  description = "Deployment environment. Used for namespacing an some configurations"
  type        = string
  default     = "dev"
}

provider "google" {
  project = var.gcp_project
  region  = var.gcp_region
}

terraform {
  backend "gcs" {
    bucket = "terraform-states-01432"
  }
}

resource "google_bigquery_dataset" "dataset" {
  dataset_id                  = "facebook_ads_test"
  description                 = "This is a test description"
  location                    = var.gcp_region
  default_table_expiration_ms = 3600000

  labels = {
    env = "default"
  }
}
