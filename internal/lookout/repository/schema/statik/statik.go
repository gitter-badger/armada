// Code generated by statik. DO NOT EDIT.

package statik

import (
	"github.com/rakyll/statik/fs"
)

const LookoutSql = "lookout/sql" // static asset namespace

func init() {
	data := "PK\x03\x04\x14\x00\x08\x00\x00\x00&kGQ\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x16\x00	\x00001_initial_schema.sqlUT\x05\x00\x01\xb9\xc1}_CREATE TABLE job\n(\n    job_id    varchar(32)  NOT NULL PRIMARY KEY,\n    queue     varchar(512) NOT NULL,\n    owner     varchar(512) NOT NULL,\n    jobset    varchar(512) NOT NULL,\n\n    priority  float        NULL,\n    submitted timestamp    NULL,\n    cancelled timestamp    NULL,\n\n    job       jsonb        NULL\n);\n\nCREATE TABLE job_run\n(\n    run_id    varchar(36)  NOT NULL PRIMARY KEY,\n    job_id    varchar(32)  NOT NULL,\n\n    cluster   varchar(512) NULL,\n    node      varchar(512) NULL,\n\n    created   timestamp    NULL,\n    started   timestamp    NULL,\n    finished  timestamp    NULL,\n\n    succeeded bool         NULL,\n    error     varchar(512) NULL\n);\n\nCREATE TABLE job_run_container\n(\n    run_id         varchar(32) NOT NULL,\n    container_name varchar(512) NOT NULL,\n    exit_code      int         NOT NULL,\n    PRIMARY KEY (run_id, container_name)\n)\n\n\nPK\x07\x08\x04\x84R1`\x03\x00\x00`\x03\x00\x00PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00&kGQ\x04\x84R1`\x03\x00\x00`\x03\x00\x00\x16\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x00\x00\x00\x00001_initial_schema.sqlUT\x05\x00\x01\xb9\xc1}_PK\x05\x06\x00\x00\x00\x00\x01\x00\x01\x00M\x00\x00\x00\xad\x03\x00\x00\x00\x00"
	fs.RegisterWithNamespace("lookout/sql", data)
}
