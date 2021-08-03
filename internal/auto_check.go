package internal

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AutoCheck struct {
	db         *gorm.DB
	config     MinersConfigJson
	repository *Repository
}

func NewAutoCheck(config MinersConfigJson) (*AutoCheck, error) {
	if config.AutoCheck.Enabled {
		switch config.AutoCheck.DBConfig.Type {
		case "postgres":
			db, err := gorm.Open(postgres.Open(config.AutoCheck.DBConfig.URL), &gorm.Config{})
			if err != nil {
				return nil, err
			}
			return &AutoCheck{
				db:         db,
				config:     config,
				repository: NewRepository(db),
			}, nil
		default:
			return nil, fmt.Errorf("DB type %s not implemented", config.AutoCheck.DBConfig.Type)
		}

	} else {
		return nil, nil
	}

}

func (a *AutoCheck) Start() {
	// Update DDL of db
	err := a.db.AutoMigrate(&MinerResult{})
	if err != nil {
		log.Panicf("Error on autmigrate DB %#v", err)
	}

	for {
		time.Sleep(time.Second * time.Duration(a.config.AutoCheck.Interval))

		results, err := DoMinersCheck(a.config)
		if err != nil {
			log.Printf("Error auto-check %#v", err)
		}
		for _, result := range results.Miners {
			_, err := a.repository.Create(result)
			if err != nil {
				log.Printf("Error on saving on DB %#v\n", err)
			}
		}
	}
}
