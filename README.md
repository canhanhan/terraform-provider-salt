# terraform-provider-salt

Terraform provider for SaltStack OSS. This provider requires CherryPy NetAPI module configured.

The intended use-case is to integrate Terraform and SaltStack:
1. Create minion keys on master using `salt_minion` resource

2. Configure and initialize Salt Minion using [cloud-init](https://cloudinit.readthedocs.io/en/latest/topics/modules.html#salt-minion).

3. Use [salt-highstate](https://github.com/finarfin/terraform-provisioner-salt-highstate) provisioner to your manifests.

4. Terraform will create the machines, execute the provisioner and will only complete when the highstate is successfully applied. 
 

## Usage

### Installation
1. Download the binary for your platform to [Terraform plugins path](https://www.terraform.io/docs/plugins/basics.html#installing-plugins).
    - Windows: `%AppData%\terraform.d\plugins`
    - Linux: `~/.terraform.d/plugins`

2. Configure CherryPy NetAPI module on SaltStack master. See [setup section of Saltstack documentation](https://docs.saltstack.com/en/latest/ref/netapi/all/salt.netapi.rest_cherrypy.html#a-rest-api-for-salt) for instructions.

3. Add the provider to your manifest. [See the example manifest](example/main.tf)

### Configuration
| Name | Type | Required | Remarks |
|-|-|-|-|
| `address` | String | **Required** | URL to the CherryPy NetAPI endpoint (e.g.: https://saltmaster:8000) | `username`| String | **Required** | Username |
| `password`| String | **Required** | Password |
| `backend` | String | **Required** | External authentincation backend (eauth) (e.g.: pam) |
| `skip_verify` | Bool | *Optional* | Skip TLS/SSL verification (Default: false) |

### Resources

#### salt_minion 

#### Inputs
| Name | Type | Usage | Remarks |
|-|-|-|-|
| `name` | String | **Required** | Minion ID |
| `key_size` | Int | *Optional* | Key size (Default: 2048) |

#### Attributes
| Name | Type | Remarks |
|-|-|-|
| `private_key` | String | Minion private key |
| `public_key` | String | Minion public key |

## Example
```terraform
provider "salt" {
    address = "http://192.168.50.10:8000"
    username = "test_user"
    password = "test_pwd"
    backend = "pam"    
}

resource "salt_minion" "test" {
    name = "minion1"
    key_size = 4096
}
```
