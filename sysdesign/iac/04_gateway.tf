# Create an Internet Gateway.
# An Internet Gateway is a horizontally scaled, redundant,
# and highly available VPC component that allows communication
# between instances in your VPC and the internet.
# It allows the VPC to connect to the internet.
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/internet_gateway
resource "aws_internet_gateway" "internet_gateway" {
  vpc_id = aws_vpc.tsg_vpc.id

  tags = {
    Name = "tsg-internet-gateway"
  }
}

# Create an Elastic IP (EIP) for the NAT Gateway below.
# An EIP is a static IPv4 address designed for dynamic cloud computing.
# Cloud IPs are not static and can change when you stop and start your instance,
# but an EIP is static and will not change.
# https://docs.aws.amazon.com/vpc/latest/userguide/VPC_NAT_Gateway.html
resource "aws_eip" "nat_eip_a" {
  network_border_group = "us-east-1"

  # OR
  # You can specify a specific availability zone like us-east-1a.
  #network_border_group = "us-east-1a"

  # OR
  # You can specify a specific local zone like Boston.
  #network_border_group = "us-east-1-bos-1a"

  # OR
  # You can specify a specific wavelength zone like Boston.
  #network_border_group = "us-east-1-wl1-bos-wlz-1"

  tags = {
    Name = "tsg-nat-eip-a"
  }
}

resource "aws_eip" "nat_eip_b" {
  network_border_group = "us-east-1"

  tags = {
    Name = "tsg-nat-eip-b"
  }
}


# Create a NAT Gateway.
# A NAT Gateway is a managed NAT service that provides
# outbound internet access to instances and services in a private subnet.
# This can be used to download software updates, access the internet, etc.
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/nat_gateway
resource "aws_nat_gateway" "nat_gateway_a" {
  allocation_id = aws_eip.nat_eip_a.id
  subnet_id = aws_subnet.public_subnet_a.id

  tags = {
    Name = "tsg-nat-gateway-a"
  }
}

# We'll declare a second NAT Gateway in a different availability zone.
# Certain resources, like load balancers, require multiple NAT Gateways in different availability zones.
# Also, this is a good practice for high availability.
resource "aws_nat_gateway" "nat_gateway_b" {
  allocation_id = aws_eip.nat_eip_b.id
  subnet_id = aws_subnet.public_subnet_b.id

  tags = {
    Name = "tsg-nat-gateway-b"
  }
}

# Create a public route table.
# https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html
resource "aws_route_table" "public_route_table" {
  vpc_id = aws_vpc.tsg_vpc.id
  
  route {
    gateway_id = aws_internet_gateway.internet_gateway.id
    cidr_block = "0.0.0.0/0" # This allows all traffic to the internet
  }

  tags = {
    Name = "tsg-public-route-table"
  }
}

resource "aws_route_table" "private_route_table" {
  vpc_id = aws_vpc.tsg_vpc.id

  route {
    nat_gateway_id = aws_nat_gateway.nat_gateway_a.id
    cidr_block = "0.0.0.0/0" # This allows all traffic to the internet
  }

  tags = {
    Name = "tsg-private-route-table"
  }
}

# Create a route table association.
# A route table association is a connection between a route table and a subnet.
# It lets the route table know which subnets to use and we can reuse the same route table for multiple subnets.
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route_table_association
resource "aws_route_table_association" "public_subnet_a" {
  subnet_id = aws_subnet.public_subnet_a.id
  route_table_id = aws_route_table.public_route_table.id
}

resource "aws_route_table_association" "public_subnet_b" {
  subnet_id = aws_subnet.public_subnet_b.id
  route_table_id = aws_route_table.public_route_table.id
}

resource "aws_route_table_association" "private_subnet_a" {
  subnet_id = aws_subnet.private_subnet_a.id
  route_table_id = aws_route_table.private_route_table.id
}

resource "aws_route_table_association" "private_subnet_b" {
  subnet_id = aws_subnet.private_subnet_b.id
  route_table_id = aws_route_table.private_route_table.id
}
