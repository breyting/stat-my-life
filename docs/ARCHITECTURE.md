# 🏛️ Architecture Technique — Stat My Life

## Objectif
Suivre la fréquence d’actions quotidiennes via une app mobile et un site web, avec un backend unique.

---

## Schéma général

```
React Native App
       |
       v
     API REST (Go) <--> PostgreSQL
       ^
       |
  Next.js Site Web
```

---

## Communication

- Les frontends (app et site) appellent l’API Go.  
- L’API : validation → accès BDD → réponse JSON.  

---

## Modèles de données

### Action
| champ     | type    |
|-----------|---------|
| id        | int     |
| name      | string  |
| createdAt | timestamp |

### Occurrence
| champ     | type      | note               |
|-----------|-----------|--------------------|
| id        | int       |                    |
| actionId  | int       | FK vers **Action** |
| timestamp | timestamp | date/heure         |

---

## Technologies

- **Backend** : Go (net/http, chi ou gin)  
- **BDD** : PostgreSQL (SQLite possible en local)  
- **Web** : Next.js (React)  
- **Mobile** : React Native (Expo conseillé)  
- **Déploiement** : Raspberry Pi ou cloud léger (Render, Fly.io…)  
- **HTTPS** : Nginx + Let’s Encrypt  

---

## Bonnes pratiques

1. MVP centré sur le CRUD Actions / Occurrences.  
2. Valider toutes les entrées côté API.  
3. Tenir cette doc à jour (README, API.md, ARCHITECTURE.md).  

---

## À venir

- Authentification JWT  
- Dashboard avec graphiques  
- Notifications push sur mobile  
