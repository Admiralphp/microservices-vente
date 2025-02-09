# Auth Service

## 📌 Description
**Auth Service** est un microservice d'authentification basé sur **Node.js, Express, MongoDB** et **JWT**. Il permet la gestion des utilisateurs avec des fonctionnalités comme l'inscription, la connexion et l'authentification sécurisée via des tokens JWT.

---

## 🚀 Technologies utilisées
- **Node.js** & **Express** – Backend API
- **MongoDB** – Base de données NoSQL
- **JWT (JSON Web Tokens)** – Authentification sécurisée
- **Swagger** – Documentation API
- **Docker** – Conteneurisation
- **Nginx Gateway** – API Gateway

---

## 📂 Arborescence du projet
```
📦 auth-service
├── 📂 src
│   ├── 📂 config        # Configuration (MongoDB, JWT...)
│   ├── 📂 controllers   # Logique métier des routes
│   ├── 📂 middlewares   # Middleware JWT
│   ├── 📂 models        # Modèles Mongoose
│   ├── 📂 routes        # Définition des endpoints
│   ├── 📂 utils         # Fonctions utilitaires
│   ├── swagger.yaml     # Documentation API
│   ├── index.js         # Point d'entrée du serveur
├── .env                # Variables d'environnement
├── Dockerfile          # Fichier Docker
├── docker-compose.yml  # Orchestration des services
├── package.json        # Dépendances Node.js
├── README.md           # Documentation
```

---

## ⚙️ Installation & Exécution

### 1️⃣ Prérequis
- **Node.js** v18+
- **MongoDB** (local ou via Docker)
- **Docker** (optionnel pour exécuter via conteneurs)

### 2️⃣ Installation des dépendances
```bash
npm install
```

### 3️⃣ Configuration de l'environnement
Crée un fichier `.env` :
```ini
PORT=3000
MONGO_URI=mongodb://localhost:27017/authdb
JWT_SECRET=super_secret_key
```

### 4️⃣ Lancer le service en local
```bash
npm run dev  # Mode développement avec Nodemon
npm start    # Exécution en production
```

### 5️⃣ Exécution avec Docker
```bash
docker-compose up --build
```

---

## 📌 API Endpoints

### 📝 **Inscription**
```http
POST /api/auth/register
```
#### Exemple de requête :
```json
{
  "username": "testuser",
  "email": "test@mail.com",
  "password": "password123"
}
```
#### Réponse :
```json
{
  "message": "Utilisateur créé avec succès"
}
```

### 🔑 **Connexion**
```http
POST /api/auth/login
```
#### Réponse :
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### 🔐 **Accès au profil** (nécessite un token JWT)
```http
GET /api/auth/profile
Authorization: Bearer <TOKEN>
```

---

## 📘 Documentation API (Swagger)
Une documentation interactive est disponible sur :
[http://localhost:3000/api/docs](http://localhost:3000/api/docs)

---

## ✅ Tests
Exécuter les tests avec :
```bash
npm test
```

---

## 🛠️ Déploiement
Déploiement via **Docker + Vercel + CI/CD GitHub Actions** (à configurer selon l'infra).

---

## 📌 Contributeurs
👤 **Nom du Dev** – *Mohamed ESSID*

---

🔥 **Besoin d'améliorations ?** N'hésitez pas à proposer des PRs !

