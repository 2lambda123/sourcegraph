# Deployment Overview

 Whether you are setting up a proof of concept, or looking for a product-level deployment, this overview will provide you with an introduction to key items to consider.

For a list of all deployment topics, visit our [Deployment Table of Contents](index), and if you're just starting out, you can [try Sourcegraph Cloud](https://sourcegraph.com) or for a quick test you may consider running Sourcegraph locally via our [Docker single container](../install/docker/index.md) type. 

## Resource planning

Sourcegraph has provided the [Resource Estimator](resource_estimator.md) as a starting point to determine necessary resources based on the size of your deployment. 

As a recommendation, if you are planning deployment scenario will include very large codebases and a large number of users, our [Kubernetes](../../admin/install/kubernetes/scale) Deployment option will be your best option.

## Options and scenarios

| Deployment Type                                          | Suggested for                                           | Setup time      | Resource isolation | Auto-healing | Multi-machine |
| -------------------------------------------------------- | ------------------------------------------------------- | --------------- | :----------------: | :----------: | :-----------: |
| [**Docker Compose**](../install/docker-compose/index.md) | **Small & medium** production deployments               | 🟢 5 minutes     |         ✅          |      ✅       |       ❌       |
| [**Kubernetes**](../install/kubernetes/index.md)         | **Medium & large** highly-available cluster deployments | 🟠 30-90 minutes |         ✅          |      ✅       |       ✅       |
| [**Single-container**](../install/docker/index.md)       | Local testing                                           | 🟢 1 minute      |         ❌          |      ❌       |       ❌       |

Each of the types listed in the table above provide a different level of capability and should be approached based on the needs of your business as well as the technical expertise you have access to.

Given this, we have a few recommendations.

