
# Create a security group for the load balancer.
# Security groups act as virtual firewalls for your instances to control inbound and outbound traffic.
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group
resource "aws_security_group" "lb_security_group" {
  name = "tsg-lb-security-group"
  vpc_id = aws_vpc.tsg_vpc.id

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "lb-security-group"
  }
}

# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lb
resource "aws_lb" "lb" {
  name = "tsg-lb"
  internal = false
  load_balancer_type = "application"
  security_groups = [aws_security_group.lb_security_group.id]
  subnets = [aws_subnet.public_subnet_a.id, aws_subnet.public_subnet_b.id]

  enable_deletion_protection = false

  tags = {
    Name = "tsg-lb"
  }
}

resource "aws_lb_listener" "lb_http_listener" {
  load_balancer_arn = aws_lb.lb.arn
  port = 80
  protocol = "HTTP"

  default_action {
    type = "fixed-response"

    fixed_response {
      content_type = "text/plain"
      message_body = "HTTP 200 OK"
      status_code  = "200"
    }
  }

  tags = {
    Name = "lb-http-listener"
  }
}

/*
resource "aws_acm_certificate" "lb_https_cert" {
  domain_name = "example.com"
  validation_method = "DNS"

  tags = {
    Name = "lb-https-cert"
  }
}
*/

# Optional: HTTPS listener for secure traffic
/*
resource "aws_lb_listener" "https_listener" {
  load_balancer_arn = aws_lb.app_load_balancer.arn  # Attach the listener to the ALB.
  port              = 443                         # Listen on HTTPS (port 443).
  protocol          = "HTTPS"                      # HTTPS traffic.

  ssl_policy        = "ELBSecurityPolicy-2016-08"  # SSL policy for the listener.
  certificate_arn   = "arn:aws:acm:region:account-id:certificate/certificate-id" # Replace with your SSL certificate ARN.

  default_action {
    type = "fixed-response"
    fixed_response {
      status_code = 200
      content_type = "text/plain"
      message_body = "OK"
    }
  }
}
*/

/*
resource "aws_lb_target_group" "lb_target_group" {
  
}

resource "aws_lb_target_group_attachment" "lb_target_group_attachment" {
  
}
*/