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

resource "google_bigquery_dataset" "facebook_ads_dataset" {
  dataset_id  = "dw_dev_staging_facebook_ads_historical"
  description = "This is a test description"
  location    = var.gcp_region
  labels = {
    env = "default"
  }
}

resource "google_bigquery_table" "ad_accounts_table" {
  dataset_id = google_bigquery_dataset.facebook_ads_dataset.dataset_id
  table_id   = "ad_accounts"

  time_partitioning {
    type = "DAY"
  }

  labels = {
    env = "default"
  }

  schema              = file("${path.module}/schemas/ad_accounts.json")
  deletion_protection = false
}

resource "google_bigquery_table" "ad_sets_table" {
  dataset_id = google_bigquery_dataset.facebook_ads_dataset.dataset_id
  table_id   = "adsets"

  time_partitioning {
    type = "DAY"
  }

  labels = {
    env = "default"
  }

  schema              = file("${path.module}/schemas/adsets.json")
  deletion_protection = false
}

resource "google_bigquery_table" "campaigns_table" {
  dataset_id = google_bigquery_dataset.facebook_ads_dataset.dataset_id
  table_id   = "campaigns"

  time_partitioning {
    type = "DAY"
  }

  labels = {
    env = "default"
  }

  schema              = file("${path.module}/schemas/campaigns.json")
  deletion_protection = false
}

resource "google_bigquery_table" "ads_table" {
  dataset_id = google_bigquery_dataset.facebook_ads_dataset.dataset_id
  table_id   = "ads"

  time_partitioning {
    type = "DAY"
  }

  labels = {
    env = "default"
  }

  schema              = file("${path.module}/schemas/ads.json")
  deletion_protection = false
}

resource "google_bigquery_table" "adcreatives_table" {
  dataset_id = google_bigquery_dataset.facebook_ads_dataset.dataset_id
  table_id   = "adcreatives"

  time_partitioning {
    type = "DAY"
  }

  labels = {
    env = "default"
  }

  schema              = file("${path.module}/schemas/adcreatives.json")
  deletion_protection = false
}

resource "google_bigquery_table" "ad_insights_table" {
  dataset_id = google_bigquery_dataset.facebook_ads_dataset.dataset_id
  table_id   = "ad_insights"

  time_partitioning {
    type = "DAY"
  }

  labels = {
    env = "default"
  }

  schema              = file("${path.module}/schemas/ad_insights.json")
  deletion_protection = false
}

resource "google_bigquery_table" "user_lead_gen_info_table" {
  dataset_id = google_bigquery_dataset.facebook_ads_dataset.dataset_id
  table_id   = "user_lead_gen_info"

  time_partitioning {
    type = "DAY"
  }

  labels = {
    env = "default"
  }

  schema              = file("${path.module}/schemas/user_lead_gen_info.json")
  deletion_protection = false
}
