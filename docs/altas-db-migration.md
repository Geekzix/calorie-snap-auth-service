# Altas DB Migration

1. Check schema
    ```bash
    atlas schema inspect --env gorm --url "env://src"                                                       
    ```
2. Generate migration file
    ```bash
    atlas migrate diff --env gorm  
    ```
3. Apply the migration
    ```bash
    atlas migrate apply --url "$DATABASE_URL" --dir "file://migrations"
    ```
4. Check migration status
    ```bash
    atlas migrate status --url "$DATABASE_URL" --dir "file://migrations"
    ```

