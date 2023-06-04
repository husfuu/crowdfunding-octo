```bash
docker run --rm \
--name husfuu-dev \
-e POSTGRES_DB=crowdfunding_dev  \
-e POSTGRES_USER=assignment \
-e POSTGRES_PASSWORD=mysecretpassword \
-e PGDATA=/var/lib/postgresql/data/pgdata \
-v "$PWD/husfuu-dev:/var/lib/postgresql/data" \
-p 5432:5432 \
postgres:13
```

```bash
docker run \
--name husfuu-dev \
-e POSTGRES_DB=crowdfunding_dev  \
-e POSTGRES_USER=assignment \
-e POSTGRES_PASSWORD=mysecretpassword \
-e PGDATA=/var/lib/postgresql/data/pgdata \
-v "$PWD/husfuu-dev:/var/lib/postgresql/data" \
-p 5432:5432 \
postgres:13
```

FLOW
1. Repository -> define Contructor
2. API.go (one of repository type, ex: cms, core, general)
    - define Contructor
    - register all repository (domain) in the contructor
3. Wrapper Repository
    - define contructor
    - register all type of repository