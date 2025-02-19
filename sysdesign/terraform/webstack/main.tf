terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }

  required_version = ">= 1.2.0"
}

# TODO: CloudWatch, Lambda (for functions), Route53, Code Pipeline, EKS, KMS, etc.

# Provider configuration for AWS cloud provider
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs
provider "aws" {
  region = "us-east-1"

  # Optional: add access_key and secret_key for authentication
  # Optional: add profile, shared_credentials_file, and shared_credentials_profile for shared credentials file
}

# CIDR (Classless Inter-Domain Routing) block for the VPC come up a lot
# https://aws.amazon.com/what-is/cidr/#:~:text=A%20CIDR%20block%20is%20a,regional%20internet%20registries%20(RIR).
# CIDR blocks are like 10.0.0.0/16 which represent a range of IP addresses
# The /16 means the first 16 bits are fixed and the remaining 32-16=16 bits can vary
# We'll see 0.0.0.0/0 which represents all IP addresses. You'll see this in ingress/egress rules.
# CIDR blocks are used to define the IP address range for the VPC
# This is opposed to traditional IP address classes (e.g., Class A, B, C)

# Create a Virtual Private Cloud (VPC) with a /16 CIDR block
# A VPC is a virtual network in the cloud that you can launch your instances into
# It is a logical isolation of the AWS cloud dedicated to your account
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc
resource "aws_vpc" "main" {
  cidr_block           = "10.0.0.0/16" # About 65536 IP addresses
  enable_dns_support   = true          # Enable internal DNS resolution for the VPC
  enable_dns_hostnames = true          # Enable DNS hostnames for the VPC

  tags = {
    Name = "main-vpc"
  }
}

# Create a public subnet in the VPC
# A subnet is a range of IP addresses in your VPC
# Subnets can be public (accessible from the internet) or private (not accessible from the internet)
# Public subnets typically have an internet gateway (IGW) for internet access
# Resources in the public subnet include load balancers, web servers, internet gateway, elastic IP, etc.
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/subnet
resource "aws_subnet" "public_subnet_1" {
  vpc_id                  = aws_vpc.main.id # Reference the VPC created above
  cidr_block              = "10.0.1.0/24"   # About 256 IP addresses
  availability_zone       = "us-east-1a"    # AZ for the subnet
  map_public_ip_on_launch = true            # Enable public IP addresses for instances in this subnet

  tags = {
    Name = "public-sn-1"
  }
}

# Create another public subnet in the VPC
# We need at least two public subnets for the Application Load Balancer (ALB)
resource "aws_subnet" "public_subnet_2" {
  vpc_id                  = aws_vpc.main.id # Reference the VPC created above
  cidr_block              = "10.0.2.0/24"   # About 256 IP addresses
  availability_zone       = "us-east-1b"    # AZ for the subnet
  map_public_ip_on_launch = true            # Enable public IP addresses for instances in this subnet

  tags = {
    Name = "public-sn-2"
  }
}

# Create a private subnet in the VPC
# This subnet is not accessible from the internet
# Resources in the private subnet include databases, KV stores, application servers, etc.
resource "aws_subnet" "private_subnet_1" {
  vpc_id                  = aws_vpc.main.id # Reference the VPC created above
  cidr_block              = "10.0.3.0/24"   # About 256 IP addresses
  availability_zone       = "us-east-1a"    # AZ for the subnet
  map_public_ip_on_launch = false           # Disable public IP addresses for instances in this subnet

  tags = {
    Name = "private-sn-1"
  }
}

# Create a private subnet in the VPC with a /24 CIDR block
resource "aws_subnet" "private_subnet_2" {
  vpc_id                  = aws_vpc.main.id # Reference the VPC created above
  cidr_block              = "10.0.4.0/24"   # About 256 IP addresses
  availability_zone       = "us-east-1b"    # AZ for the subnet
  map_public_ip_on_launch = false           # Disable public IP addresses for instances in this subnet

  tags = {
    Name = "private-sn-2"
  }
}

# Create an internet gateway for the VPC
# This allows the VPC to connect to the internet
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/internet_gateway
resource "aws_internet_gateway" "internet_gateway" {
  vpc_id = aws_vpc.main.id # Reference the VPC created above

  tags = {
    Name = "main-igw"
  }
}

# Create an Elastic IP (EIP) for the NAT gateway
# This is a static IP address that can be attached to the NAT gateway
# An elastic IP address is a static IPv4 address designed for dynamic cloud computing
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/eip
resource "aws_eip" "nat_eip" {
  network_border_group = "us-east-1" # Allocate the Elastic IP in the specified region

  tags = {
    Name = "nat-eip"
  }
}

