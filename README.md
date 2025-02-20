# Сервис для интеграции Git-вебхуков с Jenkins

Этот сервис предназначен для приема вебхуков от Git-сервера (например, Gitea, GitHub, GitLab и т.д.) и перенаправления их в Jenkins для запуска соответствующих задач (jobs). Сервис настраивается с помощью YAML-конфигурации, где указываются параметры подключения к Jenkins, а также список репозиториев и задач, которые должны запускаться при получении вебхуков.

---

## Конфигурация

Конфигурация сервиса задается в файле `config.yaml`. Пример конфигурации:

```yaml
server_port: 8081
jenkins:
  url: "https://jenkins.example.com"
  user: "mylogin"
  pass: "mypass"
  token: "myjobtoken"
  allowed_webhooks:
    - repo_name: ansible
      run_jobs:
        - job_path: jobs-dsl/jobs-dsl
          parameterized_job: true
        - job_path: another-job/another-job
          parameterized_job: true
    - repo_name: "repo/repo-sandbox"
      run_jobs:
        - job_path: "job/Sandbox/job/playground/job/myjobname"
          parameterized_job: true
```

---

## Как работает сервис
Сервис ожидает POST-запросы на эндпоинт `/webhook/gitea`.

При получении вебхука сервис проверяет:

Соответствует ли имя репозитория одному из указанных в `allowed_webhooks`.

Если репозиторий разрешен, сервис отправляет запрос в Jenkins для запуска соответствующих задач.

Для параметризованных задач (`parameterized_job: true`) сервис передает параметры из вебхука в Jenkins.

---

## Запуск сервиса в Docker контейнере

```
docker build -t webhook-receiver .
docker run -ti -v $(pwd)/config.yaml:/app/config.yaml:ro --rm webhook-receiver
```

---

## License

This package is available under the [MIT license](LICENSE).