- **Docker Compose** - We recommend this path for initial production deployments. If your requirements change, you can always [migrate to a different deployment type](deployment_overview.md#migrating-to-a-new-deployment-type) later on if needed.

- **Kubernetes** - We recommend Kubernetes for large entreprises that depend or have an expecation for highly scalable deployments. It is important to note that if you're looking to deploy via the Kubernetes path, you are **expected to have a team that is familiar with operating Kubernetes clusters**, including but not limited to the use of persistent storage. If there is any doubt about your team's ability to support this, please speak to your Sourcegraph contact about using Docker Compose instead.

- **Single Container** - Finally, if you're just starting out, you can [try Sourcegraph Cloud](https://sourcegraph.com) or [run Sourcegraph locally](../install/docker/index.md).

> NOTE: The Single container option is provided for local proof-of-concepts and not intended for testing or deployed at a pre-production/production level. Some features, such as Code Insights are not available when using this deployemnt type. If you're just starting out, and want the absolute quickest setup time, [try Sourcegraph Cloud](https://sourcegraph.com).

## External services

Sourcegraph by default provides versions of services it needs to operate, including:

- A [PostgreSQL](https://www.postgresql.org/) instance for storing long-term information, such as user information when using Sourcegraph's built-in authentication provider instead of an external one.
- A second PostgreSQL instance for storing large-volume precise code intelligence data.
- A [Redis](https://redis.io/) instance for storing short-term information such as user sessions.
- A second Redis instance for storing cache data.
- A [MinIO](https://min.io/) instance that serves as a local S3-compatible object storage to hold user uploads before they can be processed. _This data is for temporary storage and content will be automatically deleted once processed._
- A [Jaeger](https://www.jaegertracing.io/) instance for end-to-end distributed tracing. 

> NOTE: Your Sourcegraph instance can be configured to use an external or managed version of these services. Using a managed version of PostgreSQL can make backups and recovery easier to manage and perform. Using a managed object storage service may decrease your hosting costs as persistent volumes are often more expensive than object storage space.

See the following guides to use an external or managed version of each service type.

- See [Using your own PostgreSQL server](../external_services/postgres.md) to replace the bundled PostgreSQL instances.
- See [Using your own Redis server](../external_services/redis.md) to replace the bundled Redis instances.
- See [Using a managed object storage service (S3 or GCS)](../external_services/object_storage.md) to replace the bundled MinIO instance.
- See [Using an external Jaeger instance](../observability/tracing.md#use-an-external-jaeger-instance) in our [tracing documentation](../observability/tracing.md) to replace the bundled Jaeger instance.Use-an-external-Jaeger-instance

> NOTE: Using Sourcegraph with an external service is a [paid feature](https://about.sourcegraph.com/pricing). [Contact us](https://about.sourcegraph.com/contact/sales) to get a trial license.

### Cloud alternatives

- Amazon Web Services: [AWS RDS for PostgreSQL](https://aws.amazon.com/rds/), [Amazon ElastiCache](https://aws.amazon.com/elasticache/redis/), and [S3](https://aws.amazon.com/s3/) for storing user uploads.
- Google Cloud: [Cloud SQL for PostgreSQL](https://cloud.google.com/sql/docs/postgres/), [Cloud Memorystore](https://cloud.google.com/memorystore/), and [Cloud Storage](https://cloud.google.com/storage) for storing user uploads.
- Digital Ocean: [Digital Ocean Managed Databases](https://www.digitalocean.com/products/managed-databases/) for [Postgres](https://www.digitalocean.com/products/managed-databases-postgresql/), [Redis](https://www.digitalocean.com/products/managed-databases-redis/), and [Spaces](https://www.digitalocean.com/products/spaces/) for storing user uploads.

## Configuration (TBD)

Configuration at the deployment level focuses on ensuring your Sourcegraph runs optimally based on the size of your repositories and number of users. Configuration options will vary based on the type of deployment you choose, so you will want to consult the specific configuration guides for additional information.

If you're looking for configuration at the Administration level, check out the [customization section TBD](TBD.


## Upgrades and migration

A new version of Sourcegraph is released every month (with patch releases in between, released as needed). Check the [Sourcegraph blog](https://about.sourcegraph.com/blog) or the site admin updates page to learn about updates. We actively maintain the two most recent monthly releases of Sourcegraph.

**Regardless of your deployment type**, the following rules apply:

- **Upgrade one minor version at a time**, e.g. v3.26 --> v3.27 --> v3.28.
  - Patches (e.g. vX.X.4 vs. vX.X.5) do not have to be adopted when moving between vX.X versions.
- **Check the [update notes for your deployment type](#update-notes) for any required manual actions** before updating.
- Check your [out of band migration status](../migration/index.md) prior to upgrade to avoid a necessary rollback while the migration finishes.

### Upgrade notes

Please see the instructions for your deployment type:

- [Sourcegraph with Docker Compose](docker_compose.md)
- [Sourcegraph with Kubernetes](kubernetes.md)
- [Single-container Sourcegraph with Docker](server.md)
- [Pure-Docker custom deployments](pure_docker.md)

### Migrating to a new deployment type

In some cases you may need to migrate from one deployment type to another. Most commonly this will be [from a Docker Single Container deployment to Docker Compose](../install/docker-compose/migrate.md), or from [Docker Compose to Kubernetes](please contact support).

You may also want to checkout our [Migration Docs area](../migration/index.md) for other scenarios and associated guides especially specific migration requirements when upgrade from/to specfic versions.


### Changelog

For product update notes, please refer to the [changelog](../../CHANGELOG.md).

## Reference repositories

For Docker Compose and Kubernetes deployments, Sourcegraph provides reference repositories with branches corresponding to the version of Sourcegraph you wish to deploy. Depending on your deployment type, the reference repository contains everything you need to spin up and configure your instance. This will also assist in your upgrade process going forward. 

For more information, follow the install and configuration docs for your specific deployment type: [Docker Compose](https://github.com/sourcegraph/deploy-sourcegraph-docker/) or [Kubernetes](https://github.com/sourcegraph/deploy-sourcegraph/) .