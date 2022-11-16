# Bee - A k8s native serverless functions framework

<img width="280" alt="Screenshot 2022-11-14 at 2 34 08 AM" src="https://user-images.githubusercontent.com/49859828/201544484-888015ab-255f-4768-80fb-5a3270011748.png">

# Readme

### Commands

1. **Creating the function**
    
    (so far python3 and golang functions are supported, feel free to contribute)
    
    ```bash
    go run main.go create <language> <function-name>
    ```
    
    this will create a folder with the function name in the current directory,
    
    edit the code in the function file to add your function code
    

1. ******************************************Building the function image******************************************
    
    ```bash
    go run main.go build <path-to-the-function-folder>
    ```
    
    this will create a docker image based on your function
    
2. ********************************Pushing the function image********************************
    
    ```bash
    go run main.go push <path-to-the-function-folder> <repo-user/repo-name>
    ```
    
    (setup your repository with docker first before using this command)
    
3. **********************************************************************Deploying on the kubernetes cluster**********************************************************************
    
    ```bash
    go run main.go deploy <path-to-the-function-folder>
    ```
    
    (again, setup the repository with the kubernetes first)
    
    this will create a deployment, and a service for the function
    

### Config.yaml

config.yaml contains the metadata for the function

you can change the number of replicas and the auto-scaling parameters here

running the deploy commands will apply these configs

```yaml
function_name: coffee
runtime: python
repository: xatriya/coffee
replicas: 3
autoscaling: true
min_replicas: 3
max_replicas: 10
cpu_percent: 50
isbuilt: true
```

### Deployment.yaml and Service.yaml

deployment.yaml and service.yaml is generated after the pushing image

you can modify these files, **although do not change the labels**
