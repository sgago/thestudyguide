
# Terraform AWS VPC
# Declare a VPC, a logical container for your AWS resources.
# https://www.terraform.io/docs/providers/aws/r/vpc.html
resource "aws_vpc" "tsg_vpc" {
  # AWS VPC and subnets make liberal use of CIDR notation
  # CIDR notation is a compact representation of an IP address and its associated network mask
  /*
  ipcalc 10.0.0.0/16
    Address:   10.0.0.0             00001010.00000000. 00000000.00000000
    Netmask:   255.255.0.0 = 16     11111111.11111111. 00000000.00000000
    Wildcard:  0.0.255.255          00000000.00000000. 11111111.11111111
    =>
    Network:   10.0.0.0/16          00001010.00000000. 00000000.00000000
    HostMin:   10.0.0.1             00001010.00000000. 00000000.00000001
    HostMax:   10.0.255.254         00001010.00000000. 11111111.11111110
    Broadcast: 10.0.255.255         00001010.00000000. 11111111.11111111
    Hosts/Net: 65534                 Class A, Private Internet
  */

  cidr_block = "10.0.0.0/16" # 65534 IPs
  enable_dns_support = true # Enable DNS support in the VPC, so that instances can resolve public DNS hostnames to IP addresses
  enable_dns_hostnames = true # Enable DNS hostnames in the VPC, so that instances can get a public DNS hostname

  # Optional
  #instance_tenancy = "default" # Marks this VPC as the default VPC for the AWS account
  enable_network_address_usage_metrics = true # Enable network address usage metrics in the VPC, helps us monitor IP address usage

  tags = {
    Name = "tsg-vpc"
  }
}