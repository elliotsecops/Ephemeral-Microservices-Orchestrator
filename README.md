# Orquestador de Microservicios Efímeros

## Descripción General

Este Orquestador de Microservicios Efímeros es un proyecto minimalista basado en Go diseñado para gestionar el ciclo de vida de microservicios efímeros. Este proyecto proporciona una forma simple pero efectiva de crear, monitorear y suprimir contenedores Docker basados en eventos detectados. También incluye un servidor API básico para manejar solicitudes de creación de microservicios.
Este proyecto está diseñado para profesionales de TI involucrados en DevOps, Cloud Computing y profesionales de la seguridad.
- Para los ingenieros de DevOps, este proyecto les proporciona un marco para gestionar microservicios efímeros (de corta duración), lo que resulta especialmente útil en escenarios de integración y despliegue continuos, en los que los servicios pueden activarse y desactivarse con frecuencia.
- Para los ingenieros de la nube, este orquestador ayuda a gestionar microservicios efímeros que podrían utilizarse para pruebas, desarrollo o aplicaciones escalables en la nube.
- Los profesionales de seguridad (security engineers/pentesters) pueden utilizar esta herramienta para desplegar microservicios efímeros durante entornos de pruebas o pruebas de penetración. Ayuda a crear rápidamente entornos aislados para probar las vulnerabilidades de los contenedores y la seguridad a nivel de orquestador.


## Características

- **Gestión de Contenedores Docker**: Crear y destruir contenedores Docker de manera dinámica.
- **Manejo de Eventos**: Escuchar eventos para desencadenar la creación de microservicios.
- **Servidor API**: Manejar solicitudes de creación de microservicios a través de una API HTTP simple.
- **Monitoreo**: Sistema de monitoreo básico para mantener el seguimiento de los microservicios.
- **Interacción con Kubernetes** (Opcional): Desplegar microservicios en un clúster de Kubernetes.

## Prerrequisitos

Antes de comenzar, asegúrate de tener instalado lo siguiente:

