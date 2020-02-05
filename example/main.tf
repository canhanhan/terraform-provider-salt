provider "salt" {
    address = "http://192.168.50.10:8000"
    username = "test_user"
    password = "test_pwd"
    backend = "pam"    
}

resource "salt_minion" "test" {
    name = "minion4"
}

output "private_key" {
    value = salt_minion.test.private_key
    sensitive = true
}

output "public_key" {
    value = salt_minion.test.public_key
}
