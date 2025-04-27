# Technical Challenge: Performance and Data Analysis via API

## Objective

You have 1 hour to create an API that receives a JSON file with 100,000 users and provides well-structured and high-performance endpoints for data analysis.

- [Examples of expected API responses](https://github.com/codecon-dev/desafio-1-1s-vs-3j/blob/main/exemplos-endpoints.json)
- [File with 100,000 users to import](https://drive.google.com/file/d/1zOweCB2jidgHwirp_8oBnFyDgJKkWdDA/view?usp=sharing)
- [File with 1,000 users for testing](https://drive.google.com/file/d/1BX03cWxkvB_MbZN8_vtTJBDGiCufyO92/view?usp=sharing)

---

## Input JSON

The JSON contains a list of users with the following structure:

```json
{
  "id": "uuid",
  "name": "string",
  "age": "int",
  "score": "int",
  "active": "bool",
  "country": "string",
  "team": {
    "name": "string",
    "leader": "bool",
    "projects": [{ "name": "string", "completed": "bool" }]
  },
  "logs": [{ "date": "YYYY-MM-DD", "action": "login/logout" }]
}
```
---
## Required Endpoints

### `POST /users`
- Receives and stores the users in memory. You can simulate an in-memory database.

### `GET /superusers`
- Filter: score >= 900 and active = true
- Returns the data and the request processing time.

### `GET /top-countries`
- Groups the superusers by country.
- Returns the 5 countries with the highest number of superusers.

### `GET /team-insights`
- Groups by team.name.
- Returns: total members, leaders, completed projects, and % of active members.

### `GET /active-users-per-day`
- Counts how many logins occurred per date.
- Optional query param: ?min=3000 to filter days with at least 3,000 logins.

### `GET /evaluation`
It must perform a self-evaluation of the main API endpoints and return a scoring report.
The evaluation should test:

- If the returned status is 200
- The response time in milliseconds
- If the return is a valid JSON

This endpoint can run built-in test scripts within the project itself and return a JSON with the results. It will be used to validate the delivery quickly and automatically.

---

## Technical Requirements

- Response time: < 1s per endpoint.
- All endpoints must return the processing time (in milliseconds) and the request timestamp.
- Clean, modular code with well-defined functions.
- You can use any language/framework.
- Documentation or final explanation earns bonus points.
- AI usage is not allowed.