# Create a NAT gateway in the public subnet
# This allows instances in the private subnet to access the internet
# However, no incoming connections are allowed to the private subnet
# This is helpful for updates, package installations, etc. without exposing the private subnet
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/nat_gateway
resource "aws_nat_gateway" "nat_gateway" {
  allocation_id = aws_eip.nat_eip.id            # Reference the Elastic IP created above
  subnet_id     = aws_subnet.public_subnet_1.id # Reference the public subnet created above

  tags = {
    Name = "nat-gw"
  }
}

# Create a public route table for the public subnet
# This routes all traffic to the internet gateway
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route_table
resource "aws_route_table" "public_route_table" {
  vpc_id = aws_vpc.main.id # Reference the VPC created above

  route {
    cidr_block = "0.0.0.0/0"                              # Route all traffic
    gateway_id = aws_internet_gateway.internet_gateway.id # To the internet gateway
  }

  tags = {
    Name = "public-rt"
  }
}

# Associate the public route table with the public subnet
# This allows the public subnet to route traffic to the internet gateway
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route_table_association
resource "aws_route_table_association" "public_route_table_association_1" {
  subnet_id      = aws_subnet.public_subnet_1.id         # Reference the public subnet created above
  route_table_id = aws_route_table.public_route_table.id # Reference the route table created above
}

# Associate another public route table with the second public subnet
resource "aws_route_table_association" "public_route_table_association_2" {
  subnet_id      = aws_subnet.public_subnet_2.id         # Reference the public subnet created above
  route_table_id = aws_route_table.public_route_table.id # Reference the route table created above
}

# Create a private route table for the private subnet
# This routes all traffic to the NAT gateway
resource "aws_route_table" "private_route_table" {
  vpc_id = aws_vpc.main.id # Reference the VPC created above

  route {
    cidr_block     = "0.0.0.0/0"                    # Route all traffic
    nat_gateway_id = aws_nat_gateway.nat_gateway.id # To the NAT gateway
  }

  tags = {
    Name = "private-rt"
  }
}

# Associate the private route table with the private subnet
# This allows the private subnet to route traffic to the NAT gateway
resource "aws_route_table_association" "private_route_table_association_1" {
  subnet_id      = aws_subnet.private_subnet_1.id         # Reference the private subnet created above
  route_table_id = aws_route_table.private_route_table.id # Reference the route table created above
}

resource "aws_route_table_association" "private_route_table_association_2" {
  subnet_id      = aws_subnet.private_subnet_2.id         # Reference the private subnet created above
  route_table_id = aws_route_table.private_route_table.id # Reference the route table created above
}

# Create a security group for the Application Load Balancer (ALB)
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lb
resource "aws_lb" "app_load_balancer" {
  name               = "app-lb"
  internal           = false         # The load balancer is public facing
  load_balancer_type = "application" # The type of load balancer, a layer 7 application load balancer
  security_groups    = [aws_security_group.alb_sg.id]
  subnets            = [aws_subnet.public_subnet_1.id, aws_subnet.public_subnet_2.id] # Reference the public subnets

  enable_deletion_protection = false # Disable deletion protection for the load balancer, for demo purposes

  tags = {
    Name = "app-lb"
  }
}

# Create a security group for the Application Load Balancer (ALB)
# This allows HTTP and HTTPS traffic to the load balancer
# Security groups are like firewalls for your instances. They control inbound and outbound traffic.
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group
resource "aws_security_group" "alb_sg" {
  vpc_id = aws_vpc.main.id # Reference the VPC created above

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "TCP"
    cidr_blocks = ["0.0.0.0/0"] # Allow all HTTP traffic on port 80
  }

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "TCP"
    cidr_blocks = ["0.0.0.0/0"] # Allow all HTTPS traffic on port 443
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"          # Allow all outbound traffic
    cidr_blocks = ["0.0.0.0/0"] # Allow all outbound traffic
  }

  tags = {
    Name = "alb-sg"
  }
}

# Create a target group for the Application Load Balancer
# This defines the target for the load balancer, e.g., EC2 instances, IP addresses, etc.
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lb_target_group
resource "aws_lb_target_group" "app_target_group" {
  name     = "app-tg"
  port     = 80
  protocol = "HTTP"
  vpc_id   = aws_vpc.main.id # Reference the VPC created above

  health_check {
    path                = "/health" # The URL path to check for health
    protocol            = "HTTP"    # The protocol to use for health checks
    port                = "80"      # The port to use for health checks
    interval            = 30        # Health check interval, every 30 seconds
    timeout             = 5         # Health check timeout, 5 seconds
    healthy_threshold   = 2         # Number of consecutive successful health checks to consider the target healthy
    unhealthy_threshold = 2         # Number of consecutive failed health checks to consider the target unhealthy
  }

  tags = {
    Name = "app-tg"
  }
}