- **Go**: Versión 1.16 o superior. [Descargar Go](https://golang.org/dl/)
- **Docker**: Versión 20.10 o superior. [Descargar Docker](https://www.docker.com/get-started)
- **Minikube** (Opcional, para Kubernetes): Versión 1.20 o superior. [Descargar Minikube](https://minikube.sigs.k8s.io/docs/start/)
- **kubectl** (Opcional, para Kubernetes): Versión 1.20 o superior. [Descargar kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

## Instalación

### 1. Clonar el Repositorio

Clona el repositorio en tu máquina local:

```bash
git clone https://github.com/elliotsecops/ephemeral-microservices-orchestrator.git
cd ephemeral-microservices-orchestrator.git
```

### 2. Inicializar el Módulo de Go

Inicializa el módulo de Go:

```bash
go mod init github.com/elliotsecops/ephemeral-microservices-orchestrator
```

### 3. Instalar Dependencias

Instala los paquetes necesarios de Go:

```bash
go mod tidy
```

### 4. Ejecutar el Script de Configuración

Ejecuta el script de configuración para instalar Go, Docker, Minikube y kubectl:

```bash
./scripts/setup.sh
```

## Estructura del Proyecto

El proyecto está estructurado de la siguiente manera:

```
ephemeral-microservices-orchestrator/
├── cmd/
│   └── orquestador/
│       └── main.go
├── internal/
│   ├── api/
│   │   └── server.go
│   ├── container/
│   │   └── manager.go
│   ├── kubernetes/
│   │   └── manager.go
│   ├── monitor/
│   │   └── system.go
│   └── templater/
│       └── templater.go
├── pkg/
│   ├── events/
│   │   └── handler.go
│   └── models/
│       └── microservice.go
├── configs/
│   └── config.yaml
├── deployments/
│   └── kubernetes.yaml
└── scripts/
    └── setup.sh
```

### Archivos y Directorios Clave

- **`cmd/orquestador/main.go`**: El punto de entrada principal de la aplicación.
- **`internal/api/server.go`**: Implementación del servidor API.
- **`internal/container/manager.go`**: Gestión de contenedores Docker.
- **`internal/kubernetes/manager.go`**: Interacción con Kubernetes.
- **`internal/monitor/system.go`**: Sistema de monitoreo para microservicios.
- **`internal/templater/templater.go`**: Manejo de plantillas para microservicios.
- **`pkg/events/handler.go`**: Manejo de eventos del sistema.
- **`pkg/models/microservice.go`**: Definición de modelos de datos.
- **`configs/config.yaml`**: Archivo de configuración.
- **`deployments/kubernetes.yaml`**: Configuración de despliegue en Kubernetes.
- **`scripts/setup.sh`**: Script de configuración del entorno.

## Uso

### 1. Iniciar el Orquestador

Ejecuta el programa principal para iniciar el orquestador:

```bash
go run cmd/orquestador/main.go
```

### 2. Probar la Gestión de Contenedores Docker

- **Crear un Contenedor Docker**:
  - El orquestador creará automáticamente contenedores Docker basados en eventos detectados.
  - Observa la salida para asegurarte de que el contenedor se crea correctamente.

- **Eliminar un Contenedor Docker**:
  - Usa la CLI de Docker para listar los contenedores en ejecución:
    ```bash
    docker ps
    ```
  - Anota el ID del contenedor creado.
  - Ejecuta el siguiente comando para eliminar el contenedor:
    ```bash
    docker rm -f <container_id>
    ```
  - Verifica que el contenedor ha sido eliminado.

### 3. Probar la Interacción con Kubernetes (Opcional)

- **Desplegar en Kubernetes**:
  - Asegúrate de que Minikube esté en ejecución:
    ```bash
    minikube start
    ```
  - Aplica la configuración de despliegue de Kubernetes:
    ```bash
    kubectl apply -f deployments/kubernetes.yaml
    ```
  - Verifica el despliegue:
    ```bash
    kubectl get deployments
    ```

### 4. Probar el Servidor API

- **Iniciar el Servidor API**:
  - El servidor API debería iniciarse automáticamente cuando ejecutes el programa principal.
  - El servidor escucha en el puerto 8080.

- **Enviar una Solicitud de Creación**:
  - Abre otra terminal y envía una solicitud al servidor API:
    ```bash
    curl http://localhost:8080/create
    ```
  - Observa la respuesta para asegurarte de que la solicitud se maneja correctamente.

### 5. Probar el Sistema de Monitoreo

- **Iniciar el Sistema de Monitoreo**:
  - El sistema de monitoreo debería iniciarse automáticamente cuando ejecutes el programa principal.
  - Observa la salida para asegurarte de que el sistema de monitoreo está en ejecución.

### 6. Probar el Manejo de Eventos

- **Iniciar el Escuchador de Eventos**:
  - El escuchador de eventos debería iniciarse automáticamente cuando ejecutes el programa principal.
  - Observa la salida para asegurarte de que los eventos se detectan y los microservicios se crean.

## Configuración

El proyecto utiliza un archivo de configuración simple ubicado en `configs/config.yaml`. Este archivo contiene configuraciones básicas para Docker y Kubernetes:

```yaml
docker:
  image_name: "my_microservice_image"
  network: "default"
```

## Contribuciones

Las contribuciones son bienvenidas. Puedes seguir estos pasos:

1. Haz un fork del repositorio.
2. Crea una nueva branch (`git checkout -b feature-branch`).
3. Realiza tus cambios.
4. Confirma tus cambios (`git commit -am 'Add some feature'`).
5. Envía los cambios a la branch (`git push origin feature-branch`).
6. Crea una nueva pull request.

## Contacto

Para cualquier pregunta o comentario, puedes abrir un issue.

--- 

# Ephemeral Microservices Orchestrator

## Overview

This Ephemeral Microservices Orchestrator is a minimalist Go-based project designed to manage the lifecycle of ephemeral microservices. This project provides a simple yet effective way to create, monitor, and destroy Docker containers based on detected events. It also includes a basic API server for handling microservice creation requests.
This project is designed for IT professionals involved in DevOps, Cloud Computing and security professionals.
- For DevOps engineers, this project provides them with a framework for managing ephemeral (short-lived) microservices, which is especially useful in continuous integration and deployment scenarios where services may be turned on and off frequently.
- For cloud engineers, this orchestrator helps manage ephemeral microservices that could be used for testing, development or scalable cloud applications.
- Security engineers/pentesters can use this tool to deploy ephemeral microservices during testing or penetration testing environments. It helps to quickly create isolated environments for testing container vulnerabilities and orchestrator-level security.

## Features

- **Docker Container Management**: Create and destroy Docker containers dynamically.
- **Event Handling**: Listen for events to trigger microservice creation.
- **API Server**: Handle microservice creation requests via a simple HTTP API.
- **Monitoring**: Basic monitoring system to keep track of microservices.
- **Kubernetes Interaction** (Optional): Deploy microservices to a Kubernetes cluster.

## Prerequisites

Before you begin, ensure you have the following installed:

- **Go**: Version 1.16 or later. [Download Go](https://golang.org/dl/)
- **Docker**: Version 20.10 or later. [Download Docker](https://www.docker.com/get-started)
- **Minikube** (Optional, for Kubernetes): Version 1.20 or later. [Download Minikube](https://minikube.sigs.k8s.io/docs/start/)
- **kubectl** (Optional, for Kubernetes): Version 1.20 or later. [Download kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

## Installation

### 1. Clone the Repository

Clone the repository to your local machine:

```bash
git clone https://github.com/elliotsecops/ephemeral-microservices-orchestrator.git
cd ephemeral-microservices-orchestrator.git
```

### 2. Initialize the Go Module

Initialize the Go module:

```bash
go mod init github.com/elliotsecops/ephemeral-microservices-orchestrator
```

### 3. Install Dependencies

Install the necessary Go packages:

```bash
go mod tidy
```

### 4. Run the Setup Script

Run the setup script to install Go, Docker, Minikube, and kubectl:

```bash
./scripts/setup.sh
```

## Project Structure

The project is structured as follows:

```
ephemeral-microservices-orchestrator/
├── cmd/
│   └── orquestador/
│       └── main.go
├── internal/
│   ├── api/
│   │   └── server.go
│   ├── container/
│   │   └── manager.go
│   ├── kubernetes/
│   │   └── manager.go
│   ├── monitor/
│   │   └── system.go
│   └── templater/
│       └── templater.go
├── pkg/
│   ├── events/
│   │   └── handler.go
│   └── models/
│       └── microservice.go
├── configs/
│   └── config.yaml
├── deployments/
│   └── kubernetes.yaml
└── scripts/
    └── setup.sh
```

### Key Files and Directories

- **`cmd/orquestador/main.go`**: The main entry point of the application.
- **`internal/api/server.go`**: Implementation of the API server.
- **`internal/container/manager.go`**: Management of Docker containers.
- **`internal/kubernetes/manager.go`**: Interaction with Kubernetes.
- **`internal/monitor/system.go`**: Monitoring system for microservices.
- **`internal/templater/templater.go`**: Handling of templates for microservices.
- **`pkg/events/handler.go`**: Handling of system events.
- **`pkg/models/microservice.go`**: Definition of data models.
- **`configs/config.yaml`**: Configuration file.
- **`deployments/kubernetes.yaml`**: Kubernetes deployment configuration.
- **`scripts/setup.sh`**: Environment setup script.

## Usage

### 1. Start the Orchestrator

Run the main program to start the orchestrator:

```bash
go run cmd/orquestador/main.go
```

### 2. Test Docker Container Management

- **Create a Docker Container**:
  - The orchestrator will automatically create Docker containers based on detected events.
  - Observe the output to ensure the container is created successfully.

- **Remove a Docker Container**:
  - Use the Docker CLI to list running containers:
    ```bash
    docker ps
    ```
  - Note the container ID of the created container.
  - Run the following command to remove the container:
    ```bash
    docker rm -f <container_id>
    ```
  - Verify the container is removed.

### 3. Test Kubernetes Interaction (Optional)

- **Deploy to Kubernetes**:
  - Ensure Minikube is running:
    ```bash
    minikube start
    ```
  - Apply the Kubernetes deployment configuration:
    ```bash
    kubectl apply -f deployments/kubernetes.yaml
    ```
  - Verify the deployment:
    ```bash
    kubectl get deployments
    ```

### 4. Test the API Server

- **Start the API Server**:
  - The API server should start automatically when you run the main program.
  - The server listens on port 8080.

- **Send a Create Request**:
  - Open another terminal and send a request to the API server:
    ```bash
    curl http://localhost:8080/create
    ```
  - Observe the response to ensure the request is handled correctly.

### 5. Test the Monitoring System

- **Start the Monitoring System**:
  - The monitoring system should start automatically when you run the main program.
  - Observe the output to ensure the monitoring system is running.

### 6. Test the Event Handling

- **Start the Event Listener**:
  - The event listener should start automatically when you run the main program.
  - Observe the output to ensure events are detected and microservices are created.

## Configuration

The project uses a simple configuration file located at `configs/config.yaml`. This file contains basic settings for Docker and Kubernetes:

```yaml
docker:
  image_name: "my_microservice_image"
  network: "default"
```

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add some feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a new Pull Request.

## Contact

For any questions or feedback, please open an issue.
