# InventoryPro Projekt

Dieses Projekt ist eine vollständige Webanwendung zur Verwaltung von Lagerbeständen, die aus einem Frontend, einem Backend und einer MariaDB-Datenbank besteht. Diese Anleitung beschreibt die Schritte zur Installation und Ausführung des Projekts.

## Datenbank

1. **MariaDB Installation**:
   - Stellen Sie sicher, dass MariaDB auf Ihrem System installiert ist. Anleitungen zur Installation finden Sie auf der offiziellen [MariaDB-Website](https://mariadb.org/download/).

2. **Datenbank einrichten**:
   - Navigieren Sie in das Backend-Verzeichnis und führen Sie das Skript `./db-setup.sh` aus, um die Datenbank einzurichten.
   ```sh
   cd Backend
   ./db-setup.sh
   ```

## Backend

1. **Golang Installation**:
   - Installieren Sie Golang auf Ihrem System. Anleitungen finden Sie auf der offiziellen [Golang-Website](https://golang.org/doc/install).

2. **Dependencies installieren**:
   - Navigieren Sie in das Backend-Verzeichnis und führen Sie den folgenden Befehl aus, um alle Abhängigkeiten zu installieren.
   ```sh
   go mod tidy
   ```

3. **Backend starten**:
   - Führen Sie das Backend mit folgendem Befehl aus:
   ```sh
   go run main.go
   ```

## Frontend

1. **Node.js Installation**:
   - Installieren Sie Node.js, das npm (Node Package Manager) enthält. Anleitungen finden Sie auf der offiziellen [Node.js-Website](https://nodejs.org/).

2. **pnpm Installation**:
   - Installieren Sie pnpm global auf Ihrem System.
   ```sh
   npm install -g pnpm
   ```

3. **Frontend einrichten und starten**:
   - Navigieren Sie in das Frontend-Verzeichnis und installieren Sie die Abhängigkeiten.
   ```sh
   cd Frontend
   pnpm i
   ```

   - Starten Sie das Frontend mit dem folgenden Befehl:
   ```sh
   pnpm run dev
   ```

## Zusammenfassung

Nach dem Einrichten der Datenbank, des Backends und des Frontends sollte die Anwendung nun lokal auf Ihrem System laufen. Sie können die Webanwendung in Ihrem Browser öffnen und mit der Verwaltung von Lagerbeständen beginnen.

Falls Sie Fragen haben oder auf Probleme stoßen, zögern Sie nicht, einen von uns beiden zu kontaktieren.