# Create a listener for the load balancer for HTTP traffic on port 80
# The listener forwards traffic to the target group
resource "aws_lb_listener" "http_listener" {
  load_balancer_arn = aws_lb.app_load_balancer.arn # Reference the load balancer created above
  port              = 80
  protocol          = "HTTP"

  default_action {
    type = "fixed-response"
    fixed_response {
      status_code  = "200"
      content_type = "text/plain"
      message_body = "OK"
    }
  }
}

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

# Register EC2 instances to the target group
# This allows the load balancer to route traffic to the EC2 instances.
# Attachments are used to associate a target with a target group.
# They are used to connect IAM, networking, and storage resources to AWS components
# like EC2 instances, load balancers, etc.
# In this case, we're targeting the EC2 instance with the load balancer.
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lb_target_group_attachment
resource "aws_lb_target_group_attachment" "app_target_group_attachment" {
  target_group_arn = aws_lb_target_group.app_target_group.arn
  target_id        = aws_instance.my_ec2.id # Reference your EC2 instance.
  port             = 80                     # Port the EC2 instance is listening on.
}

# Create a subnet group for the RDS database.
# Subnet groups are collections of subnets that define the network where the RDS database will be created.
# In this case, it defines the subnets where the RDS database will be created.
# Typically, AWS services with high availability and failover require multiple subnets: RDS, ElastiCache, Redshift, etc.
# Our subnets are in different availability zones.
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db_subnet_group
resource "aws_db_subnet_group" "db_subnet_group" {
  name       = "db-subnet-group"
  subnet_ids = [aws_subnet.private_subnet_1.id, aws_subnet.private_subnet_2.id] # Reference the private subnets

  tags = {
    Name = "MyDBSubnetGroup"
  }
}

# Create a security group for the RDS database
# This allows inbound traffic on port 5432 (PostgreSQL default port).
resource "aws_security_group" "db_sg" {
  vpc_id      = aws_vpc.main.id
  name        = "db-security-group"
  description = "Allow inbound traffic to RDS"

  ingress {
    from_port   = 5432 # PostgreSQL default port
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = ["10.0.0.0/16"] # Allow traffic from the VPC (adjust as needed)
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"] # Allow all outbound traffic
  }

  tags = {
    Name = "db-sg"
  }
}

resource "aws_db_instance" "my_db" {
  allocated_storage      = 20                            # Storage in GB
  storage_type           = "gp2"                         # General Purpose SSD storage
  engine                 = "postgres"                    # Database engine (PostgreSQL)
  engine_version         = "17.3"                        # PostgreSQL version
  instance_class         = "db.t4g.micro"                # Instance type (Free Tier eligible)
  db_name                = "mydatabase"                  # Database name
  username               = "thestudyguide"               # Master username
  password               = "mypassword"                  # Master password (preferably use secrets manager)
  parameter_group_name   = "default.postgres17"          # Default parameter group for PostgreSQL 17
  multi_az               = false                         # Single Availability Zone
  publicly_accessible    = false                         # Not publicly accessible
  skip_final_snapshot    = true                          # Skip final snapshot on deletion
  vpc_security_group_ids = [aws_security_group.db_sg.id] # Attach security group

  db_subnet_group_name = aws_db_subnet_group.db_subnet_group.name # Reference subnet group

  deletion_protection = false # Disable deletion protection for the database, for demo purposes

  tags = {
    Name = "MyRDSInstance"
  }

  # Enable automated backups
  backup_retention_period = 7 # Days to retain backups
}

resource "aws_security_group" "redis_sg" {
  vpc_id      = aws_vpc.main.id
  name        = "redis-security-group"
  description = "Allow inbound traffic to ElastiCache"

  ingress {
    from_port       = 6379 # Redis default port
    to_port         = 6379
    protocol        = "tcp"
    security_groups = [aws_security_group.ec2_sg.id] # Allow traffic from the EC2 security group
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"          # Allow all outbound traffic
    cidr_blocks = ["0.0.0.0/0"] # Allow all outbound traffic
  }
}

