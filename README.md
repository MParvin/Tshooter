# Tshooter
(This project is an ongoing work in progress; more features and improvements are coming soon.)


This is a troubleshooter for a deployment in a Kubernetes cluster. 
It will check all steps in troubleshooting from DNS resolve to check ingress, services, deployments, and pods. 

## Features
    - Easy to use command-line interface, just run tshooter command and endpoint of your project
        e.g: `tshooter api.example.com/category/23`
    - Tests all steps in troubleshooting from DNS resolve to checking ingress, services, deployments, pods and etc. 
    - Collects all the data in a report to document the state of troubleshooting.
    - This will run the troubleshooting test and produce a report.

## Workflow
1. Get url from user
2. Do tshoot steps

## Tshoot Steps:

1. Extract hostname from URL
2. Validate hostname
3. Resolve it
4. Check http and https port (telnet)
5. ping the host
6. trace to host
7. Check ingress
8. Check service
9. Check Deployment status
10. Check Pods status

## Commands:

1. `config`: to configure kubernetes cluster credentials
    * It just need a `get`
    Usage:
        ```bash
            tshooter config 'kube_config_file'
        ```

2. `config list`: list all configurations
    Usage:
        ```bash
            tshooter config list
        ```

3. `check`: to check the url
    Usage:

        ```bash
            tshooter check "URL"
        ```
