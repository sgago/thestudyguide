
variable "my_number" {
  description = "A number variable"
  type        = number
  default     = 42
}

output "my_number_output" {
  value = var.my_number
  description = "The number variable output"
}

output "my_number_condition_output" {
  value = var.my_number > 40 ? "Greater than 40" : "Less than or equal to 40"
  description = "The condition output example"
}

variable "my_string" {
  description = "A string variable"
  type        = string
  default     = "Hello, World!"
}

output "my_string_output" {
  value = var.my_string
  description = "The string variable output"
}

variable "my_secret_string" {
  description = "value of my secret string"
  default = "my secret string"
  sensitive = true
}

output "my_secret_string_output" {
  value = var.my_secret_string
  description = "The secret string variable output"
  sensitive = true
}

variable "my_list" {
  description = "A list variable"
  type        = list(string)
  default     = ["one", "two", "three"]
}

output "my_list_output" {
  value = var.my_list
  description = "The list variable output"
}

output "my_list_element_output" {
  value = var.my_list[0]
  description = "The list variable output"
}

output "my_list_length_output" {
  value = length(var.my_list)
  description = "The count of the list variable"
}

output "my_list_for_loop_element_output" {
  // For each element in var.my_list, output the element
  value = [for element in var.my_list : element]
  description = "The list variable output"
}

variable "my_map" {
  type        = map(string)
  description = "A map variable"
  default     = {
    key1 = "value1"
    key2 = "value2"
    key3 = "value"
  }
}

output "my_map_element_output" {
  value = var.my_map["key1"]
  description = "The map key1 output"
}