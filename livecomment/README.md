# **LiveComment System**

This project implements a simple **Live Comment System**, inspired by the exercise described on [SystemDesign.one](https://systemdesign.one/live-comment-system-design/). It demonstrates how to build a distributed system to handle real-time comments for live video streams using **Go**, **RabbitMQ**, **Server-Sent Events (SSE)**, and **Docker**.

---

## **Overview**

The system is composed of the following components:

1. **Gateway Service** (Go):
    - Exposes an SSE endpoint (`/subscribe`) for clients to receive live comments. When a subscription is established for a video, this service create a subscription in dispatcher service
    - Manages connections using channels.
    - Listens for messages from the Dispatcher through RabbitMQ and forwards them to connected clients of the same video

2. **Dispatcher Service**:
    - Receives comments by video
    - Routing of every comment to the gateway according to it subscription

3. **RabbitMQ**:
    - Acts as a message broker, facilitating communication between the Dispatcher and Gateway services.

4. **Simulated Comment Generator**:
    - Produces random comments for videos, and sends them to the Dispatcher.

---

## **Getting Started**

### **Prerequisites**

- Docker and Docker Compose installed on your machine.
- A browser or tool like `curl` to test the SSE endpoint.
- Go 1.23 installed to execute the simulation of comments

---

### **Running the System**

1. **Clone the repository**:
   ```
   git clone https://github.com/mayusGomez/realtime-projects.git
   cd realtime_projects/livecomment
   ```
2. **Start the services**:
   Run the following command to start all services using Docker Compose:
   ```bash
   make build
   make run
   ```
   
3. **Generate comments:**
   Run the command to generate random comments
   ```bash
   make generate-comments
   ```
   Open the link which the console shows, something similar to  ```http://localhost:8080/subscribe?video=[uuid]```
   When the browser is closed, the connection to the server will automatically terminate.

### **Stopping the Services**
To stop all running services, use the following command: `make stop`

### **License** 
This project is open-source and available under the MIT License.
