# 変数定義
DOCKER_COMPOSE_FILE = docker/docker-compose.yml
ENV_FILE = docker/.env

## 環境設定
setup: ## 初期セットアップ (.env ファイルの作成)
	cp docker/.env.example $(ENV_FILE); \

## Docker操作
up: ## Docker コンテナを起動
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

down: ## Docker コンテナを停止・削除
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

build: ## Docker イメージをビルド
	docker-compose -f $(DOCKER_COMPOSE_FILE) build

rebuild: ## Docker イメージを再ビルドして起動
	docker-compose -f $(DOCKER_COMPOSE_FILE) build --no-cache
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

## ログ確認
logs: ## 全コンテナのログを表示
	docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f

logs-go: ## Goアプリのログを表示
	docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f go-app

logs-db: ## PostgreSQLのログを表示
	docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f postgres

## コンテナ操作
shell-go: ## Goコンテナにシェルで接続
	docker-compose -f $(DOCKER_COMPOSE_FILE) exec go-app sh

shell-db: ## PostgreSQLコンテナにシェルで接続
	docker-compose -f $(DOCKER_COMPOSE_FILE) exec postgres bash
