# Rancher and Go api tests for MacOS

## Running Rancher via Docker

**Run the Rancher server using Docker**:
```bash
   docker run -d --name rancher-server --privileged -p 80:80 -p 443:443 rancher/rancher:latest
```

## Create a password for logging into the Rancher UI

**Find the container id**
```bash
   docker ps
```
**Retrieve the generated password**
```bash
   docker logs <container-id> 2>&1 | grep "Bootstrap Password:"
```
**Copy the password**
**Access the Rancher UI**
Open https://localhost/ in your browser
**Change the password**
Use the retrieved password to log in. Once logged in change the password to Test!2345678, which is used in the tests to avoid making changes in the repository.

## Cloning the Repository
```bash
   git clone https://github.com/igTkachov/go-rancher-api-tests
```

## Running Tests
**Navigate to the root directory and run the tests in headless mode**
```bash
go test -v
```