# Create a Redis ElastiCache subnet group for the private subnets
resource "aws_elasticache_subnet_group" "redis_subnet_group" {
  name        = "redis-subnet-group"
  description = "Subnet group for ElastiCache"
  subnet_ids  = [aws_subnet.private_subnet_1.id, aws_subnet.private_subnet_2.id] # Reference the private subnets
}

# Create a Redis ElastiCache cluster
resource "aws_elasticache_cluster" "redis" {
  cluster_id           = "my-redis-cluster"
  engine               = "redis"
  node_type            = "cache.t3.micro"                                     # Node type
  port                 = 6379                                                 # Redis default port
  num_cache_nodes      = 1                                                    # Number of cache nodes, 1 for free tier
  subnet_group_name    = aws_elasticache_subnet_group.redis_subnet_group.name # Reference the subnet group
  security_group_ids   = [aws_security_group.redis_sg.id]                     # Attach the security group
  parameter_group_name = "default.redis7"                                     # Default parameter group for Redis 7
}

# Create a security group for the EC2 instances
# This allows SSH and HTTP traffic
# You may want to restrict this in production
resource "aws_security_group" "ec2_sg" {
  vpc_id      = aws_vpc.main.id
  name        = "ec2-security-group"
  description = "Allow SSH and HTTP traffic"

  # ingress block allows incoming traffic
  # from_port and to_port define the port range
  # protocol defines the protocol (TCP in this case)
  # port 22 is for SSH, port 80 is for HTTP, port 443 for HTTPS
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # Allow SSH from anywhere (you may want to restrict this in production)
  }

  # This ingress rule allows HTTP traffic on port 80.
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # Allow HTTP traffic from anywhere
  }

  # Egress block allows outgoing traffic.
  # This rule allows all outbound traffic to anywhere using any protocol.
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1" # Allow all outbound traffic
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "ec2-sg"
  }
}

# Create an IAM role for the EC2 instance
# This role defines who--which services--can assume the role. In this case, only EC2 instances.
# IAM roles are used to delegate permissions to entities that you trust.
# This role, must be attached to something! In this case, the EC2 instance.
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role
resource "aws_iam_role" "ec2_s3_role" {
  name = "ec2-s3-role"

  # Defines the trust policy for the role
  # The trust policy allows the EC2 instance to assume the role
  assume_role_policy = jsonencode({
    Version = "2012-10-17", # Version of the policy language
    Statement = [
      {
        Effect = "Allow", # Allows an entity to assume the role, in this case, the EC2 instance
        Principal = {
          Service = "ec2.amazonaws.com" # Only EC2 instances can assume this role
        },
        Action = "sts:AssumeRole" # Allow assuming the role, using AWS Security Token Service (STS)
      }
    ]
  })
}

# Attach the AmazonS3FullAccess policy to the IAM role
# This policy allows full access to Amazon S3.
# Without this, the EC2 instance won't be able to access S3.
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_policy_attachment
resource "aws_iam_policy_attachment" "s3_attachment" {
  name = "s3-access-attachment"
  roles = [aws_iam_role.ec2_s3_role.name] # Attach the policy to the IAM role
  policy_arn = "arn:aws:iam::aws:policy/AmazonS3FullAccess" # Full S3 access, this defines the permissions
}

# Create an instance profile for the EC2 instance
# An instance profile is a container for an IAM role that you can
# use to pass role information to an EC2 instance when the instance starts.
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_instance_profile
resource "aws_iam_instance_profile" "ec2_instance_profile" {
  name = "ec2-instance-profile"
  role = aws_iam_role.ec2_s3_role.name # Attach the IAM role to the instance profile
}

resource "aws_instance" "my_ec2" {
  ami           = "ami-0533b66ceb689fe0b"        # Replace with your desired AMI ID
  instance_type = "t2.micro"                     # Instance type
  subnet_id     = aws_subnet.private_subnet_1.id # Reference the private subnet
  #key_name        = aws_key_pair.ec2_key_pair.key_name  # Reference the key pair
  vpc_security_group_ids = [aws_security_group.ec2_sg.id] # Attach the security group
  iam_instance_profile = aws_iam_instance_profile.ec2_instance_profile.name # Attach the instance profile

  tags = {
    Name = "MyEC2Instance"
  }

  # Optional: Enable CloudWatch monitoring
  monitoring = true

  # Optional: Set up user data for EC2 initialization (e.g., installing Apache)
  user_data = <<-EOF
              #!/bin/bash
              yum install -y httpd
              systemctl start httpd
              systemctl enable httpd
              echo "Hello from EC2!" > /var/www/html/index.html
            EOF
}