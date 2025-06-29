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


### 📌 Occurrences

#### `POST /actions/{id}/events`
Ajouter une occurrence pour une action donnée.

URL params :

- `id`: ID de l’action.

Body JSON :
```json
{
  "timestamp": "2025-06-29T17:15:00Z"
}
```
(ou timestamp laissé vide pour enregistrer « maintenant » côté serveur)

Réponse 201 :
```json
{
  "id": 1,
  "actionId": 1,
  "timestamp": "2025-06-29T17:15:00Z"
}
```

#### `GET /actions/{id}/events`
Lister toutes les occurrences d’une action.

Réponse 200 :
```json
[
  {
    "id": 1,
    "actionId": 1,
    "timestamp": "2025-06-29T17:15:00Z"
  }
]
```
