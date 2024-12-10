# **LiveComment System**

This project implements a simple **Live Comment System**, inspired by the exercise described on [SystemDesign.one](https://systemdesign.one/live-comment-system-design/). It demonstrates how to build a distributed system to handle real-time comments for live video streams using **Go**, **RabbitMQ**, **Server-Sent Events (SSE)**, and **Docker**.

---

## **Overview**

The system is composed of the following components:

1. **Gateway Service** (Go):
    - Exposes an SSE endpoint (`/subscribe`) for clients to receive live comments.
    - Manages connections using channels.
    - Listens for messages from the Dispatcher via RabbitMQ and forwards them to connected clients.

2. **Dispatcher Service**:
    - Simulates the generation of comments for multiple videos (e.g., videos `A`, `B`, and `C`).
    - Publishes comments to RabbitMQ for the Gateway to distribute.

3. **RabbitMQ**:
    - Acts as a message broker, facilitating communication between the Dispatcher and Gateway services.

4. **Simulated Comment Generator**:
    - Produces random comments for videos `A`, `B`, and `C`, and sends them to the Dispatcher.

---

## **Features**

- **Real-Time Comments**:
    - Clients can subscribe to comments for specific videos using SSE.

- **RabbitMQ Integration**:
    - Decouples the Dispatcher and Gateway services, ensuring scalability and reliability.

- **Automatic Connection Handling**:
    - The Gateway service closes the SSE connection when the browser or client disconnects.

---

## **Getting Started**

### **Prerequisites**

- Docker and Docker Compose installed on your machine.
- A browser or tool like `curl` to test the SSE endpoint.

---

### **Running the System**

1. **Clone the repository**:
   ```
   git clone <repository-url>
   cd realtime_projects/livecomment
   ```
2. **Start the services**:
   Run the following command to start all services using Docker Compose:
   ```bash
   docker-compose up -d
   ```
3. **Access the Gateway SSE Endpoint:**
   Open a browser and navigate to the following URL to subscribe to a video stream
   ```http://localhost:8080/subscribe?video=a```
   Replace `a` with `b` or `c` to subscribe to comments for other videos.
   When the browser is closed, the connection to the server will automatically terminate.


### **Simulated Comments**
A background service simulates comment generation for videos A, B, and C. This ensures a steady stream of comments to demonstrate the functionality.

### **Stopping the Services**
To stop all running services, use the following command: `docker-compose down`

### **License** 
This project is open-source and available under the MIT License.
