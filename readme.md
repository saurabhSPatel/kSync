# 🚀 kSync: Kafka with KRaft Mode + Scalable Producer & Consumer in Go

kSync is a lightweight, scalable, and ZooKeeper-free **Kafka cluster** powered by **KRaft (Kafka Raft Mode)**. This project sets up a **single-node Kafka cluster** using **Docker Compose** and implements a **scalable producer-consumer architecture in Go**, leveraging **hashed keys for partitioning** and **consumer groups** for parallel processing.

## 🌟 Why kSync?
- **No ZooKeeper** 🦁 → Uses **KRaft mode**, simplifying Kafka's architecture.
- **Scalable Producer & Consumer** 🏗️ → Ensures efficient load distribution.
- **Partition-based Ordering** 🔑 → Messages with the same key always go to the same partition.
- **Lightweight & Fast** ⚡ → Ideal for local development and production scaling.
- **Persistent Storage** 💾 → Keeps Kafka logs even after container restarts.

---

## 📌 Features
✅ **Single-node Kafka Cluster** with **KRaft Mode**
✅ **Dockerized Setup** for quick deployment
✅ **Multi-listener configuration** (PLAINTEXT, CONTROLLER, INTERNAL)
✅ **Scalable Producer-Consumer Architecture**
✅ **Hashed Key-based Partitioning**
✅ **Consumer Groups for Load Balancing**
✅ **Persistent Storage for Kafka Logs**

---

## 🛠️ Getting Started
### **1️⃣ Clone the Repository**
```sh
git clone https://github.com/saurabhSPatel/kSync.git
cd kSync
```

### **2️⃣ Start Kafka in KRaft Mode**
```sh
docker-compose up -d
```
This will start a single-node Kafka broker using **KRaft** instead of ZooKeeper.

### **3️⃣ Create a Kafka Topic**
```sh
docker exec -it kafka kafka-topics.sh --create --topic my-topic --bootstrap-server localhost:9092 --partitions 3 --replication-factor 1
```

### **4️⃣ Build and Run the Go Producer**
```sh
cd producer
go build -o producer main.go
./producer
```
This will send messages with **hashed keys**, ensuring they are consistently assigned to the same partition.

### **5️⃣ Build and Run the Go Consumer**
```sh
cd consumer
go build -o consumer main.go
./consumer
```
Consumers in the **same consumer group** will automatically share partitions for parallel processing.

### **6️⃣ Verify Message Flow**
You can verify message flow by producing and consuming messages, ensuring that:
- Messages with the same key always go to the same partition.
- Multiple consumers in a group distribute partitions among themselves.

---

## 🔥 What Are We Achieving?
With kSync, we are **simplifying** the traditional Kafka architecture by removing ZooKeeper and enabling **KRaft mode**. In addition, we are implementing a **scalable Go-based producer-consumer system** that ensures:

1️⃣ **Partition-based Ordering** – Messages with the same key go to the same partition.
2️⃣ **Efficient Load Balancing** – Consumer groups distribute messages among multiple consumers.
3️⃣ **High Throughput** – Enables scaling producers and consumers independently.
4️⃣ **Simplified Deployment** – Easy setup with Docker & Go.
5️⃣ **Resilient Architecture** – Consumers can be added dynamically without affecting ongoing processing.

---

## 🏗️ Deployment Steps
### **Option 1: Dockerized Deployment**
You can containerize the producer and consumer using Docker for easy scaling.

#### **1️⃣ Build and Run Producer**
```sh
docker build -t kSync-producer ./producer
docker run --network=host kSync-producer
```

#### **2️⃣ Build and Run Consumer**
```sh
docker build -t kSync-consumer ./consumer
docker run --network=host kSync-consumer
```

### **Option 2: Kubernetes Deployment**
For large-scale production deployment, use **Kubernetes**:
1. **Create Kafka as a StatefulSet** (or use Confluent Operator)
2. **Deploy Producer as a Deployment** (with multiple replicas)
3. **Deploy Consumer as a Scaled Deployment** (consumer group-based auto-scaling)
4. **Expose Kafka via Service**

---

## 🏗️ Future Enhancements
🔹 Add support for **multi-node Kafka clusters** 🔹 Integrate **Kafka Connect** for data streaming 🔹 Implement **security (SASL/SSL)** 🔹 Add **Grafana + Prometheus monitoring** 🔹 Kubernetes Helm charts for auto-deployment

---

## 📜 License
This project is licensed under the **MIT License**.

---

## 🤝 Contributing
We welcome contributions! Feel free to submit issues or pull requests to improve kSync.

🚀 **Enjoy streaming with kSync!** 🎉

