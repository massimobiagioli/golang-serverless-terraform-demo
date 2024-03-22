provider "aws" {
  region = "eu-west-1"
}

data "archive_file" "lambda_main_zip" {
  type        = "zip"
  source_file = "${path.module}/../bin/bootstrap"
  output_path = "${path.module}/../bin/bootstrap.zip"
}
