
# Create a public subnet
# A public subnet is a subnet that has a route to the internet.
# We put our public-facing resources in this subnet: load balancers, NAT gateways, Elastic IP, etc.
# https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html
resource "aws_subnet" "public_subnet_a" {
  vpc_id = aws_vpc.tsg_vpc.id # Reference the VPC we created in a previous step
  cidr_block = "10.0.0.0/24" # 254 IPs
  availability_zone = "us-east-1a" # Availability zone is a physical location within a region

  # Assign a public IP address to instances launched in this subnet
  # This is what makes this a public subnet
  map_public_ip_on_launch = true

  tags = {
    Name = "tsg-public-subnet-a"
  }
}

# We'll declare a second public subnet in a different availability zone.
# Certain resources, like load balancers, require multiple subnets in different availability zones.
resource "aws_subnet" "public_subnet_b" {
  vpc_id = aws_vpc.tsg_vpc.id
  cidr_block = cidrsubnet("10.0.0.0/16", 8, 1) # 254 IPs, another way to declare a subnet cidr block
  availability_zone = "us-east-1b"
  map_public_ip_on_launch = true

  tags = {
    Name = "tsg-public-subnet-b"
  }
}

# Create a private subnet
# A private subnet is a subnet that does not have a route to the internet.
# We put our private resources in this subnet: EC2 instances, RDS databases, etc.
resource "aws_subnet" "private_subnet_a" {
  vpc_id = aws_vpc.tsg_vpc.id
  cidr_block = cidrsubnet(aws_vpc.tsg_vpc.cidr_block, 8, 2) # Yet another way
  availability_zone = "us-east-1a"

  tags = {
    Name = "tsg-private-subnet-a"
  }
}

resource "aws_subnet" "private_subnet_b" {
  vpc_id = aws_vpc.tsg_vpc.id
  cidr_block = cidrsubnet(aws_vpc.tsg_vpc.cidr_block, 8, 3)
  availability_zone = "us-east-1b"

  tags = {
    Name = "tsg-private-subnet-b"
  }
}