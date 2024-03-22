resource "aws_lambda_function" "lambda_main_func" {
  filename         = data.archive_file.lambda_main_zip.output_path
  function_name    = local.app_id
  source_code_hash = base64sha256(data.archive_file.lambda_main_zip.output_path)
  runtime          = local.runtime
  handler          = local.handler
  role             = aws_iam_role.lambda_main_exec.arn

  tags = local.common_tags
}

# Assume role setup
resource "aws_iam_role" "lambda_main_exec" {
  name_prefix = local.app_id

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF

  tags = local.common_tags
}

# Attach role to Managed Policy
variable "iam_policy_arn" {
  description = "IAM Policy to be attached to role"
  type        = list(string)

  default = [
    "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  ]
}

resource "aws_iam_policy_attachment" "role_attach" {
  name       = "policy-${local.app_id}"
  roles      = [aws_iam_role.lambda_main_exec.id]
  count      = length(var.iam_policy_arn)
  policy_arn = element(var.iam_policy_arn, count.index)
}