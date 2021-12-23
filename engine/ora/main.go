package ora

/* Would that Oracle had taken a page from the PostgreSQL book and done
something like making ORA_ a reserved prefix and then used it consistently for
naming the DB system schemas/users...

The following does appear to catch most, maybe all, of them for versions > 11?

SELECT du.username
    FROM dba_users du
    WHERE du.oracle_maintained = 'Y'
        OR EXISTS (
            SELECT 1
                FROM sys.ku_noexp_tab s
                WHERE s.obj_type = 'SCHEMA'
                    AND s.name = du.username
        )
        OR EXISTS (
            SELECT 1
                FROM system.logstdby$skip_support s
                WHERE s.action IN ( 0, -1 )
                    AND s.name = du.username
        )
        OR EXISTS (
            SELECT 1
                FROM v$sysaux_occupants s
                WHERE s.schema_name = du.username
        )
;

*/

const systemTables = `
    'ANONYMOUS', 'APEX_190200', 'APEX_PUBLIC_USER', 'APPQOSSYS',
    'AUDSYS', 'AWR_STAGE', 'CSMIG', 'CTXSYS', 'DBSFWUSER', 'DBSNMP',
    'DIP', 'DMSYS', 'DSSYS', 'EXFSYS', 'FLOWS_FILES', 'GGSYS',
    'GSMADMIN_INTERNAL', 'GSMCATUSER', 'GSMUSER', 'LBACSYS', 'MDSYS',
    'OJVMSYS', 'OLAPSYS', 'ORACLE_OCM', 'ORDPLUGINS', 'ORDSYS',
    'OUTLN', 'PERFSTAT', 'PUBLIC', 'REMOTE_SCHEDULER_AGENT',
    'SQLTXPLAIN', 'SYS', 'SYSBACKUP', 'SYSDG', 'SYSKM', 'SYSMAN',
    'SYSRAC', 'SYSTEM', 'SYS$UMF', 'TRACESVR', 'TSMSYS', 'WMSYS',
    'XDB', 'XS$NULL'
`
