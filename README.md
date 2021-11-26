<h1>
<!--   <img src="https://raw.githubusercontent.com/create-go-app/cli/master/.github/images/cgapp_logo%402x.png" width="224px"/><br/> -->
  beehive
</h1>
<p><a href="https://github.com/quarterblue/beehive/releases" target="_blank"><img src="https://img.shields.io/badge/version-v1.0.1-blue?style=for-the-badge&logo=none" alt="cli version" /></a>&nbsp;<a href="https://pkg.go.dev/github.com/quarterblue/beehive/v1?tab=doc" target="_blank"><img src="https://img.shields.io/badge/Go-1.17+-00ADD8?style=for-the-badge&logo=go" alt="go version" /></a>&nbsp;<img src="https://img.shields.io/badge/license-mit-blue?style=for-the-badge&logo=none" alt="license" /></p>
<p>Distribute hundreds of repeatable and periodic <b>cron based jobs</b> over multiple worker nodes by encapsulating jobs in <b>Docker Containers</b>. Coordinate worker nodes to redistribute work in case of crash failures. At-least-once-semnatics guarantee for all jobs scheduled using beautiful and rich frontend UI.</p>

🚧 The project is a <b>work in progress</b>, expect bugs, safety issues, and components that don't work. Refer to Todo list for progress.


## Features

- Repeatable, cron based job scheduler
- Docker containerized jobs
- Geographically distributed workers
- Full and Rich UI interface, authentication and multi user capability
- Horizontally Scalable to hundreds of nodes
- Fault tolerant worker architecture*

## Systems Architecture

<p align="center">
        <img width="100%" src="https://raw.githubusercontent.com/quarterblue/beehive/19039281bea6ffdc6f1fd5963405597825358fcb/static/systemdesigndiagram.svg?token=ANKI23KF75RZLWD4MTKJGADBUDGOW" alt="Parsec logo">
</p>

<b><a href="https://github.com/quarterblue/pulse">*Pulse<a></b>: is a hybrid failure detector library

## Quick start

### 🦖 Starting Coordinator

First of all, [download](https://www.docker.com/products/docker-desktop) and install 🐳 **Docker**. This is required since all moving parts of the application is containerized.

```bash
$ git clone github.com/quarterblue/beehive
```

```bash
$ cd /beehive
$ docker-compose build
$ docker-compose up -d

```

> This will start Frontend, Coordinator, Single Cluster Node of Cockroach DB, RabbitMQ, Updater, and Dispatcher all on the same machine.

If you want to separate these services on different machines, refer to docker-compose directory for various options.

```bash
$ cd beehive/docker-compose/
$ cat options
```

### 🐝 Starting Worker Bees

Worker Bees are the individual worker nodes usually hosted on a Cloud virtual machine.

```bash
$ git clone github.com/quarterblue/beehive
```

```bash
$ cd /beehive
$ docker-compose -f docker-compose-bee.yml build
$ docker-compose -f docker-compose-bee.yml up -d

```

Alternatively, you can host Cockroach DB Cluster along with your worker nodes to provide even more fault-tolerance!:

```bash
$ cd /beehive
$ docker-compose -f docker-compose-bee-cdb.yml build
$ docker-compose -f docker-compose-bee-cdb.yml up -d
```

That's all you need to know to start! 🎉

## Docker Components

Frontend

```yml
frontend:
  build: ./frontend
  container_name: 'frontend'
```

> Provide the user interface in forms of SPA application created in VueJS.

Coordinator

```yml
coordinator:
  build: ./src/coordinator
  container_name: 'coordinator'
```

> Provide coordinator node with REST API that schedules & (re)distributes jobs.

Kafka

```yml
kafka:
  image: confluentinc/cp-zookeeper:latest
  container_name: 'kafka'
```

> Provide the kafka topic queue to communicate between worker bees and updater.

Cockroach DB Clusters

```yml
cdb_node_1:
  container_name: cdb_node_1
  image: cockroachdb/cockraoch:latest
```

> Provide the DB cluster distributed over multiple nodes.

Updater

```yml
updater:
  image: ./src/updater
  container_name: 'updater'
```

> Provide the microservice that pulls updates from worker bees and commits to the DB.

Worker Bee

```yml
workerbee:
  image: ./src/worker
  container_name: 'workerbee'
```

> Provide the individual worker nodes on commodity hardware (or cloud virtual machine).


## License
Beehive is a free and open-source software licensed under the MIT License.
