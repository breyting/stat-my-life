# 📑 API Documentation — Stat My Life

Cette API permet de créer, consulter et mettre à jour les actions et leurs occurrences.

## 🛠️ Base URL

- Pour le développement : `http://localhost:8080`

---

## 🔗 Endpoints

---

### 📌 Actions

#### `POST /actions`
Créer une nouvelle action.

- **Body JSON :**
```json
{
  "name": "Prendre le tram"
}
```
- **Réponse 201 :**
```json
{
  "id": 1,
  "name": "Prendre le tram",
  "createdAt": "2025-06-29T15:30:00Z"
}
```

#### `GET /actions`
Lister toutes les actions existantes.

- **Réponse 200 :**
```json
[
  {
    "id": 1,
    "name": "Prendre le tram",
    "createdAt": "2025-06-29T15:30:00Z"
  }
]
```
