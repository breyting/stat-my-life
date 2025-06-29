# 📑 API Documentation — Stat My Life

Cette API permet de créer, consulter et mettre à jour les actions et leurs occurrences.

## 🛠️ Base URL
Pour le développement : `http://localhost:8080`

---

## 🔗 Endpoints

### Actions

#### POST /actions
Créer une nouvelle action.

Body JSON :

```json
{
  "name": "Prendre le tram"
}
```

Réponse 201 :

```json
{
  "id": 1,
  "name": "Prendre le tram",
  "createdAt": "2025-06-29T15:30:00Z"
}
```

---

#### GET /actions
Lister toutes les actions existantes.

Réponse 200 :

```json
[
  {
    "id": 1,
    "name": "Prendre le tram",
    "createdAt": "2025-06-29T15:30:00Z"
  }
]
```

---

### Occurrences

#### POST /actions/{id}/events
Ajouter une occurrence pour l’action dont l’ID est `{id}`.

Body JSON :

```json
{
  "timestamp": "2025-06-29T17:15:00Z"
}
```

*(ou laisse vide pour enregistrer “maintenant” côté serveur)*

Réponse 201 :

```json
{
  "id": 1,
  "actionId": 1,
  "timestamp": "2025-06-29T17:15:00Z"
}
```

---

#### GET /actions/{id}/events
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

---

## 🔐 Authentification
*Non prévue dans le MVP, mais des JWT pourront être ajoutés plus tard.*

---

## ⚠️ Codes d’erreur

- 400 Bad Request — JSON invalide ou données manquantes  
- 404 Not Found — action inexistante  
- 500 Internal Server Error — erreur serveur  

---

## ✅ Notes

- Tous les timestamps sont en UTC (ISO 8601).  
- Les IDs sont des entiers auto-incrémentés.  
