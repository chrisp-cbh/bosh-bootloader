variable "isolation_segments" {
  type        = "string"
  default     = "0"
  description = "Optionally create a LB and DNS entries for a single isolation segment. Valid values are 0 or 1."
}

variable "iso_to_bosh_ports" {
  type    = "list"
  default = [22,6868,2555,4222,25250]
}

variable "iso_to_shared_tcp_ports" {
  type    = "list"
  default = [9090,9091,8082,8300,8301,8889,8443,3000,4443,8080,3457,9023,9022,4222]
}

variable "iso_to_shared_udp_ports" {
  type    = "list"
  default = [8301,8302,8600]
}

locals {
  iso_az_count    = "${var.isolation_segments > 0 ? length(var.availability_zones) : 0}"
}

resource "aws_subnet" "iso_subnets" {
  count             = "${local.iso_az_count}"
  vpc_id            = "${aws_vpc.vpc.id}"
  cidr_block        = "${cidrsubnet(var.vpc_cidr, 4, count.index + length(var.availability_zones) + 1)}"
  availability_zone = "${element(var.availability_zones, count.index)}"

  tags {
    Name = "${var.env_id}-iso-subnet${count.index}"
  }
}

resource "aws_route_table_association" "route_iso_subnets" {
  count          = "${local.iso_az_count}"
  subnet_id      = "${element(aws_subnet.iso_subnets.*.id, count.index)}"
  route_table_id = "${aws_route_table.internal_route_table.id}"
}

output "iso_az_subnet_id_mapping" {
  value = "${
    zipmap("${aws_subnet.iso_subnets.*.availability_zone}", "${aws_subnet.iso_subnets.*.id}")
  }"
}

output "iso_az_subnet_cidr_mapping" {
  value = "${
    zipmap("${aws_subnet.iso_subnets.*.availability_zone}", "${aws_subnet.iso_subnets.*.cidr_block}")
  }"
}

resource "aws_elb" "iso_router_lb" {
  count = "${var.isolation_segments}"

  name                      = "bbl-iso-router-lb"
  cross_zone_load_balancing = true

  health_check {
    healthy_threshold   = 5
    unhealthy_threshold = 2
    interval            = 12
    target              = "TCP:80"
    timeout             = 2
  }

  listener {
    instance_port     = 80
    instance_protocol = "http"
    lb_port           = 80
    lb_protocol       = "http"
  }

  listener {
    instance_port      = 80
    instance_protocol  = "http"
    lb_port            = 443
    lb_protocol        = "https"
    ssl_certificate_id = "${aws_iam_server_certificate.lb_cert.arn}"
  }

  listener {
    instance_port      = 80
    instance_protocol  = "tcp"
    lb_port            = 4443
    lb_protocol        = "ssl"
    ssl_certificate_id = "${aws_iam_server_certificate.lb_cert.arn}"
  }

  security_groups = ["${aws_security_group.cf_router_lb_security_group.id}"]
  subnets         = ["${aws_subnet.lb_subnets.*.id}"]
}

output "cf_iso_router_lb_name" {
  value = "${aws_elb.iso_router_lb.name}"
}

resource "aws_security_group" "iso_security_group" {
  count = "${var.isolation_segments}"

  description = "iso"
  vpc_id      = "${aws_vpc.vpc.id}"

  ingress {
    self        = true
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port       = 0
    to_port         = 0
    protocol        = "-1"
    security_groups = ["${aws_security_group.bosh_security_group.id}"]
  }

  tags {
    Name = "${var.env_id}-iso-security-group"
  }
}

output "iso_security_group_id" {
  value = "${aws_security_group.iso_security_group.id}"
}

resource "aws_security_group" "iso_shared_security_group" {
  count       = "${var.isolation_segments}"
  description = "iso-shared"
  vpc_id      = "${aws_vpc.vpc.id}"
}

output "iso_shared_security_group_id" {
  value = "${aws_security_group.iso_shared_security_group.id}"
}

resource "aws_security_group_rule" "isolation_segments_to_bosh_rule" {
  count = "${var.isolation_segments * length(var.iso_to_bosh_ports)}"

  security_group_id        = "${aws_security_group.bosh_security_group.id}"
  type                     = "ingress"
  protocol                 = "tcp"
  to_port                  = "${element(var.iso_to_bosh_ports, count.index)}"
  from_port                = "${element(var.iso_to_bosh_ports, count.index)}"
  source_security_group_id = "${aws_security_group.iso_security_group.id}"
}

resource "aws_security_group_rule" "isolation_segments_to_shared_tcp_rule" {
  count = "${var.isolation_segments * length(var.iso_to_shared_tcp_ports)}"

  security_group_id        = "${aws_security_group.iso_shared_security_group.id}"
  type                     = "ingress"
  protocol                 = "tcp"
  to_port                  = "${element(var.iso_to_shared_tcp_ports, count.index)}"
  from_port                = "${element(var.iso_to_shared_tcp_ports, count.index)}"
  source_security_group_id = "${aws_security_group.iso_security_group.id}"
}

resource "aws_security_group_rule" "isolation_segments_to_shared_udp_rule" {
  count = "${var.isolation_segments * length(var.iso_to_shared_udp_ports)}"

  security_group_id        = "${aws_security_group.iso_shared_security_group.id}"
  type                     = "ingress"
  protocol                 = "udp"
  to_port                  = "${element(var.iso_to_shared_udp_ports, count.index)}"
  from_port                = "${element(var.iso_to_shared_udp_ports, count.index)}"
  source_security_group_id = "${aws_security_group.iso_security_group.id}"
}

resource "aws_security_group_rule" "isolation_segments_to_bosh_all_traffic_rule" {
  count = "${var.isolation_segments}"

  security_group_id        = "${aws_security_group.bosh_security_group.id}"
  type                     = "ingress"
  protocol                 = "-1"
  from_port                = 0
  to_port                  = 0
  source_security_group_id = "${aws_security_group.iso_security_group.id}"
}

resource "aws_security_group_rule" "shared_diego_bbs_to_isolated_cells_rule" {
  count = "${var.isolation_segments}"

  security_group_id        = "${aws_security_group.iso_security_group.id}"
  type                     = "ingress"
  protocol                 = "-1"
  from_port                = 1801
  to_port                  = 1801
  source_security_group_id = "${aws_security_group.iso_shared_security_group.id}"
}

resource "aws_security_group_rule" "nat_to_isolated_cells_rule" {
  count = "${var.isolation_segments}"

  security_group_id        = "${aws_security_group.nat_security_group.id}"
  type                     = "ingress"
  protocol                 = "-1"
  from_port                = 0
  to_port                  = 0
  source_security_group_id = "${aws_security_group.iso_security_group.id}"
}

