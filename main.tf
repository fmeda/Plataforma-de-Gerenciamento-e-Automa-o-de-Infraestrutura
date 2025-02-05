provider "aws" {
  region = "us-east-1"
}

resource "aws_instance" "example" {
  ami           = "ami-0c55b159cbfafe1f0" # ID da AMI específica
  instance_type = "t2.micro"
  
  tags = {
    Name = "MinhaInstancia"
  }
}

output "instance_public_ip" {
  value = aws_instance.example.public_ip
}
