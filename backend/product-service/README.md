# Product Service (Go + PostgreSQL)

## 📌 Description
Ce microservice gère les produits de la boutique en ligne. Il permet de :
✅ Ajouter, modifier, supprimer et lister des produits.
✅ Utiliser PostgreSQL pour stocker les produits.
✅ Sécuriser les routes avec JWT.
✅ Documenter l'API avec Swagger.

---

## 🚀 Installation
### 1️⃣ Prérequis
- **Go 1.20+**
- **PostgreSQL**
- **Docker & Docker Compose (optionnel)**

### 2️⃣ Cloner le projet
```bash
git clone https://github.com/votre-repo/product-service.git
cd product-service
```

### 3️⃣ Configurer les variables d'environnement
Créer un fichier `.env` :
```ini
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=products
JWT_SECRET=supersecret
PORT=4000
```

### 4️⃣ Installer les dépendances
```bash
go mod tidy
```

---

## 🏗️ Lancer l'Application
### 1️⃣ Démarrer PostgreSQL
Si PostgreSQL est installé localement :
```bash
sudo -u postgres psql -c "CREATE DATABASE products;"
```

Ou avec Docker Compose :
```bash
docker-compose up -d
```

### 2️⃣ Démarrer le serveur Go
```bash
go run main.go
```

---

## 🔥 Documentation API (Swagger)
Une fois le service lancé, accède à la documentation interactive :
📌 **http://localhost:4000/swagger/index.html**

---

## 📜 Routes principales
| Méthode | Endpoint | Description |
|---------|---------|-------------|
| GET | `/api/products` | Récupérer tous les produits |
| POST | `/api/products` | Ajouter un produit (JWT requis) |
| GET | `/api/products/{id}` | Récupérer un produit par ID |
| PUT | `/api/products/{id}` | Modifier un produit (JWT requis) |
| DELETE | `/api/products/{id}` | Supprimer un produit (JWT requis) |

---

## 🐳 Utilisation avec Docker
```bash
docker build -t product-service .
docker run -p 4000:4000 --env-file .env product-service
```

Ou avec **Docker Compose** :
```bash
docker-compose up --build
```

---

## 📌 Contributeurs
👤 **Nom du Dev** – *Mohamed ESSID*

