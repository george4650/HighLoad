package oracle

import (
	"gitlabnew.nextcontact.ru/nextcontactcenter/services/highload-testers/db-highload-tester/config"
	oracle "gitlabnew.nextcontact.ru/r.alfimov/go-oracle"
	"time"
)

func GenerateNewConnection(cfg config.DatabaseOracle) (*oracle.Oracle, error) {
	// Инициализация Oracle
	oracleDb, err := oracle.New(&oracle.Config{
		Host:     cfg.Host,
		Port:     cfg.Port,
		Service:  cfg.Service,
		User:     cfg.User,
		Password: cfg.Password,
		NLSQueries: []string{
			"ALTER SESSION SET NLS_DATE_FORMAT = 'dd.mm.yyyy'",
			"ALTER SESSION SET NLS_TIMESTAMP_FORMAT = 'dd.mm.yyyy hh24:mi:ss'",
			"ALTER SESSION SET NLS_TIMESTAMP_TZ_FORMAT = 'dd.mm.yyyy hh24:mi:ss tzr'",
			"ALTER SESSION SET NLS_NUMERIC_CHARACTERS = '. '",
			"ALTER SESSION SET NLS_SORT = 'RUSSIAN'",
			"ALTER SESSION SET NLS_LANGUAGE = 'RUSSIAN'",
			"ALTER SESSION SET NLS_COMP = 'BINARY'",
		},
		CheckConnection: &oracle.CheckConnectionConfig{
			ReconnectTryCount:       5,
			ReconnectTryInterval:    3 * time.Second,
			CheckConnectionInterval: 1 * time.Second,
			HealthQuery:             "select 1 from dual",
		},
		CustomParams: map[string]string{
			"FAILOVER": "3",
			//"TRACE FILE": "trace.log",
			//"DBA PRIVILEGE": "SYSDBA",
		},
	})
	if err != nil {
		return nil, err
	}
	// END Инициализация Oracle

	return oracleDb, nil
}
