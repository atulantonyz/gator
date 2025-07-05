
---

# 🐊 Gator: A CLI RSS Blog Aggregator

**Gator** is a command-line tool that aggregates RSS feeds with multi-user support. Users can register, follow feeds (including those added by other users), and browse the latest posts from the feeds they follow. It is designed for **local use only**.

---

## 🚀 Features

* Multi-user support
* Add, follow, and unfollow RSS feeds
* Automatic background aggregation
* Lightweight and easy to use

---

## ⚙️ Requirements

* [Go](https://golang.org/doc/install) (v1.24+ recommended)
* [PostgreSQL](https://www.postgresql.org/download/) (v15 or later)

---

## 🛠️ Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/yourusername/gator.git
   cd gator
   ```

2. Install the CLI:

   ```bash
   go install
   ```

---

## 🧰 Setup

### 1. Install PostgreSQL (v15+)

> If you don't have PostgreSQL installed, follow the [official installation guide](https://www.postgresql.org/download/).

### 2. Make sure PostgreSQL is running

Ensure your PostgreSQL server is started and accepting connections. On most systems, this happens automatically at startup. To verify:

```bash
psql -U yourusername -c '\l'
```

If this connects successfully and shows your list of databases, the server is running.

### 3. Create the `gator` database

Start the `psql` shell:

```bash
psql -U yourusername
```

Then run the following SQL command:

```sql
CREATE DATABASE gator;
```

Exit the shell with:

```sql
\q
```

### 4. Configure your database connection

Create a configuration file at `~/.gatorconfig.json` with the following contents:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable"
}
```

Replace `username`, `password`, and other fields with your actual database credentials.

---

## 📚 Available Commands

All commands follow the pattern:

```bash
gator <command> [args]
```

| Command     | Description                                        | Usage Example                                       |
| ----------- | -------------------------------------------------- | --------------------------------------------------- |
| `register`  | Add a new user                                     | `gator register alice`                              |
| `login`     | Set the current user                               | `gator login alice`                                 |
| `users`     | List all users and show the current user           | `gator users`                                       |
| `reset`     | Reset the database (⚠️ deletes all data)           | `gator reset`                                       |
| `addfeed`   | Add a new feed                                     | `gator addfeed tildes https://example.com/feed.xml` |
| `feeds`     | List all available feeds                           | `gator feeds`                                       |
| `follow`    | Follow a feed by URL                               | `gator follow https://example.com/feed.xml`         |
| `following` | List all feeds followed by the current user        | `gator following`                                   |
| `unfollow`  | Unfollow a feed                                    | `gator unfollow https://example.com/feed.xml`       |
| `agg`       | Fetch new posts from all feeds at a given interval | `gator agg 1m` (every 1 minute)                     |
| `browse`    | List latest posts from followed feeds              | `gator browse` or `gator browse 10`                 |

---

## 💡 Example Workflow

```bash
gator register alice
gator login alice
gator addfeed myfeed https://blog.example.com/rss
gator follow https://blog.example.com/rss
gator agg 1m        # Run in a separate terminal window
gator browse 5      # View 5 latest posts
```

---
