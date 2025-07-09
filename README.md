# ☁️ Distributed Cloud Storage – Scalable File Management System

A cloud-native backend project that simulates Dropbox-like file storage using Go, AWS S3, and PostgreSQL. Built with extensibility and distributed principles in mind, this project enables secure file upload/download, metadata tracking, and forms the foundation for chunked, multi-node distributed storage.

---

## 🧩 Overview

This project implements a backend storage system with the following structure:

🧠 [Storage Core](./file_storage): File upload/download logic with AWS S3 integration
🗃️ [Metadata Service](./models): PostgreSQL-backed metadata tracking
⚙️ \[Utilities & Routing]\(./utils + /router): S3 helpers and API routing via Gin

The system is designed to scale into a fully distributed architecture, with chunking and synchronization logic in the roadmap.

---

## 🔍 Core Features

📤 File Upload: Upload files via REST API; stores data in AWS S3
📥 File Download: Download files by unique ID or name
📑 Metadata Tracking: Filename, size, timestamps tracked in PostgreSQL
🔗 Clean API Design: Built using Gin (Go web framework)
🧱 Extensible Backend: Chunking and node distribution logic ready to integrate
🔐 Secure Access: Environment-based credential loading via `.env`
📊 Logging: CloudWatch or local logs available for tracing file operations

---

## 🛠 Tech Stack

| Component        | Tech Details                                             |
| ---------------- | -------------------------------------------------------- |
| 🧠 Backend Core  | Go (Golang), Gin, GORM                                   |
| ☁️ Cloud Storage | AWS S3, AWS IAM, AWS SDK v2                              |
| 🗃️ Metadata DB  | PostgreSQL (via GORM ORM)                                |
| 🔐 Auth/Security | `.env` file (ignored in Git), IAM-managed credentials    |
| 🧰 Utilities     | S3 utility functions, local testing CLI, structured logs |

---

## 🗺 Roadmap (Planned Enhancements)

* [ ] File Chunking and Distributed Storage across multiple nodes
* [ ] File Versioning & Conflict Resolution
* [ ] Multi-user Support with Authentication
* [ ] Shareable Access Links and Permissions
* [ ] Web Dashboard for uploads, status, and file management

---
![image](https://github.com/user-attachments/assets/97ebf4f4-e557-489a-925d-c74565c896e9)
